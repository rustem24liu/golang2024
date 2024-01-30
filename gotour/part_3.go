package main

import "my_photo-5.jpg"

import(
	"fmt"
	// "math"
)



type info struct{
	name string 
	age int
	ok bool
}


func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p // divide j through the pointer
	fmt.Println(j) // see the new value of j



	a, b, c := 5, 10, 15

	z := & a



	fmt.Println(a, b, c, *z)


	fmt.Println(info{"rustem", 18, true})

	v := info{"aslan", 11, false}

	v.name = "ersultan"
	
	fmt.Println(v)

	s := &v.name


	fmt.Println(*s)


	var arr [10]int	

	arr[1] = 99

	arr2 := [2]string{"first", "second"}

	fmt.Println(arr2)


	primes := []int{2, 3, 5, 7, 11}

	var q []int = primes[1:4]

	fmt.Println(q)


	str := []struct{
		name string
		age int
	}{
		{"rustem", 18},
		{"aslan", 11},
		{"ersultan", 5},
	}

	fmt.Println(str)


	f := []int{1,2, 3,4, 5, 6, 7, 8, 9, 10}
	
	g := f[1:]

	// c := f[1:]


	fmt.Println(len(f), len(g), cap(g))


	board := [][] string{
		[]string{"rus", "ers", "asl"},
		[]string{"18", "11", "5"},
		[]string{"", "", " "},
}

	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"

	fmt.Println(board)


	for _, i := range f{
		fmt.Println(i, " ")
	}
	
	jpg.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
}

