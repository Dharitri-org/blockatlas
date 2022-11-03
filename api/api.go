package api

import (
	"github.com/Dharitri-org/blockatlas/config"
	_ "github.com/Dharitri-org/blockatlas/docs"
	"github.com/Dharitri-org/blockatlas/platform"
	"github.com/Dharitri-org/blockatlas/services/tokenindexer"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupPlatformAPI(router gin.IRouter) {
	for _, api := range platform.Platforms {
		RegisterTransactionsAPI(router, api)
		RegisterTokensAPI(router, api)
		RegisterStakeAPI(router, api)
		RegisterBlockAPI(router, api)
	}
	for _, api := range platform.CollectionsAPIs {
		RegisterCollectionsAPI(router, api)
	}

	RegisterBatchAPI(router)
	RegisterBasicAPI(router)
}

func SetupTokensIndexAPI(router gin.IRouter, instance tokenindexer.Instance) {
	RegisterTokensIndexAPI(router, instance)
}

func SetupSwaggerAPI(router gin.IRouter) {
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetupMetrics(router gin.IRouter) {
	router.GET(config.Default.Metrics.Path, gin.WrapH(promhttp.Handler()))
}
