package main

import (
	"fmt"
	"shrading/connection"
)

func main() {
	//new_db.ShardingTable()
	conn, err := connection.Excelconn()
	if err != nil {
		fmt.Println(err)
	}

	a, _ := conn.GetCellValue("DKGD", "B328")
	fmt.Println(a)
}
