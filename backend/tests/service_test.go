package tests

import (
	"strings"
	"testing"

	"github.com/arm02/bank-statement-viewer/backend/internal/repository"
	"github.com/arm02/bank-statement-viewer/backend/internal/service"
)

func TestParseAndCompute(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	svc := service.NewTransactionService(repo)

	csv := `1624507883, JOHN DOE, DEBIT, 250000, SUCCESS, restaurant
1624608050, E-COMMERCE A, DEBIT, 150000, FAILED, clothes
1624512883, COMPANY A, CREDIT, 12000000, SUCCESS, salary
1624615065, E-COMMERCE B, DEBIT, 150000, PENDING, clothes
`

	if err := svc.UploadAndStore(strings.NewReader(csv)); err != nil {
		t.Fatalf("upload failed: %v", err)
	}

	// Check balance calculation only success
	bal := svc.ComputeBalance()
	expected := int64(11750000) // 12,000,000 - 250,000
	if bal != expected {
		t.Fatalf("unexpected balance: got %d, want %d", bal, expected)
	}

	// Check issues with pagination
	issues, meta := svc.Issues(1, 10, "timestamp", "asc")
	if meta.Total != 2 {
		t.Fatalf("expected total issues = 2, got %d", meta.Total)
	}
	if len(issues) != 2 {
		t.Fatalf("expected 2 issues on page 1, got %d", len(issues))
	}

	// Check the content
	first := issues[0]
	if first.Status != "FAILED" && first.Status != "PENDING" {
		t.Fatalf("unexpected issue status: %s", first.Status)
	}
}
