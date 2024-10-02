package main

import "fmt"

func main() {

	//  declare an array
	var ProductNames [4]string

	ProductNames = [4]string{"book"}
	ProductNames[3] = "carpet"

	prices := [4]float64{1.1, 2.2, 3.3, 4.4}

	fmt.Println(prices)
	fmt.Println(ProductNames)

	// to fetch data from a index
	fmt.Println(prices[3])

}
