package main

import "fmt"

const meterToYards float64 = 1.09361
const kmToMile float64 = 1.6
func main()  {
	var meters float64
	var kilometr float64
	fmt.Print ("Please enter meters swam : ")

	fmt.Scan(&meters)
	fmt.Print ("Please enter kilometr : ")
	fmt.Scan(&kilometr)

	yards := meters * meterToYards
	miles := kilometr * kmToMile

	fmt.Println(meters, "meters is ", yards, "yards")
	fmt.Println(kilometr, "km ", miles, "Mile")

}