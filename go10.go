//map
package main
import (
    "fmt"
)

func main(){
    stocks := map[string] float64 {
        "TATA": 289.1,
        "TESLA" : 22.0,
        "GM" : 101.1,
    }
    fmt.Println(stocks["TESLA"])
    fmt.Println(stocks["GOOGLE"])

    value, found := stocks["GOOGLE"]
    if !found {
       fmt.Println("If Google is down, then don't use Google to search for Google!")
    } else {
        fmt.Println("Look Ma, St0nks", value)
    }

    stocks["GOOGLE"] = 888.0
    delete (stocks, "GM")

    for k,v := range stocks{
        fmt.Println(k,v)
    }

}
