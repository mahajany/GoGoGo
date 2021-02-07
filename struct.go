package main

import (
	"fmt"
)

func main() {
	type user struct {
	ID int
	FirstName string
	LastName string
	
	}
	var u user
	fmt.Println(u)
	u.ID = 1
	u.FirstName = "Yogesh"
	u.LastName = "Mahajan"
	fmt.Println(u)
	fmt.Println(u.FirstName)
	u2 := user{ID: 1, FirstName: "Yogesh", LastName : "Mahajan"}
	fmt.Println(u2)

}
