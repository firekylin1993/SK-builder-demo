//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"SK-builder/internal/biz"
	"SK-builder/internal/conf"
	"SK-builder/internal/data"
	"SK-builder/internal/infrastructure"
	"SK-builder/internal/server"
	"SK-builder/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

func wireProvider(context.Context, *conf.Server, log.Logger) (func(), error) {
	panic(wire.Build(infrastructure.ProviderSet, NewProvider))
}
