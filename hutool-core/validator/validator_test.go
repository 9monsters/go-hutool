package validator

import (
	"reflect"
	"regexp"
	"testing"
)

func TestByteLength(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := ByteLength(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("ByteLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorByField(t *testing.T) {
	type args struct {
		e     error
		field string
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
			if got := ErrorByField(tt.args.e, tt.args.field); got != tt.want {
				t.Errorf("ErrorByField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorsByField(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorsByField(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorsByField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasLowerCase(t *testing.T) {
	type args struct {
		str string
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
			if got := HasLowerCase(tt.args.str); got != tt.want {
				t.Errorf("HasLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasUpperCase(t *testing.T) {
	type args struct {
		str string
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
			if got := HasUpperCase(tt.args.str); got != tt.want {
				t.Errorf("HasUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasWhitespace(t *testing.T) {
	type args struct {
		str string
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
			if got := HasWhitespace(tt.args.str); got != tt.want {
				t.Errorf("HasWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasWhitespaceOnly(t *testing.T) {
	type args struct {
		str string
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
			if got := HasWhitespaceOnly(tt.args.str); got != tt.want {
				t.Errorf("HasWhitespaceOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsASCII(t *testing.T) {
	type args struct {
		str string
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
			if got := IsASCII(tt.args.str); got != tt.want {
				t.Errorf("IsASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	type args struct {
		str string
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
			if got := IsAlpha(tt.args.str); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphanumeric(t *testing.T) {
	type args struct {
		str string
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
			if got := IsAlphanumeric(tt.args.str); got != tt.want {
				t.Errorf("IsAlphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBase64(t *testing.T) {
	type args struct {
		str string
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
			if got := IsBase64(tt.args.str); got != tt.want {
				t.Errorf("IsBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsByteLength(t *testing.T) {
	type args struct {
		str string
		min int
		max int
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
			if got := IsByteLength(tt.args.str, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("IsByteLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCIDR(t *testing.T) {
	type args struct {
		str string
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
			if got := IsCIDR(tt.args.str); got != tt.want {
				t.Errorf("IsCIDR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCRC32(t *testing.T) {
	type args struct {
		str string
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
			if got := IsCRC32(tt.args.str); got != tt.want {
				t.Errorf("IsCRC32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCRC32b(t *testing.T) {
	type args struct {
		str string
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
			if got := IsCRC32b(tt.args.str); got != tt.want {
				t.Errorf("IsCRC32b() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCreditCard(t *testing.T) {
	type args struct {
		str string
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
			if got := IsCreditCard(tt.args.str); got != tt.want {
				t.Errorf("IsCreditCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDNSName(t *testing.T) {
	type args struct {
		str string
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
			if got := IsDNSName(tt.args.str); got != tt.want {
				t.Errorf("IsDNSName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDataURI(t *testing.T) {
	type args struct {
		str string
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
			if got := IsDataURI(tt.args.str); got != tt.want {
				t.Errorf("IsDataURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDialString(t *testing.T) {
	type args struct {
		str string
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
			if got := IsDialString(tt.args.str); got != tt.want {
				t.Errorf("IsDialString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDivisibleBy(t *testing.T) {
	type args struct {
		str string
		num string
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
			if got := IsDivisibleBy(tt.args.str, tt.args.num); got != tt.want {
				t.Errorf("IsDivisibleBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsE164(t *testing.T) {
	type args struct {
		str string
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
			if got := IsE164(tt.args.str); got != tt.want {
				t.Errorf("IsE164() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	type args struct {
		str string
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
			if got := IsEmail(tt.args.str); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsExistingEmail(t *testing.T) {
	type args struct {
		email string
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
			if got := IsExistingEmail(tt.args.email); got != tt.want {
				t.Errorf("IsExistingEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFilePath(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsFilePath(tt.args.str)
			if got != tt.want {
				t.Errorf("IsFilePath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsFilePath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
	type args struct {
		str string
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
			if got := IsFloat(tt.args.str); got != tt.want {
				t.Errorf("IsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFullWidth(t *testing.T) {
	type args struct {
		str string
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
			if got := IsFullWidth(tt.args.str); got != tt.want {
				t.Errorf("IsFullWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHalfWidth(t *testing.T) {
	type args struct {
		str string
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
			if got := IsHalfWidth(tt.args.str); got != tt.want {
				t.Errorf("IsHalfWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHash(t *testing.T) {
	type args struct {
		str       string
		algorithm string
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
			if got := IsHash(tt.args.str, tt.args.algorithm); got != tt.want {
				t.Errorf("IsHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHexadecimal(t *testing.T) {
	type args struct {
		str string
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
			if got := IsHexadecimal(tt.args.str); got != tt.want {
				t.Errorf("IsHexadecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHexcolor(t *testing.T) {
	type args struct {
		str string
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
			if got := IsHexcolor(tt.args.str); got != tt.want {
				t.Errorf("IsHexcolor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHost(t *testing.T) {
	type args struct {
		str string
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
			if got := IsHost(tt.args.str); got != tt.want {
				t.Errorf("IsHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIMEI(t *testing.T) {
	type args struct {
		str string
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
			if got := IsIMEI(tt.args.str); got != tt.want {
				t.Errorf("IsIMEI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIMSI(t *testing.T) {
	type args struct {
		str string
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
			if got := IsIMSI(tt.args.str); got != tt.want {
				t.Errorf("IsIMSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIP(t *testing.T) {
	type args struct {
		str string
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
			if got := IsIP(tt.args.str); got != tt.want {
				t.Errorf("IsIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv4(t *testing.T) {
	type args struct {
		str string
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
			if got := IsIPv4(tt.args.str); got != tt.want {
				t.Errorf("IsIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv6(t *testing.T) {
	type args struct {
		str string
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
			if got := IsIPv6(tt.args.str); got != tt.want {
				t.Errorf("IsIPv6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsISBN(t *testing.T) {
	type args struct {
		str     string
		version int
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
			if got := IsISBN(tt.args.str, tt.args.version); got != tt.want {
				t.Errorf("IsISBN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsISBN10(t *testing.T) {
	type args struct {
		str string
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
			if got := IsISBN10(tt.args.str); got != tt.want {
				t.Errorf("IsISBN10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsISBN13(t *testing.T) {
	type args struct {
		str string
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
			if got := IsISBN13(tt.args.str); got != tt.want {
				t.Errorf("IsISBN13() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsISO3166Alpha2(t *testing.T) {
	type args struct {
		str string
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
			if got := IsISO3166Alpha2(tt.args.str); got != tt.want {
				t.Errorf("IsISO3166Alpha2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsISO3166Alpha3(t *testing.T) {
	type args struct {
		str string
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
			if got := IsISO3166Alpha3(tt.args.str); got != tt.want {
				t.Errorf("IsISO3166Alpha3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsISO4217(t *testing.T) {
	type args struct {
		str string
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
			if got := IsISO4217(tt.args.str); got != tt.want {
				t.Errorf("IsISO4217() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsISO693Alpha2(t *testing.T) {
	type args struct {
		str string
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
			if got := IsISO693Alpha2(tt.args.str); got != tt.want {
				t.Errorf("IsISO693Alpha2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsISO693Alpha3b(t *testing.T) {
	type args struct {
		str string
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
			if got := IsISO693Alpha3b(tt.args.str); got != tt.want {
				t.Errorf("IsISO693Alpha3b() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIn(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := IsIn(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("IsIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInRaw(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := IsInRaw(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("IsInRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	type args struct {
		str string
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
			if got := IsInt(tt.args.str); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	type args struct {
		str string
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
			if got := IsJSON(tt.args.str); got != tt.want {
				t.Errorf("IsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLatitude(t *testing.T) {
	type args struct {
		str string
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
			if got := IsLatitude(tt.args.str); got != tt.want {
				t.Errorf("IsLatitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLongitude(t *testing.T) {
	type args struct {
		str string
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
			if got := IsLongitude(tt.args.str); got != tt.want {
				t.Errorf("IsLongitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLowerCase(t *testing.T) {
	type args struct {
		str string
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
			if got := IsLowerCase(tt.args.str); got != tt.want {
				t.Errorf("IsLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMAC(t *testing.T) {
	type args struct {
		str string
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
			if got := IsMAC(tt.args.str); got != tt.want {
				t.Errorf("IsMAC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMD4(t *testing.T) {
	type args struct {
		str string
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
			if got := IsMD4(tt.args.str); got != tt.want {
				t.Errorf("IsMD4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMD5(t *testing.T) {
	type args struct {
		str string
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
			if got := IsMD5(tt.args.str); got != tt.want {
				t.Errorf("IsMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMagnetURI(t *testing.T) {
	type args struct {
		str string
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
			if got := IsMagnetURI(tt.args.str); got != tt.want {
				t.Errorf("IsMagnetURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMatch(t *testing.T) {
	type args struct {
		reg     *regexp.Regexp
		content string
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
			if got := IsMatch(tt.args.reg, tt.args.content); got != tt.want {
				t.Errorf("IsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMongoID(t *testing.T) {
	type args struct {
		str string
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
			if got := IsMongoID(tt.args.str); got != tt.want {
				t.Errorf("IsMongoID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMultibyte(t *testing.T) {
	type args struct {
		str string
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
			if got := IsMultibyte(tt.args.str); got != tt.want {
				t.Errorf("IsMultibyte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotNull(t *testing.T) {
	type args struct {
		str string
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
			if got := IsNotNull(tt.args.str); got != tt.want {
				t.Errorf("IsNotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNull(t *testing.T) {
	type args struct {
		str string
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
			if got := IsNull(tt.args.str); got != tt.want {
				t.Errorf("IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	type args struct {
		str string
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
			if got := IsNumeric(tt.args.str); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPort(t *testing.T) {
	type args struct {
		str string
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
			if got := IsPort(tt.args.str); got != tt.want {
				t.Errorf("IsPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrintableASCII(t *testing.T) {
	type args struct {
		str string
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
			if got := IsPrintableASCII(tt.args.str); got != tt.want {
				t.Errorf("IsPrintableASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRFC3339(t *testing.T) {
	type args struct {
		str string
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
			if got := IsRFC3339(tt.args.str); got != tt.want {
				t.Errorf("IsRFC3339() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRFC3339WithoutZone(t *testing.T) {
	type args struct {
		str string
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
			if got := IsRFC3339WithoutZone(tt.args.str); got != tt.want {
				t.Errorf("IsRFC3339WithoutZone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRGBcolor(t *testing.T) {
	type args struct {
		str string
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
			if got := IsRGBcolor(tt.args.str); got != tt.want {
				t.Errorf("IsRGBcolor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRegex(t *testing.T) {
	type args struct {
		str string
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
			if got := IsRegex(tt.args.str); got != tt.want {
				t.Errorf("IsRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRequestURI(t *testing.T) {
	type args struct {
		rawurl string
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
			if got := IsRequestURI(tt.args.rawurl); got != tt.want {
				t.Errorf("IsRequestURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRequestURL(t *testing.T) {
	type args struct {
		rawurl string
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
			if got := IsRequestURL(tt.args.rawurl); got != tt.want {
				t.Errorf("IsRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRipeMD128(t *testing.T) {
	type args struct {
		str string
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
			if got := IsRipeMD128(tt.args.str); got != tt.want {
				t.Errorf("IsRipeMD128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRipeMD160(t *testing.T) {
	type args struct {
		str string
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
			if got := IsRipeMD160(tt.args.str); got != tt.want {
				t.Errorf("IsRipeMD160() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRsaPub(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := IsRsaPub(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("IsRsaPub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRsaPublicKey(t *testing.T) {
	type args struct {
		str    string
		keylen int
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
			if got := IsRsaPublicKey(tt.args.str, tt.args.keylen); got != tt.want {
				t.Errorf("IsRsaPublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSHA1(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSHA1(tt.args.str); got != tt.want {
				t.Errorf("IsSHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSHA256(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSHA256(tt.args.str); got != tt.want {
				t.Errorf("IsSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSHA3224(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSHA3224(tt.args.str); got != tt.want {
				t.Errorf("IsSHA3224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSHA3256(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSHA3256(tt.args.str); got != tt.want {
				t.Errorf("IsSHA3256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSHA3384(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSHA3384(tt.args.str); got != tt.want {
				t.Errorf("IsSHA3384() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSHA3512(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSHA3512(tt.args.str); got != tt.want {
				t.Errorf("IsSHA3512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSHA384(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSHA384(tt.args.str); got != tt.want {
				t.Errorf("IsSHA384() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSHA512(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSHA512(tt.args.str); got != tt.want {
				t.Errorf("IsSHA512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSSN(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSSN(tt.args.str); got != tt.want {
				t.Errorf("IsSSN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSemver(t *testing.T) {
	type args struct {
		str string
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
			if got := IsSemver(tt.args.str); got != tt.want {
				t.Errorf("IsSemver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTiger128(t *testing.T) {
	type args struct {
		str string
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
			if got := IsTiger128(tt.args.str); got != tt.want {
				t.Errorf("IsTiger128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTiger160(t *testing.T) {
	type args struct {
		str string
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
			if got := IsTiger160(tt.args.str); got != tt.want {
				t.Errorf("IsTiger160() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTiger192(t *testing.T) {
	type args struct {
		str string
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
			if got := IsTiger192(tt.args.str); got != tt.want {
				t.Errorf("IsTiger192() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTime(t *testing.T) {
	type args struct {
		str    string
		format string
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
			if got := IsTime(tt.args.str, tt.args.format); got != tt.want {
				t.Errorf("IsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsType(t *testing.T) {
	type args struct {
		v      interface{}
		params []string
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
			if got := IsType(tt.args.v, tt.args.params...); got != tt.want {
				t.Errorf("IsType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsULID(t *testing.T) {
	type args struct {
		str string
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
			if got := IsULID(tt.args.str); got != tt.want {
				t.Errorf("IsULID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	type args struct {
		str string
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
			if got := IsURL(tt.args.str); got != tt.want {
				t.Errorf("IsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUTFDigit(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUTFDigit(tt.args.str); got != tt.want {
				t.Errorf("IsUTFDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUTFLetter(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUTFLetter(tt.args.str); got != tt.want {
				t.Errorf("IsUTFLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUTFLetterNumeric(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUTFLetterNumeric(tt.args.str); got != tt.want {
				t.Errorf("IsUTFLetterNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUTFNumeric(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUTFNumeric(tt.args.str); got != tt.want {
				t.Errorf("IsUTFNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUUID(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUUID(tt.args.str); got != tt.want {
				t.Errorf("IsUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUUIDv3(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUUIDv3(tt.args.str); got != tt.want {
				t.Errorf("IsUUIDv3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUUIDv4(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUUIDv4(tt.args.str); got != tt.want {
				t.Errorf("IsUUIDv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUUIDv5(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUUIDv5(tt.args.str); got != tt.want {
				t.Errorf("IsUUIDv5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnixFilePath(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUnixFilePath(tt.args.str); got != tt.want {
				t.Errorf("IsUnixFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnixTime(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUnixTime(tt.args.str); got != tt.want {
				t.Errorf("IsUnixTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUpperCase(t *testing.T) {
	type args struct {
		str string
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
			if got := IsUpperCase(tt.args.str); got != tt.want {
				t.Errorf("IsUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsVariableWidth(t *testing.T) {
	type args struct {
		str string
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
			if got := IsVariableWidth(tt.args.str); got != tt.want {
				t.Errorf("IsVariableWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWinFilePath(t *testing.T) {
	type args struct {
		str string
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
			if got := IsWinFilePath(tt.args.str); got != tt.want {
				t.Errorf("IsWinFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxStringLength(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := MaxStringLength(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("MaxStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStringLength(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := MinStringLength(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("MinStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := Range(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneLength(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := RuneLength(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("RuneLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetFieldsRequiredByDefault(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetFieldsRequiredByDefault(tt.args.value)
		})
	}
}

func TestSetNilPtrAllowedByRequired(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetNilPtrAllowedByRequired(tt.args.value)
		})
	}
}

func TestStringLength(t *testing.T) {
	type args struct {
		str    string
		params []string
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
			if got := StringLength(tt.args.str, tt.args.params...); got != tt.want {
				t.Errorf("StringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringMatches(t *testing.T) {
	type args struct {
		s      string
		params []string
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
			if got := StringMatches(tt.args.s, tt.args.params...); got != tt.want {
				t.Errorf("StringMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnsupportedTypeError_Error(t *testing.T) {
	type fields struct {
		Type reflect.Type
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &UnsupportedTypeError{
				Type: tt.fields.Type,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateArray(t *testing.T) {
	type args struct {
		array    []interface{}
		iterator ConditionIterator
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
			if got := ValidateArray(tt.args.array, tt.args.iterator); got != tt.want {
				t.Errorf("ValidateArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateMap(t *testing.T) {
	type args struct {
		s map[string]interface{}
		m map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateMap(tt.args.s, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateMapAsync(t *testing.T) {
	type args struct {
		s map[string]interface{}
		m map[string]interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  <-chan bool
		want1 <-chan error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateMapAsync(tt.args.s, tt.args.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateMapAsync() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateMapAsync() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestValidateStruct(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateStruct(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateStruct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateStructAsync(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  <-chan bool
		want1 <-chan error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateStructAsync(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateStructAsync() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateStructAsync() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkRequired(t *testing.T) {
	type args struct {
		v       reflect.Value
		t       reflect.StructField
		options tagOptionsMap
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkRequired(tt.args.v, tt.args.t, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkRequired() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkRequired() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEmptyValue(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isEmptyValue(tt.args.v); got != tt.want {
				t.Errorf("isEmptyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidTag(t *testing.T) {
	type args struct {
		s string
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
			if got := isValidTag(tt.args.s); got != tt.want {
				t.Errorf("isValidTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTagIntoMap(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want tagOptionsMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTagIntoMap(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTagIntoMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prependPathToErrors(t *testing.T) {
	type args struct {
		err  error
		path string
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
			if err := prependPathToErrors(tt.args.err, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("prependPathToErrors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_stringValues_Len(t *testing.T) {
	tests := []struct {
		name string
		sv   stringValues
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sv.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringValues_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		sv   stringValues
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sv.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringValues_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		sv   stringValues
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sv.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_stringValues_get(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		sv   stringValues
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sv.get(tt.args.i); got != tt.want {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stripParams(t *testing.T) {
	type args struct {
		validatorString string
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
			if got := stripParams(tt.args.validatorString); got != tt.want {
				t.Errorf("stripParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toJSONName(t *testing.T) {
	type args struct {
		tag string
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
			if got := toJSONName(tt.args.tag); got != tt.want {
				t.Errorf("toJSONName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeCheck(t *testing.T) {
	type args struct {
		v       reflect.Value
		t       reflect.StructField
		o       reflect.Value
		options tagOptionsMap
	}
	tests := []struct {
		name        string
		args        args
		wantIsValid bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsValid, err := typeCheck(tt.args.v, tt.args.t, tt.args.o, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("typeCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("typeCheck() gotIsValid = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}
