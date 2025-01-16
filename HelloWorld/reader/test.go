package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any greeting : ")
	greeting, _ := reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)
	fmt.Println(greeting)
}
