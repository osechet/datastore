package query

import (
	"reflect"
	"testing"

	"github.com/golang/protobuf/descriptor"
	test "github.com/osechet/go-datastore/_proto/osechet/test"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

func TestApply(t *testing.T) {
	data := []descriptor.Message{
		&test.Tested{Int32Value: 55, Int64Value: 35},
		&test.Tested{Int32Value: 68, Int64Value: 43},
		&test.Tested{Int32Value: 42, Int64Value: 56},
		&test.Tested{Int32Value: 47, Int64Value: 43},
	}
	type args struct {
		storage Storage
		query   datastore.Query
		t       reflect.Type
		results ResultSet
	}
	tests := []struct {
		name string
		args args
		want []descriptor.Message
	}{
		{
			"empty storage",
			args{
				SliceStorage{},
				datastore.Query{},
				reflect.TypeOf(test.Tested{}),
				NewSliceResultSet(),
			},
			[]descriptor.Message{},
		},
		{
			"no order - no filter",
			args{
				SliceStorage{data},
				datastore.Query{},
				reflect.TypeOf(test.Tested{}),
				NewSliceResultSet(),
			},
			[]descriptor.Message{
				&test.Tested{Int32Value: 55, Int64Value: 35},
				&test.Tested{Int32Value: 68, Int64Value: 43},
				&test.Tested{Int32Value: 42, Int64Value: 56},
				&test.Tested{Int32Value: 47, Int64Value: 43},
			},
		},
		{
			"order on ascending int32_value - no filter",
			args{
				SliceStorage{data},
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "int32_value"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
					},
				},
				reflect.TypeOf(test.Tested{}),
				NewSliceResultSet(),
			},
			[]descriptor.Message{
				&test.Tested{Int32Value: 42, Int64Value: 56},
				&test.Tested{Int32Value: 47, Int64Value: 43},
				&test.Tested{Int32Value: 55, Int64Value: 35},
				&test.Tested{Int32Value: 68, Int64Value: 43},
			},
		},
		{
			"no order - filter on int64_value",
			args{
				SliceStorage{data},
				datastore.Query{
					Filter: &datastore.Filter{
						FilterType: &datastore.Filter_CompositeFilter{
							CompositeFilter: &datastore.CompositeFilter{
								Filters: []*datastore.Filter{
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "int64_value"},
												Op:       datastore.PropertyFilter_GREATER_THAN_OR_EQUAL,
												Value: &datastore.Value{
													ValueType: &datastore.Value_IntegerValue{IntegerValue: 43},
												},
											},
										},
									},
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "int64_value"},
												Op:       datastore.PropertyFilter_LESS_THAN,
												Value: &datastore.Value{
													ValueType: &datastore.Value_IntegerValue{IntegerValue: 50},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				reflect.TypeOf(test.Tested{}),
				NewSliceResultSet(),
			},
			[]descriptor.Message{
				&test.Tested{Int32Value: 68, Int64Value: 43},
				&test.Tested{Int32Value: 47, Int64Value: 43},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Apply(tt.args.storage, tt.args.query, tt.args.t, tt.args.results)
			if got := tt.args.results.(*SliceResultSet).Items; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
