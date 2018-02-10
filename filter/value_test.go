package filter

import (
	"reflect"
	"testing"
)

type Tested1 struct {
	Field1 int `protobuf:"varint,1,opt,name=field1"`
	Field2 int `protobuf:"varint,2,opt,name=field2"`
}

type Tested2 struct {
	Field1 int
}

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
		{"unknwon", args{reflect.TypeOf(Tested1{}), "none"}, -1},
		{"valid", args{reflect.TypeOf(Tested1{}), "field1"}, 0},
		{"not protobuf", args{reflect.TypeOf(Tested2{}), "field1"}, -1},
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
		{"valid", args{Tested1{42, 35}, 0}, 42},
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
		{"valid", args{Tested1{42, 5}, "field1"}, 42},
		{"invalid property", args{Tested1{42, 35}, "none"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valueOf(tt.args.message, tt.args.property); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valueOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
