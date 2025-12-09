package main

import (
	"SummerVactionSQL/config"
	"SummerVactionSQL/database"
	"SummerVactionSQL/function"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
)

var db *gorm.DB

func init() {
	var err error
	db, err = database.Init()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&config.RegisterInfo{}, &config.MottoInfo{})
	if err != nil {
		panic(err)
	}
}

func main() {
	// 设置为 release 模式，禁用 debug 输出
	gin.SetMode(gin.ReleaseMode)

	// 禁用 Gin 所有输出（包括启动信息）
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard

	// 使用 gin.New() 代替 gin.Default()，避免默认加 Logger
	engine := gin.New()

	// 如需防止崩溃可添加 recovery 中间件（不会打印）
	engine.Use(gin.Recovery())
	engine.Static("/static", "./static")
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		var mottos []config.MottoInfo
		db.Find(&mottos)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Mottos": mottos,
		})
	})
	engine.POST("/", func(c *gin.Context) {
		cookie, err := c.Cookie("yliken_cookie")
		var req config.PostMotto
		c.ShouldBind(&req)
		var nickname string
		if err != nil {
			htmlContent := `<script>alert("请先登录后再写Motto");window.location.href="/login"</script>`
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
			return
		} else {
			jwt, _ := function.DeJwt(cookie)
			nickname = (*jwt)["nickname"].(string)

		}
		if err := function.AddMotto(db, nickname, req.Motto); err != nil {
			htmlContent := `<script>alert("Motto添加失败!");window.location.href="/"</script>`
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
			return
		}
		htmlContent := `<script>alert("Motto添加成功!");window.location.href="/"</script>`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
	//登录页面
	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	engine.POST("/login", func(c *gin.Context) {
		var loginInfo config.LoginInfo
		var result config.RegisterInfo
		c.ShouldBind(&loginInfo)
		tx := db.Where("username = ?", loginInfo.Username).First(&result)
		if tx.Error != nil {
			c.JSON(200, gin.H{
				"success": false,
				"message": "用户名或密码输入错误",
			})
		}
		if loginInfo.Password == result.Password {
			jwt, _ := function.EnJwt(result.Username, result.Nickname)
			c.SetCookie("yliken_cookie", jwt, 1800, "/", "", false, true)
			c.JSON(200, gin.H{
				"success": true,
			})
		} else {
			c.JSON(200, gin.H{
				"success": false,
				"message": "用户名或密码输入错误",
			})
		}
	})
	//注册页面
	engine.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})
	engine.POST("/register", func(c *gin.Context) {
		var registerInfo config.RegisterInfo
		c.ShouldBind(&registerInfo)
		fmt.Println(registerInfo.Username)
		tx := db.Where("username = ?", registerInfo.Username).First(&config.RegisterInfo{})
		if tx.Error == nil {
			if tx.RowsAffected != 0 {
				c.JSON(200, gin.H{
					"success": false,
					"message": "该用户已存在",
				})
				return
			}
		}
		fmt.Println(registerInfo)
		db.Save(&registerInfo)
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	//个人信息展示页面
	engine.GET("/myinfo", func(c *gin.Context) {
		cookie, err := c.Cookie("yliken_cookie")
		if err != nil {
			htmlContent := `<script>alert("请先登录");window.location.href="/login"</script>`
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
			return
		}
		jwt, _ := function.DeJwt(cookie)
		nickname := (*jwt)["nickname"].(string)
		username := (*jwt)["username"].(string)
		c.HTML(http.StatusOK, "myinfo.html", gin.H{
			"nickname": nickname,
			"username": username,
		})
	})
	//更新nickname
	engine.POST("/changeNickName", func(c *gin.Context) {
		cookie, err := c.Cookie("yliken_cookie")
		if err != nil {
			htmlContent := `<script>alert("请先登录");window.location.href="/login"</script>`
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
			return
		}
		var newNickname config.PostNickname
		c.ShouldBind(&newNickname)
		jwt, _ := function.DeJwt(cookie)
		username := (*jwt)["username"].(string)
		db.Model(&config.RegisterInfo{}).Where("username = ?", username).Update("nickname", newNickname.NickName)
		enJwt, _ := function.EnJwt(username, newNickname.NickName)
		c.SetCookie("yliken_cookie", enJwt, 1800, "/", "", false, true)
		htmlContent := `<script>alert("修改成功");window.location.href="/myinfo"</script>`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
		return
	})
	//查看我的motto
	engine.GET("/mymottos", func(c *gin.Context) {
		cookie, err := c.Cookie("yliken_cookie")
		if err != nil {
			htmlContent := `<script>alert("请先登录");window.location.href="/login"</script>`
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
			return
		}
		jwt, _ := function.DeJwt(cookie)
		nickname := (*jwt)["nickname"].(string)
		var mottos []config.MottoInfo
		sql := "SELECT * FROM motto_infos where nick_name = '" + nickname + "'"
		fmt.Println(sql)
		tx := db.Raw(sql).Scan(&mottos)
		if tx.Error != nil {
			var motto config.MottoInfo
			motto.NickName = nickname
			motto.Motto = "db error"
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Mottos": motto,
			})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Mottos": mottos,
		})
	})

	engine.Run(":9090")
}
