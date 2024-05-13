package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	user := User{
		Name:        "John",
		Age:         30,
		Password:    "password",
		Permissions: []string{"admin", "user"},
	}

	jsonData := `{"name":"John","age":30,"roles":["admin","user"]}`

	var person User
	err := json.Unmarshal([]byte(jsonData), &person)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the struct data is ", person, user)

	server := NewAPIServer(":8080")

	server.Run()

	// router := http.NewServeMux()

	// router.HandleFunc("/api/create", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("post user data is", r.Body)
	// })
	// fmt.Println("Server is running on port 8080")
	// http.ListenAndServe(":8080", nil)

}
