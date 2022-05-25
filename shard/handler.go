package shard

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/sharding"
	"github.com/gofiber/fiber/v2"
	"shrading/helper"
	"shrading/model"
	"sync"
)

var Cluster *sharding.Cluster
var Nshards int

// createShard creates database schema for a given shard.
func createShard(shard *pg.DB) error {
	queries := []string{
		`DROP SCHEMA IF EXISTS ?shard CASCADE`,
		`CREATE SCHEMA ?shard`,
		sqlFuncs,
		CreatTableTKB,
		CreateTableRegisterSubject,
	}

	for _, q := range queries {
		_, err := shard.Exec(q)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	db := pg.Connect(&pg.Options{
		Addr:     "13.215.49.1:5432",
		User:     "root",
		Password: "r1UeSNbJ",
		Database: "sharding1",
	})
	db2 := pg.Connect(&pg.Options{
		Addr:     "13.215.49.1:5433",
		User:     "root",
		Password: "ZrEFc5tR",
		Database: "sharding2",
	})

	db3 := pg.Connect(&pg.Options{
		Addr:     "13.215.49.1:5434",
		User:     "root",
		Password: "q2XLM027",
		Database: "sharding3",
	})

	dbs := []*pg.DB{db, db2, db3} // list of physical PostgreSQL servers
	nshards := 3                  // 2 logical shards

	cluster := sharding.NewCluster(dbs, nshards)
	Cluster = cluster
	Nshards = nshards

	fmt.Println("Create cluster successful")
}

func DoShard(c *fiber.Ctx) error {

	for i := 0; i < Nshards; i++ {
		if err := createShard(Cluster.Shard(int64(i))); err != nil {
			return c.JSON(helper.Response{
				Status:  false,
				Message: "Create shard fail",
				Data:    nil,
				Error:   helper.Error{},
			})
		}
	}

	err := transferDataToShard(Cluster)
	if err != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Message: "Fail to transfer data",
			Data:    nil,
			Error:   helper.Error{},
		})
	}

	fmt.Println("shrad database successful")
	return nil
}

func transferDataToShard(cluster *sharding.Cluster) error {
	var wg sync.WaitGroup

	listTKB, count, err := model.ListAll(2000, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

	for i, idx := range listTKB {
		fmt.Println(i)
		wg.Add(1)
		go func(idx model.TKB, wg *sync.WaitGroup) {
			defer wg.Done()
			err := CreateTKB(cluster, changeDataType(idx))
			if err != nil {
				panic(err)
			}
		}(idx, &wg)
		wg.Wait()
	}

	fmt.Println("shard thanh cong")
	return nil
}
