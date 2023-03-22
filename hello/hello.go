package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(quote.Go())

	var x string = "Hey there"
	var y int = 3
	var z float32 = 3.4343

	var arr = [5]int{1, 0, 1, 2, 3}
	arr2 := [3]int{0: 1, 2: 15}
	arr2Slice := arr2[1:2]

	x = "Boi"
	w := 2

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(y * w)

	fmt.Println(arr)
	fmt.Println(arr2)

	fmt.Println(arr2Slice)
}
