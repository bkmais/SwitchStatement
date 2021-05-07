package main

import (
	"fmt"
	"time"
)

func main() {

	switch t := time.Now().Hour(); {

	case t > 0 && t < 12:
		{
			fmt.Println("Good Morning")
		}
	case t > 12 && t < 18:
		{
			fmt.Println("Goor Afternoon")
		}
	case t > 18:
		{
			fmt.Println("Good evening")
		}
	default:
		{
			fmt.Println("Good day!!")
		}
	}
}
