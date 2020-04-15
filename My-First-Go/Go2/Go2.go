package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	arr[0]=1
	arr[1]=2
	arr[2]=3
	fmt.Println(arr)
	fmt.Println(arr[1])
	arr1:=[3]int{1,2,3}
	fmt.Println(arr1)

	slice:=arr[:]
	fmt.Println(arr,slice)
	arr[1]=22;
	slice[2]=33;
	fmt.Println(arr,slice)

	slice1:=[]int{1,2,3}
	slice1=append(slice1,4)
	fmt.Println(slice1)
	slice2:=slice1[1:]
	slice3:=slice1[:2]
	slice4:=slice1[1:2]
	fmt.Println(slice2,slice3,slice4)

	map1:=map[string]int{"utkarsh":18,"sandeep":20}
	fmt.Println(map1)
	fmt.Println(map1["utkarsh"])
	map1["utkarsh"]=22;
	fmt.Println(map1["utkarsh"])
	delete(map1,"utkarsh")
	fmt.Println(map1)

	type user struct{
		ID int
		Fname string
		Lname string
	}
	var u user
	fmt.Println(u)
	u.ID=10
	u.Fname="utkarsh"
	u.Lname="singh"
	fmt.Println(u)
	u2:=user{
		ID:11,
		Fname:"sandeep",
		Lname:"singh",//insert comma after completing data
	}
	fmt.Println(u2)
}