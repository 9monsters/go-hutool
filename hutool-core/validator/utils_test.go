package validator

import (
	"reflect"
	"testing"
)

func TestBlackList(t *testing.T) {
	type args struct {
		str   string
		chars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlackList(tt.args.str, tt.args.chars); got != tt.want {
				t.Errorf("BlackList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelCaseToUnderscore(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCaseToUnderscore(tt.args.str); got != tt.want {
				t.Errorf("CamelCaseToUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		str       string
		substring string
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
			if got := Contains(tt.args.str, tt.args.substring); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLine(t *testing.T) {
	type args struct {
		s     string
		index int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLine(tt.args.s, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLines(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLines(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftTrim(t *testing.T) {
	type args struct {
		str   string
		chars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftTrim(tt.args.str, tt.args.chars); got != tt.want {
				t.Errorf("LeftTrim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	type args struct {
		str     string
		pattern string
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
			if got := Matches(tt.args.str, tt.args.pattern); got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeEmail(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeEmail(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadBoth(t *testing.T) {
	type args struct {
		str    string
		padStr string
		padLen int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadBoth(tt.args.str, tt.args.padStr, tt.args.padLen); got != tt.want {
				t.Errorf("PadBoth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLeft(t *testing.T) {
	type args struct {
		str    string
		padStr string
		padLen int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadLeft(tt.args.str, tt.args.padStr, tt.args.padLen); got != tt.want {
				t.Errorf("PadLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	type args struct {
		str    string
		padStr string
		padLen int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRight(tt.args.str, tt.args.padStr, tt.args.padLen); got != tt.want {
				t.Errorf("PadRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveTags(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveTags(tt.args.s); got != tt.want {
				t.Errorf("RemoveTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplacePattern(t *testing.T) {
	type args struct {
		str     string
		pattern string
		replace string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplacePattern(tt.args.str, tt.args.pattern, tt.args.replace); got != tt.want {
				t.Errorf("ReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightTrim(t *testing.T) {
	type args struct {
		str   string
		chars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RightTrim(tt.args.str, tt.args.chars); got != tt.want {
				t.Errorf("RightTrim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeFileName(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeFileName(tt.args.str); got != tt.want {
				t.Errorf("SafeFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripLow(t *testing.T) {
	type args struct {
		str          string
		keepNewLines bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripLow(tt.args.str, tt.args.keepNewLines); got != tt.want {
				t.Errorf("StripLow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		str   string
		chars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.str, tt.args.chars); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		str    string
		length int
		ending string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.args.str, tt.args.length, tt.args.ending); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncatingErrorf(t *testing.T) {
	type args struct {
		str  string
		args []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TruncatingErrorf(tt.args.str, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("TruncatingErrorf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnderscoreToCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnderscoreToCamelCase(tt.args.s); got != tt.want {
				t.Errorf("UnderscoreToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhiteList(t *testing.T) {
	type args struct {
		str   string
		chars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhiteList(tt.args.str, tt.args.chars); got != tt.want {
				t.Errorf("WhiteList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addSegment(t *testing.T) {
	type args struct {
		inrune  []rune
		segment []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSegment(tt.args.inrune, tt.args.segment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildPadStr(t *testing.T) {
	type args struct {
		str      string
		padStr   string
		padLen   int
		padLeft  bool
		padRight bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildPadStr(tt.args.str, tt.args.padStr, tt.args.padLen, tt.args.padLeft, tt.args.padRight); got != tt.want {
				t.Errorf("buildPadStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
