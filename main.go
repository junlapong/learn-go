package main

import (
	"fmt"
	"learn-go/hello"
)

func init() {
	fmt.Println("Initial block")
}

func main() {
	fmt.Println(hello.Hello("Jun", "spanish"))
}
