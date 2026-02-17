package config

import (
	"os"
)

type SMTPConfig struct {
	Host string
	Port string
	User string
	Pass string
	From string
}

var SMTP SMTPConfig

func LoadSMTPConfig() {
	SMTP = SMTPConfig{
		Host: os.Getenv("SMTP_HOST"),
		Port: os.Getenv("SMTP_PORT"),
		User: os.Getenv("SMTP_USER"),
		Pass: os.Getenv("SMTP_PASS"),
		From: os.Getenv("SMTP_FROM"),
	}
}
