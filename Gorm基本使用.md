# Gorm基本使用

## 模型定义

​		通常我们需要在代码中定义模型（Models）与数据库中的数据表进行映射，在GORM中模型（Models）通常是正常定义的结构体、基本的go类型或它们的指针。

​		为了方便模型定义，GORM内置了一个`gorm.Model`结构体。`gorm.Model`是一个包含了`ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`四个字段的Golang结构体。

```
// gorm.Model 定义
type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}

// 将 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`字段注入到`User`模型中
type User struct {
  gorm.Model
  Name string
}
```

- **CreatedAt**:

  ```
  db.Create(&user) // `CreatedAt`将会是当前时间
  
  user2 := User{Name: "ha", CreatedAt: time.Now()}
  db.Create(&user2) // user2 的 `CreatedAt` 不会被修改
  
  // 可以使用`Update`方法来改变`CreateAt`的值
  db.Model(&user).Update("CreatedAt", time.Now())
  ```

- **UpdatedAt**:

  ```
  db.Save(&user) // `UpdatedAt`将会是当前时间
  
  db.Model(&user).Update("name", "ha") // `UpdatedAt`将会是当前时间
  
  user2 := User{Name: "ha", UpdatedAt: time.Now()}
  db.Save(&user2) // 更新时，user2 的 `UpdatedAt` 会修改为当前时间
  ```

- **DeletedAt**:

​		如果模型有`DeletedAt`字段，调用`Delete`删除该记录时，将会设置`DeletedAt`字段为当前时间，而不是直接将记录从数据库中删除。

**模型定义示例：**

```
type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64	//零值类型
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // 设置字段大小为255
  MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
  Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
  IgnoreMe     int     `gorm:"-"` // 忽略本字段
}
```

## 创建表

```
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB //后续用于操作数据库


type User struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`
}

func init() {
	//配置MySQL连接参数
	username := "root"       //账号
	password := "***"        //密码
	host := "***" //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "db1"         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
}

func main() {
    //创建表
	DB.AutoMigrate(&User{}) //在这里传指针的原因是指针占用的空间少
	
	
}
```

**注**：表的名字与结构体名字不同因为 ——gorm采用的命名策略是：表名是蛇形复数，字段名是身形单数

`DB.AutoMigrate(&Usr{})` 是 GORM 中用于自动迁移（Auto Migrate）的一个方法调用

gorm常用标签：

- `primaryKey`: 标记主键字段。
- `column`: 指定数据库中的列名（如果与结构体字段名不同）。
- `size`: 设置字符串类型的字段大小。
- `unique`: 确保字段值唯一。
- `default`: 设置字段默认值。
- `not null`: 确保字段非空。

