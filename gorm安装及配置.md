# gorm安装及配置

**1.初始化一个go模块**

`go mod init project_name`

**2.添加gorm依赖**

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

3.配置和连接MySQL数据库

```
package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Wrn@18530164628@tcp(192.168.171.10:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local"
	//charset=utf8mb4&parseTime=true&loc=Local: 连接参数，指定字符集、时间解析选项及时区设置。
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
		log.Fatal("failed to ping database ", err)
	}

	// 如果一切正常，打印成功消息
	log.Println("Successfully connected and pinged the database!")

}
```

**dsn格式**：`username:password@protocol(address数据库服务器地址：端口号)/dbname?param1=value1&param2=value2`

**gorm.Open函数使用方法**：`func Open(dialector Dialector, opts ...Option) (*DB, error)`

- **dialector**：这是一个实现了 `gorm.Dialector` 接口的对象，用于指定要使 用的数据库驱动和连接信息。对MySQL而言 `mysql.Open(dsn)`返回的就是该接口的对象。
- **opts...Option**：这是一个可选参数，允许你传入多个 GORM 配置选项。

**db.DB**:获取底层的sql.DB对象

**sqlDB.Ping(**):用于测试与数据库的连接是否仍然有效



一个bug:运行后告诉我MySQL 服务器不允许从指定的主机地址进行连接

解决：在MySQL上为特定地址授权

```
CREATE USER 'your_username'@'192.168.171.10' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON *.* TO 'your_username'@'192.168.171.10' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```



