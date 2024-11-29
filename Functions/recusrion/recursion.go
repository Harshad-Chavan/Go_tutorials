package recursion

import "fmt"

func main() {

	fmt.Println(factorial(5))

}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	number = number * factorial(number-1)
	return number
}
