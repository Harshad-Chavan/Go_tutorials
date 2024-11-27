package maps

import "fmt"

func main() {

	//maps are a structure that hold key:value pairs
	// in go we have to define the datatypes of the key and value beforehand

	websites := map[string]string{

		"Google":               "https://www.google.com",
		"Amazone Web Services": "https://www.aws.com",
	}

	// to access a key value pair
	fmt.Println(websites["Google"])

	// add key value pairs
	websites["LinkedIn"] = "www.linkedin.com"

	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	//why use maps ?
	// 1) for maps we can use anaything as key (string,int,array,struct)
	// 2) Maps solve a different problem with stuct we have predefined data structure
	//    we can't just add a key value pair or cant delete this we can do with maps
	//    struct is greate when we have cleary defined data

}
