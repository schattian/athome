package server

import (
	"github.com/athomecomar/athome/backend/mailer/pb/pbmailer"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type Server struct {
	pbmailer.UnimplementedMailerServer
	Dialer *gomail.Dialer

	Sender hermes.Hermes
}
