//Slices - built on the top of array, noone uses arrays
package main
import (
    "fmt"
)

func main(){
    loons := [] string {"bugs", "daffy", "dan"}

    fmt.Println(len(loons))
    fmt.Println(loons[1])
    fmt.Println(loons[:1])

    loons = append(loons, "taz") //<=============


    fmt.Println(loons[1:])

    for i, name := range loons {
        fmt.Println(i, name)
    }

    
    for _, name := range loons {
        fmt.Println( name)
    }
}
