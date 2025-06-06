/*
Package main apaxrag

Swagger：https://github.com/swaggo/swag#declarative-comments-format

Usage：

	go get -u github.com/swaggo/swag/main/swag
	swag init --generalInfo ./main/main.go --output ./app/interfaces/api/swagger
*/
package main

import (
	"context"
	"os"

	"github.com/tommylike/apaxrag/injector/api"

	"github.com/tommylike/apaxrag/injector"

	"github.com/tommylike/apaxrag/pkg/logger"
	"github.com/urfave/cli/v2"
)

// VERSION You can specify the version number by compiling：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "0.5.0"

//go:generate go env -w GO111MODULE=on
//go:generate go mod tidy
//go:generate go mod download

// @title apaxrag
// @version 0.2.0
// @description Apax RAG Application.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
func main() {
	logger.SetVersion(VERSION)
	ctx := logger.NewTagContext(context.Background(), "__main__")
	app := cli.NewApp()
	app.Name = "ApaxRAG"
	app.Version = VERSION
	app.Usage = "ApaxRAG Application."
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Error(err)
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "Run web server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "server config files(.json,.yaml,.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return injector.RunServer(ctx,
				api.SetConfigFile(c.String("conf")),
				api.SetVersion(VERSION))
		},
	}
}
