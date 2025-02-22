## MySQL基本使应用

```
mysql -u root -p //进入
```

## 数据库基本操作

**创建数据库**：`create database 数据库名称(假设为db1_);`

**查看数据库信息**：`show create database db1;`

**删除数据库**：`drop database db1;`

**查询所有数据库**：`show database;`

**切换数据库**：`use db1;`

**查看当前数据库名称**：`select database()`；

## 数据表结构

1. **表名**：数据表唯一标识符，用于在数据库引用该表。

2. **列**：表中字段，每一列都有名称和数据类型。

3. **数据类型**：定义了列中可以存储的数据类型，常见的数据类型包括：

   - **整数类型**：`INT`、`TINYINT`、`SMALLINT`、`MEDIUMINT`、`BIGINT`

   - **浮点数类型**：`FLOAT`、`DOUBLE`、`DECIMAL`
   - **字符串类型**：`CHAR`、`VARCHAR`、`TEXT`、`BLOB`
   - **日期和时间类型**：`DATE`、`TIME`、`DATETIME`、`TIMESTAMP`
   - **布尔类型**：`BOOLEAN` 或 `TINYINT(1)`
   - **枚举类型**：`ENUM`（用于存储预定义的值列表）

4. **约束**：限制列中数据，保证数据完整性和一致性。常见的约束包括：

   - **主键约束 (Primary Key)**：唯一标识表中的每一行，不允许重复和空值。

   - **外键约束 (Foreign Key)**：用于建立表与表之间的关系。
   - **唯一约束 (Unique)**：确保列中的值唯一，但允许空值。
   - **非空约束 (Not Null)**：确保列中的值不能为空。
   - **默认值 (Default)**：为列指定默认值,后面直接引出默认值。
   - **检查约束 (Check)**：确保列中的值满足特定条件。

5. **索引**：加快数据检索速度，常见的索引类型包括：

   - **主键索引 (Primary Key Index)**：自动为主键列创建索引。

   - **唯一索引 (Unique Index)**：为唯一约束的列创建索引。
   - **普通索引 (Index)**：为普通列创建索引。
   - **全文索引 (Fulltext Index)**：用于全文搜索。

6. **表选项**：指定表的存储引擎、字符集、排序规则。常见的表选项包括：

   - **存储引擎(Storage Engine)**：`InnoDB`(支持事务、外键，适合大多数应用场景）、`MyISAM`（不支持事务，适合读多写少的场景） `Memory`（数据存储在内存中，速度快但数据不持久）等。

   - **字符集Character Set**：决定了表中可以存储的字符类，如：`utf8`、`utf8mb4` 等。
   - **排序规则Collation**：决定了字符的比较和排序方式，如：`utf8_general_ci`、`utf8mb4_unicode_ci` 等。



示例：

```
CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,  -- 用户ID，主键，自增
    username VARCHAR(50) NOT NULL UNIQUE,    -- 用户名，唯一且不能为空
    password VARCHAR(255) NOT NULL,         -- 密码，不能为空
    email VARCHAR(100) UNIQUE,              -- 邮箱，唯一
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间，默认为当前时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  -- 更新时间，自动更新
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 外键——表间建立关联

在创建表时定义外键：

```
CREATE TABLE 表名 (
    列名 数据类型,
    ...
    FOREIGN KEY (外键列名) REFERENCES 关联表名(关联列名)
);
```

在已存在的表中添加外键：

```
ALTER TABLE 表名
ADD CONSTRAINT 外键名称
FOREIGN KEY (外键列名) REFERENCES 关联表名(关联列名);
```

示例：

```
create table student05(
id int primary key,
name varchar(20)
); --创建学生表

create table class(
classid int primary key,
studentid int
);  --创建班级表

alter table class
add constraint fk_class_studentid
foreign key(studentid) references student05(id);
--学生表作为主表，班级表作为关联表设置外键
```

外键可以定义一些约束行为，以确保数据的一致性：

- **ON DELETE CASCADE**：当关联表中的记录被删除时，自动删除当前表中的相关记录。
- **ON UPDATE CASCADE**：当关联表中的记录被更新时，自动更新当前表中的相关记录。
- **ON DELETE SET NULL**：当关联表中的记录被删除时，将当前表中的外键列设置为 `NULL`。
- **ON DELETE RESTRICT**：阻止删除关联表中的记录（默认行为）。

## 数据表基本操作

**创建数据表**：

````
creat table studeut(
	id int,
	name varchar(20),
	gender varchar(10),
	birthday date
);
````

**查看数据库中所有数据表**：`show tables;`

**查看数据表基本信息**：`show create table student;`

**查看表的字段具体信息**：`desc srudent;  `

**修改表名**：`alter table student rename to stu;`

**修改字段名**：`alter table stu change name sname varchar(10);`//必须要有数据类型

**修改字段数据类型**：`alter table stu modify sname int;`

**增加字段**：`alter table stu add address varchar(50);`

**删除字段**：`alter table stu drop address;`

**删除数据表**：`drop table stu;`

**删除外键**：`alter table 从表名 drop foreign key 外键名;`

`alter table class drop foreign key fk_class_studentid;`

## CURD

### 创建

**为表中字段插入数据**：字段与值严格一一对应。

````
INSERT INTO 表名（字段名1,字段名2,...)
VALUES (值 1,值 2,...);
````

- 示例

```
 create table student(
 id int,
 name varchar(30),
 age int,
 gender varchar(30)
 );
insert into student (id,name,age,gender)
values (1,'bob',16,'male');

```

**同时插入多条数据**：

```
insert into student (id,name,age,gender)
values
	(2,'lucy',17,'female'),
	(3,'jack',19,'male'),
	(4,'tom',18,'male');
```

### 读取

```
SELECT 列1, 列2, ...
FROM 表名
WHERE 条件
ORDER BY 列名 [ASC|DESC]
LIMIT 数量;
```

**查询所有字段**：`select * from student;`

**查询指定字段**：`select sid,sname from student;`

**带条件的查询**：`select * from users where gender = 'male';`

**排序查询**：`select * from users order by ASC;`

**限制查询结果数量**：`select * from users limit 2;`

**模糊查询（用like）**：`select * from users where username like 'j%';  --查询以j开头的用户名`

**常数的查询**：select中书写常数（字符串、数字或日期等），用于向结果中添加额外信息或标记特定的记录。可以为查询结果添加自定义信息，这些信息并不直接来源于数据库表中的数据。

`select sid,sname,'2021-03-02' from student;`

- 其他例子：

```
SELECT name, 'Active' AS status FROM employees;
--添加静态文本

SELECT id, name, 1 AS is_vip FROM customers WHERE vip = 'Yes';
--标记记录
```

**从查询结果中过滤重复数据**：`select distinct gender from student;`

**可以使用加减乘除运算符**：` select sname,age+10 from student;`

### 更新

**更新数据**：

```
UPDATE 表名
SET 字段名1=值1,字段名2 =值2,…
WHERE 条件表达式;
```

- 示例：

```
update student
set age=20,gender='female'
where name='tom';  --更新部分数据

update student set age=18;  --更新所有数据

```

### 删除：

```
DELETE FROM 表名
WHERE 条件表达式;
```

- 示例：

```
delete from student
where age=14;  --删除age=14的数据

delete from student;  --删除表中所有数据
```