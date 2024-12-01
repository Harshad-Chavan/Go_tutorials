package prices

import (
	"example/price_calculator/conversion"
	"example/price_calculator/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]float64 `json:"tax_inlcuded_prices"`
	// the json tag says that it will be ignored in output json
	IoManager iomanager.IoManager `json:"-"`
}

// contructor for the price job obj
func NewTaxIncludedPriceJob(fm iomanager.IoManager, taxrate float64) *TaxIncludedPriceJob {
	// returns a taxincluded price job

	newStruc := TaxIncludedPriceJob{
		TaxRate:     taxrate,
		InputPrices: []float64{},
		IoManager:   fm,
	}
	//sending a pointer
	return &newStruc

}

// Load prices from a file one price on each line
func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IoManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

// takes pointer to the job and deos not make a copy
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = result
	job.IoManager.WriteJson(job)

}
