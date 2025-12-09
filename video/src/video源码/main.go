package main

import (
	"SNCTF_SQL_Video/hander"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Static("/static", "./static")
	engine.LoadHTMLGlob("static/html/*")
	store := cookie.NewStore([]byte("akeycalledRedBean"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24 * 30,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	engine.Use(sessions.Sessions("jokeryliken", store))
	//2025年8月4日10:23:40 完工
	engine.GET("/", hander.GetIndexfunc)
	//2025年8月4日10:23:37 完工
	engine.GET("/getvideo", hander.GetVideo)
	//2025年8月4日10:23:33 完工
	engine.GET("/login", hander.GetLogin)
	//2025年8月4日10:23:26 完工
	engine.POST("/login", hander.PostLogin)
	//2025年8月3日21:38:09 register 路由已完工
	engine.POST("/register", hander.PostRegister)
	//2025年8月4日10:42:34 完工
	engine.POST("/resetrequest", hander.PostResetrequest)
	engine.POST("/resetconfirm", hander.PostResetconfirm)
	engine.GET("/getflag", hander.Getflag)
	engine.Run(":9090")
}
