package filter

import (
	"reflect"
	"testing"

	test "github.com/osechet/go-datastore/_proto/osechet/test"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

func TestNewComparator(t *testing.T) {
	type args struct {
		t     reflect.Type
		order *datastore.PropertyOrder
	}
	tests := []struct {
		name string
		args args
		want Comparator
	}{
		{"invalid order", args{reflect.TypeOf(test.Tested{}), nil}, nil},
		{"invalid property", args{reflect.TypeOf(test.Tested{}), &datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "none"}, Direction: datastore.PropertyOrder_ASCENDING}}, (*PropertyComparator)(nil)},
		{"ascending", args{reflect.TypeOf(test.Tested{}), &datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "int32_value"}, Direction: datastore.PropertyOrder_ASCENDING}}, &PropertyComparator{2, Ascending}},
		{"descending", args{reflect.TypeOf(test.Tested{}), &datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "int32_value"}, Direction: datastore.PropertyOrder_DESCENDING}}, &PropertyComparator{2, Descending}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComparator(tt.args.t, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
