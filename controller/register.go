package controller

import (
	"ginarticle/model"
	"ginarticle/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowRegisters(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
func RegisterController(c *gin.Context) {
	db := mysql.Dbinit()
	account := c.PostForm("account")
	password := c.PostForm("password")
	newUser := model.User{Account: account, Password: password}
	db.Create(&newUser)
	// c.JSON(200, "注册成功")

	// if account == "" || password == "" {

	// 	c.JSON(200, "请输入账号和密码")

	// } else {

	// 	newUser := model.User{Account: account, Password: password}
	// 	db.Create(&newUser)
	// 	c.JSON(200, "注册成功")
	// }

	c.Redirect(http.StatusMovedPermanently, "/login")
	defer db.Close()

}
