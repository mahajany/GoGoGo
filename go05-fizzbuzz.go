//Fizbuzz
package main
import (
    "fmt"
)

func main(){
    for x:=1; x< 20; x++ {
		switch {
		case x%3==0 && x%5==0:
			fmt.Printf("fizz-buzz\n")
		case x%3==0:
			fmt.Printf("fizz\n")
		case x%5==0:
				fmt.Printf("buzz\n")
		default:
			fmt.Printf("%v\n",x)
		}
	}
}
