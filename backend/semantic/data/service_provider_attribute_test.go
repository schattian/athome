package data

import (
	"testing"

	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func newServiceProviderAttributeData(t *testing.T, ty value.Type) *ServiceProviderAttributeData {
	t.Helper()
	pc, err := NewServiceProviderAttributeData(ty)
	if err != nil {
		t.Fatal(err)
	}
	return pc
}

func TestServiceProviderAttributeData_SetValue(t *testing.T) {
	tests := []struct {
		name string
		att  *ServiceProviderAttributeData
		v    interface{}

		wantAtt *ServiceProviderAttributeData
		wantErr bool
	}{
		{
			name:    "nil val given",
			att:     newServiceProviderAttributeData(t, value.TypeBool),
			wantErr: true,
		},
		{
			name:    "all nil values",
			att:     &ServiceProviderAttributeData{},
			wantErr: true,
		},
		{
			name:    "bool",
			att:     newServiceProviderAttributeData(t, value.TypeBool),
			v:       true,
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{BoolValue: &value.Bool{Bool: true, Valid: true}},
		},
		{
			name:    "f64",
			att:     newServiceProviderAttributeData(t, value.TypeFloat64),
			v:       3.02,
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{Float64Value: &value.Float64{Float64: 3.02, Valid: true}},
		},
		{
			name:    "i64",
			att:     newServiceProviderAttributeData(t, value.TypeInt64),
			v:       int64(3),
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{Int64Value: &value.Int64{Int64: 3, Valid: true}},
		},
		{
			name:    "string",
			att:     newServiceProviderAttributeData(t, value.TypeString),
			v:       "foo",
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{StringValue: &value.String{String: "foo", Valid: true}},
		},
		{
			name:    "slstring",
			att:     newServiceProviderAttributeData(t, value.TypeSlString),
			v:       []string{"foo", "bar"},
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{SlStringValue: &value.SlString{
				&value.String{String: "foo", Valid: true},
				&value.String{String: "bar", Valid: true},
			}},
		},
		{
			name:    "slint64",
			att:     newServiceProviderAttributeData(t, value.TypeSlInt64),
			v:       []int64{1, 2},
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{SlInt64Value: &value.SlInt64{
				&value.Int64{Int64: 1, Valid: true},
				&value.Int64{Int64: 2, Valid: true},
			}},
		},
		{
			name:    "slfloat64",
			att:     newServiceProviderAttributeData(t, value.TypeSlFloat64),
			v:       []float64{1.1, 2.2},
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{SlFloat64Value: &value.SlFloat64{
				&value.Float64{Float64: 1.1, Valid: true},
				&value.Float64{Float64: 2.2, Valid: true},
			}},
		},
		{
			name:    "slfloat32",
			att:     newServiceProviderAttributeData(t, value.TypeSlFloat64),
			v:       []float32{1.1, 2.2},
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{SlFloat64Value: &value.SlFloat64{
				&value.Float64{Float64: 1.1, Valid: true},
				&value.Float64{Float64: 2.2, Valid: true},
			}},
		},
		{
			name:    "slint",
			att:     newServiceProviderAttributeData(t, value.TypeSlInt64),
			v:       []int{1, 2},
			wantErr: false,
			wantAtt: &ServiceProviderAttributeData{SlInt64Value: &value.SlInt64{
				&value.Int64{Int64: 1, Valid: true},
				&value.Int64{Int64: 2, Valid: true},
			}},
		},
	}
	for _, tt := range tests {
		// tt := tt
		t.Run(tt.name, func(t *testing.T) {

			// t.Parallel()
			old := *tt.att
			err := tt.att.SetValue(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceProviderAttributeData.SetValue() error = %v, wantErr %v", err, tt.wantErr)
			}

			msg := "ServiceProviderAttributeData.SetValue()"
			if err != nil {
				if diff := cmp.Diff(&old, tt.att, cmpopts.IgnoreUnexported(ServiceProviderAttributeData{})); diff != "" {
					t.Fatalf("%s errored mismatch (-want +got): %s", msg, diff)
				}
			}

			if diff := cmp.Diff(tt.wantAtt, tt.att, cmpopts.IgnoreUnexported(ServiceProviderAttributeData{})); diff != "" {
				t.Errorf("%s mismatch (-want +got): %s", msg, diff)
			}

		})
	}
}
