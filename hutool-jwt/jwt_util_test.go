package jwt

import "testing"

func TestCreateToken(t *testing.T) {
	type args struct {
		payload map[string]any
		key     []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{payload: map[string]any{"1": "1"}, key: []byte("123")}, "eyJhbGciOiJIUzI1NiJ9.eyIxIjoiMSJ9.gi00We2cXmQSIhmXdssCBYLbsRD2ZCjwzOeE4q4qKwA="},
		{"test1", args{payload: map[string]any{"typ": "JWT"}, key: []byte("1234567890")}, "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiYWRtaW4iOnRydWUsIm5hbWUiOiJsb29seSJ9.U2aQkC2THYV9L0fTN-yBBI7gmo5xhmvMhATtu8v0zEA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateToken(tt.args.payload, tt.args.key); got != tt.want {
				t.Errorf("CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
