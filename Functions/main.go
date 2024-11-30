package main

import "fmt"

// variadic fucntion are fucntiosn that can acccpet arbitary numnber of parameters

func sumup(one int, numbers ...int) int {

	sum := 0
	for _, value := range numbers {

		sum = sum + value

	}

	return sum

}

func main() {

	fmt.Println(sumup(1, 2, 3, 4, 5))
}
