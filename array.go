package main

import (
	"fmt"
)

func main() {
	var arr [3] int
	arr[0] = 1

	arr[1] = 11

	arr[2] = 13
	fmt.Println(arr)
	
	ar2 := [3] int {1, 2, 333}
	fmt.Println(ar2)
	
	slice := arr[:]
	fmt.Println(slice)
	
	slice2 := []int {1, 222, 333, 34555}
	slice2 = 	append(slice2, 89)
	fmt.Println(slice2)
}
