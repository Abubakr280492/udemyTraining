package main

import "fmt"

func main()  {
	n := average(43,56,77,65,11,45)
	fmt.Println(n)

}

func average(sf ...float64) float64{
	total := 0.0
	for _, v :=range sf{
		total +=v
	}
	return total/float64(len(sf))
}