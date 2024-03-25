package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	msg := "Namaskaram, Duniya!"

	fmt.Println(msg)
	fmt.Println(reverse.String(msg))
}
