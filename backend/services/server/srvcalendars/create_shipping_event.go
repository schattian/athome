package srvcalendars

import (
	"context"
	"math"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/server"
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

func (s *Server) CreateShippingEvent(ctx context.Context, in *pbservices.CreateShippingEventRequest) (*pbservices.CreateEventResponse, error) {
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
	defer authCloser()

	users, usersCloser, err := pbutil.ConnUsersViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()
	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()
	sem, semCloser, err := pbutil.ConnSemanticServiceProviders(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()
	check, checkCloser, err := pbutil.ConnCheckoutPurchases(ctx)
	if err != nil {
		return nil, err
	}
	defer checkCloser()
	return s.createShippingEvent(ctx, db,
		prods, users, sem, check,
		server.GetUserFromAccessToken(auth, in.GetAccessToken()), in)
}

func (s *Server) createShippingEvent(
	ctx context.Context,
	db *sqlx.DB,
	prods pbproducts.ViewerClient,
	users pbusers.ViewerClient,
	sem pbsemantic.ServiceProvidersClient,
	check pbcheckout.PurchasesClient,
	authFn server.AuthFunc,
	in *pbservices.CreateShippingEventRequest,
) (*pbservices.CreateEventResponse, error) {
	claimantId, err := authFn(ctx)
	if err != nil {
		return nil, err
	}
	order, err := check.RetrieveCurrent(ctx, &pbcheckout.RetrieveCurrentRequest{AccessToken: in.GetAccessToken()})
	if err != nil {
		return nil, err
	}
	svc, err := ent.FindService(ctx, db, in.GetServiceId())
	if err != nil {
		return nil, err
	}

	maxVolWeight, err := totalMaxVolWeight(ctx, order.GetPurchase(), prods)
	if err != nil {
		return nil, err
	}

	cats, err := sem.RetrieveShippingCategories(ctx, &pbsemantic.RetrieveShippingCategoriesRequest{MaxVolWeight: maxVolWeight})
	if err != nil {
		return nil, err
	}
	claimant, err := users.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: svc.UserId})
	if err != nil {
		return nil, err
	}
	if !containsCategoryId(claimant.GetUser().GetCategoryId(), cats.GetCategories()) {
		return nil, status.Errorf(xerrors.InvalidArgument, "invalid category for user that holds the svc. Max vol weight is out he's range")
	}
	totalDuration := order.GetPurchase().GetDistanceInKilometers() * float64(svc.DurationInMinutes)
	start, err := pbutil.RestTimeOfDay(in.GetEnd(), int64(math.Ceil(totalDuration)))
	if err != nil {
		return nil, err
	}

	event, err := ent.EventFromPb(&pbservices.Event{
		Start:   start,
		OrderId: order.GetPurchaseId(),
		End:     in.GetEnd(),
		Dow:     in.GetDow(),
	}, claimantId, svc.CalendarId)

	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "EventFromPb: %v", err)
	}

	return insertEvent(ctx, db, claimant.GetUser(), event)
}

func containsCategoryId(i uint64, sl map[uint64]*pbsemantic.Category) bool {
	for s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

func totalMaxVolWeight(ctx context.Context, p *pbcheckout.Purchase, prods pbproducts.ViewerClient) (float64, error) {
	var itemIds []uint64
	for id := range p.GetItems() {
		itemIds = append(itemIds, id)
	}
	resp, err := prods.RetrieveProductsMaxVolWeight(ctx, &pbproducts.RetrieveProductsRequest{Ids: itemIds})
	if err != nil {
		return 0, err
	}
	var totalMaxVolWeight float64
	for prodId, maxVolWeight := range resp.GetMaxVolWeights() {
		totalMaxVolWeight += float64(p.GetItems()[prodId]) * maxVolWeight
	}
	return totalMaxVolWeight, nil
}
