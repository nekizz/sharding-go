package constant

var ErrorCode = map[string]int{
	"ERROR_MISSING_PARAMS":                4000,
	"ERROR_GET_ELASTICSEARCH":             4001,
	"ERROR_DUPLICATE_SUBJECT":             4002,
	"ERROR_SUBJECT_DO_NOT_EXIST":          4003,
	"ERROR_FULL_SLOT":                     4004,
	"ERROR_INVALID_SLOT":                  4005,
	"ERROR_LOCK_RECORD":                   4006,
	"ERROR_INSERT_DATABASE":               4007,
	"ERROR_INSERT_ELASTICSEARCH":          4008,
	"ERROR_UPDATE_TKB":                    4009,
	"ERROR_ROLLBACK_DB_FAIL":              4010,
	"ERROR_ROLLBACK_ES_FAIL":              4011,
	"ERROR_COMMIT_FAIL":                   4012,
	"ERROR_DELETE_DATABASE":               4013,
	"ERROR_DELETE_ELASTICSEARCH":          4014,
	"ERROR_UNREGIST_SUBJECT_DO_NOT_EXIST": 4015,
	"ERROR_FOR_UPDATE":                    4016,
}
