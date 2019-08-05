package main

import "fmt"

func main()  {
	switch "abdi" {
	case "daniel":
		fmt.Println("Wassup Daniel")//no need to break 
	case "abdi":
		fmt.Println("Wassup Abubakr")
	case "jenny" :
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends")
		
	}
}