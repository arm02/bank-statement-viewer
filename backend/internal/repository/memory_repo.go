package repository

import (
	"sync"
	"time"

	"github.com/arm02/bank-statement-viewer/backend/internal/model"
)

type InMemoryRepo struct {
	mu           sync.RWMutex
	transactions []model.Transaction
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{}
}

func (r *InMemoryRepo) StoreMany(txs []model.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.transactions = append(r.transactions, txs...)
	return nil
}

func (r *InMemoryRepo) ListAll() []model.Transaction {
	r.mu.RLock()
	defer r.mu.RUnlock()
	copyTxs := make([]model.Transaction, len(r.transactions))
	copy(copyTxs, r.transactions)
	return copyTxs
}

func (r *InMemoryRepo) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.transactions = nil
}

func (r *InMemoryRepo) SeedSample() {
	r.StoreMany([]model.Transaction{
		{Timestamp: time.Now(), Name: "JOHN DOE", Type: model.Debit, Amount: 250000, Status: model.Success, Description: "restaurant"},
		{Timestamp: time.Now(), Name: "E-COMMERCE A", Type: model.Debit, Amount: 150000, Status: model.Failed, Description: "clothes"},
		{Timestamp: time.Now(), Name: "COMPANY A", Type: model.Credit, Amount: 12000000, Status: model.Success, Description: "salary"},
	})
}
