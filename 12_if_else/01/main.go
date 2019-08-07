package main

import "fmt"
func main() {
	if !true {
		fmt.Println("This not gonna run ")
	}
	if !false {
		fmt.Println("This  gonna run ")
	}

	if true {
		fmt.Println("This gonna run ")
	}
	if false {
		fmt.Println("This not gonna run ")
	}
}