package utils

import (
	"log"
	"strconv"
)

const (
	AUTHORIZATION = "Authorization"
)

func ConvertQueryStringToInt(query string) (int, error) {
	value, err := strconv.Atoi(query)
	if err != nil {
		log.Println("Error: ", err)
		return 0, err
	}
	return value, nil
}
