// Copyright 2023 © Tokenomy. All rights reserved.
package model

import "sort"

type (
	Data struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	dataSorter struct {
		lessFunc []dataLessFunc
		slice    List
	}
	dataLessFunc func(a, b *Data) bool
	List         []*Data
)

func (l List) Filter(list List, f func(data *Data) bool) List {
	result := make(List, 0)
	for _, data := range list {
		if f(data) {
			result = append(result, data)
		}
	}
	return result
}

func (ds *dataSorter) Len() int { return len(ds.slice) }

func (ds *dataSorter) Less(i, j int) bool {
	var n int
	p, q := ds.slice[i], ds.slice[j]
	for n = 0; n < len(ds.lessFunc)-1; n++ {
		f := ds.lessFunc[n]
		switch {
		case f(p, q):
			return true
		case f(q, p):
			return false
		}
	}
	return ds.lessFunc[n](p, q)
}

func (ds *dataSorter) Sort(slice List) {
	ds.slice = slice
	sort.Sort(ds)
}

func (ds *dataSorter) Swap(i, j int) { ds.slice[i], ds.slice[j] = ds.slice[j], ds.slice[i] }

func NewData(id int, name string) *Data { return &Data{id, name} }
