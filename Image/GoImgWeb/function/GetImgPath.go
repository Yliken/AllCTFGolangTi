package function

import (
	"io"
	"net/http"
	"regexp"
)

type ImgPath struct {
	Path []string `json:"img"`
}

func GetImgPath(url string) []string {
	var wrongjps = []string{"https://yliken-images-test.oss-cn-beijing.aliyuncs.com/images/error.png"}
	resp, err := http.Get(url)
	if err != nil {
		return wrongjps
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	reg := regexp.MustCompile(`https?://[^\s"']+\.(jpg|jpeg|png|gif|bmp|ico)`)

	matches := reg.FindAllString(string(body), -1)

	return matches
}
