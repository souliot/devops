package routes

import (
	"devops/api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// metrics touter
	root := r.Group("/")
	{
		root.GET("/metrics", (&controllers.Metrics{}).Metrics)
	}
	// v1 version
	v1 := r.Group("/v1")
	{
		// export
		export := v1.Group("/export")
		{
			export.POST("/", (&controllers.Export{}).Add)
			export.GET("/node", (&controllers.Export{}).Node)
		}
	}
}
