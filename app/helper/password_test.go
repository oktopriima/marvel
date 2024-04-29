package helper

import "testing"

func TestGeneratePassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "generate password",
			args: args{
				password: "123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GeneratePassword(tt.args.password)
			t.Logf(got)
		})
	}
}
