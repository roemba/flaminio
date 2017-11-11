package flaminio

import (
	"net/http"

	"log"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	privKeyPath = "app.rsa"
	pubKeyPath = "app.rsa.pub"
)

var db *gorm.DB

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func startServer(){
	router := gin.Default()
	setRoutes(router)

	router.Run(":8080")
}

func connectToDatabase(){
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=flaminio dbname=flaminio sslmode=disable password=ZzS08RNyosHD2xg49k9Z")
	fatal(err)

	db.AutoMigrate(&User{})

	if db.Find(&User{}, User{Username:"admin"}).RecordNotFound() {
		log.Println("Default admin user not found. Creating one...")
		hashedPassword, err := hashPassword("admin")
		fatal(err)
		db.Create(User{FirstName:"admin", LastName:"admin", Username:"admin", Password:hashedPassword})
	}
}

func Main() {
	initKeys()
	connectToDatabase()
	startServer()
}

func JsonResponse(response interface{}, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
