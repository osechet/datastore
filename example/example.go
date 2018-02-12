package example

import (
	"reflect"
	"sort"

	"github.com/golang/protobuf/descriptor"
	test "github.com/osechet/go-datastore/_proto/osechet/test"
	"github.com/osechet/go-datastore/filter"
	"github.com/osechet/go-datastore/query"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

func Apply(q datastore.Query, books []*test.Book) ([]*test.Book, error) {
	results := NewBookResultSet()
	err := query.Apply(NewBookStorage(books), q, reflect.TypeOf(test.Book{}), results)
	return results.Books, err
}

type BookScanner struct {
	storage BookStorage
	current int
}

func NewBookScanner(storage BookStorage) *BookScanner {
	return &BookScanner{
		storage,
		0,
	}
}

func (s BookScanner) HasNext() bool {
	return s.current < len(s.storage.books)
}

func (s *BookScanner) Next() interface{} {
	ret := s.current
	s.current++
	return ret
}

func (s *BookScanner) Err() error {
	return nil
}

type BookStorage struct {
	books []*test.Book
}

func NewBookStorage(books []*test.Book) *BookStorage {
	return &BookStorage{
		books,
	}
}

func (s BookStorage) Scanner() query.Scanner {
	return NewBookScanner(s)
}

func (s BookStorage) ItemFor(key interface{}) (descriptor.Message, error) {
	return s.books[key.(int)], nil
}

type BookResultSet struct {
	Books []*test.Book
}

func NewBookResultSet() *BookResultSet {
	return &BookResultSet{
		make([]*test.Book, 0),
	}
}

func (rs *BookResultSet) Len() int {
	return len(rs.Books)
}

func (rs *BookResultSet) At(index int) interface{} {
	return rs.Books[index]
}

func (rs *BookResultSet) Insert(book interface{}, index int) {
	rs.Books = append(rs.Books, nil)
	copy(rs.Books[index+1:], rs.Books[index:])
	rs.Books[index] = book.(*test.Book)
}

func (rs *BookResultSet) Append(book interface{}) {
	rs.Books = append(rs.Books, book.(*test.Book))
}

// Query is an example on how to use the comparators.
func Query(query datastore.Query, dbBooks []*test.Book) []*test.Book {
	comparator := filter.MakeComparator(query, reflect.TypeOf(test.Book{}))
	books := make([]*test.Book, 0)
	for _, book := range dbBooks {
		if filter.Match(query.Filter, book) {
			if comparator.HasNested() {
				index := sort.Search(len(books), func(i int) bool {
					return comparator.Less(book, books[i])
				})
				books = append(books, nil)
				copy(books[index+1:], books[index:])
				books[index] = book
			} else {
				books = append(books, book)
			}
		}
	}
	return books
}
