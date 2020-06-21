package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveShippingMethods(ctx context.Context, in *pbcheckout.RetrieveShippingMethodsRequest) (*pbcheckout.RetrieveShippingMethodsResponse, error) {
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
	o, err := server.FindLatestPurchase(ctx, db, uid)
	if err != nil {
		return nil, err
	}

	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()

	sem, semCloser, err := pbutil.ConnSemanticProducts(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	addrs, addrsCloser, err := pbutil.ConnAddresses(ctx)
	if err != nil {
		return nil, err
	}
	defer addrsCloser()

	users, usersCloser, err := pbutil.ConnUsersViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()

	svcs, svcsCloser, err := pbutil.ConnServicesViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer svcsCloser()

	return s.retrieveShippingMethods(ctx, db,
		sem, users, prods, svcs, addrs,
		in, o,
	)
}

func (s *Server) retrieveShippingMethods(
	ctx context.Context,
	db *sqlx.DB,

	sem pbsemantic.ProductsClient,
	users pbusers.ViewerClient,
	prods pbproducts.ViewerClient,
	svcs pbservices.ViewerClient,
	addr pbaddress.AddressesClient,

	in *pbcheckout.RetrieveShippingMethodsRequest,
	order *order.Purchase,
) (*pbcheckout.RetrieveShippingMethodsResponse, error) {
	start, err := pbutil.RestTimeOfDay(in.GetTime(), dispatchDelayMinutes)
	if err != nil {
		return nil, err
	}
	mch, err := order.Merchant(ctx, users)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Merchant: %v", err)
	}

	dist, err := addr.MeasureDistance(ctx, &pbaddress.MeasureDistanceRequest{AAddressId: order.AddressId, BAddressId: mch.GetUser().GetAddressId()})
	if err != nil {
		return nil, err
	}

	shippings, err := svcs.SearchAvailableShippings(ctx, &pbservices.SearchAvailableShippingsRequest{
		MaxVolWeight:         2, // TODO: Add products count by cat
		DistanceInKilometers: dist.GetManhattanInKilometers(),

		Dow:   in.GetDow(),
		Start: start,
		End:   in.GetTime(),
	})
	if err != nil {
		return nil, err
	}

	resp := &pbcheckout.RetrieveShippingMethodsResponse{}
	resp.ShippingMethods = make(map[uint64]*pbcheckout.ShippingMethod)
	for id, ship := range shippings.GetServices() {
		resp.ShippingMethods[id] = serviceSearchResultToShippingMethod(ship) // TODO: Add price
	}
	return nil, nil
}

func serviceSearchResultToShippingMethod(ship *pbservices.ServiceSearchResult) *pbcheckout.ShippingMethod {
	svc, user := ship.GetService(), ship.GetUser()
	return &pbcheckout.ShippingMethod{
		Service: &pbcheckout.Service{
			Title: svc.GetTitle(),
			// Price: svc.GetPrice(),
			DurationInMinutes: svc.GetDurationInMinutes(),
		},
		User: &pbcheckout.User{
			Name:      user.GetName(),
			Surname:   user.GetSurname(),
			ImageUrl:  user.GetImageUrl(),
			AddressId: user.GetAddressId(),
		},
	}

}
func (s *Server) RetrieveCurrent(ctx context.Context, in *pbcheckout.RetrieveCurrentRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
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
	o, err := server.FindLatestPurchase(ctx, db, uid)
	if err != nil {
		return nil, err
	}

	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()
	return s.retrievePurchase(ctx, db, prods, o)
}

func (s *Server) retrievePurchase(
	ctx context.Context,
	db *sqlx.DB,
	prods pbproducts.ViewerClient,
	order *order.Purchase,
) (*pbcheckout.RetrievePurchaseResponse, error) {
	o, err := order.ToPbWrapped(ctx, db, prods)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "o.ToPbWrapped")
	}

	return &pbcheckout.RetrievePurchaseResponse{PurchaseId: order.Id, Purchase: o}, nil
}
