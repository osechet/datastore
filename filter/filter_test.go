package filter

import (
	"os"
	"testing"

	"github.com/golang/protobuf/descriptor"
	test "github.com/osechet/datastore/_proto/osechet/test"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

func TestMatch(t *testing.T) {
	type args struct {
		filter  datastore.Filter
		message descriptor.Message
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"no filter", args{datastore.Filter{}, &test.Tested{}}, false},
		{"property filter - no property", args{datastore.Filter{FilterType: &datastore.Filter_PropertyFilter{}}, &test.Tested{}}, false},
		{"property filter - no value", args{datastore.Filter{FilterType: &datastore.Filter_PropertyFilter{PropertyFilter: &datastore.PropertyFilter{Property: &datastore.PropertyReference{Name: "int32_value"}, Op: datastore.PropertyFilter_EQUAL}}}, &test.Tested{Int32Value: 42}}, false},
		{"property filter", args{datastore.Filter{FilterType: &datastore.Filter_PropertyFilter{PropertyFilter: &datastore.PropertyFilter{Property: &datastore.PropertyReference{Name: "int32_value"}, Op: datastore.PropertyFilter_EQUAL, Value: makeIntegerValue(42)}}}, &test.Tested{Int32Value: 42}}, true},
		{"composite filter - no filters", args{datastore.Filter{FilterType: &datastore.Filter_CompositeFilter{}}, &test.Tested{}}, false},
		{"composite filter - 1 filter", args{datastore.Filter{FilterType: &datastore.Filter_CompositeFilter{CompositeFilter: &datastore.CompositeFilter{Filters: []*datastore.Filter{&datastore.Filter{FilterType: &datastore.Filter_PropertyFilter{PropertyFilter: &datastore.PropertyFilter{Property: &datastore.PropertyReference{Name: "int32_value"}, Op: datastore.PropertyFilter_EQUAL, Value: makeIntegerValue(42)}}}}}}}, &test.Tested{Int32Value: 42}}, true},
		{"composite filter - 2 filters - no match", args{datastore.Filter{FilterType: &datastore.Filter_CompositeFilter{CompositeFilter: &datastore.CompositeFilter{Filters: []*datastore.Filter{&datastore.Filter{FilterType: &datastore.Filter_PropertyFilter{PropertyFilter: &datastore.PropertyFilter{Property: &datastore.PropertyReference{Name: "int32_value"}, Op: datastore.PropertyFilter_EQUAL, Value: makeIntegerValue(42)}}}, &datastore.Filter{FilterType: &datastore.Filter_PropertyFilter{PropertyFilter: &datastore.PropertyFilter{Property: &datastore.PropertyReference{Name: "int64_value"}, Op: datastore.PropertyFilter_EQUAL, Value: makeIntegerValue(35)}}}}}}}, &test.Tested{Int32Value: 35, Int64Value: 35}}, false},
		{"composite filter - 2 filters", args{datastore.Filter{FilterType: &datastore.Filter_CompositeFilter{CompositeFilter: &datastore.CompositeFilter{Filters: []*datastore.Filter{&datastore.Filter{FilterType: &datastore.Filter_PropertyFilter{PropertyFilter: &datastore.PropertyFilter{Property: &datastore.PropertyReference{Name: "int32_value"}, Op: datastore.PropertyFilter_EQUAL, Value: makeIntegerValue(42)}}}, &datastore.Filter{FilterType: &datastore.Filter_PropertyFilter{PropertyFilter: &datastore.PropertyFilter{Property: &datastore.PropertyReference{Name: "int64_value"}, Op: datastore.PropertyFilter_EQUAL, Value: makeIntegerValue(35)}}}}}}}, &test.Tested{Int32Value: 42, Int64Value: 35}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.args.filter, tt.args.message); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compare(t *testing.T) {
	type args struct {
		message   descriptor.Message
		property  string
		op        datastore.PropertyFilter_Operator
		reference datastore.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"lt", args{&test.Tested{DoubleValue: 35}, "double_value", datastore.PropertyFilter_LESS_THAN, *makeDoubleValue(42)}, true},
		{"lte", args{&test.Tested{DoubleValue: 35}, "double_value", datastore.PropertyFilter_LESS_THAN_OR_EQUAL, *makeDoubleValue(42)}, true},
		{"gt", args{&test.Tested{DoubleValue: 35}, "double_value", datastore.PropertyFilter_GREATER_THAN, *makeDoubleValue(42)}, false},
		{"gte", args{&test.Tested{DoubleValue: 35}, "double_value", datastore.PropertyFilter_GREATER_THAN_OR_EQUAL, *makeDoubleValue(42)}, false},
		{"equal", args{&test.Tested{DoubleValue: 35}, "double_value", datastore.PropertyFilter_EQUAL, *makeDoubleValue(42)}, false},
		{"hasAncestor", args{&test.Tested{DoubleValue: 35}, "double_value", datastore.PropertyFilter_HAS_ANCESTOR, *makeDoubleValue(42)}, false},
		{"unspecified", args{&test.Tested{DoubleValue: 35}, "double_value", datastore.PropertyFilter_OPERATOR_UNSPECIFIED, *makeDoubleValue(42)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(tt.args.message, tt.args.property, tt.args.op, tt.args.reference); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lt(t *testing.T) {
	type args struct {
		value     interface{}
		reference datastore.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"float32 - less", args{float32(1), *makeDoubleValue(2)}, true},
		{"float32 - equal", args{float32(2), *makeDoubleValue(2)}, false},
		{"float32 - greater", args{float32(2), *makeDoubleValue(1)}, false},
		{"float64 - less", args{1.0, *makeDoubleValue(2)}, true},
		{"float64 - equal", args{2.0, *makeDoubleValue(2)}, false},
		{"float64 - greater", args{2.0, *makeDoubleValue(1)}, false},
		{"int32 - less", args{int32(1), *makeIntegerValue(2)}, true},
		{"int32 - equal", args{int32(2), *makeIntegerValue(2)}, false},
		{"int32 - greater", args{int32(2), *makeIntegerValue(1)}, false},
		{"int64 - less", args{int64(1), *makeIntegerValue(2)}, true},
		{"int64 - equal", args{int64(2), *makeIntegerValue(2)}, false},
		{"int64 - greater", args{int64(2), *makeIntegerValue(1)}, false},
		{"string - less", args{"abc", *makeStringValue("def")}, true},
		{"string - equal", args{"abc", *makeStringValue("abc")}, false},
		{"string - greater", args{"def", *makeStringValue("abc")}, false},
		{"other", args{os.File{}, datastore.Value{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lt(tt.args.value, tt.args.reference); got != tt.want {
				t.Errorf("lt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lte(t *testing.T) {
	type args struct {
		value     interface{}
		reference datastore.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"float32 - less", args{float32(1), *makeDoubleValue(2)}, true},
		{"float32 - equal", args{float32(2), *makeDoubleValue(2)}, true},
		{"float32 - greater", args{float32(2), *makeDoubleValue(1)}, false},
		{"float64 - less", args{1.0, *makeDoubleValue(2)}, true},
		{"float64 - equal", args{2.0, *makeDoubleValue(2)}, true},
		{"float64 - greater", args{2.0, *makeDoubleValue(1)}, false},
		{"int32 - less", args{int32(1), *makeIntegerValue(2)}, true},
		{"int32 - equal", args{int32(2), *makeIntegerValue(2)}, true},
		{"int32 - greater", args{int32(2), *makeIntegerValue(1)}, false},
		{"int64 - less", args{int64(1), *makeIntegerValue(2)}, true},
		{"int64 - equal", args{int64(2), *makeIntegerValue(2)}, true},
		{"int64 - greater", args{int64(2), *makeIntegerValue(1)}, false},
		{"string - less", args{"abc", *makeStringValue("def")}, true},
		{"string - equal", args{"abc", *makeStringValue("abc")}, true},
		{"string - greater", args{"def", *makeStringValue("abc")}, false},
		{"other", args{os.File{}, datastore.Value{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lte(tt.args.value, tt.args.reference); got != tt.want {
				t.Errorf("lte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gt(t *testing.T) {
	type args struct {
		value     interface{}
		reference datastore.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"float32 - less", args{float32(1), *makeDoubleValue(2)}, false},
		{"float32 - equal", args{float32(2), *makeDoubleValue(2)}, false},
		{"float32 - greater", args{float32(2), *makeDoubleValue(1)}, true},
		{"float64 - less", args{1.0, *makeDoubleValue(2)}, false},
		{"float64 - equal", args{2.0, *makeDoubleValue(2)}, false},
		{"float64 - greater", args{2.0, *makeDoubleValue(1)}, true},
		{"int32 - less", args{int32(1), *makeIntegerValue(2)}, false},
		{"int32 - equal", args{int32(2), *makeIntegerValue(2)}, false},
		{"int32 - greater", args{int32(2), *makeIntegerValue(1)}, true},
		{"int64 - less", args{int64(1), *makeIntegerValue(2)}, false},
		{"int64 - equal", args{int64(2), *makeIntegerValue(2)}, false},
		{"int64 - greater", args{int64(2), *makeIntegerValue(1)}, true},
		{"string - less", args{"abc", *makeStringValue("def")}, false},
		{"string - equal", args{"abc", *makeStringValue("abc")}, false},
		{"string - greater", args{"def", *makeStringValue("abc")}, true},
		{"other", args{os.File{}, datastore.Value{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gt(tt.args.value, tt.args.reference); got != tt.want {
				t.Errorf("gt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gte(t *testing.T) {
	type args struct {
		value     interface{}
		reference datastore.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"float32 - less", args{float32(1), *makeDoubleValue(2)}, false},
		{"float32 - equal", args{float32(2), *makeDoubleValue(2)}, true},
		{"float32 - greater", args{float32(2), *makeDoubleValue(1)}, true},
		{"float64 - less", args{1.0, *makeDoubleValue(2)}, false},
		{"float64 - equal", args{2.0, *makeDoubleValue(2)}, true},
		{"float64 - greater", args{2.0, *makeDoubleValue(1)}, true},
		{"int32 - less", args{int32(1), *makeIntegerValue(2)}, false},
		{"int32 - equal", args{int32(2), *makeIntegerValue(2)}, true},
		{"int32 - greater", args{int32(2), *makeIntegerValue(1)}, true},
		{"int64 - less", args{int64(1), *makeIntegerValue(2)}, false},
		{"int64 - equal", args{int64(2), *makeIntegerValue(2)}, true},
		{"int64 - greater", args{int64(2), *makeIntegerValue(1)}, true},
		{"string - less", args{"abc", *makeStringValue("def")}, false},
		{"string - equal", args{"abc", *makeStringValue("abc")}, true},
		{"string - greater", args{"def", *makeStringValue("abc")}, true},
		{"other", args{os.File{}, datastore.Value{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gte(tt.args.value, tt.args.reference); got != tt.want {
				t.Errorf("gte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equal(t *testing.T) {
	type args struct {
		value     interface{}
		reference datastore.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"float32 - less", args{float32(1), *makeDoubleValue(2)}, false},
		{"float32 - equal", args{float32(2), *makeDoubleValue(2)}, true},
		{"float32 - greater", args{float32(2), *makeDoubleValue(1)}, false},
		{"float64 - less", args{1.0, *makeDoubleValue(2)}, false},
		{"float64 - equal", args{2.0, *makeDoubleValue(2)}, true},
		{"float64 - greater", args{2.0, *makeDoubleValue(1)}, false},
		{"int32 - less", args{int32(1), *makeIntegerValue(2)}, false},
		{"int32 - equal", args{int32(2), *makeIntegerValue(2)}, true},
		{"int32 - greater", args{int32(2), *makeIntegerValue(1)}, false},
		{"int64 - less", args{int64(1), *makeIntegerValue(2)}, false},
		{"int64 - equal", args{int64(2), *makeIntegerValue(2)}, true},
		{"int64 - greater", args{int64(2), *makeIntegerValue(1)}, false},
		{"string - less", args{"abc", *makeStringValue("def")}, false},
		{"string - equal", args{"abc", *makeStringValue("abc")}, true},
		{"string - greater", args{"def", *makeStringValue("abc")}, false},
		{"other", args{os.File{}, datastore.Value{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equal(tt.args.value, tt.args.reference); got != tt.want {
				t.Errorf("equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasAncestor(t *testing.T) {
	type args struct {
		value     interface{}
		reference datastore.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not implemented", args{float32(1), *makeDoubleValue(2)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAncestor(tt.args.value, tt.args.reference); got != tt.want {
				t.Errorf("hasAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
