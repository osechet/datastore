package query

import (
	"log"
	"reflect"
	"sort"

	"github.com/osechet/go-datastore/filter"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

func AutoQuery(storage Storage, query datastore.Query, t reflect.Type, results ResultSet) error {
	scanner := storage.Scanner()
	comparator := filter.MakeComparator(query, t)
	for scanner.HasNext() {
		key := scanner.Next()
		item, err := storage.ItemFor(key)
		if err != nil {
			log.Printf("Cannot get item for %v: %v", key, err)
		} else if filter.Match(query.Filter, item) {
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
	return scanner.Err()
}
