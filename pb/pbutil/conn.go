package pbutil

import (
	"context"

	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbidentifier"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbmailer"
	"github.com/athomecomar/athome/pb/pbmessager"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func conn(ctx context.Context, host string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, host)
	}
	return conn, nil
}

func ConnMessager(ctx context.Context) (pbmessager.MessagesClient, func() error, error) {
	host := pbconf.Messager.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbmessager.NewMessagesClient(conn)
	return c, conn.Close, nil
}

func ConnNotifier(ctx context.Context) (pbnotifier.NotificationsClient, func() error, error) {
	host := pbconf.Notifier.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbnotifier.NewNotificationsClient(conn)
	return c, conn.Close, nil
}

func ConnAuth(ctx context.Context) (pbauth.AuthClient, func() error, error) {
	host := pbconf.Auth.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbauth.NewAuthClient(conn)
	return c, conn.Close, nil
}

func ConnMailer(ctx context.Context) (pbmailer.MailerClient, func() error, error) {
	host := pbconf.Mailer.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbmailer.NewMailerClient(conn)
	return c, conn.Close, nil
}

func ConnIdentifier(ctx context.Context) (pbidentifier.IdentifierClient, func() error, error) {
	host := pbconf.Mailer.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbidentifier.NewIdentifierClient(conn)
	return c, conn.Close, nil
}

func ConnUsersConfig(ctx context.Context) (pbusers.ConfigClient, func() error, error) {
	host := pbconf.Users.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbusers.NewConfigClient(conn)
	return c, conn.Close, nil
}

func ConnUsersViewer(ctx context.Context) (pbusers.ViewerClient, func() error, error) {
	host := pbconf.Users.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbusers.NewViewerClient(conn)
	return c, conn.Close, nil
}

func ConnImages(ctx context.Context) (pbimages.ImagesClient, func() error, error) {
	host := pbconf.Images.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbimages.NewImagesClient(conn)
	return c, conn.Close, nil
}

func ConnProductsCreator(ctx context.Context) (pbproducts.CreatorClient, func() error, error) {
	host := pbconf.Products.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbproducts.NewCreatorClient(conn)
	return c, conn.Close, nil
}
func ConnProductsViewer(ctx context.Context) (pbproducts.ViewerClient, func() error, error) {
	host := pbconf.Products.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbproducts.NewViewerClient(conn)
	return c, conn.Close, nil
}

func ConnServicesViewer(ctx context.Context) (pbservices.ViewerClient, func() error, error) {
	host := pbconf.Services.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbservices.NewViewerClient(conn)
	return c, conn.Close, nil
}

func ConnServicesRegister(ctx context.Context) (pbservices.RegisterClient, func() error, error) {
	host := pbconf.Services.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbservices.NewRegisterClient(conn)
	return c, conn.Close, nil
}

func ConnAddresses(ctx context.Context) (pbaddress.AddressesClient, func() error, error) {
	host := pbconf.Addresses.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbaddress.NewAddressesClient(conn)
	return c, conn.Close, nil
}

func ConnSemanticProducts(ctx context.Context) (pbsemantic.ProductsClient, func() error, error) {
	host := pbconf.Semantic.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbsemantic.NewProductsClient(conn)
	return c, conn.Close, nil
}

func ConnSemanticServiceProviders(ctx context.Context) (pbsemantic.ServiceProvidersClient, func() error, error) {
	host := pbconf.Semantic.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbsemantic.NewServiceProvidersClient(conn)
	return c, conn.Close, nil
}

func ConnSemanticMerchants(ctx context.Context) (pbsemantic.MerchantsClient, func() error, error) {
	host := pbconf.Semantic.GetHost()
	conn, err := conn(ctx, host)
	if err != nil {
		return nil, nil, err
	}
	c := pbsemantic.NewMerchantsClient(conn)
	return c, conn.Close, nil
}
