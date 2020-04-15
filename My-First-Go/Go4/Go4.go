package main

import "fmt"

func main() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 2 {
			continue
		}
		println("continuing...")
	}

	var j int
	for ; j < 5; j++ {
		fmt.Println(j)
	}

	for k := 0; k < 5; k++ {
		fmt.Println(k)
	}

	var l int
	for l = 0; l < 5; l++ {
		fmt.Println(l)
	}

	var m int
	for {
		if m == 5 {
			break
		}
		fmt.Println(m)
		m++
	}

	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	} //similar for maps as well

	wellPorts := map[string]int{"http": 8080, "https": 3000}
	for k, v := range wellPorts {
		println(k, v)
	}

	for k := range wellPorts {
		println(k)
	}
	for _, v := range wellPorts {
		println(v)
	}

	fmt.Println("Hello")
	//panic("Dump")
	fmt.Println("Bye")
}
