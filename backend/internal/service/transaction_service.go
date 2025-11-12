package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/arm02/bank-statement-viewer/backend/internal/model"
	"github.com/arm02/bank-statement-viewer/backend/internal/repository"
)

type TransactionService struct {
	repo *repository.InMemoryRepo
}

func NewTransactionService(r *repository.InMemoryRepo) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) ParseCSV(r io.Reader) ([]model.Transaction, error) {
	cr := csv.NewReader(r)
	cr.TrimLeadingSpace = true
	var out []model.Transaction
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("csv read error: %w", err)
		}
		if len(rec) < 6 {
			return nil, fmt.Errorf("invalid record length: %v", rec)
		}

		tsInt, err := strconv.ParseInt(strings.TrimSpace(rec[0]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp '%s': %w", rec[0], err)
		}
		ts := time.Unix(tsInt, 0)
		name := strings.TrimSpace(rec[1])
		typeStr := strings.ToUpper(strings.TrimSpace(rec[2]))
		var txType model.TxType
		if typeStr == string(model.Credit) {
			txType = model.Credit
		} else {
			txType = model.Debit
		}
		amt, err := strconv.ParseInt(strings.TrimSpace(rec[3]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid amount '%s': %w", rec[3], err)
		}
		status := model.Status(strings.ToUpper(strings.TrimSpace(rec[4])))
		desc := strings.TrimSpace(rec[5])
		out = append(out, model.Transaction{
			Timestamp:   ts,
			Name:        name,
			Type:        txType,
			Amount:      amt,
			Status:      status,
			Description: desc,
		})
	}
	return out, nil
}

func (s *TransactionService) UploadAndStore(r io.Reader) error {
	txs, err := s.ParseCSV(r)
	if err != nil {
		return err
	}
	return s.repo.StoreMany(txs)
}

func (s *TransactionService) ComputeBalance() int64 {
	all := s.repo.ListAll()
	var bal int64 = 0
	for _, t := range all {
		if t.Status != model.Success {
			continue
		}
		if t.Type == model.Credit {
			bal += t.Amount
		} else {
			bal -= t.Amount
		}
	}
	return bal
}

func (s *TransactionService) Issues(page, limit int, sortBy, sortOrder string) ([]model.Transaction, model.Meta) {
	all := s.repo.ListAll()
	filtered := make([]model.Transaction, 0, len(all))
	for _, t := range all {
		if t.Status != model.Success {
			filtered = append(filtered, t)
		}
	}

	sortBy = strings.ToLower(strings.TrimSpace(sortBy))
	sortOrder = strings.ToLower(strings.TrimSpace(sortOrder))

	if sortBy == "" {
		sortBy = "timestamp"
	}

	if sortOrder != "desc" && sortOrder != "asc" {
		sortOrder = "asc"
	}

	sort.Slice(filtered, func(i, j int) bool {
		a, b := filtered[i], filtered[j]
		switch sortBy {
		case "amount":
			if sortOrder == "desc" {
				return a.Amount > b.Amount
			}
			return a.Amount < b.Amount
		case "name":
			la := strings.ToLower(a.Name)
			lb := strings.ToLower(b.Name)
			if sortOrder == "desc" {
				return la > lb
			}
			return la < lb
		case "type":
			la := strings.ToLower(string(a.Type))
			lb := strings.ToLower(string(b.Type))
			if sortOrder == "desc" {
				return la > lb
			}
			return la < lb
		case "status":
			sa := strings.ToLower(string(a.Status))
			sb := strings.ToLower(string(b.Status))
			if sortOrder == "desc" {
				return sa > sb
			}
			return sa < sb
		case "timestamp":
			fallthrough
		default:
			if sortOrder == "desc" {
				return a.Timestamp.After(b.Timestamp)
			}
			return a.Timestamp.Before(b.Timestamp)
		}
	})

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	total := len(filtered)
	start := (page - 1) * limit
	if start >= total {
		return make([]model.Transaction, 0), model.NewMeta(page, limit, total)
	}
	end := start + limit
	if end > total {
		end = total
	}

	paged := filtered[start:end]
	meta := model.NewMeta(page, limit, total)
	return paged, meta
}

func (s *TransactionService) Reset() {
	s.repo.Reset()
}
