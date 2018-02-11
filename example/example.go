package example

import (
	"reflect"
	"sort"

	test "github.com/osechet/go-datastore/_proto/osechet/test"
	"github.com/osechet/go-datastore/filter"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

// Query is an example on how to use the comparators.
func Query(query datastore.Query, dbBooks []*test.Book) []*test.Book {
	comparators := make([]filter.Comparator, 0)
	for _, order := range query.Order {
		property := order.Property.Name
		if order.Direction == datastore.PropertyOrder_DESCENDING {
			c := filter.NewPropertyComparator(reflect.TypeOf(test.Book{}), property, filter.Descending)
			comparators = append(comparators, c)
		} else {
			c := filter.NewPropertyComparator(reflect.TypeOf(test.Book{}), property, filter.Ascending)
			comparators = append(comparators, c)
		}
	}
	comparator := filter.NewCompositeComparator(comparators)
	books := make([]*test.Book, 0)
	for _, book := range dbBooks {
		if filter.Match(query.Filter, book) {
			if len(comparators) > 0 {
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
