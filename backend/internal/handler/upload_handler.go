package handler

import (
	"net/http"

	"github.com/arm02/bank-statement-viewer/backend/internal/service"
)

type UploadHandlers struct {
	Svc *service.TransactionService
}

func NewUploadHandlers(svc *service.TransactionService) *UploadHandlers {
	return &UploadHandlers{Svc: svc}
}

func (h *UploadHandlers) UploadStatement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file required", http.StatusBadRequest)
		return
	}
	defer file.Close()
	if err := h.Svc.UploadAndStore(file); err != nil {
		respondErr(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]any{
		"status":  "ok",
		"message": "file uploaded successfully",
	})
}

func (h *UploadHandlers) Reset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	h.Svc.Reset()
	respondJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
		"data":   "reset successful",
	})
}
