package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arm02/bank-statement-viewer/backend/internal/handler"
	"github.com/arm02/bank-statement-viewer/backend/internal/repository"
	"github.com/arm02/bank-statement-viewer/backend/internal/service"
)

func main() {
	repo := repository.NewInMemoryRepo()
	svc := service.NewTransactionService(repo)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux, svc)

	handler := cors(mux)

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	server := &http.Server{Addr: addr, Handler: handler, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}
	log.Printf("listening on %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
