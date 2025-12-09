package config

type LoginInfo struct {
	Username   string `json:"username"`
	StudentNum string `json:"studentNum"`
}
type JwtForm struct {
	Username   string `json:"username"`
	StudentNum string `json:"studentNum"`
	ImgUrl     string `json:"imgUrl"`
}

var WhiteStuNum = []string{
	"5013230222",        //admin
	"5013240225",        //李静蕾
	"4810230103",        //何浩
	"5013240129",        //齐一蕾
	"5013240101",        //陈栋泽
	"5013240223",        //康佳悦
	"5013240112",        //吴成信
	"5013240136",        //游梦瑶
	"5013240115",        //张润潇
	"5013240103",        //付东朔
	"this is beiyong 1", //备用1
	"this is beiyong 2", //备用2
	"this is beiyong 3", //备用3
	"this is beiyong 4", //备用4
}
