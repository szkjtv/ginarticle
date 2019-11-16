package controller

import (
	"ginarticle/model"
	"ginarticle/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Showaddarticle(c *gin.Context) {
	c.HTML(http.StatusOK, "addcontent.html", nil)

}

func Addarticle(c *gin.Context) {
	db := mysql.Dbinit()
	title := c.PostForm("title")
	content := c.PostForm("content")
	newArticle := model.Article{Title: title, Content: content}
	db.Create(&newArticle)
	c.Redirect(http.StatusMovedPermanently, "/admin")
}
