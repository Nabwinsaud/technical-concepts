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

	router.HandleFunc("GET /user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		log.Println("post user data is", r.Body)
		userId := r.PathValue("userId")

		w.Write([]byte("user id is " + userId))
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
