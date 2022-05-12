package migration

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"shrading/connection"
	"time"
)

func MigrateDB() {
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
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("migration successful")
}
