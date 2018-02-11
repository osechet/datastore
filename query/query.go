package query

import (
	"reflect"
	"sort"

	"github.com/osechet/go-datastore/filter"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

func AutoQuery(storage Storage, query datastore.Query, t reflect.Type, results ResultSet) {
	scanner := storage.Scanner()
	comparator := filter.MakeComparator(query, t)
	for scanner.HasNext() {
		key := scanner.Next()
		item := storage.ItemFor(key)
		if filter.Match(query.Filter, item) {
			if comparator.HasNested() {
				index := sort.Search(results.Len(), func(i int) bool {
					return comparator.Less(item, results.At(i))
				})
				results.Insert(item, index)
			} else {
				results.Append(item)
			}
		}
	}
}
