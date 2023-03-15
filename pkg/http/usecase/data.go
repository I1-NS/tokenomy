// Copyright 2023 © Tokenomy. All rights reserved.
package usecase

import (
	"github.com/i1-ns/tokenomy/pkg/http/domain/model"
	"github.com/i1-ns/tokenomy/pkg/http/domain/repository"
	"github.com/i1-ns/tokenomy/pkg/http/domain/service"
)

type (
	Data interface {
		GetByIDs([]int) (model.List, int, error)
	}
	data struct {
		repository repository.Data
		service    *service.Data
	}
)

func (d *data) GetByIDs(ids []int) (model.List, int, error) { return d.repository.ReadByIDs(ids) }

func NewData(repository repository.Data, service service.DataFunc) *data {
	return &data{repository, service(repository)}
}