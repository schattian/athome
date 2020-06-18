package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) createPurchase(ctx context.Context, db *sqlx.DB,
	in *pbcheckout.CreatePurchaseRequest,
	prods pbproducts.ViewerClient,
	userId uint64,
) (*pbcheckout.CreatePurchaseResponse, error) {
	o := order.NewPurchase(ctx, in.GetItems(), userId)
	products, err := o.Products(ctx, prods)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Products")
	}
	err = o.ValidateStock(ctx, products)
	if err != nil {
		return nil, status.Errorf(xerrors.ResourceExhausted, "ValidateStock")
	}
	err = storeql.InsertIntoDB(ctx, db, o)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "o InsertIntoDB")
	}

	sc, err := order.NewPurchaseStateChange(ctx, o.Id)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB")
	}
	err = storeql.InsertIntoDB(ctx, db, sc)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "sc InsertIntoDB")
	}

	amount := o.AmountFromProducts(ctx, products)
	oPb, err := o.ToPb([]order.StateChange{sc}, amount)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "ToPbWrapped")
	}

	return &pbcheckout.CreatePurchaseResponse{Purchase: oPb, PurchaseId: o.Id}, nil
}

func (s *Server) CreatePurchase(ctx context.Context, in *pbcheckout.CreatePurchaseRequest) (*pbcheckout.CreatePurchaseResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = authCloser()
	if err != nil {
		return nil, err
	}

	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()
	return s.createPurchase(ctx, db, in, prods, uid)
}

func (s *Server) StateMachine(ctx context.Context, _ *emptypb.Empty) (*pbcheckout.StateMachineResponse, error) {
	return sm.PurchaseStateMachine.ToPb(), nil
}

func (s *Server) RetrieveShippingMethods(context.Context, *pbcheckout.RetrieveShippingMethodsRequest) (*pbcheckout.RetrieveShippingMethodsResponse, error) {
	return nil, nil
}

func (s *Server) RetrieveCurrent(ctx context.Context, in *pbcheckout.RetrieveCurrentRequest) (*pbcheckout.Purchase, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = authCloser()
	if err != nil {
		return nil, err
	}

	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()
	return s.retrieveCurrent(ctx, db, in, prods, uid)
}

func (s *Server) retrieveCurrent(
	ctx context.Context,
	db *sqlx.DB,
	in *pbcheckout.RetrieveCurrentRequest,
	prods pbproducts.ViewerClient,
	userId uint64,
) (*pbcheckout.Purchase, error) {
	order, err := order.FindLatestPurchase(ctx, db, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindLatestPurchase")
	}
	o, err := order.ToPbWrapped(ctx, db, prods)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "o.ToPbWrapped")
	}
	return o, nil
}