![](https://raw.githubusercontent.com/StrayPumpkin/img/refs/heads/main/QQ%E5%9B%BE%E7%89%8720250222015721.png)

## 简单操作

### 判断表是否存在

```
// 检查模型是否存在
db.HasTable(&User{})
// 检查表是否存在
db.HasTable("users")
```

### 删除表

```
// 删除表
db.DropTable(&User{})
db.DropTable("users")
```

### 修改表

```
// 修改列(修改字段类型)，将description列的字段类型改为text
db.Model(&User{}).ModifyColumn("description", "text")
// 删除列，删除description列
db.Model(&User{}).DropColumn("description")
```

## CURD操作

### 插入

**直接插入**：

```
	user := User{Name: "ksy", Age: 18}
	db.Create(&user) 
```

**插入多项**：

```
	users := []User{
	{ID: 2, Name: "xi", Age: 12}
	{ID: 23, Name: "ha", Age: 12}
	}
	db.Create(&users)
```

**选择性插入**：

```
	user := User{ID: 4, Name: "wu", Age: 12}
	db.Select("id", "name").Create(&user)
```

**忽略字段**：

```
	user := User{ID: 3, Name: "kuangsiyuan", Age: 12}
	DB.Omit("name").Create(&user)
	//输出结果的name将会是默认值
```

### 更新

`Save()`默认会更新该对象的所有字段，即使你没有赋值

```
user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
```

**更新单个列：**

```
// Update with conditions 把active字段为ture的数据其中的name字段更新为hello。
针对的是整个User表中符合条件的所有记录进行更新。
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")

//直接更新
db.Model(&user).Update("name", "hello")

// Update with conditions and model value
针对一个具体的用户（通过user对象引用）。
db.Model(&user).Where("active = ?", true).Update("name", "hello")
```

**更新多列**：

```
// Update attributes with `struct`, will only update non-zero fields
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})

// Update attributes with `map`
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})

```

**批量更新**：

```
// Update with struct
db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

// Update with map
db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

```

### 删除

**删除记录**：**警告** 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录。

```
// 删除现有记录,此操作会根据 email 结构体中的主键（假设为 id=10）删除对应的记录。
db.Delete(&email)
// DELETE from emails where id=10;

//通过主键删除用户
db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM users WHERE id = 10;

//批量删除
db.Delete(&users, []int{1,2,3})
// DELETE FROM users WHERE id IN (1,2,3);

// 带额外条件的删除
db.Where("name = ?", "jinzhu").Delete(&email)
// DELETE from emails where id = 10 AND name = "jinzhu";
```

### 批量删除

```
db.Where("email LIKE ?", "%jinzhu%").Delete(email{})
 DELETE from emails where email LIKE "%jinzhu%";

db.Delete(email{}, "email LIKE ?", "%jinzhu%")
 DELETE from emails where email LIKE "%jinzhu%";

```

### 软删除

​		如果一个 model 有 `DeletedAt` 字段，他将自动获得软删除的功能！ 当调用 `Delete` 方法时， 记录不会真正的从数据库中被删除， 只会将`DeletedAt` 字段的值会被设置为当前时间

```
db.Delete(&user)
//UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// 批量删除
db.Where("age = ?", 20).Delete(&User{})
//UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// 查询记录时会忽略被软删除的记录
db.Where("age = 20").Find(&user)
//SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

// Unscoped 方法可以查询被软删除的记录
db.Unscoped().Where("age = 20").Find(&users)
//SELECT * FROM users WHERE age = 20;

```

### 永久删除

```
// Unscoped 方法可以物理删除记录
db.Unscoped().Delete(&order)
//DELETE FROM orders WHERE id=10;

```

## 查询

### 一般查询(前4个也是读取)

```
// 根据主键查询第一条记录
 db.First(&user)
 SELECT * FROM users ORDER BY id LIMIT 1;
 result := db.First(&user)
 result.RowsAffected // 返回找到的记录数
 result.Error        // returns error or nil
 
// 随机获取一条记录
db.Take(&user)
 SELECT * FROM users LIMIT 1;

// 根据主键查询最后一条记录
db.Last(&user)
 SELECT * FROM users ORDER BY id DESC LIMIT 1;

// 查询所有的记录
db.Find(&users)
 SELECT * FROM users;

// 查询指定的某条记录(仅当主键为整型时可用)
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

db.First(&user, "10")
// SELECT * FROM users WHERE id = 10;

db.Find(&users, []int{1,2,3})
// SELECT * FROM users WHERE id IN (1,2,3);

```

### 条件查询

#### where条件

```
// 得到第一条符合的记录
db.Where("name = ?", "jinzhu").First(&user)
//SELECT * FROM users WHERE name = 'jinzhu' limit 1;

// 得到所有符合的记录
db.Where("name = ?", "jinzhu").Find(&users)
//SELECT * FROM users WHERE name = 'jinzhu';

//<> 不等于的所有记录
db.Where("name <> ?", "jinzhu").Find(&users)
//SELECT * FROM users WHERE name <> 'jinzhu';

// IN子查询 在给定列表中的所有记录
db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

// LIKE模糊查询	包含 "jin" 的所有记录
db.Where("name LIKE ?", "%jin%").Find(&users)
//SELECT * FROM users WHERE name LIKE '%jin%';

// AND同时满足两个条件的所有记录
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
//SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// Time  updated_at 字段值大于 lastWeek 的所有记录。
db.Where("updated_at > ?", lastWeek).Find(&users)
//SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN 范围查询
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
//SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
```

#### Struct&Map条件

```
// Struct
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

// Map
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

// Slice of primary keys
db.Where([]int64{20, 21, 22}).Find(&users)
// SELECT * FROM users WHERE id IN (20, 21, 22);
```

若要在查询条件中包含零值，可以使用映射，该映射将包含所有键值作为查询条件，例如:

```
db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
```

#### Not条件

用法与Where类似，不过是取反

```
db.Not([]int64{}).First(&user)
//SELECT * FROM users;
```

#### Or条件 或者

```
db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// Struct
db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

// Map
db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

```

#### 内联条件

作用与Where类似，格式更为简单

```
// 根据主键获取记录 (只适用于整形主键)
db.First(&user, 23)
 SELECT * FROM users WHERE id = 23 LIMIT 1;
 
// 根据主键获取记录, 如果它是一个非整形主键
db.First(&user, "id = ?", "string_primary_key")
 SELECT * FROM users WHERE id = 'string_primary_key' LIMIT 1;
 
db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
 SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

```

### 排序

 ```
 //先按年龄降序，年龄一样再按名字升序
 db.Order("age desc, name").Find(&users)
 // SELECT * FROM users ORDER BY age desc, name;
 
 // Multiple orders链式调用
 db.Order("age desc").Order("name").Find(&users)
 // SELECT * FROM users ORDER BY age desc, name;
 ```

### 分页

**Limit:**指定从数据库检索出的最大记录数。

**Offset**:指定开始返回记录前要跳过的记录数。

```

db.Limit(3).Find(&users)
// SELECT * FROM users LIMIT 3;

// Cancel limit condition with -1
db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
// SELECT * FROM users LIMIT 10; (users1)
// SELECT * FROM users; (users2)
//先对users1的查询设置了跳过前10条记录（OFFSET 10）。之后，通过调用.Offset(-1)取消了Offset条件，使得对users2的查询不会跳过任何记录，即返回所有匹配的记录而不跳过任何行

db.Offset(3).Find(&users)
// SELECT * FROM users OFFSET 3;

db.Limit(10).Offset(5).Find(&users)
// SELECT * FROM users OFFSET 5 LIMIT 10;

// Cancel offset condition with -1
db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
// SELECT * FROM users OFFSET 10; (users1)
// SELECT * FROM users; (users2)
```

