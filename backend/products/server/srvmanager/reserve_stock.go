package srvmanager

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type reserveStockAuthorizationFunc func(order order, uid uint64) bool

func confirmReserveAuth(order order, uid uint64) bool { return order.GetMerchantId() == uid }
func createReserveAuth(order order, uid uint64) bool  { return order.GetUserId() == uid }
func deleteReserveAuth(order order, uid uint64) bool {
	return createReserveAuth(order, uid) || confirmReserveAuth(order, uid)
}

type reserveStockFunc func(p *ent.Product, qt uint64) error

func createReserve(p *ent.Product, qt uint64) error  { return p.ReserveStock(qt) }
func deleteReserve(p *ent.Product, qt uint64) error  { return p.UndoStockReserve(qt) }
func confirmReserve(p *ent.Product, qt uint64) error { return p.ConsumeStockReserve(qt) }

func (s *Server) CreateReserveStock(ctx context.Context, in *pbproducts.ReserveStockRequest) (*emptypb.Empty, error) {
	return s.ReserveStock(ctx, in, createReserve, createReserveAuth)
}
func (s *Server) ConfirmReserveStock(ctx context.Context, in *pbproducts.ReserveStockRequest) (*emptypb.Empty, error) {
	return s.ReserveStock(ctx, in, confirmReserve, confirmReserveAuth)
}

func (s *Server) DeleteReserveStock(ctx context.Context, in *pbproducts.ReserveStockRequest) (*emptypb.Empty, error) {
	return s.ReserveStock(ctx, in, deleteReserve, deleteReserveAuth)
}

func (s *Server) ReserveStock(ctx context.Context, in *pbproducts.ReserveStockRequest, reserveStockFn reserveStockFunc, authFn reserveStockAuthorizationFunc) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	c, closer, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()
	uid, err := pbutil.GetUserFromAccessToken(ctx, c, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	order, err := retrieveOrder(ctx, in.GetOrder(), in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	if !authFn(order, uid) {
		return nil, status.Error(xerrors.PermissionDenied, "you are not allowed to reserve stock")
	}
	return s.reserveStock(ctx, db, in, reserveStockFn, order)
}

func (s *Server) reserveStock(ctx context.Context, db *sqlx.DB, in *pbproducts.ReserveStockRequest,
	reserveStockFn reserveStockFunc,
	order order,
) (*emptypb.Empty, error) {
	prods, err := ent.FindProductsById(ctx, db, getItemIds(order))
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductsById")
	}
	var storables []storeql.Storable
	for _, prod := range prods {
		err = reserveStockFn(prod, order.GetItems()[prod.Id])
		if err != nil {
			return nil, status.Errorf(xerrors.InvalidArgument, "ReserveStock")
		}
		storables = append(storables, prod)
	}
	err = storeql.UpdateIntoDB(ctx, db, storables...)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB")
	}
	return &emptypb.Empty{}, nil
}
