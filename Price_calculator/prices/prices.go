package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// Load prices from a file one price on each line
func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices/prices.txt")

	if err != nil {
		fmt.Println("Error while opening the file")
		fmt.Println(err)
		return
	}

	// to read the content of the file
	scanner := bufio.NewScanner(file)

	//
	var lines []string
	//read line by line scanner.scan return bool value..if value is there it will return true
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())

	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error while reading the file")
		fmt.Println(err)
		return
	}

	prices := make([]float64, len(lines))

	// coverting string to float
	for index, line := range lines {

		floatprice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Error while converting prices")
			fmt.Println(err)
			return
		}
		prices[index] = floatprice

	}
	job.InputPrices = prices

}

func NewTaxIncludedPriceJob(taxrate float64) *TaxIncludedPriceJob {
	// returns a taxincluded price job

	newStruc := TaxIncludedPriceJob{
		TaxRate:     taxrate,
		InputPrices: []float64{},
	}
	//sending a pointer
	return &newStruc

}

// takes pointer to the job and deos not make a copy
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}
