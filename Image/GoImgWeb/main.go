package main

import (
	"GoImgWeb/OSSConfig"
	"GoImgWeb/config"
	"GoImgWeb/function"
	"GoImgWeb/utils"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Static("/static", "./static")
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		cookie, err := c.Cookie("yliken_cookie")
		if err != nil {
			utils.ToLogin(c)
			return
		}
		decode, err := function.JwtDecode(cookie)
		if err != nil {
			utils.ToLogin(c)
			return
		}
		imgUrl := (*decode)["ImgUrl"].(string)
		if strings.Contains(imgUrl, "127.0.0.1") || strings.Contains(imgUrl, "localhost") {
			fmt.Println("非法")
			var wrongjps = []string{"https://yliken-images-test.oss-cn-beijing.aliyuncs.com/images/wrongip.jpg", "https://yliken-images-test.oss-cn-beijing.aliyuncs.com/images/error.png"}
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Images": wrongjps,
			})
			return
		}

		path := function.GetImgPath(imgUrl)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Images": path,
		})
	})
	engine.POST("/upload", func(context *gin.Context) {
		cookie, err2 := context.Cookie("yliken_cookie")
		if err2 != nil {
			context.JSON(http.StatusOK, gin.H{
				"success": false,
			})
			return
		}
		jwt, _ := function.JwtDecode(cookie)
		fmt.Println("cookie:", cookie)
		file, err := context.FormFile("file")
		if err != nil {
			panic(err)
		}
		stuNum := (*jwt)["stuNum"].(string)
		ext := filepath.Ext(file.Filename)
		newFilename := base64.StdEncoding.EncodeToString([]byte(stuNum)) + "_" + time.Now().Format("20060102150405") + ext
		path := "SummerOSS/" + newFilename
		open, _ := file.Open()
		if OSSConfig.OssuploadImg(path, open) {
			context.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"success": false,
			})
		}

	})
	engine.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})
	engine.POST("/login", func(context *gin.Context) {
		var login config.LoginInfo
		if err := context.ShouldBindBodyWithJSON(&login); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"success": false,
			})
		}
		if !utils.Contains(config.WhiteStuNum, login.StudentNum) {
			context.JSON(http.StatusOK, gin.H{
				"success": false,
			})
			return
		}

		encode, _ := function.JwtEncode(login.Username, login.StudentNum)
		context.SetCookie("yliken_cookie", encode, 1800, "/", "", false, true)
		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})

	})
	engine.GET("/.env", func(c *gin.Context) {
		c.String(http.StatusOK, "JwtKey=Can not lose RedBean")
	})

	engine.Run(":9090")
}
