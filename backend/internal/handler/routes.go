package handler

import (
	"net/http"

	"github.com/arm02/bank-statement-viewer/backend/internal/service"
)

func RegisterRoutes(mux *http.ServeMux, svc *service.TransactionService) {
	uploadHandlers := NewUploadHandlers(svc)
	issuesHandlers := NewIssuesHandlers(svc)
	balanceHandlers := NewBalanceHandlers(svc)

	mux.HandleFunc("/api/upload", uploadHandlers.UploadStatement)
	mux.HandleFunc("/api/balance", balanceHandlers.GetBalance)
	mux.HandleFunc("/api/issues", issuesHandlers.GetAllIssues)
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, _ *http.Request) {
		respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
}
