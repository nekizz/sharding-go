package migration

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"shrading/connection"
	"shrading/excel"
	"shrading/helper"
	"time"
)

func MigrateDB() error {
	orm, err := gorm.Open(postgres.New(postgres.Config{DSN: connection.DSN}), &gorm.Config{})
	postgresDB, err := orm.DB()
	if err != nil {
		panic(err)
	}

	postgresDB.SetConnMaxLifetime(300 * time.Minute)
	postgresDB.SetMaxIdleConns(10)
	postgresDB.SetMaxOpenConns(15)

	defer func() {
		if err := postgresDB.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Println("mysql connection established")

	err = Migrate(orm)
	if err != nil {
		return err
		os.Exit(1)
	}

	fmt.Println("migration successful")

	return nil
}

func MigrateAndSync(c *fiber.Ctx) error {
	errMS := MigrateDB()
	if errMS != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Migrate fail.",
			Error:   helper.Error{},
		})
	}

	errEx := excel.ReadDataFromExcel()
	if errEx != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Transfer data to database fail.",
			Error:   helper.Error{},
		})
	}

	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Transfer data to database successful.",
		Error:   helper.Error{},
	})
}
