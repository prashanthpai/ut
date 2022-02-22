package main

import (
	"net/http"

	"github.com/prashanthpai/ut/pkg/db"
	"github.com/prashanthpai/ut/pkg/log"
)

func main() {
	logger := log.New()

	dbClient, err := db.NewClient("db endpoint")
	if err != nil {
		logger.Errorf("db.NewClient() failed: error: %s", err.Error())
		return
	}

	svc := New(dbClient, logger)

	http.HandleFunc("/", svc.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
