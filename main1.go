package main

import (
	"fmt"
	"os"
)

func main() {

	city := os.Args[1]

	switch city {
	case "Tokyo":
		fmt.Println("Japan")
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Imphal":
		fmt.Println("India")
	default:
		fmt.Println("Unknown Place.")
	}

	fmt.Println("Shortswitch:")

}
