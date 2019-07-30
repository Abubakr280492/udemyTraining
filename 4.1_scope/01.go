package main

import "fmt"	

var x = 5  // Every time it gonna be updated and will be incremented 

func increment () int {
	//var x = 2
	x++
	return x 
}

func main() {

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
 	fmt.Println(increment())
	fmt.Println(increment())
	
}