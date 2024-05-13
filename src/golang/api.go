package main

import (
	"log"
	"net/http"
)

type User struct {
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`
	Permissions []string `json:"roles"`
}

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}

}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		log.Println("post user data is", r.Body)
		userId := r.PathValue("userId")
		w.Write([]byte("user id is:  " + userId))
	})
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Println("server running on port 8080")
	return server.ListenAndServe()

}
