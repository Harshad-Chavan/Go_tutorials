package main

import "fmt"

type transformFn func(int) int

func main() {
	fmt.Println("Fucntions deeep dive")

	//Fucntions are first class values
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Print(dNumbers(&numbers))

	fmt.Println(transformNumbers(&numbers, double))
	fmt.Println(transformNumbers(&numbers, triple))

}

// funciton that uses for loop to double the number
func dNumbers(pnumbers *[]int) []int {

	dnumbers := []int{}
	for _, value := range *pnumbers {
		dnumbers = append(dnumbers, double(value))
	}
	return dnumbers

}

// making the fucntion more generic by using passing fucntion as value
func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}

	for _, value := range *numbers {

		tNumbers = append(tNumbers, transform(value))

	}

	return tNumbers

}

// we can have a double and a triple fucntion
func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
