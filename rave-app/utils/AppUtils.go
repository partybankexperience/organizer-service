package utils

import (
	"encoding/base64"
	"github.com/google/uuid"
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

func isDateValid(date string) bool {
	return false
}

// TODO: implement me
func GenerateEventReference() string {
	s := uuid.New()
	v := base64.RawURLEncoding.EncodeToString([]byte(s.String()))
	return "evt-" + v
}
