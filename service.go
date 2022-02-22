package main

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/prashanthpai/ut/pkg/db"
	"github.com/prashanthpai/ut/pkg/log"
)

type Store interface {
	GetUserByID(ctx context.Context, id int) (*db.User, error)
}

type Service struct {
	store  Store
	logger log.Logger
}

func New(store Store, logger log.Logger) *Service {
	return &Service{
		store:  store,
		logger: logger,
	}
}

func (s *Service) HandleRequest(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := s.store.GetUserByID(context.TODO(), userID)
	if err != nil {
		s.logger.Errorf("store.GetUserByID() failed", err.Error())
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		s.logger.Errorf("json.Marshal() failed", err.Error())
		http.Error(w, "server malfunctioned", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		s.logger.Errorf("w.Write() failed", err.Error())
	}
}
