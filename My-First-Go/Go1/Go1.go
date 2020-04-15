package main

import (
	"fmt"
)

const(
	first=1
	second="sec"

	a1=iota
	a2=iota
	a3=iota+6
	//a4=2<<iota
	a5//it will give next value from previous value onwards
)
const(
	a4=iota
)

func main() {
	var i int
	i=42
	fmt.Println(i)
	var f float32=3.14
	fmt.Println(f)
	fname:="Utkarsh"
	fmt.Println(fname)
	c:=complex(3,4)
	fmt.Println(c)
	r,im:=real(c),imag(c)
	fmt.Println(r,im)

	var fn *string=new(string)
	*fn="Utkarsh"
	fmt.Println(fn)

	fna:="Sandeep"
	var1:=&fna
	fmt.Println(var1,*var1)
	fna="Singh"
	var1=&fna
	fmt.Println(var1,*var1)

	const pi =3
	fmt.Println(pi)
	//pi=1.2
	fmt.Println(pi+4)
	fmt.Println(pi+2.2)

	const pic int =3
	fmt.Println(pic)
	//pi=1.2
	fmt.Println(pic+4)
	fmt.Println(float32(pic)+2.2)

	const pid ="hello"
	fmt.Println(pid)
	//pi=1.2
	fmt.Println(pid+"a")
	fmt.Println(pid)

	fmt.Println(first,second)
	fmt.Println(a1,a2,a3,a4,a5)
}