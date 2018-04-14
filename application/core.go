package application

import "github.com/m0t0k1ch1/metamask-login-sample/domain"

type Core struct {
	config    *Config
	container *domain.Container
}

func NewCore(config *Config, container *domain.Container) *Core {
	return &Core{
		config:    config,
		container: container,
	}
}
