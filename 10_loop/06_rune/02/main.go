package main

import "fmt"

func main()  {
	for i:=50; i<=150; i++{
		fmt.Println(i, "-" , string(i), "-", []byte(string(i) ))
		fmt.Printf("%v - %v -%v \n ", i, string(i), []byte (string(i)))
	}
	fmt.Printf("%v", string(65))
	fmt.Printf("%v", string(98))
	fmt.Printf("%v", string(100))
	fmt.Printf("%v \n", string(105))
	

}