package flaminio

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"flaminio/database"
	"flaminio/utility"
)



func startServer(){
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	setRoutes(router)

	router.Run(":8080")
}

func Main() {
	utility.InitKeys()
	database.ConnectToDatabase()
	database.GetEnums()
	startServer()
}
