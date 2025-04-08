package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/FelipeSoft/filesync-cloud/internal/application/handler"
	"github.com/FelipeSoft/filesync-cloud/internal/application/middleware"
	"github.com/FelipeSoft/filesync-cloud/internal/application/service"
	jwt_adapter "github.com/FelipeSoft/filesync-cloud/internal/infrastructure/jwt"
	rmysql "github.com/FelipeSoft/filesync-cloud/internal/infrastructure/repository/mysql"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatalf("Could not load the .env file: %v", err)
	}

	mysqlConfig := mysql.Config{
		Addr:   os.Getenv("MYSQL_ADDRESS"),
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
	}
	mysqlConnector, err := mysql.NewConnector(&mysqlConfig)
	if err != nil {
		log.Fatalf("MySQL connector error: %v", err)
	}
	mysqlConnection, err := mysqlConnector.Connect(context.Background())
	if err != nil {
		log.Fatalf("MySQL connection error: %v", err)
	}
	log.Printf("MySQL connection established successfully")

	mysqlFingerprintRepository := rmysql.NewMySQLFingerprintRepository(&mysqlConnection)

	httpServer := http.NewServeMux()

	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")

	httpUrl := fmt.Sprintf("%s:%s", httpHost, httpPort)

	tokenManager := jwt_adapter.NewJwtTokenManager(jwt.SigningMethodRS256)
	fingerprintService := service.NewFingerprintService(tokenManager, mysqlFingerprintRepository)
	fingerprintHandler := handler.NewFingerprintHandler(fingerprintService)

	authMiddleware := middleware.NewAuthMiddleware(tokenManager)

	httpServer.HandleFunc("/fingerprint/install", fingerprintHandler.SetInstallationKey)
	httpServer.HandleFunc("/fingerprint/check", func(w http.ResponseWriter, r *http.Request) {
		authMiddleware.Handle(w, r, fingerprintHandler.TestBearerToken)
	})

	go func() {
		log.Printf("HTTP server listening on http://%s", httpUrl)
		if err := http.ListenAndServe(httpUrl, httpServer); err != nil {
			log.Fatalf("Could not start http server on %s caused by error: %v", httpUrl, err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Print("Exited")
}
