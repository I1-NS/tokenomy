// Copyright 2023 © Tokenomy. All rights reserved.
package rfc

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (r *Response) Error(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

func (r *Response) Get() (int, interface{}, error) {
	f := func(s string) error {
		if r.Message == "" { return nil }
		return errors.New(r.Message)
	}
	return r.Code, r.Data, f(r.Message)
}

func NewResponse(code int, data interface{}, err error) *Response {
	response := Response{
		Code: code,
		Data: data,
	}
	if err != nil { response.Message = err.Error() }
	return &response
}