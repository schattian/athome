package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) StateMachine(ctx context.Context, _ *emptypb.Empty) (*pbcheckout.StateMachineResponse, error) {
	return sm.PurchaseStateMachine.ToPb(), nil
}
