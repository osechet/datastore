package filter

import (
	"reflect"

	"github.com/golang/protobuf/proto"
)

// FieldIndex returns the index of the field for the given property in the
// given type. The property name is based on the field name in the protobuf
// description.
func FieldIndex(t reflect.Type, property string) int {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup("protobuf")
		if ok {
			props := proto.Properties{}
			props.Parse(tag)
			if props.OrigName == property {
				return i
			}
		}
	}
	return -1
}

// Pre: The fieldIndex is considered valid.
func valueOfField(message interface{}, fieldIndex int) interface{} {
	v := reflect.Indirect(reflect.ValueOf(message))
	return v.Field(fieldIndex).Interface()
}

func valueOf(message interface{}, property string) interface{} {
	t := reflect.Indirect(reflect.ValueOf(message)).Type()
	index := FieldIndex(t, property)
	if index >= 0 {
		v := reflect.Indirect(reflect.ValueOf(message))
		return v.Field(index).Interface()
	}
	return nil
}
