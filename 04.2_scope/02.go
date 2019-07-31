package main

import "fmt"

func main()  {
	x:=4



	increment := func() int {  //func expression //anonamus func     // taking function assigning to increment name variable  
		x++
		return x

	}

	fmt.Println(increment())
	fmt.Println(increment())
}