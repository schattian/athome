package payment

import "github.com/athomecomar/athome/backend/checkout/ent/sm"

func (p *Payment) GetMerchantId() uint64        { return 0 }
func (p *Payment) GetConsumerId() uint64        { return p.UserId }
func (p *Payment) GetServiceProviderId() uint64 { return 0 }

func (p *Payment) StateChange() sm.StateChange { return &PaymentStateChange{} }
