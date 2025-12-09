package function

import (
	"SummerVactionSQL/config"
	"gorm.io/gorm"
)

func FindMotto(db *gorm.DB, nickname string) []config.MottoInfo {
	var mottos []config.MottoInfo
	err := db.Model(&config.MottoInfo{}).Where("nickname = ?", nickname).Find(&mottos).Error
	if err != nil {
		panic(err)
	}
	return mottos
}
