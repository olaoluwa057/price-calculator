package prices

import (
	"fmt"
)
import ("example.com/price-calculator/conversion"
    "example.com/price-calculator/filemanager"
)
type TaxIncludedPriceJob struct {
	IOM 			filemanager.FileManager `json:"-"`
	TaxRate          float64 `json:"tax_rate"`
	InputPrice       []float64 `json:"input_price"`
	TaxIncludedPrice map[string]string `json:"tax_included_price"`
}

func NewTaxIncludedPriceJob(inputPath, outputPath string, taxRate float64) *TaxIncludedPriceJob {
	fm := filemanager.New(inputPath, outputPath)
	return &TaxIncludedPriceJob{
		IOM: fm,
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}


//// Methods //////////////////////////////////////////////////////////////////

func (job *TaxIncludedPriceJob) Process(done chan bool, errorChan chan error) {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrice {

		taxAmount := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxAmount)

	}
	job.TaxIncludedPrice = result 
	err := filemanager.FileManager.WriteJsonFile(job.IOM, job)

	if err != nil {
		errorChan <- err
	}

	done <- true

}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.FileManager.ReadFile(job.IOM)	

	
    if err!= nil {
        fmt.Println("Error reading file: ", err)
        return
    }

	floatValue, err := conversion.StringToFloat(lines)

		if err!= nil {
            fmt.Println("Error parsing number: ", err)
            return
        }
		
	
	job.InputPrice = floatValue

	
}
