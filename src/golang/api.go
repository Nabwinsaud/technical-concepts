package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type User struct {
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`
	Permissions []string `json:"roles"`
}

type APIServer struct {
	addr string
	db   *sql.DB
}

// var db *sql.DB

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}

}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("POST /user/auth", func(w http.ResponseWriter, r *http.Request) {
		var userData User

		body := json.NewDecoder(r.Body)

		if err := body.Decode(&userData); err != nil {
			fmt.Println("error while parsing the data")
			http.Error(w, "Error while reading the data ", http.StatusBadRequest)
		}

		accessToken, err := CreateToken(userData.Name)

		if err != nil {
			fmt.Println("Invalid token")
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		w.Write([]byte(accessToken))

	})

	router.HandleFunc("GET /user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		log.Println("post user data is", r.Body)
		userId := r.PathValue("userId")

		row := s.db.QueryRow("SELECT name,password,roles,age FROM users WHERE id=$1", userId)

		// if err != nil {
		// 	http.Error(w, "error while fetching the data", http.StatusInternalServerError)
		// 	log.Fatal("error while fetching the data", err)
		// 	return
		// }

		var user User
		// err := row.Scan(&user.Name, &user.Password, &user.Permissions, &user.Age)
		err := row.Scan(&user.Name, &user.Password, pq.Array(&user.Permissions), &user.Age)

		if err != nil {
			http.Error(w, "error while scanning the data", http.StatusInternalServerError)
			log.Fatal("error while scanning the data", err)
			return
		}
		userData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "error while marshalling the data", http.StatusInternalServerError)
			log.Fatal("error while marshalling the data", err)
			return
		}
		w.Write(userData)
		// w.Write([]byte("user id is " + userId))
	})

	// create the post

	router.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {

		var userData User

		// decoder := json.NewDecoder(r.Body)
		// if err := decoder.Decode(&userData); err != nil {
		// 	http.Error(w, "Error reading request body", http.StatusBadRequest)
		// 	log.Fatal("error while decoding the body", err)
		// 	return
		// }
		body, err := io.ReadAll(r.Body) // instead of reading all encoder can be the good choice
		if err != nil {
			log.Fatal("error while reading the body", err)
		}

		err = json.Unmarshal(body, &userData)
		if err != nil {
			log.Fatal("error while parsing the body", err)
		}
		rolesString := "{" + strings.Join(userData.Permissions, ",") + "}"
		_, err = s.db.Exec("INSERT INTO users (name, age, password, roles) VALUES ($1, $2, $3, $4) ", userData.Name, userData.Age, userData.Password, rolesString)
		if err != nil {
			http.Error(w, "error while inserting the data", http.StatusInternalServerError)
			log.Fatal("error while inserting the data", err)
			return
		}
		fmt.Println("user data is ", userData)
	})

	router.HandleFunc("PUT /user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		var userData User

		body := json.NewDecoder(r.Body)

		if err := body.Decode(&userData); err != nil {
			fmt.Println("error while decoding the body", err)
			http.Error(w, "error parsing the json body", http.StatusBadRequest)
		}

		_, err := s.db.Exec("UPDATE users SET name=$1,age=$2,roles=$3 where id=$4", userData.Name, userData.Age, pq.Array(userData.Permissions), r.PathValue("userId"))

		if err != nil {
			fmt.Println("error while updating the data", err)
			http.Error(w, "error while updating the data", http.StatusInternalServerError)
			return
		}

		// return the success message with json
		w.Write([]byte("user data is updated successfully"))

	})

	router.HandleFunc("DELETE /user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		_, err := s.db.Exec("DELETE FROM users WHERE id=$1", r.PathValue("userId"))

		if err != nil {
			fmt.Println("error while deleting the data", err)
			http.Error(w, "error while deleting user ", http.StatusInternalServerError)
		}
		w.Write([]byte("user deleted successfully"))
	})
	server := http.Server{
		Addr: s.addr,
		// Handler: router,
		// Handler: RequestLoggerMiddleware(router),
		Handler: AuthMiddleWare(RequestLoggerMiddleware((router))),
	}

	// working with various types of incoming method like pathValue,...

	router.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
		// page,perPage,limit
		page := r.URL.Query().Get("page")
		perPage := r.URL.Query().Get("perPage")
		limit := r.URL.Query().Get("limit")

		fmt.Println("page ,perPage,limit", page, perPage, limit)

		rows, err := s.db.Query("SELECT name,age,roles FROM users")

		var users []User

		defer rows.Close()

		var usr User

		for rows.Next() {
			if err := rows.Scan(&usr.Name, &usr.Age, pq.Array(&usr.Permissions)); err != nil {
				fmt.Println("errors scanning the data ", err)
				http.Error(w, "Error reading the data from databases", http.StatusInternalServerError)
				return
			}

			users = append(users, usr)
		}

		if err != nil {
			fmt.Println("error while getting user ", err)
			http.Error(w, "Error while fetching from db", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "applications/json")

		if err = json.NewEncoder(w).Encode(users); err != nil {
			fmt.Println("error parsing the json ", err)
			http.Error(w, "error parsing the json", http.StatusBadRequest)
		}

		fmt.Println("user details  is", users)
		// w.Write)
	})
	log.Println("server running on port 8080")
	return server.ListenAndServe()

}

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("request is ", r.URL.Path)
		log.Printf("method is %s,path is %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		tokenParts := strings.Split(token, " ")

		fmt.Println("the access token is ", tokenParts)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "invalid authorization", http.StatusUnauthorized)
		}
		err := VerifyToken(tokenParts[1])
		if err != nil {
			http.Error(w, "Invalid token or jwr malformed", http.StatusBadRequest)
			return
		}
		// if token != `Bearer ` {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }
		log.Printf("auth middleware is called ....")
		next.ServeHTTP(w, r)
	})
}
