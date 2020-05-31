package mailerconf

import (
	"log"
	"net/mail"
	"strconv"

	"github.com/athomecomar/envconf"
)

func GetSMTP_CONFIG() *smtpConfig {
	smtpPort := envconf.Get("SMTP_PORT", "1025")
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatalf("couldnt cast port env variable to int64: %v", err)
	}
	return &smtpConfig{
		Port:     port,
		User:     envconf.Get("SMTP_USER", ""),
		Host:     envconf.Get("SMTP_HOST", "localhost"),
		Password: envconf.Get("SMTP_PASSWORD", ""),
		Sender:   GetSMTP_SENDER(),
	}
}

func GetSMTP_SENDER_NAME() (name string) {
	switch envconf.GetENV() {
	case envconf.Development:
		name = "atHome development"
	case envconf.Staging:
		name = "atHome staging"
	case envconf.Production:
		name = "atHome production"
	}
	return
}

func GetSMTP_SENDER_EMAIL() string {
	return "athome@athome.com.ar"
}

func GetSMTP_SENDER() mail.Address {
	return mail.Address{
		Name:    GetSMTP_SENDER_NAME(),
		Address: GetSMTP_SENDER_EMAIL(),
	}
}

type smtpConfig struct {
	Host     string
	Port     int
	Sender   mail.Address
	User     string
	Password string
}
