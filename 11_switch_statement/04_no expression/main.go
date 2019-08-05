package main

import "fmt"
func main()  {

	myFriendsName :="Ab";
	switch {
	case len(myFriendsName)==2:
		fmt.Println("Wassup my friend with name of lengs 2") 
	case myFriendsName=="Ab":
		fmt.Println("wassup mr Abdi") //Even though we have two matches, First one gonna be taken which is name with a lengs 2  
	case myFriendsName == "tim":
		fmt.Println("wassup tim")
	case myFriendsName == "mehd", myFriendsName == "madina":
		fmt.Println("Your name is either  Mehdi or Madina ")
	default :
		fmt.Println("nothing matched this is default ")

	}
}