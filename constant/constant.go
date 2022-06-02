package constant

var ErrorCode = map[string]int{
	"ERROR_MISSING_PARAMS": 4000,
}

const Secretkey = "aisoft27052022"

var Sharding1 = "host=13.215.49.1 port=5432 user=root password=r1UeSNbJ dbname=sharding1 sslmode=disable"

var Sharding2 = "host=13.215.49.1 port=5433 user=root password=ZrEFc5tR dbname=sharding2 sslmode=disable"

var Sharding3 = "host=13.215.49.1 port=5434 user=root password=q2XLM027 dbname=sharding3 sslmode=disable"

//var ELASTIC_SEARCH_LIVECHAT_CONFIG = []string{"localhost:9200"}

var ELASTIC_SEARCH_LIVECHAT_CONFIG = []string{"http://13.215.49.1:9200/"}
