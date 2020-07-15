package purchase

import (
	"context"
	"testing"

	"github.com/athomecomar/athome/pb/pbproducts"
)

func TestPurchase_AmountFromProducts(t *testing.T) {
	type fields struct {
	}
	type args struct {
		ctx   context.Context
		prods map[uint64]*pbproducts.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		items   Items
		want    float64
		wantErr bool
	}{
		{
			name:  "basic",
			items: Items{1: 3, 2: 1},
			args: args{
				prods: map[uint64]*pbproducts.Product{
					1: {Price: 100, Stock: 29},
					2: {Price: 33, Stock: 29},
				},
			},
			want: 100*3 + 33*1,
		},
		{
			name:  "ran out of stock",
			items: Items{1: 3, 2: 1000},
			args: args{
				prods: map[uint64]*pbproducts.Product{
					1: {Price: 100, Stock: 29},
					2: {Price: 33, Stock: 29},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Purchase{Items: tt.items}
			got, err := o.AmountFromProducts(tt.args.ctx, tt.args.prods)
			if (err != nil) != tt.wantErr {
				t.Errorf("Purchase.AmountFromProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Purchase.AmountFromProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
