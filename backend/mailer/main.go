package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/mailer/mailerconf"
	"github.com/athomecomar/athome/backend/mailer/server"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbmailer"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"

	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Mailer
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listening on port " + port)

	smtpConfig := mailerconf.GetSMTP_CONFIG()
	d := gomail.NewDialer(smtpConfig.Host, smtpConfig.Port, smtpConfig.User, smtpConfig.Password)

	h := hermes.Hermes{
		Theme: &hermes.Flat{},
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: mailerconf.GetHERMES_NAME(),
			Link: mailerconf.GetHERMES_URL(),
			Logo: mailerconf.GetHERMES_LOGO(),
		},
	}

	s := grpc.NewServer()
	pbmailer.RegisterMailerServer(s, &server.Server{Sender: h, Dialer: d})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
