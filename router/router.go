package router

import (
	"ginarticle/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	router.GET("/index", controller.Index)
	router.GET("/indexcontent/:id", controller.IndexOne)

	router.GET("/showregister", controller.ShowRegisters)
	router.POST("/register", controller.RegisterController)

	router.GET("/showadd", controller.Showaddarticle)
	router.POST("/addarticle", controller.Addarticle)

	router.GET("/admin", controller.Admin)
	router.GET("/contentt/:id", controller.QueryOne)

	router.GET("/editor/:id", controller.ShowEditor)
	router.POST("/editort/:id", controller.Editor)

	router.GET("/delete/:id", controller.Delete)

	router.Run(":4040")

}
