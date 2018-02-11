package filter

import (
	"reflect"
	"testing"

	test "github.com/osechet/datastore/_proto/osechet/test"
)

func TestFieldIndex(t *testing.T) {
	type args struct {
		t        reflect.Type
		property string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"unknwon", args{reflect.TypeOf(test.Tested{}), "none"}, -1},
		{"valid", args{reflect.TypeOf(test.Tested{}), "int32_value"}, 2},
		{"not protobuf", args{reflect.TypeOf(NotProtobufType{}), "int32_value"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldIndex(tt.args.t, tt.args.property); got != tt.want {
				t.Errorf("FieldIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valueOfField(t *testing.T) {
	type args struct {
		message    interface{}
		fieldIndex int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"valid", args{test.Tested{Int32Value: 42, Int64Value: 35}, 2}, int32(42)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valueOfField(tt.args.message, tt.args.fieldIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valueOfField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valueOf(t *testing.T) {
	type args struct {
		message  interface{}
		property string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"valid", args{test.Tested{Int32Value: 42, Int64Value: 5}, "int32_value"}, int32(42)},
		{"invalid property", args{test.Tested{Int32Value: 42, Int64Value: 35}, "none"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valueOf(tt.args.message, tt.args.property); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valueOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
