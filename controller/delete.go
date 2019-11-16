package controller

import (
	"ginarticle/model"
	"ginarticle/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	db := mysql.Dbinit()
	id := c.Param("id")
	var deletearticle model.Article
	db.Where(id).Delete(&deletearticle)
	c.Redirect(http.StatusMovedPermanently, "/admin")
}
