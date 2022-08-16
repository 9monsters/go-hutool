package net

import (
	"net"
	"reflect"
	"sort"
	"testing"
)

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

func TestIsIpV4(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"trueCase", args{"127.0.0.1"}, true},
		{"falseCase", args{"::7f00:0001"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIpV4(tt.args.ip); got != tt.want {
				t.Errorf("IsIpV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIpV6(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"trueCase", args{"127.0.0.1"}, false},
		{"falseCase", args{"::7f00:0001"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIpV6(tt.args.ip); got != tt.want {
				t.Errorf("IsIpV6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllIp(t *testing.T) {
	type args struct {
		predicate  func(ip net.IP) bool
		ifaceNames []string
	}
	tests := []struct {
		name    string
		args    args
		want    []net.IP
		wantErr bool
	}{
		{"test1", args{
			predicate:  nil,
			ifaceNames: nil,
		}, []net.IP{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListAllIp(tt.args.predicate, tt.args.ifaceNames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllIp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllIpV4(t *testing.T) {
	type args struct {
		ifaceNames []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test1", args{ifaceNames: nil}, []string{"192.168.16.153"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListAllIpV4(tt.args.ifaceNames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllIpV4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllIpV4() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllIpV6(t *testing.T) {
	type args struct {
		ifaceNames []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test1", args{ifaceNames: nil}, make([]string, 0), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListAllIpV6(tt.args.ifaceNames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllIpV6() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllIpV6() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListIfaceNames(t *testing.T) {
	tests := []struct {
		name      string
		wantNames []string
	}{
		{"test1", []string{"en0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNames := ListIfaceNames()
			if !anyIn(tt.wantNames, gotNames) {
				t.Errorf("ListIfaceNames() = %v, want %v", gotNames, tt.wantNames)
			}
		})
	}
}

func TestOutbound(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test1", "192.168.16.153"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Outbound(); got != tt.want {
				t.Errorf("Outbound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainIP(t *testing.T) {
	type args struct {
		ifaceName []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []string
	}{
		{"test1", args{}, "192.168.16.153", []string{"192.168.16.153"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MainIP(tt.args.ifaceName...)
			if got != tt.want {
				t.Errorf("MainIP() got = %v, want %v", got, tt.want)
			}
			if !anyIn(got1, tt.want1) {
				t.Errorf("MainIP() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func in(target string, strArr []string) bool {
	sort.Strings(strArr)
	index := sort.SearchStrings(strArr, target)

	if index < len(strArr) && strArr[index] == target {
		return true
	}
	return false
}

func anyIn(targetArr []string, strArr []string) bool {
	if targetArr == nil || len(targetArr) == 0 {
		return false
	}
	if strArr == nil || len(strArr) == 0 {
		return false
	}

	for _, target := range targetArr {
		sort.Strings(strArr)
		index := sort.SearchStrings(strArr, target)
		if index < len(strArr) && strArr[index] == target {
			return true
		}
	}
	return false
}
