package main

import "fmt"

func main() {
	var i interface{}

	i = 4
	fmt.Printf("i=%v (%T)\n", i, i)

	i = "hi"
	fmt.Printf("i=%v (%T)\n", i, i)

	s := i.(string) // type assertion
	fmt.Println("s =", s)
}
