package excel

import (
	"fmt"
	"strconv"
)

func ConvertStrToInt(a string) uint {
	if a == "" {
		return 0
	}
	i, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Loi chuyen kieu du lieu")
	}

	return uint(i)
}
