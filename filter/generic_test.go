package filter

import (
	"os"
	"testing"
)

func TestLess(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"float32_1", args{float32(1), float32(2)}, true},
		{"float32_2", args{float32(2), float32(1)}, false},
		{"float64_1", args{1.0, 2.0}, true},
		{"float64_2", args{2.0, 1.0}, false},
		{"int_1", args{1, 2}, true},
		{"int_2", args{2, 1}, false},
		{"string_1", args{"abc", "def"}, true},
		{"string_2", args{"def", "abc"}, false},
		{"other", args{os.File{}, os.File{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Less(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"float32_1", args{float32(1), float32(1)}, true},
		{"float32_2", args{float32(2), float32(1)}, false},
		{"float64_1", args{1.0, 1.0}, true},
		{"float64_2", args{2.0, 1.0}, false},
		{"int_1", args{1, 1}, true},
		{"int_2", args{2, 1}, false},
		{"string_1", args{"abc", "abc"}, true},
		{"string_2", args{"def", "abc"}, false},
		{"other", args{os.File{}, os.File{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreater(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"float32_1", args{float32(1), float32(2)}, false},
		{"float32_2", args{float32(2), float32(1)}, true},
		{"float64_1", args{1.0, 2.0}, false},
		{"float64_2", args{2.0, 1.0}, true},
		{"int_1", args{1, 2}, false},
		{"int_2", args{2, 1}, true},
		{"string_1", args{"abc", "def"}, false},
		{"string_2", args{"def", "abc"}, true},
		{"other", args{os.File{}, os.File{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greater(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Greater() = %v, want %v", got, tt.want)
			}
		})
	}
}
