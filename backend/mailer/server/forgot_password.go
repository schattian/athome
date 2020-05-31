package server

import (
	"context"
	"net/mail"

	"github.com/athomecomar/athome/backend/mailer/email"
	"github.com/athomecomar/athome/backend/mailer/mailerconf"
	"github.com/athomecomar/athome/backend/mailer/pb/pbmailer"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
)

func (s *Server) ForgotPassword(ctx context.Context, in *pbmailer.ForgotPasswordRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.forgotPassword(ctx, in)
}

func (s *Server) forgotPassword(ctx context.Context, in *pbmailer.ForgotPasswordRequest) (*emptypb.Empty, error) {
	tokenByRole := make(map[string]string)
	for _, tkUser := range in.GetTokenizedUsers() {
		tokenByRole[tkUser.GetRole()] = tkUser.GetToken()
	}

	e := email.ForgotPassword(in.GetName(), tokenByRole)
	htmlBody, err := s.Sender.GenerateHTML(e)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "hermes.GenerateHTML: %v", err)
	}
	plainBody, err := s.Sender.GeneratePlainText(e)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "hermes.GenerateHTML: %v", err)
	}

	msg := s.forgotPasswordMessage(in)
	msg.SetBody("text/plain", plainBody)
	msg.AddAlternative("text/html", htmlBody)

	err = s.Dialer.DialAndSend(msg)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "dialer.DialAndSend: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) forgotPasswordMessage(in *pbmailer.ForgotPasswordRequest) *gomail.Message {
	m := gomail.NewMessage()
	from := mailerconf.GetSMTP_INSTITUTIONAL_SENDER()
	to := mail.Address{
		Name:    in.GetName(),
		Address: in.GetEmail(),
	}
	m.SetHeaders(email.ForgotPasswordHeaders(from, to))
	return m
}
