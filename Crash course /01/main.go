package main

// import (
// 	"fmt"
// 	"time"
// 	"sync"

// )

// func main ()  {
// 	//go count("sheep")
// 	//count ("fish")
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func (){
// 		count("sheep")
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }

// func count(thing string )  {
	
// //	for i:=1; true; i++{
// 	for i:=1; i<=5; i++{
// 		fmt.Println(i, thing)
// 		time.Sleep(time.Millisecond *500)
// 	}
// }


import (
	"fmt"
	"time"
)
func	main(){

	c:=make(chan string)
	go count ("sheep", c)
	msg := <- c
	fmt.Println(msg)

}

func count (thing string, c chan string){
	for i:=1; i<=5; i++ {
		c <- thing
		time.Sleep(time.Millisecond*500)
	}
}