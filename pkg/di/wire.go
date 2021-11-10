//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/contract"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/controller/http"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/controller/http/middleware"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/search"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/usecase"
)

func InitializeHttpServer() (*http.HttpServer, func(), error) {
	panic(wire.Build(
		pkg.ProvideConfig,
		pkg.ProvideZap,
		search.ProvideElastic,
		search.ProvideElasticMeowRepository,
		wire.Bind(new(contract.MeowRepository), new(*search.ElasticMeowRepository)),
		usecase.ProvideMeowUsecase,
		wire.Bind(new(contract.MeowUsecase), new(*usecase.MeowUsecase)),
		http.ProvideMeowController,
		middleware.ProvideErrorHandlerMiddleware,
		middleware.ProvideRecoveryHandlerMiddleware,
		http.ProvideHttpServer,
	))
}
