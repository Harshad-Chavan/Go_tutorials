// Every go file must be part of package
package main

import "fmt"

func presentOptions() {
	fmt.Println("What do you want to do with you account ?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

}
