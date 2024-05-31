package common

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	ConfigRoot struct {
		Debug bool `envconfig:"POSTFORM_DEBUG" default:"false"`
		Smtp  SmtpConfigRoot
	}

	SmtpConfigRoot struct {
		Server      string `envconfig:"POSTFORM_SMTP_SERVER" required:"true"`
		Username    string `envconfig:"POSTFORM_SMTP_USERNAME" required:"true"`
		Password    string `envconfig:"POSTFORM_SMTP_PASSWORD" required:"true"`
		FromAddress string `envconfig:"POSTFORM_SMTP_FROM_ADDRESS" required:"true"`
	}
)

var (
	Config *ConfigRoot
)

func InitConfig() error {
	Config = &ConfigRoot{}
	return errors.Join(
		tryReadEnv(),
	)
}

func tryReadEnv() error {
	godotenv.Load(".env.local")
	if err := envconfig.Process("", Config); err != nil {
		return err
	}

	return nil
}
