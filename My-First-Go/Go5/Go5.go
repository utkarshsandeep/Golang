package main

import "fmt"

type User struct {
	ID    int
	Fname string
	Lname string
}

type HTTPRequest struct {
	Method string
}

func main() {
	u1 := User{
		ID:    1,
		Fname: "Utkarsh",
		Lname: "Singh",
	}
	u2 := User{
		ID:    2,
		Fname: "Sandeep",
		Lname: "Singh",
	}
	if u1 == u2 {
		fmt.Println("Same User")
	} else {
		fmt.Println("Different User")
	}

	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		fmt.Println("Get Request")
		fallthrough
	case "DELETE":
		fmt.Println("Delete Request")
	case "POST":
		fmt.Println("Post Request")
	case "PUT":
		fmt.Println("Put Request")
	default:
		fmt.Println("Unhandled")
	}
}
