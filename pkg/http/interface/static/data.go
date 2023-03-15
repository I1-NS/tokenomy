// Copyright 2023 © Tokenomy. All rights reserved.
package static

import (
	"fmt"

	"github.com/i1-ns/tokenomy/pkg/http/domain/model"
)

type data struct {
	list model.List
}

var dummy = model.List{model.NewData(1, "A"), model.NewData(2, "B"), model.NewData(3, "C")}

func (d *data) ReadByIDs(ids []int) (model.List, int, error) {
	if len(ids) < 1 { return d.list, 3, nil }
	list := d.list[:0]
	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(d.list); j++ {
			if d.list[j].ID == ids[i] { list = append(list, d.list[j]) }
		}
	}
	if len(list) < 1 { return nil, 0, fmt.Errorf("resource with ID %v doesn't exists", ids) }
	return list, len(list), nil
}

func NewData() *data { return &data{dummy} }
