package main

import "fmt"

const(
	A int = 1
	B = 3.14
	C = "Cobi Kurt"
)

func main(){
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Printf("%T", B, "\n")
}