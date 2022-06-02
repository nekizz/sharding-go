package connection

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"math"
	"shrading/constant"
	"time"
)

var ESLivechatConnection *elasticsearch.Client

func init() {
	ESLivechatConnection = ConnectLivechatElastic()
}

func ConnectElastic() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			constant.ELASTIC_SEARCH_LIVECHAT_CONFIG[0],
		},
		RetryOnStatus: []int{429, 502, 503, 504},
		RetryBackoff: func(i int) time.Duration {
			// A simple exponential delay
			d := time.Duration(math.Exp2(float64(i))) * time.Second
			fmt.Printf("Attempt: %d | Sleeping for %s...\n", i, d)
			return d
		},
	}
	elastic, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("CONNECT ELASTIC ERROR! ", err.Error())
	}
	if elastic == nil {
		log.Println("CONNECT ELASTIC ERROR! ")
	}

	return elastic, err
}

func ConnectLivechatElastic() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses:     constant.ELASTIC_SEARCH_LIVECHAT_CONFIG,
		RetryOnStatus: []int{429, 502, 503, 504},
		RetryBackoff: func(i int) time.Duration {
			// A simple exponential delay
			d := time.Duration(math.Exp2(float64(i))) * time.Second
			fmt.Printf("Attempt: %d | Sleeping for %s...\n", i, d)
			return d
		},
	}
	elastic, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("ConnectLivechatElastic ERROR! ", err.Error())
	}
	if elastic == nil {
		log.Println("ConnectLivechatElastic ERROR! CONNECT IS EMPTY!")
	}

	fmt.Println("Connect elasticsearch successful")

	return elastic
}
