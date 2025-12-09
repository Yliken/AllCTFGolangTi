package utils

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

func GetVideo() string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get("http://api.yujn.cn/api/xjj.php?type=video")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Header.Get("Location")
}

// 返回值id == 0 的时候表示用户不存在
func SearchUserId(username string, db *gorm.DB) int {
	fmt.Println("search")
	var id int
	db.Raw("SELECT id FROM userinfos WHERE username = ?", username).Scan(&id)
	return id
}
