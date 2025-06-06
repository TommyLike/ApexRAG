package injector

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tommylike/apaxrag/configs"

	"github.com/gin-gonic/gin"
	"github.com/tommylike/apaxrag/injector/api"

	"github.com/tommylike/apaxrag/pkg/logger"
)

func NewApiInjector(
	engine *gin.Engine,
) *ApiInjector {
	return &ApiInjector{
		engine: engine,
	}
}

type ApiInjector struct {
	engine *gin.Engine
}

func initHttpServer(ctx context.Context, opts ...api.Option) (func(), error) {
	var o api.Options
	for _, opt := range opts {
		opt(&o)
	}

	configs.MustLoad(o.ConfigFile)
	configs.PrintWithJSON()

	logger.WithContext(ctx).Printf("starting server，run mode：%s，ver：%s，pid：%d", configs.C.RunMode, o.Version, os.Getpid())

	loggerCleanFunc, err := InitLogger()
	if err != nil {
		return nil, err
	}

	monitorCleanFunc := api.InitMonitor(ctx)

	injector, injectorCleanFunc, err := BuildApiInjector()
	if err != nil {
		return nil, err
	}
	httpServerCleanFunc := api.InitHTTPServer(ctx, injector.engine)

	return func() {
		httpServerCleanFunc()
		injectorCleanFunc()
		monitorCleanFunc()
		loggerCleanFunc()
	}, nil
}

func RunServer(ctx context.Context, opts ...api.Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := initHttpServer(ctx, opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("catched signal[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.WithContext(ctx).Infof("stopping server")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
