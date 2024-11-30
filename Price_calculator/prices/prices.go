package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxrate float64) *TaxIncludedPriceJob {
	// returns a taxincluded price job

	newStruc := TaxIncludedPriceJob{
		TaxRate:     taxrate,
		InputPrices: []float64{1, 2, 3},
	}
	return &newStruc

}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}
