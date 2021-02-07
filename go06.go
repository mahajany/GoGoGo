//Fizbuzz
package main
import (
    "fmt"
)

func main(){
	book := "Life of Pi"
	fmt.Println(book  , len(book))
	fmt.Printf("%T, book[0] = %v (type %T)\n",book, book[0], book[0])
	fmt.Printf("\n"+ book + ", " + book[1:3] + ", " + book[:8] + "," + book [1:])

	n := 42
	s := fmt.Sprintf("%d", n)
	fmt.Printf("n", "s = %s or %q?"+ "\n", s,s)

}
