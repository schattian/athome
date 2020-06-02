package signsrv

import (
	"context"

	"github.com/athomecomar/athome/backend/users/internal/userjwt"
	"github.com/athomecomar/athome/backend/users/pb/pbmailer"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ForgotPassword(ctx context.Context, in *pbuser.ForgotPasswordRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()

	conn, err := grpc.Dial(userconf.GetMAILER_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetMAILER_ADDR())
	}
	defer conn.Close()

	return s.forgotPassword(ctx, db, conn, in)
}

func (s *Server) forgotPassword(ctx context.Context, db *sqlx.DB, conn *grpc.ClientConn, in *pbuser.ForgotPasswordRequest) (*emptypb.Empty, error) {
	rows, err := db.QueryxContext(ctx, `SELECT id, role FROM users WHERE email=$1 limit 3`, in.GetEmail())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}

	var tokenizedUsers []*pbmailer.TokenizedUser
	defer rows.Close()
	for rows.Next() {
		var userId uint64
		tokenizedUser := &pbmailer.TokenizedUser{}
		err = rows.Scan(&userId, &tokenizedUser.Role)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "rows.Scan: %v", err)
		}
		tokenizedUser.Token, err = userjwt.CreateForgotToken(userId)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "createForgotToken: %v", err)
		}
		tokenizedUsers = append(tokenizedUsers, tokenizedUser)
	}

	mailer := pbmailer.NewMailerClient(conn)
	_, err = mailer.ForgotPassword(ctx, &pbmailer.ForgotPasswordRequest{TokenizedUsers: tokenizedUsers, Email: in.GetEmail()})
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "rows.Err: %v", err)
	}

	return &emptypb.Empty{}, nil
}
