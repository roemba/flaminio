package utility

import (
	"log"
	"github.com/satori/go.uuid"
	"errors"
)

const (
	ISO8601DATE      = "2006-01-02"
	ISO8601TIME = "15:04:05"
	ISO8601DATE_TIME = "2006-01-02 15:04:05"
)

func Fatal(err error) {
	log.SetFlags(log.Llongfile)
	if err != nil {
		//TODO set log flags for all log functions
		log.Fatal(err)
	}
}

func GetUUIDFromMapSafely(key string, uuidMap map[string]uuid.UUID) uuid.UUID {
	uuidValue := uuidMap[key]
	if uuidValue == uuid.Nil {
		Fatal(errors.New("Invalid key: '" + key + "' is being requested from map"))
	}
	return uuidValue
}
