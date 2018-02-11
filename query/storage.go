package query

import (
	"github.com/golang/protobuf/descriptor"
)

type Scanner interface {
	HasNext() bool
	Next() interface{}
	Err() error
}

type Storage interface {
	Scanner() Scanner
	ItemFor(key interface{}) (descriptor.Message, error)
}

type SliceStorageScanner struct {
	storage SliceStorage
	current int
}

func NewSliceStorageScanner(storage SliceStorage) *SliceStorageScanner {
	return &SliceStorageScanner{
		storage,
		0,
	}
}

func (s SliceStorageScanner) HasNext() bool {
	return s.current < len(s.storage.items)
}

func (s *SliceStorageScanner) Next() interface{} {
	ret := s.current
	s.current++
	return ret
}

func (s SliceStorageScanner) Err() error {
	return nil
}

type SliceStorage struct {
	items []descriptor.Message
}

func (s SliceStorage) Scanner() Scanner {
	return NewSliceStorageScanner(s)
}

func (s SliceStorage) ItemFor(key interface{}) (descriptor.Message, error) {
	return s.items[key.(int)], nil
}
