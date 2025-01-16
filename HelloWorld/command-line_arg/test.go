package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	greeting := fmt.Sprintf("Hello, %s!", name)
	fmt.Print(greeting)

}
