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
		{"float32 - less", args{float32(1), float32(2)}, true},
		{"float32 - equal", args{float32(2), float32(2)}, false},
		{"float32 - greater", args{float32(2), float32(1)}, false},
		{"float64 - less", args{1.0, 2.0}, true},
		{"float64 - equal", args{2.0, 2.0}, false},
		{"float64 - greater", args{2.0, 1.0}, false},
		{"int32 - less", args{int32(1), int32(2)}, true},
		{"int32 - equal", args{int32(2), int32(2)}, false},
		{"int32 - greater", args{int32(2), int32(1)}, false},
		{"int64 - less", args{int64(1), int64(2)}, true},
		{"int64 - equal", args{int64(2), int64(2)}, false},
		{"int64 - greater", args{int64(2), int64(1)}, false},
		{"string - less", args{"abc", "def"}, true},
		{"string - equal", args{"abc", "abc"}, false},
		{"string - greater", args{"def", "abc"}, false},
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
		{"float32 - less", args{float32(1), float32(2)}, false},
		{"float32 - equal", args{float32(2), float32(2)}, true},
		{"float32 - greater", args{float32(2), float32(1)}, false},
		{"float64 - less", args{1.0, 2.0}, false},
		{"float64 - equal", args{2.0, 2.0}, true},
		{"float64 - greater", args{2.0, 1.0}, false},
		{"int32 - less", args{int32(1), int32(2)}, false},
		{"int32 - equal", args{int32(2), int32(2)}, true},
		{"int32 - greater", args{int32(2), int32(1)}, false},
		{"int64 - less", args{int64(1), int64(2)}, false},
		{"int64 - equal", args{int64(2), int64(2)}, true},
		{"int64 - greater", args{int64(2), int64(1)}, false},
		{"string - less", args{"abc", "def"}, false},
		{"string - equal", args{"abc", "abc"}, true},
		{"string - greater", args{"def", "abc"}, false},
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
		{"float32 - less", args{float32(1), float32(2)}, false},
		{"float32 - equal", args{float32(2), float32(2)}, false},
		{"float32 - greater", args{float32(2), float32(1)}, true},
		{"float64 - less", args{1.0, 2.0}, false},
		{"float64 - equal", args{2.0, 2.0}, false},
		{"float64 - greater", args{2.0, 1.0}, true},
		{"int32 - less", args{int32(1), int32(2)}, false},
		{"int32 - equal", args{int32(2), int32(2)}, false},
		{"int32 - greater", args{int32(2), int32(1)}, true},
		{"int64 - less", args{int64(1), int64(2)}, false},
		{"int64 - equal", args{int64(2), int64(2)}, false},
		{"int64 - greater", args{int64(2), int64(1)}, true},
		{"string - less", args{"abc", "def"}, false},
		{"string - equal", args{"abc", "abc"}, false},
		{"string - greater", args{"def", "abc"}, true},
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
