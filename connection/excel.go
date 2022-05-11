package connection

import (
	"github.com/xuri/excelize/v2"
)

func Excelconn() (*excelize.File, error) {
	f, err := excelize.OpenFile("./connection/TKB.xlsx")
	if err != nil {
		return &excelize.File{}, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()
	return f, nil
}
