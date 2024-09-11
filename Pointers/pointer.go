package main

import "fmt"

func main() {

	age := 32 // regular variable

	//var agePointer *int
	agePointer := &age

	fmt.Println("Age : ", age)
	fmt.Println("Age address : ", agePointer)
	//fmt.Println("Age value : ", *agePointer)
	// adultYears := GetAdultYears(agePointer)
	// fmt.Println("Adult years : ", adultYears)

	GetAdultYears(agePointer)
	fmt.Println(" age after fucntion : ", age)

}

// func GetAdultYears(age int64) int64 {
// 	// The address will be different from the on outside as we are creating a copy
// 	fmt.Println("Age address in func : ", &age)

// 	adultYears := age - 18

// 	return adultYears

// }

// here i am passing a pointer
// func GetAdultYears(age *int) int {
// 	// The address will be different from the on outside as we are creating a copy
// 	fmt.Println("Age address in func : ", &age)

// 	adultYears := *age - 18

// 	return adultYears

// }

// mutating the value in place
func GetAdultYears(age *int) {
	// The address will be different from the on outside as we are creating a copy
	fmt.Println("Age address in func : ", age)

	*age = *age - 18

}
