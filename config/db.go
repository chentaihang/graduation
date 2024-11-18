package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库连接实例
var DB *gorm.DB

// 初始化数据库连接
// InitDB initializes the database connection using GORM with MySQL.
// It sets up the connection string, opens the connection, and performs
// automatic migration for the User model.
//
// Returns an error if the connection or migration fails.
func InitDB() error {
	// MySQL数据库连接信息
	dsn := "endless:451124@tcp(20.2.80.131:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移
	// err = DB.AutoMigrate(&models.User{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
