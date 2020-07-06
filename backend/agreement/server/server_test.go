package server

import (
	"testing"
)

func Test_randString(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "7 digits",
			args: args{s: 7},
		},
		{
			name: "8 digits",
			args: args{s: 8},
		},
		{
			name: "6 digits",
			args: args{s: 6},
		},
		{
			name: "3 digits",
			args: args{s: 3},
		},
		{
			name: "2 digits",
			args: args{s: 2},
		},
		{
			name: "4 digits",
			args: args{s: 4},
		},
		{
			name: "1 digits",
			args: args{s: 1},
		},
		{
			name: "5 digits",
			args: args{s: 5},
		},
		{
			name: "18 digits",
			args: args{s: 18},
		},
		{
			name: "19 digits",
			args: args{s: 19},
		},
		{
			name:    "negative",
			args:    args{s: -10},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := randString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("randString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if len(got) != tt.args.s {
				t.Errorf("randString() = %v, want length %v", got, tt.args.s)
			}
		})
	}
}
