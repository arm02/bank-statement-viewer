package handler

import (
	"net/http"
	"strconv"

	"github.com/arm02/bank-statement-viewer/backend/internal/service"
)

type BalanceHandlers struct {
	Svc *service.TransactionService
}

func NewBalanceHandlers(svc *service.TransactionService) *BalanceHandlers {
	return &BalanceHandlers{Svc: svc}
}

func (h *BalanceHandlers) GetBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	bal := h.Svc.ComputeBalance()
	resp := map[string]string{"balance": strconv.FormatInt(bal, 10)}
	respondJSON(w, http.StatusOK, resp)
}
