package server

import (
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type Server struct {
	Dialer *gomail.Dialer

	Sender hermes.Hermes
}
