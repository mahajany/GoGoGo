//For Loop
package main
import (
    "fmt"
)

func main(){
    for i:=1; i< 10; i++ {
		if i>5{
			break
		}
		if i>3{
			continue
		}
		fmt.Printf("%v\n",i)
	}

	a:=1
	for a < 4{
		fmt.Printf("%v..",a)
		a++
	}
}
