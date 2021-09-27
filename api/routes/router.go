package routes

import (
	"devops/api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.NoRoute((&controllers.Base{}).HandlerNoRouter)

	// metrics touter
	root := r.Group("")
	{
		root.GET("/metrics", (&controllers.Metrics{}).Metrics)
	}
	// v1 version
	v1 := r.Group("/v1")
	{
		// user
		user := v1.Group("/user")
		{
			user.POST("/login", (&controllers.User{}).Login)
			user.Use((&controllers.Auth{}).Authorize())
			user.GET("", (&controllers.User{}).All)
			user.GET("/:id", (&controllers.User{}).One)
			user.POST("", (&controllers.User{}).Add)
			user.PUT("", (&controllers.User{}).Update)
			user.PUT("/password", (&controllers.User{}).UpdatePassword)
			user.DELETE("/:id", (&controllers.User{}).Delete)
			user.POST("/logout", (&controllers.User{}).Logout)
			user.GET("/getUserInfo", (&controllers.User{}).GetUserInfo)
		}

		// export
		export := v1.Group("/export")
		{
			export.GET("/type/:type", (&controllers.Export{}).Node)
			export.Use((&controllers.Auth{}).Authorize())
			export.GET("", (&controllers.Export{}).All)
			export.GET("/:id", (&controllers.Export{}).One)
			export.POST("", (&controllers.Export{}).Add)
			export.PUT("", (&controllers.Export{}).Update)
			export.DELETE("/:id", (&controllers.Export{}).Delete)
		}

		v1.Use((&controllers.Auth{}).Authorize())

		// env
		env := v1.Group("/env")
		{
			env.GET("", (&controllers.Environment{}).All)
			env.GET("/:id", (&controllers.Environment{}).One)
			env.POST("", (&controllers.Environment{}).Add)
			env.PUT("", (&controllers.Environment{}).Update)
			env.DELETE("/:id", (&controllers.Environment{}).Delete)
		}

		// service
		service := v1.Group("/service")
		{
			service.GET("/all", (&controllers.Service{}).All)
			service.DELETE("/:env/:path/:type/:id", (&controllers.Service{}).DeleteNode)
			service.PUT("/outAddress", (&controllers.Service{}).SetOutAddress)
		}

		// promjob
		promjob := v1.Group("/promjob")
		{
			promjob.GET("", (&controllers.PromJob{}).All)
			promjob.GET("/:id", (&controllers.PromJob{}).One)
			promjob.POST("", (&controllers.PromJob{}).Add)
			promjob.PUT("", (&controllers.PromJob{}).Update)
			promjob.DELETE("/:id", (&controllers.PromJob{}).Delete)
		}

		// prom
		prom := v1.Group("/prom")
		{
			prom.GET("/conf", (&controllers.Prom{}).BuildConfiger)
			prom.POST("/reload", (&controllers.Prom{}).Reload)
		}

		// globalSetting
		globalSetting := v1.Group("/globalSetting")
		{
			globalSetting.GET("/:env", (&controllers.GlobalSetting{}).GetGlobalSetting)
			globalSetting.PUT("", (&controllers.GlobalSetting{}).SetGlobalSetting)
		}

		// typSetting
		typSetting := v1.Group("/typSetting")
		{
			typSetting.GET("/:env/:typ", (&controllers.TypSetting{}).GetTypSetting)
			typSetting.PUT("", (&controllers.TypSetting{}).SetTypSetting)
		}

		// appSetting
		appSetting := v1.Group("/appSetting")
		{
			appSetting.GET("/:env/:typ/:id", (&controllers.AppSetting{}).GetAppSetting)
			appSetting.PUT("", (&controllers.AppSetting{}).SetAppSetting)
		}

		// controller
		ctrl := v1.Group("/controller")
		{
			ctrl.PUT("", (&controllers.Controller{}).Controller)
		}

		// metricsProm
		metricsProm := v1.Group("/metricsProm")
		{
			metricsProm.GET("hostinfo", (&controllers.MetricsProm{}).HostInfo)
		}
	}
}
