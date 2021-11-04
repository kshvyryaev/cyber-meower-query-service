//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg"
	controller "github.com/kshvyryaev/cyber-meower-query-service/pkg/controller/http"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/query"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/search"
)

func InitializeHttpServer() (*controller.HttpServer, func(), error) {
	panic(wire.Build(
		pkg.ProvideConfig,
		pkg.ProvideZap,
		search.ProvideElastic,
		search.ElasticMeowRepositorySet,
		query.ProvideMeowQueries,
		controller.ProvideMeowController,
		controller.ProvideErrorHandlerMiddleware,
		controller.ProvideRecoveryHandlerMiddleware,
		controller.ProvideHttpServer,
	))
}
