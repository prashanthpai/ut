package main

import (
	stdlog "log"
	"net/http"

	"github.com/ppai-plivo/ut/pkg/db"

	"github.com/plivo/pkg/log"
)

func main() {
	logger, err := log.NewLogger(&log.Config{
		Environment: "prod",
	})
	if err != nil {
		stdlog.Fatalf("log.NewLogger() failed: %s", err.Error())
	}
	defer logger.Sync()

	dbClient, err := db.NewClient("db endpoint")
	if err != nil {
		logger.Fatalw("db.NewClient() failed", "error", err.Error())
	}

	svc := New(dbClient, logger)

	http.HandleFunc("/", svc.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
