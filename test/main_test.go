package test

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func TestAppLoads(t *testing.T) {

}
