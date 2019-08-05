package main

import "fmt"

func main ()  {
	switch "marcus" {
	case "tim": 
		fmt.Println("hi tim")
	case "Janny":
		fmt.Println("hi Janny")
	case "marcus":
		fmt.Println("hi marcus")
		fallthrough
	case "mehdi":
		fmt.Println("hi mehdi ")
		fallthrough
	case "julian":
		fmt.Println("hi julian")
	case "suchant": 
		fmt.Println("hi Suchant")
	}
}