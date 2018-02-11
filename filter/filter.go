package filter

import (
	"github.com/golang/protobuf/descriptor"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

// Match checks if the given message matches the given filter.
func Match(filter datastore.Filter, message descriptor.Message) bool {
	switch filter.FilterType.(type) {
	case *datastore.Filter_PropertyFilter:
		impl := filter.GetPropertyFilter()
		property := impl.GetProperty()
		if property == nil {
			return false
		}
		value := impl.GetValue()
		if value == nil {
			return false
		}
		return compare(message, property.Name, impl.GetOp(), *value)
	case *datastore.Filter_CompositeFilter:
		filters := filter.GetCompositeFilter().GetFilters()
		matched := len(filters) > 0
		for _, filter := range filters {
			matched = matched && Match(*filter, message)
		}
		return matched
	}
	return false
}

func compare(message descriptor.Message, property string, op datastore.PropertyFilter_Operator, reference datastore.Value) bool {
	value := valueOf(message, property)
	switch op {
	case datastore.PropertyFilter_LESS_THAN:
		return lt(value, reference)
	case datastore.PropertyFilter_LESS_THAN_OR_EQUAL:
		return lte(value, reference)
	case datastore.PropertyFilter_GREATER_THAN:
		return gt(value, reference)
	case datastore.PropertyFilter_GREATER_THAN_OR_EQUAL:
		return gte(value, reference)
	case datastore.PropertyFilter_EQUAL:
		return equal(value, reference)
	case datastore.PropertyFilter_HAS_ANCESTOR:
		return hasAncestor(value, reference)
	}
	return false
}

func lt(value interface{}, reference datastore.Value) bool {
	switch v := value.(type) {
	case float32:
		return v < float32(reference.GetDoubleValue())
	case float64:
		return v < reference.GetDoubleValue()
	case int32:
		return v < int32(reference.GetIntegerValue())
	case int64:
		return v < reference.GetIntegerValue()
	case string:
		return v < reference.GetStringValue()
	}
	return false
}

func lte(value interface{}, reference datastore.Value) bool {
	switch v := value.(type) {
	case float32:
		return v <= float32(reference.GetDoubleValue())
	case float64:
		return v <= reference.GetDoubleValue()
	case int32:
		return v <= int32(reference.GetIntegerValue())
	case int64:
		return v <= reference.GetIntegerValue()
	case string:
		return v <= reference.GetStringValue()
	}
	return false
}

func gt(value interface{}, reference datastore.Value) bool {
	switch v := value.(type) {
	case float32:
		return v > float32(reference.GetDoubleValue())
	case float64:
		return v > reference.GetDoubleValue()
	case int32:
		return v > int32(reference.GetIntegerValue())
	case int64:
		return v > reference.GetIntegerValue()
	case string:
		return v > reference.GetStringValue()
	}
	return false
}

func gte(value interface{}, reference datastore.Value) bool {
	switch v := value.(type) {
	case float32:
		return v >= float32(reference.GetDoubleValue())
	case float64:
		return v >= reference.GetDoubleValue()
	case int32:
		return v >= int32(reference.GetIntegerValue())
	case int64:
		return v >= reference.GetIntegerValue()
	case string:
		return v >= reference.GetStringValue()
	}
	return false
}

func equal(value interface{}, reference datastore.Value) bool {
	switch v := value.(type) {
	case float32:
		return v == float32(reference.GetDoubleValue())
	case float64:
		return v == reference.GetDoubleValue()
	case int32:
		return v == int32(reference.GetIntegerValue())
	case int64:
		return v == reference.GetIntegerValue()
	case string:
		return v == reference.GetStringValue()
	}
	return false
}

func hasAncestor(value interface{}, reference datastore.Value) bool {
	return false
}
