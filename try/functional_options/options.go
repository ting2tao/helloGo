package ctlFunctionalOptions

import (
	"crypto/tls"
	"time"
)

type FunctionalOptionsCtl struct {
}

type Server struct {
	Addr string
	Port int
	Conf *Config
}

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

//Using the default configuratrion
