package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"foo": 42, "haapy": 1}
	fmt.Println(m["foo"])
	
	m["foo"] = 298
	fmt.Println(m["foo"])
	
	delete(m, "foo")
	fmt.Println(m["foo"])
	fmt.Println(m)
	
	
}
