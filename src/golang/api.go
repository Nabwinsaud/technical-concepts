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
			log.Fatal("erorr while parsing the body", err)
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
	server := http.Server{
		Addr: s.addr,
		// Handler: router,
		// Handler: RequestLoggerMiddleware(router),
		Handler: AuthMiddleWare(RequestLoggerMiddleware((router))),
	}

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

		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		log.Printf("auth middleware is called ....")
		next.ServeHTTP(w, r)
	})
}
