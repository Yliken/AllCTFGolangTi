package settings

type Userinfo struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"size:255";not null;unique""`
	Password string `gorm:"size:255";not null""`
}
type Resetpassword struct {
	Username string `gorm:"size:255"`
	Token    string `gorm:"size:255"`
}
