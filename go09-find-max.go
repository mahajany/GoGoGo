//Slices - built on the top of array, noone uses arrays
package main
import (
    "fmt"
)

func main(){
    loons := [] int {2, 3, 66, 4, 2, 241 , 32}

    max := loons[0]
    for _, value := range loons {
        if value > max{
            max = value
        }
    }
    fmt.Println("Max value is", max)
}
