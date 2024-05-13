package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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

		w.Write([]byte("user id is " + userId))
	})

	// create the post

	router.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {

		var userData User

		body, err := io.ReadAll(r.Body) // instead of reading all encoder can be the good choice
		if err != nil {
			log.Fatal("error while reading the body", err)
		}

		err = json.Unmarshal(body, &userData)
		if err != nil {
			log.Fatal("erorr while parsing the body", err)
		}
		s.db.Exec("INSERT INTO users (name, age, password, roles) VALUES ($1, $2, $3, $4) returning *", userData.Name, userData.Age, userData.Password, userData.Permissions)
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
