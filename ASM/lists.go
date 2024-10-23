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

	//slicing array (first included 2nd excluded)
	fmt.Println(prices[1:3])

	// slicing array
	fmt.Println(prices[:3])

	fmt.Println(prices[:3][1:2])

	// cant use negative slicng index..cant pick higher indexx than orignal aray size

	//modifying usinf slices
	changeprices := prices[2:4] //2.2, 3.3
	changeprices[0] = 10.5      // 10.5, 3.3

	fmt.Println(prices)

	fmt.Print(len(changeprices), cap(changeprices))
}
