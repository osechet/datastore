// Utility functions to run tests.

package filter

import datastore "google.golang.org/genproto/googleapis/datastore/v1"

func makeDoubleValue(value float64) *datastore.Value {
	v := &datastore.Value{}
	v.ValueType = &datastore.Value_DoubleValue{DoubleValue: value}
	return v
}

func makeIntegerValue(value int64) *datastore.Value {
	v := &datastore.Value{}
	v.ValueType = &datastore.Value_IntegerValue{IntegerValue: value}
	return v
}

func makeStringValue(value string) *datastore.Value {
	v := &datastore.Value{}
	v.ValueType = &datastore.Value_StringValue{StringValue: value}
	return v
}
