package rsa

import (
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
		{name: "TestGenRsaKeyHex",
			args: args{1024},
			want: AsymmetricRsaKey{
				"3082025c02010002818100bddf45cc6d8302aa09957fae9cbe00c39353ab0bfdaa0f1ed73f4d511c0bc77066356881ec30c12dab14e80f15c7a6d21586b76ff29a97e6fc47ea14751dfb51843e2a6e3005cdd6937296aaf98f4a5ddf8708aa966e18cd395e93da4607115e2f455de6f74f1056b08a11d8ea5bf8804c02fabbc740cbf8e7d226c7ff6135e5020301000102818053db17e881352b280db389008519251d3b23e8971320d82cb6c3ae51f420ffa4147fe1f7ff6848e8f275f94730474cd5dbcbf397ed7d7938fa92883f20ba6d13807c52810a0adfba715323da5b44358996271953d5df8ba9e826385368b7d9739e6e44e5b2e71db6e5d928791fe414f80c808b6699ef1c06228a14ca57381801024100e5b47fc4d636a99ffe6fe53732d9ebdec2c433068bb0ccf67149f2c7568c56b09e006a7ca610330d3c38e022ff955926c53bfd6a1da4c4b5ec844f47ce010e61024100d39b7896b390c9e13d4af61be744ffd7bd1752c8d7c3a69b983f5cc75b7f3d0d96106b62770ee2ae9b177896d9f5d86a1217ddb2b1c44eae1ef6bed68d56ae05024034bc59d110f4ffc071a3b2d617609c9c49cecd802535f3f8684df7e4f7e3c44b4c1b005799b08267f9797d4bcbadb80804a41f2beb9e566305e7e263b3d1de4102405602657b6bb5383952fe7b1650645d51454c8b9b9307b55a9d0174269bec15d6c5de1aa7c518c9ce8abf798667456417890f264f91ade7052314cbf0cd362439024100cf2419d2ac3c04823c2f582951d55ffa32b261afe98de78876c700a01c8868b6a0e92a7f005909ba4bbab827ae72122bf706be0b5766dc2b51bf88c548b0103a",
				"30818902818100bddf45cc6d8302aa09957fae9cbe00c39353ab0bfdaa0f1ed73f4d511c0bc77066356881ec30c12dab14e80f15c7a6d21586b76ff29a97e6fc47ea14751dfb51843e2a6e3005cdd6937296aaf98f4a5ddf8708aa966e18cd395e93da4607115e2f455de6f74f1056b08a11d8ea5bf8804c02fabbc740cbf8e7d226c7ff6135e50203010001"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenRsaKeyHex(tt.args.bits)
			t.Logf("GenRsaKeyHex() = %v, want %v", got, tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenRsaKeyHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.PrivateKey) == 0 || len(got.PublicKey) == 0 {
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
			name: "TestGenRsaKeyBase64",
			args: args{2048},
			want: AsymmetricRsaKey{
				"MIICXgIBAAKBgQC+LqAHUOS9H73Ckv9LEDUlfxD1t4q1bwU6xAlawHoVTYDtZUy8cUOSAsvs0dvaG4q8gdomHl6aNS4d8Ykjp52JPA1KBMt0schm/PE3V27hiGIdN31JhzVuzUlj5CsE4oWe4+vAcIsDaCDqOapF95ZckWlJ6sWzDBGhIDxo4o3wlQIDAQABAoGBAIKP/80F/ALah6vMvvcJaXMUiNAn8VIdpblyW/mxqh511s0jfkrgd0MGKHmsmQJjZhuECf+hWdEEMWTdLQkNSVh+jF04xPY07N6FgQCL+70ReLIcfI8ZFPBsU0T5EuS7r6f8XMkWJWHoC5dzl9nOO+wFcCRjFfFay1v8QdL9D/ABAkEA4vB1DUWKBA6bamhBZ1jJ6zDVuFzI3pdeFhP1ebv9yRNtFW7z5twA2e6uMn+Fh/K9uX2SxzCHk1I6xzQSJQGSRQJBANaJMRGMNEW+De0Hpte1ij9MkgcODAGhSeFt9XJPLUConVh5fZdBzyIl2yry9b/+U9HvHnMAHVtFGL1ejnZW8hECQQC5jIEn6MN/rbHEvpk7No1hjvwvUS2rUXfL9WWvstU9onR+Icmp6BeKGGy1PS6giQg2sUVhN3yiJ4mHdjufpjlBAkBX4xwW9Dj12UYbNGdFNznLdLLd8QM6J7j8gO9sRMNlGa5b8Gli7bLNanS+w4mpfacY+byfoAxmt4fLDFGQaU4BAkEA1SGajgPkBnA5wDy//ip8QkQTl6zW0oZPxHBvcqOE/qHOJEkzr47zfuDu33UUwXfCO7nlceKtiAaD3qQ7GUuddg==",
				"MIGJAoGBAL4uoAdQ5L0fvcKS/0sQNSV/EPW3irVvBTrECVrAehVNgO1lTLxxQ5ICy+zR29obiryB2iYeXpo1Lh3xiSOnnYk8DUoEy3SxyGb88TdXbuGIYh03fUmHNW7NSWPkKwTihZ7j68BwiwNoIOo5qkX3llyRa",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateRsaKeyBase64(tt.args.bits)
			t.Logf("GenRsaKeyHex() = %v, want %v", got, tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRsaKeyBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.PrivateKey) == 0 || len(got.PublicKey) == 0 {
				t.Errorf("GenerateRsaKeyBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
