package config

type MottoInfo struct {
	MottoId  int    `gorm:"primaryKey;autoIncrement" form:"user_id"`
	NickName string `gorm:"type:varchar(25);size:500" form:"nickname" binding:"required"`
	Motto    string `gorm:"type:varchar(50);size:500;not null;" form:"password" binding:"required"`
}
type PostMotto struct {
	Motto string `form:"motto" binding:"required"`
}
type PostNickname struct {
	NickName string `form:"nickname" binding:"required"`
}
