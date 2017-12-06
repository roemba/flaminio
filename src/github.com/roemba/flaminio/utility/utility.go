package utility

import (
	"log"
	"github.com/satori/go.uuid"
	"errors"
)

func Fatal(err error) {
	if err != nil {
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
