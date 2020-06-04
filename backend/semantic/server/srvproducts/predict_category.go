package srvproducts

import (
	"context"
	"io"
	"time"

	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/predictor"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) PredictCategory(srv pbsemantic.Products_PredictCategoryServer) error {
	ctx := srv.Context()
	context.WithTimeout(ctx, 300*time.Second)

	db, err := server.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()
	return s.predictCategory(ctx, db, srv)
}

func (s *Server) predictCategory(ctx context.Context, db *sqlx.DB, srv pbsemantic.Products_PredictCategoryServer) error {
	predictor := predictor.NewPredictor(ctx)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default: // no-op
		}

		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		cat, err := predictor.Predict(in.GetTitle())
		if err != nil {
			return status.Errorf(xerrors.Internal, "predictor.Predict: %v", err)
		}

		err = srv.Send(&pbsemantic.PredictCategoryResponse{Score: 1, Category: &pbsemantic.Category{Name: cat.CategoryName}})
		if err != nil {
			return err
		}
	}
}
