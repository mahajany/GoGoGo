//Fuction
package main
import (
    "fmt"
    "math"
)

func bong( a int , b int) {
    // c := a^b
    fmt.Println( a^b )
}

func add( a int , b int) int{
    return a + b
}

func divide( a int , b int) (int, int){
    return a/b, a%b
}

func sqrt(n float64)(float64, error){
    if n<0.0{
        return 0.0, fmt.Errorf("Square of negative numbers - not now!", n)
    }
    return math.Sqrt(n), nil
}

func main(){
    sum := add(3,44)
    fmt.Println("sum:", sum)

    bong(3, 444)

    div, mod := divide(55,9)
    fmt.Println( "2 values returned: ", div, " x ", mod)    

    n, err  := sqrt(-2)
    if err != nil{fmt.Println(err)} else { 
        fmt.Println("sqrt is ", n)}

    n, err  = sqrt(9)
    if err != nil{fmt.Println("...ouch!")} else { fmt.Println("sqrt is ", n)}

}
