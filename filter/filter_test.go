package filter

import (
	"testing"

	"github.com/golang/protobuf/descriptor"
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAncestor(tt.args.value, tt.args.reference); got != tt.want {
				t.Errorf("hasAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
