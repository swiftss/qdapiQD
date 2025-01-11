package sign

import (
	"errors"
	"slices"
)

var ErrNotFound error = errors.New("key not found")
var ErrOutOfRange error = errors.New("OutOfRange")

type MetaRW struct {
	data       []string
	dataStruct []string
}

func (rw *MetaRW) SDkGet(key string) (string, error) {
	index := slices.Index(rw.dataStruct, key)
	if index < 0 {
		return "", ErrNotFound
	}
	if index < len(rw.dataStruct) {
		return rw.data[index], nil
	}
	return "", ErrOutOfRange
}
func (rw *MetaRW) Modify(key, val string) error {
	index := slices.Index(rw.dataStruct, key)
	if index < 0 {
		return ErrNotFound
	}
	if index < len(rw.dataStruct) {
		rw.data[index] = val
		return nil
	}
	return ErrOutOfRange
}
