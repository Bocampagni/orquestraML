package main

import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args) > 1) {
		fmt.Println("Hello, " + os.Args[1] + "!")
		os.Exit(0)
	}

	fmt.Println("Hello, World!")
	os.Exit(0)
}