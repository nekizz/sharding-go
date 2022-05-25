package shard

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/sharding"
	"shrading/model"
	"shrading/old_db"
	"sync"
)

const sqlFuncs = `
CREATE SEQUENCE ?shard.id_seq;

-- _next_id returns unique sortable id.
CREATE FUNCTION ?shard._next_id(tm timestamptz, shard_id int, seq_id bigint)
RETURNS bigint AS $$
DECLARE
  max_shard_id CONSTANT bigint := 2048;
  max_seq_id CONSTANT bigint := 4096;
  id bigint;
BEGIN
  shard_id := shard_id % max_shard_id;
  seq_id := seq_id % max_seq_id;
  id := (floor(extract(epoch FROM tm) * 1000)::bigint - ?epoch) << 23;
  id := id | (shard_id << 12);
  id := id | seq_id;
  RETURN id;
END;
$$
LANGUAGE plpgsql
IMMUTABLE;

CREATE FUNCTION ?shard.next_id()
RETURNS bigint AS $$
BEGIN
   RETURN ?shard._next_id(clock_timestamp(), ?shard_id, nextval('?shard.id_seq'));
END;
$$
LANGUAGE plpgsql;
`

var CLUSTER *sharding.Cluster

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

func createCluster() (*sharding.Cluster, int) {
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

	return cluster, nshards
}

func DoShard() {
	cluster, nshards := createCluster()

	for i := 0; i < nshards; i++ {
		if err := createShard(cluster.Shard(int64(i))); err != nil {
			panic(err)
		}
	}

	err := transferDataToShard(cluster)
	if err != nil {
		panic(err)
	}

	fmt.Println("shrad database successful")
}

func transferDataToShard(cluster *sharding.Cluster) error {
	var wg sync.WaitGroup

	listTKB, count, err := old_db.ListAll(2000, 0)
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

	err1 := CreateRS(cluster, &RegisterSubject{
		ID:       2,
		MaSV:     "B18DCCN405",
		MaMonHoc: "INT1446",
	})
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("shard thanh cong")
	return nil
}
