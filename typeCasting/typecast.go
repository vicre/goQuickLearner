package main

import "fmt"

func main() {
	//You can convert values from one type to another using the cast operator, syntax is as follows −
	//type_name(expression)
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("Value of mean : %f\n", mean)
}
