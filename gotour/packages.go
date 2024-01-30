package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)


func square_root(number float64) float64{
	return math.Sqrt(number)
}

func get_pi() float64 {
	return math.Pi
}

func swap(a, b string ) (string, string){
	return b, a
}

func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x
	z = 99
	return
}

var(
	z complex128 = cmplx.Sqrt(-5+12i)
)




func main() {
	fmt.Println("My fav number is the", rand.Intn(100000))

	a:=rand.Intn(10)
	fmt.Println(a)
	fmt.Println(rand.Intn(100))

	b := square_root(49)
	fmt.Println(b)

	c := get_pi()
	fmt.Println(c)

	fmt.Println(swap("rustem", "best"))

	

	fmt.Println(split(17))
	fmt.Println(z)



}
