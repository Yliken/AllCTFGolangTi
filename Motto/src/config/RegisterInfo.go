package config

type RegisterInfo struct {
	UserID   int    `gorm:"primaryKey;autoIncrement" form:"user_id"`
	Nickname string `gorm:"type:varchar(25);size:500;not null" form:"nickname" binding:"required"`
	Username string `gorm:"type:varchar(25);unique" form:"username" binding:"required"`
	Password string `gorm:"type:varchar(50);not null" form:"password" binding:"required"`
}

type LoginInfo struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
