package net

import "testing"

func TestLongToIPv41(t *testing.T) {
	type args struct {
		longIP int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test local", args{2130706433}, "127.0.0.1"},
		{"test 192", args{3232238810}, "192.168.12.218"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongToIPv4(tt.args.longIP); got != tt.want {
				t.Errorf("LongToIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}
