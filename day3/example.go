package main

import "fmt"

type student struct {
	name string
	rno float64
	dept string
}
func main() {

	st:=student{name : "jon",rno : 25 , dept : "cs"}
	fmt.Println(st.name)

}

func ifelsedemo() {
	var a, b int

	fmt.Println("Enter the two numbers:")
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	if a > b {
		fmt.Println("a is greater than b")
	} else if a == b {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println("a is lesser than b")

	}
}

func fordemo(){
	n := 1
	for n=2;n<6;n++{
		fmt.Println(n)
		
	}
}
