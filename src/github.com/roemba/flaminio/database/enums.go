package database

import (
	"github.com/satori/go.uuid"
	"github.com/roemba/flaminio/utility"
)

func GetEnums(){
	PermissionsMap = getPermissionMap()
	LogOperationTypeMap = getLogOperationMap()
}

var PermissionsMap map[string]uuid.UUID
func getPermissionMap() (permissionMap map[string]uuid.UUID) {
	permissionArray, err := GetPermissionArray()
	utility.Fatal(err)

	permissionMap = make(map[string]uuid.UUID)
	for _, e := range permissionArray {
		permissionMap[e.Name] = e.UUID
	}
	return permissionMap
}

var LogOperationTypeMap map[string]uuid.UUID
func getLogOperationMap() (logOperationTypeMap map[string]uuid.UUID) {
	operationsArray, err := GetLogOperationsArray()
	utility.Fatal(err)

	logOperationTypeMap = make(map[string]uuid.UUID)
	for _, e := range operationsArray {
		logOperationTypeMap[e.Name] = e.UUID
	}
	return logOperationTypeMap
}
