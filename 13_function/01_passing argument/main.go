package main

import "fmt"

func main()  {
	greed("Abdi")
	greed("Madina")
	greet("Abdi" , "Sami")
}



func greed(name string)  {
	fmt.Println(name)
}

func greet(name string, lname string){
	fmt.Println(name,lname)
}