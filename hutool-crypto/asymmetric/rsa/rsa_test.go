package rsa

import (
	"reflect"
	"testing"
)

func TestGenRsaKeyHex(t *testing.T) {
	type args struct {
		bits int
	}
	tests := []struct {
		name    string
		args    args
		want    AsymmetricRsaKey
		wantErr bool
	}{
		{name: "test1", args: args{1024}, want: AsymmetricRsaKey{"", ""}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenRsaKeyHex(tt.args.bits)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenRsaKeyHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenRsaKeyHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenRsaKeyBase64(t *testing.T) {
	type args struct {
		bits int
	}
	tests := []struct {
		name    string
		args    args
		want    AsymmetricRsaKey
		wantErr bool
	}{
		{
			name: "test1",
			args: args{1024},
			want: AsymmetricRsaKey{
				"MIICXgIBAAKBgQC+LqAHUOS9H73Ckv9LEDUlfxD1t4q1bwU6xAlawHoVTYDtZUy8cUOSAsvs0dvaG4q8gdomHl6aNS4d8Ykjp52JPA1KBMt0schm/PE3V27hiGIdN31JhzVuzUlj5CsE4oWe4+vAcIsDaCDqOapF95ZckWlJ6sWzDBGhIDxo4o3wlQIDAQABAoGBAIKP/80F/ALah6vMvvcJaXMUiNAn8VIdpblyW/mxqh511s0jfkrgd0MGKHmsmQJjZhuECf+hWdEEMWTdLQkNSVh+jF04xPY07N6FgQCL+70ReLIcfI8ZFPBsU0T5EuS7r6f8XMkWJWHoC5dzl9nOO+wFcCRjFfFay1v8QdL9D/ABAkEA4vB1DUWKBA6bamhBZ1jJ6zDVuFzI3pdeFhP1ebv9yRNtFW7z5twA2e6uMn+Fh/K9uX2SxzCHk1I6xzQSJQGSRQJBANaJMRGMNEW+De0Hpte1ij9MkgcODAGhSeFt9XJPLUConVh5fZdBzyIl2yry9b/+U9HvHnMAHVtFGL1ejnZW8hECQQC5jIEn6MN/rbHEvpk7No1hjvwvUS2rUXfL9WWvstU9onR+Icmp6BeKGGy1PS6giQg2sUVhN3yiJ4mHdjufpjlBAkBX4xwW9Dj12UYbNGdFNznLdLLd8QM6J7j8gO9sRMNlGa5b8Gli7bLNanS+w4mpfacY+byfoAxmt4fLDFGQaU4BAkEA1SGajgPkBnA5wDy//ip8QkQTl6zW0oZPxHBvcqOE/qHOJEkzr47zfuDu33UUwXfCO7nlceKtiAaD3qQ7GUuddg==",
				"MIGJAoGBAL4uoAdQ5L0fvcKS/0sQNSV/EPW3irVvBTrECVrAehVNgO1lTLxxQ5ICy+zR29obiryB2iYeXpo1Lh3xiSOnnYk8DUoEy3SxyGb88TdXbuGIYh03fUmHNW7NSWPkKwTihZ7j68BwiwNoIOo5qkX3llyRa",
			}, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateRsaKeyBase64(tt.args.bits)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRsaKeyBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRsaKeyBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}


