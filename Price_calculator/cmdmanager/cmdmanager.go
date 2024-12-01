package cmdmanager

import "fmt"

type CmdManager struct {
}

func New() CmdManager {
	return CmdManager{}
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Enter prices confirm each price with enter")

	var prices []string
	for {
		var price string
		fmt.Println("Price:")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil

}

func (cmd CmdManager) WriteJson(data interface{}) error {
	fmt.Println(data)
	return nil
}
