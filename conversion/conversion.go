package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToFloat(_string []string) ([]float64, error) {

	//values := make([]float64, len(_string))

	var values []float64

	for _, value := range _string {

		floatValue, err := strconv.ParseFloat(value, 64)

		if err != nil {

			 fmt.Printf("Error parsing string '%s' to float: %v", value, err)

			return nil, errors.New("could not convert string to float")
		}

		values = append(values, floatValue)
	}
	return values, nil
}
