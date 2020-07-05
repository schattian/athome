package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

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

	prodsViewer, prodsViewerCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsViewerCloser()

	prodsManager, prodsManagerCloser, err := pbutil.ConnProductsManager(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsManagerCloser()

	users, usersCloser, err := pbutil.ConnUsersViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()

	addrs, addrsCloser, err := pbutil.ConnAddresses(ctx)
	if err != nil {
		return nil, err
	}
	defer addrsCloser()

	return s.createPurchase(ctx, db, in, users, addrs, prodsViewer, prodsManager, uid)
}

func (s *Server) createPurchase(ctx context.Context, db *sqlx.DB,
	in *pbcheckout.CreatePurchaseRequest,
	users pbusers.ViewerClient,
	addr pbaddress.AddressesClient,
	prodsViewer pbproducts.ViewerClient,
	prodsManager pbproducts.ManagerClient,
	userId uint64,
) (*pbcheckout.CreatePurchaseResponse, error) {
	o := order.NewPurchase(ctx, in.GetItems(), userId)
	products, err := o.Products(ctx, prodsViewer)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Products")
	}
	err = o.AssignMerchant(ctx, products)
	if err != nil {
		return nil, status.Errorf(xerrors.ResourceExhausted, "ValidateStock")
	}
	err = o.AssignSrcAddress(ctx, users, addr)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "AssignSrcAddress")
	}
	_, err = prodsManager.CreateReserveStock(ctx, &pbproducts.ReserveStockRequest{AccessToken: in.GetAccessToken(), Order: pbutil.ToPbEntity(o)})
	if err != nil {
		return nil, err
	}
	err = storeql.InsertIntoDB(ctx, db, o)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "o InsertIntoDB")
	}

	sc, err := order.NewPurchaseStateChange(ctx, o.Id, sm.PurchaseCreated)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "NewPurchaseStateChange")
	}
	err = storeql.InsertIntoDB(ctx, db, sc)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "sc InsertIntoDB")
	}

	amount, err := o.AmountFromProducts(ctx, products)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "AmountFromProducts")
	}

	oPb, err := o.ToPb([]sm.StateChange{sc}, amount)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "ToPbWrapped")
	}

	return &pbcheckout.CreatePurchaseResponse{Purchase: oPb, PurchaseId: o.Id}, nil
}
