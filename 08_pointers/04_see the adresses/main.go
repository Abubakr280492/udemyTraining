package main

import "fmt"

func zero(z int )  {    // z variabli in his own location    passing value 5  not x  and assigning to new variable z 
	fmt.Printf("%p\n",&z)  // adress in func zero 
	fmt.Println(&z)   // adress in func zero
	z=0
}

func main(){

	x :=5
	fmt.Printf("%p\n",&x)  // adress in func main    first   // %p base 16 notation with leading 0x
	fmt.Println(&x)   // adress in func main         second 

	zero(x)  // passing value 5                3rd 4 th 
	fmt.Println(x) // x is still 5 

}