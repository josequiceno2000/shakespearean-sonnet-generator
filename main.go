package main

import (
	"fmt"
	"os"
	"strconv"
) 


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ssg <name>")
		os.Exit(1)
	}
	arg1 := os.Args[1]

	maxLength, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid integer.\n", arg1)
		os.Exit(1)
	}
	
	for i := 1; i <= maxLength; i++ {
		fmt.Printf("%v. When I do count the clock that tells the time\n", i)
	}
	
}