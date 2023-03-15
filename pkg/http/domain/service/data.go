// Copyright 2023 © Tokenomy. All rights reserved.
package service

import "github.com/i1-ns/tokenomy/pkg/http/domain/repository"

type (
	Data struct {
		repository repository.Data
	}
	DataFunc func(repository.Data) *Data
)

func NewData(repository repository.Data) *Data { return &Data{repository} }