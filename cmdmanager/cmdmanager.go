package cmdmanager

import "fmt"

type CmdManager struct{}

func (cmd CmdManager) ReadFile() ([]string, error) {
	var prices []string
	fmt.Println("Enter your prices")
	for {
		var price string
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CmdManager ) WriteJsonFile( data any) error {
	fmt.Println("Written to file")

	return nil
}
