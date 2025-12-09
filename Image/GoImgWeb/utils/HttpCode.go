package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToLogin(c *gin.Context) {
	cookie, err := c.Cookie("yliken_cookie")
	if err != nil {
		htmlContent := `<script>alert("请先登录");window.location.href="/login"</script>`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	}
	fmt.Println(cookie)
}
