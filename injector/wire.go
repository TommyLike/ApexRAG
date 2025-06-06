//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/tommylike/apaxrag/app/application"
	"github.com/tommylike/apaxrag/app/domain/user"
	userInfra "github.com/tommylike/apaxrag/app/infrastructure/user"
	"github.com/tommylike/apaxrag/app/interfaces/api/handler"
	"github.com/tommylike/apaxrag/app/interfaces/api/router"
	"github.com/tommylike/apaxrag/injector/api"

	"github.com/google/wire"
)

func BuildApiInjector() (*ApiInjector, func(), error) {
	wire.Build(
		// init,
		InitGormDB,
		api.InitGinEngine,

		// domain
		user.NewService,
		// infrastructure
		userInfra.NewRepository,
		//rbacInfra.NewRepository,

		// application
		application.NewUser,

		// handler
		handler.NewHealthCheck,
		handler.NewUser,

		// router
		router.NewRouter,

		// injector
		NewApiInjector,
	)
	return nil, nil, nil
}
