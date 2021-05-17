package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ppai-plivo/ut/mocks"
	"github.com/ppai-plivo/ut/pkg/db"

	lmocks "github.com/plivo/pkg/log/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandlerSuccess(t *testing.T) {
	assert := require.New(t)

	store := new(mocks.Store)
	store.On("GetUserByID", mock.Anything, 123).Return(&db.User{
		Name:  "John Doe",
		Email: "support@plivo.com",
	}, nil)

	logger := new(lmocks.Logger)

	svc := New(store, logger)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/123", nil)
	assert.NotNil(req)
	assert.Nil(err)

	w := httptest.NewRecorder()
	svc.HandleRequest(w, req)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(w.Header().Get("Content-Type"), "application/json")

	store.AssertExpectations(t)
}
