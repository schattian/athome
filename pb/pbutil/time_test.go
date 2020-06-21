package pbutil

import (
	"reflect"
	"testing"

	"github.com/athomecomar/athome/pb/pbshared"
)

func TestRestTimeOfDay(t *testing.T) {
	type args struct {
		t1   *pbshared.TimeOfDay
		mins int64
	}
	tests := []struct {
		name    string
		args    args
		want    *pbshared.TimeOfDay
		wantErr bool
	}{
		{
			name: "basic rest only mins",
			args: args{
				t1:   &pbshared.TimeOfDay{Hour: 3, Minute: 10},
				mins: 10,
			},
			want: &pbshared.TimeOfDay{Hour: 3, Minute: 0},
		},
		{
			name: "basic rest only hs",
			args: args{
				t1:   &pbshared.TimeOfDay{Hour: 3, Minute: 10},
				mins: 180,
			},
			want: &pbshared.TimeOfDay{Hour: 0, Minute: 10},
		},
		{
			name: "basic rest  mins overflow and hs",
			args: args{
				t1:   &pbshared.TimeOfDay{Hour: 3, Minute: 10},
				mins: 135,
			},
			want: &pbshared.TimeOfDay{Hour: 0, Minute: 55},
		},
		{
			name: "day overflow",
			args: args{
				t1:   &pbshared.TimeOfDay{Hour: 3, Minute: 10},
				mins: 10000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RestTimeOfDay(tt.args.t1, tt.args.mins)
			if (err != nil) != tt.wantErr {
				t.Errorf("RestTimeOfDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestTimeOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
