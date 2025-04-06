package handler

import (
	"fmt"
	"net/http"

	"github.com/FelipeSoft/filesync-cloud/internal/application/service"
	httputil "github.com/FelipeSoft/filesync-cloud/internal/utils/http"
)

type BackupHandler struct {
	backupService *service.BackupService
}

func NewBackupHandler(backupService *service.BackupService) *BackupHandler {
	return &BackupHandler{
		backupService: backupService,
	}
}

func (h *BackupHandler) SetInstallationKey(w http.ResponseWriter, r *http.Request) {	
	if r.Method != "POST" {
		httputil.WriteJSON(w, http.StatusMethodNotAllowed, httputil.HttpResponse{
			Error: fmt.Sprintf("Could not %s to %s", r.Method, r.Pattern),
		})
	}
	accessToken := ""
	httputil.WriteJSON(w, http.StatusOK, httputil.HttpResponse{
		Message: "Installation successfully completed!",
		Data: map[string]any{
			"access_token": accessToken,
		},
	})
}