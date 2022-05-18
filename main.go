package main

import (
	"fmt"
	"shrading/shard"
	"time"
)

func main() {
	fmt.Println(time.Now())
	shard.DoShard()
	fmt.Println(time.Now())
}
