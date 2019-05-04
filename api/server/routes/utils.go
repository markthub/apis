package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markthub/apis/api/pkg/version"
	"github.com/markthub/apis/api/server/utils"
)

const pagesize = 25

// Version returns the APIs version
func Version(c *gin.Context) {
	utils.Response(c, http.StatusOK, map[string]string{
		"version": version.LongVersion(),
	})
}
