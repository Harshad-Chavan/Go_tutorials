package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) outputFloatMap() {
	fmt.Print(m)
}

func main() {

	// we can tell go to create a bigger array behind the scenes
	// 2nd argument is the length of the array we want to
	// 3rd argument is the capacity that controls the memory that will be allocated
	userNames := make([]string, 2, 5)

	userNames[1] = "Jhon"

	userNames = append(userNames, "harshad")
	userNames = append(userNames, "Nivi")
	userNames = append(userNames, "Rahul")

	// this will be done after allocating new space
	userNames = append(userNames, "Henry")

	fmt.Println(userNames)

	// make can also be used to make maps
	//takes only one argument in case of maps as we cannnot et empty slots 5 is the max capacity
	courses := make(floatMap, 5)

	courses["go"] = 4.5
	courses["python"] = 4.7
	courses["C++"] = 4.8

	// fmt.Println(courses)
	courses.outputFloatMap()

	//Looping over arrays, slices, Maps
	// for i := 0; i < len(userNames); i++ {
	// 	fmt.Println(userNames[i])

	// }

	// range always return 2 values
	for index, values := range userNames {
		println(index, values)
	}

	for key, values := range courses {
		println(key, values)
	}

}
