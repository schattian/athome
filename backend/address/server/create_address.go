package server

import (
	"context"

	"github.com/athomecomar/athome/backend/address/ent"
	"github.com/athomecomar/athome/backend/address/pb/pbaddress"
	"github.com/athomecomar/athome/backend/address/pb/pbauth"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAddress(ctx context.Context, in *pbaddress.CreateAddressRequest) (*pbaddress.CreateAddressResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()
	return s.createAddress(ctx, db, auth, in)
}

func (s *Server) createAddress(ctx context.Context, db *sqlx.DB, auth pbauth.AuthClient, in *pbaddress.CreateAddressRequest) (*pbaddress.CreateAddressResponse, error) {
	userId, err := GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	addr := pbAddrDataToAddr(in.GetBody())
	addr.UserId = userId

	err = storeql.InsertIntoDB(ctx, db, addr)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}
	resp := &pbaddress.CreateAddressResponse{Metadata: addrToPbAddrData(addr), AddressId: addr.Id}
	return resp, nil
}

func pbAddrDataToAddr(d *pbaddress.AddressData) *ent.Address {
	return &ent.Address{
		Country:   d.GetCountry(),
		Province:  d.GetProvince(),
		Zipcode:   d.GetZipcode(),
		Street:    d.GetStreet(),
		Number:    d.GetNumber(),
		Floor:     d.GetFloor(),
		Latitude:  d.GetLatitude(),
		Longitude: d.GetLongitude(),

		Alias: d.GetAlias(),
	}
}

func addrToPbAddrData(addr *ent.Address) *pbaddress.AddressData {
	return &pbaddress.AddressData{
		Country:   addr.Country,
		Province:  addr.Province,
		Zipcode:   addr.Zipcode,
		Street:    addr.Street,
		Number:    addr.Number,
		Floor:     addr.Floor,
		Latitude:  addr.Latitude,
		Longitude: addr.Longitude,
		Alias:     addr.Alias,
	}
}
