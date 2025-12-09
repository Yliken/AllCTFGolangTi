package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() (*gorm.DB, error) {
	config := struct {
		User     string
		Password string
		Host     string
		DB       string
		Port     int
	}{
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     3306,
		DB:       "sql",
	}

	// ✅ 正确的 MySQL DSN 格式
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DB,
	)
	newLogger := logger.New(
		nil, // 输出为 nil，相当于禁用日志输出
		logger.Config{
			LogLevel: logger.Silent,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return nil, err
	}

	return db, nil
}
