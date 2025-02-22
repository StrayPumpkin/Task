package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB //后续用于操作数据库

type User struct {
	ID uint // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	Name  string `gorm:"column:username;not null"`
	Age   int    `gorm:"column:age;not null"`
	Email string `gorm:"column:email;unique"`
}

func initDB() {
	//配置MySQL连接参数
	username := "root"            //账号
	password := "Wrn@18530164628" //密码
	host := "192.168.171.10"      //数据库地址，可以是Ip或者域名
	port := 3306                  //数据库端口
	Dbname := "db1"               //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

// 创建新用户
func CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

// GetAllUsers 获取所有用户
func GetAllUsers(users *[]User) error {
	if err := db.Find(users).Error; err != nil {
		return fmt.Errorf("failed to get all users: %v", err)
	}
	return nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *User) error {
	if err := db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

// HardDeleteUser 物理删除用户
func HardDeleteUser(id uint) error {
	var user User
	if err := db.Unscoped().First(&user, id).Error; err != nil {
		return fmt.Errorf("failed to find user for hard delete: %v", err)
	}
	if err := db.Unscoped().Delete(&user).Error; err != nil {
		return fmt.Errorf("failed to hard delete user: %v", err)
	}
	return nil
}

func main() {
	// 初始化数据库
	initDB()

	// 创建用户
	newUser := User{Name: "xi", Email: "xi@example.com", Age: 18}
	if err := CreateUser(&newUser); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("Created user with ID: %d\n", newUser.ID)

	// 获取所有用户
	var users []User
	if err := GetAllUsers(&users); err != nil {
		log.Fatalf("Failed to get all users: %v", err)
	}
	for _, u := range users {
		log.Printf("User: %+v\n", u)
	}

	// 更新用户
	updatedUser := User{ID: newUser.ID, Age: 21}
	if err := UpdateUser(&updatedUser); err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	log.Println(fmt.Sprintf("Updated user age to %d", updatedUser.Age))

	// 物理删除用户
	// if err := HardDeleteUser(newUser.ID); err != nil {
	//     log.Fatalf("Failed to hard delete user: %v", err)
	// }
	// log.Println("Hard deleted user")
}
