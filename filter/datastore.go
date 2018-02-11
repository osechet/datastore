package filter

import (
	"log"
	"reflect"

	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

// NewComparator creates a new Comparator from the given query's PropertyOrder.
func NewComparator(order *datastore.PropertyOrder, t reflect.Type) *PropertyComparator {
	property := order.GetProperty().GetName()
	if len(property) == 0 {
		log.Println("Invalid order: property name is not set")
		return nil
	}
	if order.Direction == datastore.PropertyOrder_DESCENDING {
		return NewPropertyComparator(t, property, Descending)
	}
	return NewPropertyComparator(t, property, Ascending)
}
