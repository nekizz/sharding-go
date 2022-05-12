package main

import (
	"fmt"
	"shrading/excel"
	"shrading/migration"
)

func main() {
	migration.MigrateDB()
	err := excel.ReadDataFromExcel()
	if err != nil {
		fmt.Println(err)
	}
}
