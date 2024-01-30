package main

import (
	"fmt"
	// "math"
	"runtime"
	
)

// func pow(x, n)

func main(){
	sum := 0
	for i := 0; i < 10; i++{
		sum += i
	}

	n := 1
	for; n < 10;{
		n += n
	}

	fmt.Println(sum, n)

	// for i := 0; i > -1; i++{
	// 	fmt.Println("cout")
	// }
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	
default:
	fmt.Printf("%s.\n", os)
}

}