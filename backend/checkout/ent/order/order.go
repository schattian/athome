package order

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/jmoiron/sqlx"
)

type Order interface {
	sm.Stateful

	OrderClass() Class
	CanView(ctx context.Context, db *sqlx.DB, userId uint64) (bool, error)
}

func (s *Shipping) GetMerchantId() uint64        { return 0 }
func (s *Shipping) GetConsumerId() uint64        { return 0 }
func (s *Shipping) GetServiceProviderId() uint64 { return s.UserId }

func (p *Payment) GetMerchantId() uint64        { return 0 }
func (p *Payment) GetConsumerId() uint64        { return p.UserId }
func (p *Payment) GetServiceProviderId() uint64 { return 0 }

func (p *Purchase) GetMerchantId() uint64        { return p.MerchantId }
func (p *Purchase) GetConsumerId() uint64        { return p.UserId }
func (p *Purchase) GetServiceProviderId() uint64 { return 0 }
