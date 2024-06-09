package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	area := Circle{
		radius: 10,
	}

	area.radius = 10000

	area1 := Circle{
		radius: 10,
	}

	// arr := Array_stuffs()
	// fmt.Println("Array stuffs is ")
	fmt.Println("Area of the circle is ", area.Area(), area1.Area(), area.CalculateArea())

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

	db, err := InitializeDD("localhost", "5432", "postgres", "nabin", "godb")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	server := NewAPIServer(":8080", db)

	server.Run()

	// router := http.NewServeMux()

	// router.HandleFunc("/api/create", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("post user data is", r.Body)
	// })
	// fmt.Println("Server is running on port 8080")
	// http.ListenAndServe(":8080", nil)

}
