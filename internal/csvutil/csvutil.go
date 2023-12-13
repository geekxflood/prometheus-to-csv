// internal/csvutil/csvutil.go
package csvutil

import (
	"github.com/prometheus/common/model"
)

type CSVData struct {
	Headers []string
	Rows    [][]string
}

func CreateCSVDataFromResult(result model.Value) (CSVData, error) {
	// Logic to create and return CSVData from result
	// ...

	return CSVData{}, nil
}

func PrintCSVData(csvData CSVData) {
	// Logic to print CSVData
	// ...

}
