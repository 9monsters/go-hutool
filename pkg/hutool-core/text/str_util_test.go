package text

import (
	"testing"
)

func TestIsBlank(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test nil", args: args{}, want: true},
		{name: "test \"\"", args: args{""}, want: true},
		{name: "test \"  \"", args: args{"  "}, want: true},
		{name: "test \"123\"", args: args{"123"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.args.str); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test nil", args: args{}, want: true},
		{name: "test \"\"", args: args{""}, want: true},
		{name: "test \"  \"", args: args{"  "}, want: false},
		{name: "test \"123\"", args: args{"123"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.str); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
