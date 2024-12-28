package main

import (
	"fmt"
	"example.com/price-calculator/prices"
)
func main() {

	taxRate := []float64 {0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRate))
	errorChans := make([]chan error, len(taxRate))

 for index, value := range taxRate {
	doneChans[index] = make(chan bool)
	errorChans[index] = make(chan error)
    priceJob := prices.NewTaxIncludedPriceJob("prices/prices.txt", fmt.Sprintf("result_%.0f.txt", value*100), value)
	go priceJob.Process(doneChans[index], errorChans[index] )
 }



for index := range taxRate {
	select {
	case err := <-errorChans[index]:
		if err != nil {
			fmt.Println(err)
		}
	case <-doneChans[index]:
		// Process completed successfully
	}
}

}
