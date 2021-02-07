// Cacluate the average of two numbers
package main

import(
	"fmt"
)
func main(){
	x := 10.1
	if  x > 5{
		fmt.Printf("%v is greater than 5\n", x)
	}

	y:=10.2
	if (x+y)>15{
		fmt.Printf("%v is what?\n",x+y)
	}

	n:=11
	switch n{
	case 1:
		fmt.Printf("n is 1")
	case 2:
		fmt.Printf("n is 2")
	default:
		fmt.Printf("n is - unknown\n")

	}
	switch {
	case n<10:
		fmt.Printf("n is less than 10")
	case n>10:
		fmt.Printf("n is greater than 10")
	default:
		fmt.Printf("n is - unknown\n")

	}

}
