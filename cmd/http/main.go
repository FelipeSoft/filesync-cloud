package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/FelipeSoft/filesync-cloud/internal/application/handler"
	"github.com/FelipeSoft/filesync-cloud/internal/application/service"
	jwt_adapter "github.com/FelipeSoft/filesync-cloud/internal/infrastructure/jwt"
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

	httpServer := http.NewServeMux()

	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")

	httpUrl := fmt.Sprintf("%s:%s", httpHost, httpPort)

	tokenManager := jwt_adapter.NewJwtTokenManager(jwt.SigningMethodRS256)
	backupService := service.NewBackupService(tokenManager)
	backupHandler := handler.NewBackupHandler(backupService)
	httpServer.HandleFunc("/backup/install", backupHandler.SetInstallationKey)

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
