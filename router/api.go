package router

import (
	"PushServer/api/v1"

	"github.com/gin-gonic/gin"
)

func apiRouter(r *gin.Engine) {
	//推入
	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.POST("/push", v1.Index)
	}
}
