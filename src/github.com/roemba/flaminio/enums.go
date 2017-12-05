package flaminio

import (
	"github.com/satori/go.uuid"
	"errors"
)

func getEnums(){
	permissionsMap = Permission{}.getMap()
	logOperationTypeMap = LogOperationType{}.getMap()
}

func getUUIDFromMapSafely(key string, uuidMap map[string]uuid.UUID) uuid.UUID {
	uuidValue := uuidMap[key]
	if uuidValue == uuid.Nil {
		fatal(errors.New("Invalid key: '" + key + "' is being requested from map"))
	}
	return uuidValue
}

var permissionsMap map[string]uuid.UUID
func (Permission) getMap() (permissionMap map[string]uuid.UUID) {
	var permissionArray []Permission
	db.Find(&permissionArray)

	permissionMap = make(map[string]uuid.UUID)
	for _, e := range permissionArray {
		permissionMap[e.Name] = e.UUID
	}
	return permissionMap
}

var logOperationTypeMap map[string]uuid.UUID
func (LogOperationType) getMap() (logOperationTypeMap map[string]uuid.UUID) {
	var operationsArray []LogOperationType
	db.Find(&operationsArray)

	logOperationTypeMap = make(map[string]uuid.UUID)
	for _, e := range operationsArray {
		logOperationTypeMap[e.Name] = e.UUID
	}
	return logOperationTypeMap
}
