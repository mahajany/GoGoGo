// Cacluate the average of two numbers
package main

import(
	"fmt"
)
func main(){
	var x float64
	var y int

	x=1
	y=99
	fmt.Printf("x=%v, type of x %T \n", x,x)
	fmt.Printf("y=%v, type of y %T \n", y,y)
	var mean int
	mean = (x+y)/2
	fmt.Printf("Mean of %v and %v is %v, type is %T\n", x,y,mean, mean)

}
