package main

import "fmt"

func main()  {
	x := 13%3

	fmt.Println(x)
	if x==1 {      // dont do x = 1  
		fmt.Println("Odd")
	}else {
		fmt.Println("Even")
	}

	for i:=0; i<100; i++ {  //don't forget simicolon 
		if i%2==1{
			fmt.Println("Odd")
		}else{
			fmt.Println("Even baby ")
		}
	}
}