package handler

import (
	"net/http"
	"strconv"

	"github.com/arm02/bank-statement-viewer/backend/internal/service"
)

type IssuesHandlers struct {
	Svc *service.TransactionService
}

func NewIssuesHandlers(svc *service.TransactionService) *IssuesHandlers {
	return &IssuesHandlers{Svc: svc}
}

func (h *IssuesHandlers) GetAllIssues(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	sortBy := r.URL.Query().Get("sort_by")
	sortOrder := r.URL.Query().Get("sort_order")

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	issues, meta := h.Svc.Issues(page, limit, sortBy, sortOrder)

	resp := map[string]interface{}{
		"data": issues,
		"meta": meta,
	}

	respondJSON(w, http.StatusOK, resp)
}
