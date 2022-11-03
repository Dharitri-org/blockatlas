package endpoint

import (
	"github.com/Dharitri-org/blockatlas/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"build":  internal.Build,
		"date":   internal.Date,
	})
}
