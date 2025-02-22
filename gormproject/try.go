package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Wrn@18530164628@tcp(192.168.171.10:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("filed to connect database", err)

	}
	// 获取底层的sql.DB对象并Ping数据库以验证连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get generic database object", err)
	}

	// Ping数据库
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}

	// 如果一切正常，打印成功消息
	log.Println("Successfully connected and pinged the database!")

}
