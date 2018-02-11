package filter

import (
	"reflect"
	"testing"

	test "github.com/osechet/go-datastore/_proto/osechet/test"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

func TestNewComparator(t *testing.T) {
	type args struct {
		order *datastore.PropertyOrder
		t     reflect.Type
	}
	tests := []struct {
		name string
		args args
		want *PropertyComparator
	}{
		{"invalid order", args{nil, reflect.TypeOf(test.Tested{})}, nil},
		{"invalid property", args{&datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "none"}, Direction: datastore.PropertyOrder_ASCENDING}, reflect.TypeOf(test.Tested{})}, (*PropertyComparator)(nil)},
		{"ascending", args{&datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "int32_value"}, Direction: datastore.PropertyOrder_ASCENDING}, reflect.TypeOf(test.Tested{})}, &PropertyComparator{2, Ascending}},
		{"descending", args{&datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "int32_value"}, Direction: datastore.PropertyOrder_DESCENDING}, reflect.TypeOf(test.Tested{})}, &PropertyComparator{2, Descending}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComparator(tt.args.order, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
