//Repeat-number!
package main
import (
    "fmt"
)

func main(){

	count := 0
	for i := 1000; i<9999; i++{
		for j :=i; j<9999; j++{
			p := fmt.Sprintf("%v", i*j)
			// fmt.Println("%q == %q",  p[:len(p)/2] ,  p[len(p)/2:])
			if p[:len(p)/2] == p[len(p)/2:]{
				count++
			}
		}
	}
	fmt.Println("There are", count, " even-numbered number(s)")

	//Even-ended if n[0] = n[len(n)-1]
}
