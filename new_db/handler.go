package new_db

import (
	"fmt"
	"gorm.io/sharding"
	"shrading/connection"
	"shrading/model"
	"shrading/old_db"
)

type ShardingEmployee struct {
	ID          int64 `gorm:"primarykey"`
	Name        string
	Email       string
	Address     string
	Discription string
}

func ShardingTable() {
	for i := 0; i < 2; i += 1 {
		table := fmt.Sprintf("employee_%02d", i)
		connection.DB.Exec(`DROP TABLE IF EXISTS ` + table)
		connection.DB.Exec(`CREATE TABLE ` + table + ` (
			ID bigint PRIMARY KEY,
			Name text, 
			Email text,
			Address text, 
			Discription text
		)`)
	}

	middleware := sharding.Register(sharding.Config{
		ShardingKey:    "ID",
		NumberOfShards: 2,
	}, "orders")
	connection.DB.Use(middleware)

	TransferData()
}

func CreateOne(employee *model.Employee, tableName string) error {
	query := connection.DB.Model(&model.Employee{}).Table(tableName).Create(employee)

	if err := query.Error; nil != err {
		return err
	}

	return nil
}

func TransferData() error {

	listEmployee, _, err := old_db.ListAll(10, 0)
	if err != nil {
		return err
	}

	for _, idx := range listEmployee {
		CreateOne(&idx, "employee_00")
	}

	listEmployee, _, err = old_db.ListAll(10, 10)
	if err != nil {
		return err
	}

	for _, idx := range listEmployee {
		CreateOne(&idx, "employee_01")
	}

	return nil
}
