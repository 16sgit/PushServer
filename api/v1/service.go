package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "It work")
}
