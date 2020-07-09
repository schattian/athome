package purchase

import (
	"context"
	"fmt"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (p *Purchase) GetMerchantId() uint64        { return p.MerchantId }
func (p *Purchase) GetConsumerId() uint64        { return p.UserId }
func (p *Purchase) GetServiceProviderId() uint64 { return 0 }

func (o *Purchase) StateMachine() *sm.StateMachine {
	return sm.PurchaseStateMachine
}

func (o *Purchase) ValidateStateChange(ctx context.Context, db *sqlx.DB, newState *sm.State) (err error) {
	switch newState.Name {
	case sm.PurchaseAddressed:
		err = o.validateStateChangeAddressed(ctx, db)
	case sm.PurchasePaid:
		err = o.validateStateChangePaid(ctx, db)
	case sm.PurchaseFinished:
		err = o.validateStateChangeFinished(ctx, db)
	}
	return
}

func (o *Purchase) validateStateChangeFinished(ctx context.Context, db *sqlx.DB) error {
	if o.ShippingId == 0 {
		return nil
	}
	ship, err := o.Shipping(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Shipping")
	}
	sc, err := sm.LatestStateChange(ctx, db, ship)
	if err != nil {
		return errors.Wrap(err, "sm.LatestStateChange")
	}
	if sc.GetState(ship.StateMachine()).Name != sm.ShippingFinished {
		return errors.New("shipping isn't finshed yet")
	}
	return nil
}

func (o *Purchase) validateStateChangeAddressed(ctx context.Context, db *sqlx.DB) error {
	if o.DestAddressId == 0 {
		return errors.New("nil dest address id")
	}
	return nil
}

func (o *Purchase) validateStateChangePaid(ctx context.Context, db *sqlx.DB) error {
	paidAmount, err := o.TotalPaid(ctx, db)
	if err != nil {
		return errors.Wrap(err, "TotalPaid")
	}
	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return errors.Wrap(err, "pbutil.ConnProductsViewer")
	}
	defer prodsCloser()
	totalAmount, err := o.Amount(ctx, db, prods)
	if err != nil {
		return errors.Wrap(err, "Amount")
	}
	if totalAmount != paidAmount.Float64() {
		return fmt.Errorf("order total wasn't paid. Total: %v, paid: %v", totalAmount, paidAmount.Float64())
	}
	return nil
}
