package utility

import (
	"errors"
	"log"

	"github.com/satori/go.uuid"
)

const (
	ISO8601DATE         = "2006-01-02"
	ISO8601TIME         = "15:04:05"
	ISO8601DATE_TIME    = "2006-01-02 15:04:05"
)

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogDebug(error string) {
	if IsDebug {
		log.Println(error)
	}
}


func GetUUIDFromMapSafely(key string, uuidMap map[string]uuid.UUID) uuid.UUID {
	uuidValue := uuidMap[key]
	if uuidValue == uuid.Nil {
		LogFatal(errors.New("Invalid key: '" + key + "' is being requested from map"))
	}
	return uuidValue
}
