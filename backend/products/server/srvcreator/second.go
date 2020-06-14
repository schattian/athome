package srvcreator

import (
	"context"
	"io"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Second(srv pbproducts.Creator_SecondServer) error {
	ctx := srv.Context()

	db, err := server.ConnDB()
	if err != nil {
		return err
	}

	sem, semCloser, err := pbutil.ConnSemanticProducts(ctx)
	if err != nil {
		return err
	}
	defer semCloser()

	var resp *emptypb.Empty
	var draft *ent.Draft
	var access string
	for {
		select {
		case <-ctx.Done():
			err = srv.SendAndClose(resp)
			if err != nil {
				return err
			}
			return ctx.Err()
		default: // no-op
		}

		in, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(resp)
		}
		if err != nil {
			return err
		}
		err = in.Validate()
		if err != nil {
			return err
		}

		if draft == nil { // first iteration
			access = in.GetAccessToken()
			draft, err = server.RetrieveLatestDraft(ctx, db, access)
			if err != nil {
				return err
			}
			continue // see oneof def
		}

		if draft.Stage != stage.Second {
			return status.Errorf(xerrors.InvalidArgument, "stage expected %v, got %v", stage.Second, draft.Stage)
		}

		resp, err = s.second(ctx, db, sem, in.GetBody(), draft, access)
		if err != nil {
			return err
		}
	}
}

func (s *Server) second(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, in *pbproducts.SecondRequest_Body, draft *ent.Draft, access string) (*emptypb.Empty, error) {
	ln, err := draft.LineById(ctx, db, in.GetDraftLineId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "LineById: %v", err)
	}

	semSrv, err := sem.SetAttributeDatas(ctx)
	if err != nil {
		return nil, err
	}
	authReq := &pbsemantic.SetAttributeDatasRequest{
		Corpus: &pbsemantic.SetAttributeDatasRequest_Authorization_{
			Authorization: &pbsemantic.SetAttributeDatasRequest_Authorization{
				AccessToken: access,
				Entity:      pbutil.ToPbSemanticEntity(ln),
			},
		},
	}
	err = semSrv.Send(authReq)
	if err != nil {
		return nil, err
	}

	for _, att := range in.GetDraftLine().Attributes {
		req := &pbsemantic.SetAttributeDatasRequest{
			Corpus: &pbsemantic.SetAttributeDatasRequest_Data{
				Data: &pbsemantic.AttributeData{
					SchemaId: att.GetSchemaId(),
					Values:   att.GetValues(),
				},
			},
		}
		err = semSrv.Send(req)
		if err != nil {
			return nil, err
		}
		_, err := semSrv.Recv()
		if err != nil {
			return nil, err
		}
	}

	ln = applySecondRequestToDraftLine(in.GetDraftLine(), ln)
	err = storeql.UpdateIntoDB(ctx, db, ln)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func applySecondRequestToDraftLine(in *pbproducts.DraftLineSecond, ln *ent.DraftLine) *ent.DraftLine {
	ln.Price = currency.ToARS(in.GetPrice())
	ln.Stock = in.GetStock()
	return ln
}
