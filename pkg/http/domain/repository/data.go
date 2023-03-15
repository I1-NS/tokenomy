// Copyright 2023 © Tokenomy. All rights reserved.
package repository

import "github.com/i1-ns/tokenomy/pkg/http/domain/model"

type Data interface {
	ReadByIDs([]int) (model.List, int, error)
}
