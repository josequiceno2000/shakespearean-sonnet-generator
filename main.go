package main

import (
	"fmt"
	"os"
	"bufio"
) 


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("SSG > ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println(name)
}