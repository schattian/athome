package server

// func RetrieveLatestInProgressOrder(ctx context.Context, db *sqlx.DB, accessToken string) (*ent.Draft, error) {
// 	c, closer, err := pbutil.ConnAuth(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer closer()
// order, err := retrieveLatestInProgressOrder(ctx, db, c, accessToken)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return draft, nil
// }

// func retrieveLatestOrder(ctx context.Context, db *sqlx.DB, auth pbauth.AuthClient, accessToken string) (*ent.Draft, error) {
// 	userId, err := pbutil.GetUserFromAccessToken(ctx, auth, accessToken)
// 	if err != nil {
// 		return nil, err
// 	}

// 	draft, err := ent.FindOrCreateDraft(ctx, db, userId)
// 	if err != nil {
// 		return nil, status.Errorf(xerrors.Internal, "FindOrCreateDraft: %v", err)
// 	}

// 	return draft, nil
// }
