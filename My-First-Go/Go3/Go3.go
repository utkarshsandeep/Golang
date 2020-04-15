package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello Go!!!")
	startWebServer()
	port := 3000
	startWebService(port, 3)
	startWebServ("hi", port, 3)
	boolcheck(10)
	checkBool := boolcheck(20)
	fmt.Println(checkBool)
	checkE1 := erCheck(20)
	fmt.Println(checkE1)
	checkE2 := erCheck2(40)
	fmt.Println(checkE2)
	checkRv, checkRe := manyReturn(100)
	fmt.Println(checkRv, checkRe)
}
func startWebServer() {
	fmt.Println("Starting server")
	//server status
	fmt.Println("Server Started")
}
func startWebService(port int, noOfRetry int) {
	fmt.Println("Starting service")
	//server status
	fmt.Println("Service Started on port", port)
	fmt.Println("No of retries", noOfRetry)
}
func startWebServ(name string, port, noOfRetry int) {
	fmt.Println("Starting service")
	//server status
	fmt.Println(name)
	fmt.Println("Service Started on port", port)
	fmt.Println("No of retries", noOfRetry)
}
func boolcheck(p int) bool {
	fmt.Println("Hello P!!!", p)
	return true
}
func erCheck(p int) error {
	fmt.Println("Hello E!!!", p)
	return nil
}
func erCheck2(p int) error {
	fmt.Println("Hello E!!!", p)
	return errors.New("Someting went Wrong!")
}
func manyReturn(p int) (int, error) {
	fmt.Println("Hello E!!!", p)
	return p, nil
}
