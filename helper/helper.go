package helper

import (
	"crypto/sha1"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"hash/fnv"
	"strconv"
)

func ValidateStruct(user interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err3 := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err3.StructNamespace()
			element.Tag = err3.Tag()
			element.Value = err3.Value()
			errors = append(errors, &element)
		}
	}
	return errors
}

func StringToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return val
}

func HashToInt(hashString string) uint32 {
	//sha hash
	sha := sha1.New()
	sha.Write([]byte(hashString))
	bs := sha.Sum(nil)
	shaString := fmt.Sprintf("%x", bs)

	h := fnv.New32a()
	h.Write([]byte(shaString))
	return h.Sum32()
}
