package main

import "fmt"

func main()  {
	a :=43 

	fmt.Println(a) //43
	fmt.Println(&a) //0xc820074188

	var b*int = &a   // b referensing to memory adress 

	fmt.Println(b)  //0xc820074188
	fmt.Println(*b) // 43 //  Dereferessing // Do not give me memory adress give me the value at that  memory adress  

	*b = 42 // b says  The value at this memory adress change it to 42 	
	fmt.Println(a) //42 
}