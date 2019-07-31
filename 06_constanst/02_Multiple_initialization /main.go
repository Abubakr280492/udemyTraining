package main

import (
	"fmt"
	"time"
)


const (
	A = iota //0
	B = iota //1
	c = iota //2
	PI = 3.14
	Language = "GO"
	d = iota //5
    f = iota //6 
)

const (
	q = iota //0   another constant it resets and statrs again from zero 
	w = iota
	e = iota

)


const(
	 _ = iota    // 0
	 r = iota*10   //1*10
	 t = iota*10   //2*10
)

const (
	_ = iota
	KB = 1<<(iota*10)
	MB = 1<<(iota*10)
	GB = 1<<(iota*10)
	TB = 1<<(iota*10)

)
func main()  {
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	
	
	fmt.Println(q)
	fmt.Println(w)
	fmt.Println(e)

	fmt.Println(r)
	fmt.Println(t)

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t",KB)
	fmt.Printf("%d\n",KB)
	fmt.Printf("%b\t",MB)
	fmt.Printf("%d\n",MB)
	fmt.Printf("%b\t",GB)
	fmt.Printf("%d\n",GB)
	fmt.Printf("%b\t",TB)
	fmt.Printf("%d\n",TB)

	fmt.Println("Toshkentda soat", time.Now())


}