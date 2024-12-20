package lists

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

	fmt.Println(len(changeprices), cap(changeprices))

	// this is an dynamic array
	dynamic_price := []float64{100, 200}

	// cant be done as the slice has only 2 elemetns
	//dynamic_price[3] = 300

	// append fucntion go cretes a new array and does not modify the slice
	dynamic_price_2 := append(dynamic_price, 300)

	fmt.Print(dynamic_price)
	fmt.Println(dynamic_price_2)

	//practice
	type Product struct {
		id         int64
		name       string
		decription string
	}

	//creating and array
	productOne := Product{id: 101, name: "Milk", decription: "cow milk"}
	productTwo := Product{id: 102, name: "Bread", decription: "Brown Bread"}
	products := []Product{productOne}
	products = append(products, productTwo)

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}

	//we cant directly append one list to another to do that we have to use a specail syntax

	numbers := []int64{1, 2, 3}
	newNumbers := []int64{4, 5, 6}

	allNumbers := append(numbers, newNumbers...)

	fmt.Println(allNumbers)
}
