package query

import (
	"github.com/packagrio/fetchr/pkg/config"
)

type Configuration struct {
	*config.BaseConfiguration
}

func NewConfiguration() *Configuration {
	c := Configuration{}
	c.BaseConfiguration = config.New()

	return &c
}
