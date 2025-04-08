package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FelipeSoft/filesync-cloud/internal/application/dto"
	"github.com/FelipeSoft/filesync-cloud/internal/application/service"
	httputil "github.com/FelipeSoft/filesync-cloud/internal/utils/http"
	"github.com/go-playground/validator/v10"
)

type FingerprintHandler struct {
	fingerprintService *service.FingerprintService
}

func NewFingerprintHandler(fingerprintService *service.FingerprintService) *FingerprintHandler {
	return &FingerprintHandler{
		fingerprintService: fingerprintService,
	}
}

func (h *FingerprintHandler) SetInstallationKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httputil.WriteJSON(w, http.StatusMethodNotAllowed, httputil.HttpResponse{
			Error: fmt.Sprintf("Could not %s to %s", r.Method, r.Pattern),
		})
	}

	defer r.Body.Close()

	var fingerprintRequest dto.FingerprintRequest
	err := json.NewDecoder(r.Body).Decode(&fingerprintRequest)
	if err != nil {
		httputil.WriteJSON(w, http.StatusBadRequest, httputil.HttpResponse{
			Error: "Invalid JSON error: " + err.Error(),
		})
		return
	}

	validator := validator.New()
	err = validator.Struct(fingerprintRequest)
	if err != nil {
		httputil.WriteJSON(w, http.StatusBadRequest, httputil.HttpResponse{
			Error: "Validation error: " + err.Error(),
		})
		return
	}

	accessToken, err := h.fingerprintService.VerifyFingerprint(fingerprintRequest)
	httputil.WriteJSON(w, http.StatusOK, httputil.HttpResponse{
		Message: "Installation successfully completed!",
		Data: map[string]any{
			"access_token": accessToken,
		},
	})
}

func (h *FingerprintHandler) TestBearerToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httputil.WriteJSON(w, http.StatusMethodNotAllowed, httputil.HttpResponse{
			Error: fmt.Sprintf("Could not %s to %s", r.Method, r.Pattern),
		})
	}

	defer r.Body.Close()

	bearerToken := r.Header.Get("Authorization")
	accessToken, err := h.fingerprintService.Refresh(bearerToken)
	if err != nil {
		httputil.WriteJSON(w, http.StatusUnauthorized, httputil.HttpResponse{
			Error: err.Error(),
		})
		return
	}

	httputil.WriteJSON(w, http.StatusOK, httputil.HttpResponse{
		Message: "Test successfully completed!",
		Data: map[string]any{
			"bearer_token": accessToken,
		},
	})
}
