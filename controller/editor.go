package controller

import (
	"ginarticle/model"
	"ginarticle/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowEditor(c *gin.Context) {
	db := mysql.Dbinit()
	var showEditor model.Article
	id := c.Param("id")
	db.Find(&showEditor, id)

	c.HTML(http.StatusOK, "editor.html", showEditor)
}

func Editor(c *gin.Context) {
	db := mysql.Dbinit()
	id := c.Param("id")
	title := c.PostForm("title")
	content := c.PostForm("content")

	db.Model(&model.Article{}).Where(id).Update(&model.Article{Title: title, Content: content})
	c.Redirect(http.StatusMovedPermanently, "/admin")
}
