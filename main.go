package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	x := 10
	name := "okey"
	isWorking := true
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(x, name, isWorking)
}
