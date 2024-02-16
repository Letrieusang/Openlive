package utils

import (
	"log"
	"testing"
)

func TestGetExchangeRate(t *testing.T) {
	result, err := GetExchangeRates("tether", 30)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)
	log.Println(len(result))
}

func TestBuildConstainJson(t *testing.T) {
	param := "Size:M,Size:L,Color:back,Color:Red"
	result := BuildFilterJson(param, "properties_json")
	log.Println(result)
}
