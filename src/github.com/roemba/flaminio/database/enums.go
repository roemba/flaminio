package database

import (
	"github.com/satori/go.uuid"
	"github.com/roemba/flaminio/models"
)

func GetEnums(){
	PermissionsMap = getPermissionMap()
	LogOperationTypeMap = getLogOperationMap()
}

var PermissionsMap map[string]uuid.UUID
func getPermissionMap() (permissionMap map[string]uuid.UUID) {
	var permissionArray []models.Permission
	GetPermissionArray(&permissionArray)

	permissionMap = make(map[string]uuid.UUID)
	for _, e := range permissionArray {
		permissionMap[e.Name] = e.UUID
	}
	return permissionMap
}

var LogOperationTypeMap map[string]uuid.UUID
func getLogOperationMap() (logOperationTypeMap map[string]uuid.UUID) {
	var operationsArray []models.LogOperationType
	GetLogOperationsArray(&operationsArray)

	logOperationTypeMap = make(map[string]uuid.UUID)
	for _, e := range operationsArray {
		logOperationTypeMap[e.Name] = e.UUID
	}
	return logOperationTypeMap
}
