package connection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "123456"
	dbname   = "sharding"
)

var DB *gorm.DB
var DSN string

func init() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		panic(err)
		fmt.Println("CONNECT POSTGRESSQL FAILED!")
		fmt.Println(err.Error())
	}

	DB = db
	DSN = dsn

	fmt.Println(fmt.Sprintf("POSTGRESSQL connection established"))
}

func PostgresConn(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		panic(err)
		fmt.Println("CONNECT POSTGRESSQL FAILED!")
		fmt.Println(err.Error())
	}

	return db
}