//go:build wireinject
// +build wireinject

package main

import (
	"chatbot/internal/conf"
	"chatbot/internal/server"
	"chatbot/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(string, *conf.Bootstrap_DING, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
