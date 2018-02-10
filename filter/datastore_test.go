package filter

import (
	"reflect"
	"testing"

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
		{"invalid order", args{reflect.TypeOf(Tested1{}), nil}, nil},
		{"invalid property", args{reflect.TypeOf(Tested1{}), &datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "none"}, Direction: datastore.PropertyOrder_ASCENDING}}, (*PropertyComparator)(nil)},
		{"ascending", args{reflect.TypeOf(Tested1{}), &datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "field1"}, Direction: datastore.PropertyOrder_ASCENDING}}, &PropertyComparator{0, Ascending}},
		{"descending", args{reflect.TypeOf(Tested1{}), &datastore.PropertyOrder{Property: &datastore.PropertyReference{Name: "field1"}, Direction: datastore.PropertyOrder_DESCENDING}}, &PropertyComparator{0, Descending}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComparator(tt.args.t, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
