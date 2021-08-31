package routes

import (
	"devops/api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.NoRoute((&controllers.Base{}).HandlerNoRouter)
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
			export.GET("/", (&controllers.Export{}).All)
			export.GET("/:id", (&controllers.Export{}).One)
			export.POST("/", (&controllers.Export{}).Add)
			export.PUT("/", (&controllers.Export{}).Update)
			export.DELETE("/:id", (&controllers.Export{}).Delete)
			export.GET("/type/:type", (&controllers.Export{}).Node)
		}

		// export
		env := v1.Group("/env")
		{
			env.GET("/", (&controllers.Environment{}).All)
			env.GET("/:id", (&controllers.Environment{}).One)
			env.POST("/", (&controllers.Environment{}).Add)
			env.PUT("/", (&controllers.Environment{}).Update)
			env.DELETE("/:id", (&controllers.Environment{}).Delete)
		}

		// service
		service := v1.Group("/service")
		{
			service.GET("/all", (&controllers.Service{}).All)
			service.DELETE("/:env/:path/:type/:id", (&controllers.Service{}).Delete)
		}

		// export
		promjob := v1.Group("/promjob")
		{
			promjob.GET("/", (&controllers.PromJob{}).All)
			promjob.GET("/:id", (&controllers.PromJob{}).One)
			promjob.POST("/", (&controllers.PromJob{}).Add)
			promjob.PUT("/", (&controllers.PromJob{}).Update)
			promjob.DELETE("/:id", (&controllers.PromJob{}).Delete)
		}

		// prom
		prom := v1.Group("/prom")
		{
			prom.GET("/conf", (&controllers.Prom{}).BuildConfiger)
			prom.POST("/reload", (&controllers.Prom{}).Reload)
		}
	}
}
