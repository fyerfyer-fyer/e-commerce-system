package svc

import (
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
