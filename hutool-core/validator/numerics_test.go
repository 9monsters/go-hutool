package validator

import "testing"

func TestAbs(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.value); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInRange(t *testing.T) {
	type args struct {
		value interface{}
		left  interface{}
		right interface{}
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
			if got := InRange(tt.args.value, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("InRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInRangeFloat32(t *testing.T) {
	type args struct {
		value float32
		left  float32
		right float32
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
			if got := InRangeFloat32(tt.args.value, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("InRangeFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInRangeFloat64(t *testing.T) {
	type args struct {
		value float64
		left  float64
		right float64
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
			if got := InRangeFloat64(tt.args.value, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("InRangeFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInRangeInt(t *testing.T) {
	type args struct {
		value interface{}
		left  interface{}
		right interface{}
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
			if got := InRangeInt(tt.args.value, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("InRangeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNatural(t *testing.T) {
	type args struct {
		value float64
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
			if got := IsNatural(tt.args.value); got != tt.want {
				t.Errorf("IsNatural() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNegative(t *testing.T) {
	type args struct {
		value float64
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
			if got := IsNegative(tt.args.value); got != tt.want {
				t.Errorf("IsNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNonNegative(t *testing.T) {
	type args struct {
		value float64
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
			if got := IsNonNegative(tt.args.value); got != tt.want {
				t.Errorf("IsNonNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNonPositive(t *testing.T) {
	type args struct {
		value float64
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
			if got := IsNonPositive(tt.args.value); got != tt.want {
				t.Errorf("IsNonPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPositive(t *testing.T) {
	type args struct {
		value float64
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
			if got := IsPositive(tt.args.value); got != tt.want {
				t.Errorf("IsPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWhole(t *testing.T) {
	type args struct {
		value float64
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
			if got := IsWhole(tt.args.value); got != tt.want {
				t.Errorf("IsWhole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSign(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.value); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
