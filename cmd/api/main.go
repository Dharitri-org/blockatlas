package main

import (
	"context"

	"github.com/Dharitri-org/blockatlas/internal/metrics"

	golibsGin "github.com/Dharitri-org/tw-go-libs/network/gin"

	"github.com/Dharitri-org/tw-go-libs/network/middleware"

	"github.com/Dharitri-org/blockatlas/api"
	"github.com/Dharitri-org/blockatlas/config"
	"github.com/Dharitri-org/blockatlas/db"
	_ "github.com/Dharitri-org/blockatlas/docs"
	"github.com/Dharitri-org/blockatlas/internal"
	"github.com/Dharitri-org/blockatlas/platform"
	"github.com/Dharitri-org/blockatlas/services/tokenindexer"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
)

var (
	ctx            context.Context
	cancel         context.CancelFunc
	port, confPath string
	engine         *gin.Engine
	database       *db.Instance
	tokenIndexer   tokenindexer.Instance
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	ctx, cancel = context.WithCancel(context.Background())
	var err error

	internal.InitConfig(confPath)

	if err := middleware.SetupSentry(config.Default.Sentry.DSN); err != nil {
		log.Error(err)
	}

	engine = internal.InitEngine(config.Default.Gin.Mode)
	platform.Init(config.Default.Platform)

	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Log)
	if err != nil {
		log.Fatal(err)
	}

	metrics.Setup(database)

	tokenIndexer = tokenindexer.Init(database)
}

func main() {
	api.SetupTokensIndexAPI(engine, tokenIndexer)
	api.SetupSwaggerAPI(engine)
	api.SetupPlatformAPI(engine)
	api.SetupMetrics(engine)

	golibsGin.SetupGracefulShutdown(ctx, port, engine)
	cancel()
}
