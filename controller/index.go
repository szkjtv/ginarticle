package controller

import (
	"ginarticle/model"
	"ginarticle/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	db := mysql.Dbinit()
	var index []model.Article
	db.Find(&index)
	c.HTML(http.StatusOK, "index.html", index)

}

func IndexOne(c *gin.Context) {
	db := mysql.Dbinit()
	var indexone model.Article
	id := c.Param("id")
	db.Find(&indexone, id)
	c.HTML(http.StatusOK, "indexocntent.html", indexone)
}
