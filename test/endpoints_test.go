// Copyright 2023 © Tokenomy. All rights reserved.
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/i1-ns/tokenomy/api/handler"
)

func TestGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v0/tokenomy", nil)
	if err != nil { t.Fatal(err) }
	req.URL.RawQuery = req.URL.Query().Encode()
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.GetByIDs).ServeHTTP(rec, req)
	if n := rec.Code; n != http.StatusOK { t.Errorf("Wrong status code: got %v want %v", n, http.StatusOK) }
	mock := `{"code":200,"data":[{"id":1,"name":"A"},{"id":2,"name":"B"},{"id":3,"name":"C"}]}`
	resp := rec.Body.String()
	if value := resp[:len(resp) - 1]; value != mock { t.Errorf("unexpected body, got %v want %v", resp, mock) }
}

func TestGetByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v0/tokenomy?id=1", nil)
	if err != nil { t.Fatal(err) }
	req.URL.RawQuery = req.URL.Query().Encode()
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.GetByIDs).ServeHTTP(rec, req)
	if n := rec.Code; n != http.StatusOK { t.Errorf("Wrong status code: got %v want %v", n, http.StatusOK) }
	mock := `{"code":200,"data":[{"id":1,"name":"A"}]}`
	resp := rec.Body.String()
	if value := resp[:len(resp) - 1]; value != mock { t.Errorf("unexpected body, got %v want %v", resp, mock) }
}

func TestGetByIDs(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v0/tokenomy?id=1,3,4", nil)
	if err != nil { t.Fatal(err) }
	req.URL.RawQuery = req.URL.Query().Encode()
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.GetByIDs).ServeHTTP(rec, req)
	if n := rec.Code; n != http.StatusOK { t.Errorf("Wrong status code: got %v want %v", n, http.StatusOK) }
	mock := `{"code":200,"data":[{"id":1,"name":"A"},{"id":3,"name":"C"}]}`
	resp := rec.Body.String()
	if value := resp[:len(resp) - 1]; value != mock { t.Errorf("unexpected body, got %v want %v", resp, mock) }
}

func TestGetInvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v0/tokenomy?id=A", nil)
	if err != nil { t.Fatal(err) }
	req.URL.RawQuery = req.URL.Query().Encode()
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.GetByIDs).ServeHTTP(rec, req)
	if n := rec.Code; n != http.StatusBadRequest { t.Errorf("Wrong status code: got %v want %v", n, http.StatusBadRequest) }
	mock := `{"code":400,"message":"strconv.Atoi: parsing \"A\": invalid syntax"}`
	resp := rec.Body.String()
	if value := resp[:len(resp) - 1]; value != mock { t.Errorf("unexpected body, got %v want %v", resp, mock) }
}

func TestGetIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v0/tokenomy?id=4", nil)
	if err != nil { t.Fatal(err) }
	req.URL.RawQuery = req.URL.Query().Encode()
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.GetByIDs).ServeHTTP(rec, req)
	if n := rec.Code; n != http.StatusNotFound { t.Errorf("Wrong status code: got %v want %v", n, http.StatusNotFound) }
	mock := `{"code":404,"message":"resource with ID [4] doesn't exists"}`
	resp := rec.Body.String()
	if value := resp[:len(resp) - 1]; value != mock { t.Errorf("unexpected body, got %v want %v", resp, mock) }
}