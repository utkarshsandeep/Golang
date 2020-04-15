package main

import (
	"fmt"
	"inter/test/organization"
)

func main() {
	p := organization.NewFunction("Utkarsh", "Singh")

	/*var q organization.Identifiable = organization.Person{}
	fmt.Println(q.ID())*/

	fmt.Println(p.ID())
	fmt.Println(p.FullName())
	//err := p.SetTwitterHandler(organization.TwitterHandler("@utkarshsandeep"))
	//Another way of checking TwitterHandler functionality
	err := p.SetTwitterHandler("@utkarshsandeep")
	fmt.Printf("%T\n", organization.TwitterHandler("a"))
	if err != nil {
		fmt.Printf("An error occured setting twitter handler %s\n", err.Error())
	}
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectURL())
}
