// very go file must have a package at the start
// code must be organised in packages
// main specifies that this package will be the entry point of the application we are building
package main

// fmt is a built in package from which we use the print fucntion
import "fmt"

// must be named main. go will call and execute this funciton when program starts
func main() {
	fmt.Print("Hello World")

}

// one module consists of multiple packages
// -go build- command enabales us to create a executable

// intilly go deos not recognise the folder as a module for that we have to run command
// go mod init <path or name>
// post module creation we can use go build to create executable file
// ./first-app

// libraries are not executbales so they dont need to have a main fucntion
