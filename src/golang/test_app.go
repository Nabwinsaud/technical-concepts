package main

import "fmt"

func Array_stuffs() {
	var user [3]string = [3]string{"John", "Doe", "Smith"}

	for _, ele := range user {
		fmt.Println("user is ", ele)
	}
}
