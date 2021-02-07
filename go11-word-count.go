//map
package main
import (
    "fmt"
    "strings"
)

func main(){

    text := `
    And in the days that came, 
    they were always there,
    If not then it was not indeed,
    some snow and somewhere,
    but why can't I have everything,
    "wondered" why,
    'wandered' how,
    went in with it,
    I everywhere!
    `

    count := map[string] int{}
    words := strings.Fields(text)
    for _, word := range words{
        fmt.Printf(word)
        count[strings.ToLower(word)]++
    }
    for k,v := range count{
        fmt.Println(k,v)
    }

}
