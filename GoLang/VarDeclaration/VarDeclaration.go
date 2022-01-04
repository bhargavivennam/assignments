package main

import "fmt"

func main() {
	p := 25
	q := 20

	fmt.Println(p + q)

	var a int
	var b int
	a = 10
	b = 12
	fmt.Println(a + b)

	var x int = 20
	var y int = 22
	fmt.Println(x + y)

	str1 := "Hello\n"
	str2 := "Welcome to GO Language"
	fmt.Println(str1 + str2)

	var str3, str4 string
	str3 = "Bhargavi"
	str4 = "Vennam"
	fmt.Println("My full name is :", str3+" "+str4)

	var str5 string = "This program is related to variable declaration "
	var str6 string = "using GO programming language"
	fmt.Println(str5 + str6)
}
