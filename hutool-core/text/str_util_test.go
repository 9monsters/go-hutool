package text

import (
	"reflect"
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

func TestTrims(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test empty arr", args{[]string{}}, []string{}},
		{"test empty", args{[]string{"  11 "}}, []string{"11"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trims(tt.args.str)
			t.Logf("trim res: %s", got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trims() = %v, want %v", got, tt.want)
			}
		})
	}
}
