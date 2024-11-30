package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	// coverting string to float
	result := make([]float64, len(strings))
	for index, stringVal := range strings {

		floatprice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}
		result[index] = floatprice

	}
	return result, nil

}
