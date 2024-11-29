package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	//anaonymous funtions are fucntions that are defined just in time instead of being defined n advance
	//transformed := transformNumbers(&numbers, func(number int) int { return number * 2 })

	//creating diuble fucntion using factory pattern
	double := createTransformerFunc(2)
	triple := createTransformerFunc(3)

	dtransformed := transformNumbers(&numbers, double)
	fmt.Println(dtransformed)

	ttransformed := transformNumbers(&numbers, triple)
	fmt.Println(ttransformed)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// closure example
// Go functions may be closures. A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
func createTransformerFunc(factor int) func(int) int {
	return func(value int) int { return value * factor }
}
