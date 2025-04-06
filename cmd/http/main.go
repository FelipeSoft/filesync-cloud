package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FelipeSoft/filesync-cloud/internal/application/handler"
	"github.com/FelipeSoft/filesync-cloud/internal/application/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatalf("Could not load the .env file: %v", err)
	}

	httpServer := http.NewServeMux()

	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")

	httpUrl := fmt.Sprintf("%s:%s", httpHost, httpPort)

	backupService := service.NewBackupService()
	backupHandler := handler.NewBackupHandler(backupService)
	httpServer.HandleFunc("/backup/install", backupHandler.SetInstallationKey)

	go func() {
		if err := http.ListenAndServe(httpUrl, httpServer); err != nil {
			log.Fatalf("Could not start http server on %s caused by error: %v", httpUrl, err)
		}
	}()
}
