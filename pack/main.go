package main

import (
	"demo/pack/book"
	"fmt"
)

func main() {
	b := book.New()
	fmt.Printf("%T\n %v", b, b)
}
