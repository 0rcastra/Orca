package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/0rcastra/Orca/handler"
	"github.com/0rcastra/Orca/internal/data"
	"github.com/gorilla/mux"
)

func TestSetHandler(t *testing.T) {
	db := data.NewDatabase()

	reqBody := strings.NewReader("")
	req := httptest.NewRequest(http.MethodPost, "/set/mykey/myvalue", reqBody)

	res := httptest.NewRecorder()

	h := handler.NewHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/set/{key}/{value}", h.SetHandler).Methods(http.MethodPost)

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", res.Code, http.StatusOK)
	}

	if value, exists := db.Get("mykey"); !exists || value != "myvalue" {
		t.Errorf("unexpected stored value: got %s, want %s", value, "myvalue")
	}
}
