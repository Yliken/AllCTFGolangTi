package function

import (
	"SummerVactionSQL/config"
	"gorm.io/gorm"
)

func AddMotto(db *gorm.DB, nickname, Motto string) error {
	var motto config.MottoInfo
	motto.Motto = Motto
	motto.NickName = nickname
	tx := db.Create(&motto)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
