package controller

import (
	"ginarticle/model"
	"ginarticle/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	db := mysql.Dbinit()
	var querarticle []model.Article
	db.Find(&querarticle)
	c.HTML(http.StatusOK, "adminindex.html", querarticle)

}

func QueryOne(c *gin.Context) {
	db := mysql.Dbinit()
	var onearticle model.Article
	id := c.Param("id")
	db.Find(&onearticle, id)
	c.HTML(http.StatusOK, "contentt.html", onearticle)
}
