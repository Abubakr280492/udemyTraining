 package main
 
 import "fmt"

 func zero(z *int)  {   //recieve at memory adress 
	fmt.Println(z)
	*z=0   // new variable dereferensing it 
 }

 func main (){
	 x:=5
	fmt.Println(&x)
	 zero(&x)  // pass over by value  to memory adress 
	 fmt.Println(x) // x is 0
	
	}

 //all by pass by value 																	