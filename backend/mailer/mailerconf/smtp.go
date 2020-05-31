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
	}
}

func GetSMTP_INSTITUTIONAL_SENDER_NAME() (name string) {
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

func GetSMTP_INSTITUTIONAL_SENDER_EMAIL() string {
	return "athome@athome.com.ar"
}

func GetSMTP_INSTITUTIONAL_SENDER() mail.Address {
	return mail.Address{
		Name:    GetSMTP_INSTITUTIONAL_SENDER_NAME(),
		Address: GetSMTP_INSTITUTIONAL_SENDER_EMAIL(),
	}
}

type smtpConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}
