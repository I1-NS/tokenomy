// Copyright 2023 © Tokenomy. All rights reserved.
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/i1-ns/tokenomy/api/rfc"
	"github.com/i1-ns/tokenomy/pkg/http/domain/service"
	"github.com/i1-ns/tokenomy/pkg/http/interface/static"
	"github.com/i1-ns/tokenomy/pkg/http/usecase"
)

type Data struct {
	usecase usecase.Data
}

func (d Data) distinct(slice ...int) []int {
	ib, ii := make(map[int]bool), make([]int, 0, len(slice))
	for _, i := range slice {
		if _, ok := ib[i]; !ok { ib[i], ii = true, append(ii, i) }
	}
	return ii
}

func GetByIDs(w http.ResponseWriter, r *http.Request) {
	params, ch := make(map[string]string, 0), make(chan *rfc.Response, 1)
	for k := range r.URL.Query() { params[k] = r.URL.Query().Get(k) }
	ids, sid := make([]int, 0), make([]string, 0)
	for k, v := range params {
		if k == "id" { sid = strings.Split(v, ",") }
	}
	for i := 0; i < len(sid); i++ {
		id, err := strconv.Atoi(sid[i])
		if err != nil { rfc.NewResponse(http.StatusBadRequest, nil, err).Error(w); return }
		ids = append(ids, id)
	}
	go func() {
		dummy := newData()
		data, _, err := dummy.usecase.GetByIDs(dummy.distinct(ids...))
		ch <- rfc.NewResponse(http.StatusOK, data, err)
	}()
	select {
	case c := <-ch:
		_, data, err := c.Get()
		if err != nil { rfc.NewResponse(http.StatusNotFound, nil, err).Error(w); return }
		json.NewEncoder(w).Encode(rfc.NewResponse(http.StatusOK, data, err))
	case <-r.Context().Done():
		return
	}
}

func newData() *Data {
	return &Data{usecase.NewData(static.NewData(), service.NewData)}
}