package main

import (
	"context"
	"errors"
	"flag"
	"flaminio/database"
	"flaminio/utility"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// global flags
var isDebug bool
var port int
var publicKeyPath string
var privateKeyPath string
var dbUrl string

func startServer(){
	if isDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	setRoutes(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	go func() {
		log.Println("Server started!")
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Exiting")
}

func main() {
	// Setup command line flags
	flag.IntVar(&port, "port", 8080, "Specify the port to listen to.")
	flag.BoolVar(&isDebug, "isdebug", false, "Set to true to run the app in debug mode.  In debug, it may panic on some errors.")
	flag.StringVar(&publicKeyPath, "publickey", "", "Required. Enables listening on 443 with -publickey and -privatekey files specified.  This must be a full path to the certificate .pem file. See http://golang.org/pkg/net/http/#ListenAndServeTLS for more information.")
	flag.StringVar(&privateKeyPath, "privatekey", "", "Required. Enables listening on 443 with -publickey and -privatekey files specified.  This must be a full path to the key .pem file. See http://golang.org/pkg/net/http/#ListenAndServeTLS for more information.")
	flag.StringVar(&dbUrl, "dburl", "postgres://flaminio:ZzS08RNyosHD2xg49k9Z@localhost/flaminio?sslmode=disable", "Specifies the Postgres DB url.")
	flag.Parse()

	if isDebug != false {
		log.SetFlags(log.Lshortfile)
		log.Println("DEBUG mode enabled")
	}
	if publicKeyPath == "" || privateKeyPath == "" {
		utility.LogFatal(errors.New("Missing required public and/or private key paths!"))
	}
	if dbUrl != "" {
		log.Printf("Using default DB URL: %s", dbUrl)
	}

	utility.Init(privateKeyPath, publicKeyPath, isDebug)
	database.ConnectToDatabase(dbUrl)
	database.GetEnums()
	startServer()
}
