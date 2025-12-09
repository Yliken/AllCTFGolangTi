package hander

import (
	"SNCTF_SQL_Video/settings"
	"SNCTF_SQL_Video/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	db, err = settings.Init()
	if err != nil {
		panic(err)
	}
	//err = db.AutoMigrate(&settings.Userinfo{}, &settings.Resetpassword{})
	//if err != nil {
	//	panic(err)
	//}
}

func GetVideo(c *gin.Context) {
	c.JSON(200, gin.H{
		"videoUrl": utils.GetVideo(),
	})
}

func GetLogin(c *gin.Context) {
	c.HTML(200, "user.html", gin.H{})
}
func PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var id int
	err := db.Raw(
		"SELECT id FROM userinfos WHERE username = ? AND password = ?",
		username, password,
	).Scan(&id).Error

	if err != nil {
		c.String(500, "服务器错误")
		return
	}
	if id == 0 {
		c.String(400, "用户名或者密码错误")
		return
	}

	session := sessions.Default(c)
	session.Set("user", username)
	if err := session.Save(); err != nil {
		c.String(500, "登录失败")
		return
	}

	c.String(200, "登录成功")
}

func PostRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.String(400, "用户名和密码不能为空")
		return
	}
	id := utils.SearchUserId(username, db)
	if id == 0 {
		err := db.Exec("INSERT INTO userinfos (username, password) VALUES (?, ?)", username, password).Error
		token := strconv.FormatInt(time.Now().Unix(), 10) + "-" + uuid.New().String()
		db.Exec("INSERT INTO resetpasswords (username, token) VALUES (?, ?)", username, token)
		if err != nil {
			log.Println("插入用户失败:", err)
			c.String(500, "服务器内部错误")
			return
		}
		c.String(200, "注册成功")
		return
	} else {
		c.String(400, "用户名已存在")
	}
}
func GetIndexfunc(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	fmt.Println("user", user)
	if user == nil {
		c.Redirect(302, "/login")
		return
	}
	c.HTML(200, "index.html", gin.H{})
}

func PostResetrequest(c *gin.Context) {
	user := c.PostForm("username")
	// 判断一下 Reset Token 是否存在
	var token string
	err := db.Raw("SELECT token FROM resetpasswords WHERE username=?", user).Scan(&token).Error
	if err != nil {
		fmt.Println(err)
		c.String(400, "用户不存在")
		return
	}
	token = ""
	token = strconv.FormatInt(time.Now().Unix(), 10) + "-" + uuid.New().String()
	fmt.Println("user", user)
	err = db.Exec("UPDATE resetpasswords SET token=? WHERE username=?", token, user).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("token:", token)
	c.String(200, "成功发送重置密码Token")
}

func PostResetconfirm(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(400, "无效的 JSON")
		return
	}

	username, _ := data["username"].(string)
	newPassword, _ := data["newPassword"].(string)

	token, ok := data["token"]

	if !ok {
		c.String(400, "缺少 token 字段")
		return
	}
	fmt.Printf("username = > %T,token = > %T,newpasswd = > %T\n", username, token, newPassword)
	fmt.Println(username, token, newPassword)
	var re string
	db.Raw("SELECT username from resetpasswords where token=?", token).Scan(&re)
	if re == username {
		db.Exec("UPDATE userinfos SET password=? WHERE username=?", newPassword, username)
		c.String(200, "密码重置成功!")
		return
	} else {
		c.String(400, "token错误")

	}
}

func Getflag(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != "admin" {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<script>alert("只有 admin 才可以访问"); window.location.href = "/";</script>`)
		return
	}
	flag := os.Getenv("DASFLAG")
	c.String(200, "Hi admin! this is your flag : %s", flag)
}
