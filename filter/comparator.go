package filter

import (
	"reflect"

	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

// Comparator describes the behavior of a comparable.
type Comparator interface {
	Less(a, b interface{}) bool
	Equals(a, b interface{}) bool
	Greater(a, b interface{}) bool
}

// CompositeComparator is composed of several others Comparators.
// CompositeComparator is not a Comparator, mainly to avoid nesting
// CompositeComparator inside CompositeComparator.
type CompositeComparator struct {
	comparators []Comparator
}

// NewCompositeComparator creates a new CompositeComparator from the given list of Comparators.
func NewCompositeComparator(comparators []Comparator) *CompositeComparator {
	return &CompositeComparator{
		comparators,
	}
}

// HasNested returns true if the CompositeComparator has nested comparators.
func (c CompositeComparator) HasNested() bool {
	return len(c.comparators) > 0
}

// Less returns true if a is less than b based on all the inner comparators.
// The comparison starts with the first Comparator of the list. If a is less than b using this Comparator,
// the function returns. If a equals b, the next comparator is used, and so on.
func (c CompositeComparator) Less(a, b interface{}) bool {
	for _, comparator := range c.comparators {
		if comparator.Less(a, b) {
			return true
		}
		if comparator.Greater(a, b) {
			return false
		}
	}
	return false
}

// MakeComparator creates a CompositeComparator from the orders of the given query.
func MakeComparator(query datastore.Query, t reflect.Type) *CompositeComparator {
	comparators := make([]Comparator, 0)
	for _, order := range query.Order {
		comparator := NewComparator(order, t)
		if comparator != nil {
			comparators = append(comparators, comparator)
		}
	}
	return NewCompositeComparator(comparators)
}

// Direction describes the direction of a sort operation.
type Direction int32

const (
	// Ascending direction
	Ascending Direction = iota
	// Descending direction
	Descending
)

// ValueComparator compares given data on their values.
type ValueComparator struct {
	direction Direction
}

// Less returns true if a is less than b.
func (c ValueComparator) Less(a, b interface{}) bool {
	if c.direction == Ascending {
		return Less(a, b)
	}
	return Less(b, a)
}

// Equals returns true if a equals b.
func (c ValueComparator) Equals(a, b interface{}) bool {
	return Equals(a, b)
}

// Greater returns true if a is greater than b.
func (c ValueComparator) Greater(a, b interface{}) bool {
	if c.direction == Ascending {
		return Greater(a, b)
	}
	return Greater(b, a)
}

// PropertyComparator compares data according to the value of the associated property.
type PropertyComparator struct {
	fieldIndex int
	direction  Direction
}

// NewPropertyComparator creates a new PropertyComparator for the given type and property.
// If the given property does not exists in the given type, returns nil.
// The property is based on the protobuf name.
func NewPropertyComparator(t reflect.Type, property string, direction Direction) *PropertyComparator {
	fieldIndex := FieldIndex(t, property)
	if fieldIndex < 0 {
		return nil
	}
	return &PropertyComparator{
		fieldIndex,
		direction,
	}
}

// Less returns true if the value of associated property of a is less than the
// value of the same property of b.
func (c PropertyComparator) Less(a, b interface{}) bool {
	if c.direction == Ascending {
		return Less(valueOfField(a, c.fieldIndex), valueOfField(b, c.fieldIndex))
	}
	return Less(valueOfField(b, c.fieldIndex), valueOfField(a, c.fieldIndex))
}

// Equals returns true if the value of associated property of a equals the
// value of the same property of b.
func (c PropertyComparator) Equals(a, b interface{}) bool {
	return Equals(valueOfField(a, c.fieldIndex), valueOfField(b, c.fieldIndex))
}

// Greater returns true if the value of associated property of a is greater
// than the value of the same property of b.
func (c PropertyComparator) Greater(a, b interface{}) bool {
	if c.direction == Ascending {
		return Greater(valueOfField(a, c.fieldIndex), valueOfField(b, c.fieldIndex))
	}
	return Greater(valueOfField(b, c.fieldIndex), valueOfField(a, c.fieldIndex))
}
