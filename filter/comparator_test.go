package filter

import (
	"reflect"
	"testing"

	test "github.com/osechet/go-datastore/_proto/osechet/test"
)

func TestNewCompositeComparator(t *testing.T) {
	type args struct {
		comparators []Comparator
	}
	tests := []struct {
		name string
		args args
		want *CompositeComparator
	}{
		{"valid", args{make([]Comparator, 0)}, &CompositeComparator{make([]Comparator, 0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompositeComparator(tt.args.comparators); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompositeComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompositeComparator_Less(t *testing.T) {
	type fields struct {
		comparators []Comparator
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"no comparators", fields{make([]Comparator, 0)}, args{1, 1}, false},
		{"1 comparator - less", fields{[]Comparator{ValueComparator{Ascending}}}, args{int32(1), int32(2)}, true},
		{"1 comparator - equal", fields{[]Comparator{ValueComparator{Ascending}}}, args{int32(2), int32(2)}, false},
		{"1 comparator - greater", fields{[]Comparator{ValueComparator{Ascending}}}, args{int32(2), int32(1)}, false},
		{"2 comparators simple - true", fields{[]Comparator{NewPropertyComparator(reflect.TypeOf(test.Tested{}), "int32_value", Ascending), NewPropertyComparator(reflect.TypeOf(test.Tested{}), "int64_value", Ascending)}}, args{test.Tested{Int32Value: 1, Int64Value: 2}, test.Tested{Int32Value: 2, Int64Value: 3}}, true},
		{"2 comparators simple - false", fields{[]Comparator{NewPropertyComparator(reflect.TypeOf(test.Tested{}), "int32_value", Ascending), NewPropertyComparator(reflect.TypeOf(test.Tested{}), "int64_value", Ascending)}}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 1, Int64Value: 3}}, false},
		{"2 comparators composition - true", fields{[]Comparator{NewPropertyComparator(reflect.TypeOf(test.Tested{}), "int32_value", Ascending), NewPropertyComparator(reflect.TypeOf(test.Tested{}), "int64_value", Ascending)}}, args{test.Tested{Int32Value: 1, Int64Value: 2}, test.Tested{Int32Value: 1, Int64Value: 3}}, true},
		{"2 comparators composition - false", fields{[]Comparator{NewPropertyComparator(reflect.TypeOf(test.Tested{}), "int32_value", Ascending), NewPropertyComparator(reflect.TypeOf(test.Tested{}), "int64_value", Ascending)}}, args{test.Tested{Int32Value: 1, Int64Value: 3}, test.Tested{Int32Value: 1, Int64Value: 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CompositeComparator{
				comparators: tt.fields.comparators,
			}
			if got := c.Less(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CompositeComparator.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueComparator_Less(t *testing.T) {
	type fields struct {
		direction Direction
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"ascending - less", fields{Ascending}, args{int32(1), int32(2)}, true},
		{"ascending - equal", fields{Ascending}, args{int32(2), int32(2)}, false},
		{"ascending - greater", fields{Ascending}, args{int32(2), int32(1)}, false},
		{"descending - less", fields{Descending}, args{int32(2), int32(1)}, true},
		{"descending - equal", fields{Descending}, args{int32(2), int32(2)}, false},
		{"descending - greater", fields{Descending}, args{int32(1), int32(2)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ValueComparator{
				direction: tt.fields.direction,
			}
			if got := c.Less(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ValueComparator.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueComparator_Equals(t *testing.T) {
	type fields struct {
		direction Direction
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"ascending - less", fields{Ascending}, args{int32(1), int32(2)}, false},
		{"ascending - equals", fields{Ascending}, args{int32(2), int32(2)}, true},
		{"ascending - greater", fields{Ascending}, args{int32(2), int32(1)}, false},
		{"descending - less", fields{Descending}, args{int32(2), int32(1)}, false},
		{"descending - equal", fields{Descending}, args{int32(2), int32(2)}, true},
		{"descending - greater", fields{Descending}, args{int32(1), int32(2)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ValueComparator{
				direction: tt.fields.direction,
			}
			if got := c.Equals(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ValueComparator.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueComparator_Greater(t *testing.T) {
	type fields struct {
		direction Direction
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"ascending - less", fields{Ascending}, args{int32(1), int32(2)}, false},
		{"ascending - equal", fields{Ascending}, args{int32(2), int32(2)}, false},
		{"ascending - greater", fields{Ascending}, args{int32(2), int32(1)}, true},
		{"descending - less", fields{Descending}, args{int32(2), int32(1)}, false},
		{"descending - equal", fields{Descending}, args{int32(2), int32(2)}, false},
		{"descending - greater", fields{Descending}, args{int32(1), int32(2)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ValueComparator{
				direction: tt.fields.direction,
			}
			if got := c.Greater(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ValueComparator.Greater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPropertyComparator(t *testing.T) {
	type args struct {
		t         reflect.Type
		property  string
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want *PropertyComparator
	}{
		{"invalid property", args{reflect.TypeOf(test.Tested{}), "none", Ascending}, nil},
		{"valid property", args{reflect.TypeOf(test.Tested{}), "int32_value", Ascending}, &PropertyComparator{2, Ascending}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPropertyComparator(tt.args.t, tt.args.property, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPropertyComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertyComparator_Less(t *testing.T) {
	type fields struct {
		fieldIndex int
		direction  Direction
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"ascending - less", fields{2, Ascending}, args{test.Tested{Int32Value: 1, Int64Value: 1}, test.Tested{Int32Value: 2, Int64Value: 2}}, true},
		{"ascending - equal", fields{2, Ascending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 2, Int64Value: 2}}, false},
		{"ascending - greater", fields{2, Ascending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 1, Int64Value: 1}}, false},
		{"descending - less", fields{2, Descending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 1, Int64Value: 1}}, true},
		{"descending - equal", fields{2, Descending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 2, Int64Value: 2}}, false},
		{"descending - greater", fields{2, Descending}, args{test.Tested{Int32Value: 1, Int64Value: 1}, test.Tested{Int32Value: 2, Int64Value: 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := PropertyComparator{
				fieldIndex: tt.fields.fieldIndex,
				direction:  tt.fields.direction,
			}
			if got := c.Less(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("PropertyComparator.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertyComparator_Equals(t *testing.T) {
	type fields struct {
		fieldIndex int
		direction  Direction
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"ascending - less", fields{2, Ascending}, args{test.Tested{Int32Value: 1, Int64Value: 1}, test.Tested{Int32Value: 2, Int64Value: 2}}, false},
		{"ascending - equal", fields{2, Ascending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 2, Int64Value: 2}}, true},
		{"ascending - greater", fields{2, Ascending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 1, Int64Value: 1}}, false},
		{"descending - less", fields{2, Descending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 1, Int64Value: 1}}, false},
		{"descending - equal", fields{2, Descending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 2, Int64Value: 2}}, true},
		{"descending - greater", fields{2, Descending}, args{test.Tested{Int32Value: 1, Int64Value: 1}, test.Tested{Int32Value: 2, Int64Value: 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := PropertyComparator{
				fieldIndex: tt.fields.fieldIndex,
				direction:  tt.fields.direction,
			}
			if got := c.Equals(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("PropertyComparator.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertyComparator_Greater(t *testing.T) {
	type fields struct {
		fieldIndex int
		direction  Direction
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"ascending - less", fields{2, Ascending}, args{test.Tested{Int32Value: 1, Int64Value: 1}, test.Tested{Int32Value: 2, Int64Value: 2}}, false},
		{"ascending - equal", fields{2, Ascending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 2, Int64Value: 2}}, false},
		{"ascending - greater", fields{2, Ascending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 1, Int64Value: 1}}, true},
		{"descending - less", fields{2, Descending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 1, Int64Value: 1}}, false},
		{"descending - equal", fields{2, Descending}, args{test.Tested{Int32Value: 2, Int64Value: 2}, test.Tested{Int32Value: 2, Int64Value: 2}}, false},
		{"descending - greater", fields{2, Descending}, args{test.Tested{Int32Value: 1, Int64Value: 1}, test.Tested{Int32Value: 2, Int64Value: 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := PropertyComparator{
				fieldIndex: tt.fields.fieldIndex,
				direction:  tt.fields.direction,
			}
			if got := c.Greater(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("PropertyComparator.Greater() = %v, want %v", got, tt.want)
			}
		})
	}
}
