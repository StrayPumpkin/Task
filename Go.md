# Go语法

## 注释

行注释：// 快捷键：ctrl+/

块注释：/**/快捷键：shift+alt+a

## 变量

变量声明：

```go
var age int
```

变量赋值：

```
age = 18
```

变量使用

```
fmt.Println("age = ",age)
```

声明和变量合成一句

```
var age int = 18
```

**注**：

(1)变量未赋值，输出为默认值0

(2)没有写变量类型，如：

```
var num = 2
```

则会根据赋值内容自动类型判断。

(3)若省略var，则写为 ：=，如

```
num := 1
```

(4)多变量声明

```
var a,b,c int

var a,b,c = 1,"look",1.3

a,b,c :=1,2,3
```

### 全局变量与局部变量

全局变量：定义在函数外的变量。

```
//方式1
var a = 1
var b = 2
//方式2
var(
		m = 100
		n = "hello"
)
```

局部变量：定义在函数内，在{ }里定义。

## 基本数据类型

### 字符数据类型(byte)

```
var c byte ='a'
```

用byte定义，fmt.Println输出值为字符对应ASCII码值,要想输出字母，则用格式化输出`fmt.Printf("%c",c)`

### 布尔类型(bool)

只允许取值ture/false

```
var f bool = ture/false
var f boll = 5 > 8  
```

### 字符串类型(string)

```
var s string = "hello"
```

**注**：

(1)字符串中没有特殊字符，则用“ ”引起。

(2)字符串中有特殊字符，则用反引号`引起。

(3)字符串的拼接用“+”，如：`var s = "asd" + "dfg"`

### 数据类型转换

#### 数字间的转换

必须强制类型转换，如：

```
var n1 int = 100
var n2 float32 = float32(n1)
```

当有多种数据类型进行运算时，一定要将等号左右数字进行匹配，如：

```
var n1 int32 = 18
var n2 int64 = int64(n1) + 10
```

#### 基本数据类型转字符串

`fmt.Sprintf("%参数",表达式)`

```
var n1 int = 9
var n2 float32 = 9.1
var n3 bool = rute
var n4 byte = 'a'

var s1 string = fmt.Sprintf("%d",n1)
var s2 string = fmt.Sprintf("%f",n2)
var s3 string = fmt.Sprintf("%t",n3)
var s4 string = fmt.Sprintf("%c",n4)
```

#### 字符串转基本数据类型

用strconv包，如：

```
var s1 string = "ture"
var b bool
b , _ = stronv.ParseBool(s1)

var s2 string = "16"
var n int64
n , _ = stronv.ParseInt(s2,10,64)
//第一个位置写转换对象，第二个写进制，第三个写位数。

var s3 string = "2.3"
var f float64
b , _ = stronv.ParseFloat(s3,64)
//第一个位置写转换对象，第二个写位数。
```

## 指针

```
var a int = 6
var ptr *int = &age
```

## 输入

```
var age int
//法1
fmt.Scanln(&age)
//法2
fmt.scanf("%d",a)
```

## 分支结构

### if

#### 单分支

```
if 条件 {
	逻辑代码
}
//"条件"内容不建议加括号，后必须有大括号。
```

#### 双分支

```
if 条件 {
	逻辑代码1
}else{
	逻辑代码2	
}
```

#### 多分支

```
if 条件1 {
	逻辑代码1
}else if 条件2 {
	逻辑代码2	
}
...
else{
	逻辑代码n	
}
```

### switch

```
switch 表达式 {
	case 值1，值2...
		语句块1
	case 值3，值4...
		语句块2
	case 值5，值6...
		语句块3
		...
	default
		语句块
}
```

## 循环结构

### 仅有for循环

...

### for range

```
var str string = "hello 你好"

for i , value := range str{
	fmt.Printf("索引为： %d，具体值为 ：%c",i,a)
}
/*对str进行遍历，遍历的每个结果的索引值被i接收，每个结果的具体值被value接收，可以为汉字，占3字节*/
```

## 各种终止

break：停止正在执行的循环

continue：结束本次循环，继续下一次循环。

return：结束当前函数。

## 函数

```
//自定义函数
func cal (n1 int,n2 int) (int){
	var sum int = 0
	sun += n1
	sun += n2
	return sum
}
func main(){
//调用函数
	sum := cal(6,9)
	fmt.Println(sum)
	
}

```

**注**：

(1)不支持重载。

(2)支持可变参数。

(3)函数也是一种数据类型，合约赋给一个变量。

```
func cal (n1 int,n2 int) (int){
	var sum int = 0
	sun += n1
	sun += n2
	return sum
}
func main{
a := cal
}
```

(4)函数可以作为形参并调用。

(5)可以自定义数据类型（相当于起别名）,起别名后就不属于同一数据类型。

```
type myInt int

var n myInt = 30
```

（6）支持对返回值命名。

### 包

1.package进行包的声明，建议包的声明与包所在文件夹同名那个。

2.main包是程序的入口包

3.导入包的名称首字母要大写。

4.函数调用时要定位到所在包。

### init函数

先于main函数执行， 后于全局变量

### 匿名函数

在定义函数时直接调用，仅使用一次。

```
func main(){
	rusult := func (n1 int ,n2 int) int{
		rusult = n1 + n2
	}(10,20)
}
```

### 闭包

```
//getSun函数返回值为一个函数，该函数参数为int类型，返回值也是int类型
func getSum() func (int) int {
	var sum int = 0
	return func (num int) int{
		sum += num
        return sum
	} 
}
// 
func main(){
	f :=getSum()
	fmt.Println(f(1))//1
	fmt.Println(f(2))//3
	fmt.Println(f(3))//6
	fmt.Println(f(4))//10
}
```

效果类似于c语言中的static

### defer关键字

靠后执行

```
func add(n1 int ,n2 int) int{
//遇到defer 不立即执行，将其压入栈中，执行后续语句。
	defer fmt.Println("n1="，n1)
	defer fmt.Println("n2="，n2)
//栈的特点是先进后出
	var sum int = n1 + n2
	fmt.Println("sum="，sum)
	ruturn sum
}
fun main(){
	fmt.Println(add(30,60))
}
//执行结果：
sum= 90
n2= 60
n1= 30
90
//进入栈后即使数据发生变化，输出还是变化前的值
```

### 字符串函数

统计字符串长度

```
str := "go你好"
fmt.Println(len(str))//8字节
```

对字符串进行遍历

法1：用键值循环：for-range

```
var str string = "hello 你好"

for i , value := range str{
	fmt.Printf("索引为： %d，具体值为 ：%c",i,a)
}
/*对str进行遍历，遍历的每个结果的索引值被i接收，每个结果的具体值被value接收，可以为汉字，占3字节*/
```

法2：利用r:=[]rune(str)

```
r :=[]rune(str)
for i :=0; i < len(r); i++{
	fmt.Print("%c\n",r[i])
}
```

字符串转整数

```
n1,_ :stronv.Atoi("666")
fmt.Println(n1)
```

整数转字符串

```
str :=strconv.Itoa(77)
fmt.Println(str)
```

统计一个字符串有几个指定字符

```
cont :=strings.Cont("aaasdff","a")
fmt.Println(cont)
```

返回字串在字符串第一次出现的索引值，如果没有则返回-1

```
fmt.Println(strings.Index("hellohhe","he"))
```

字符串的替换

```
str1 :=strings.Replace("hhhhh","h","hi",-1)
/*第一空填原始字符串，第二空填想要替换的字符串，第三空填替换后字符串，第四空填替换数目，全部替换填-1*/
```

字符串拆分

```
arr :=strings.Split("1-2-3-4","-")
```

大小写替换

```
strings.Tolower("Go")
strings.Toupper("go")
```

 ## 错误处理

程序中出现错误后程序被中断

引入defer+recover机制处理错误

```
func main(){
	test()
	fmt.Println("执行成功")
}
func test(){
//用defer+recover捕获错误，defr后加上匿名函数的调用
	defer func(){
	//调用recover内置函数捕获错误
        err :=recover()
        //没有错误则返回值为零值：nil
        if err !=nill{
            fmt.Println("错误已捕获")
            fmt.Println("err是： "，err)
        }
	}()
	n1 :=10
	n2 :=0
	result :=n1/n2
	fmt.Println("reshult")
}
```

自定义错误

```
func main(){
	err ：=test()
	if err != nil {
	fmt.Println("执行成功")	
	}
	fmt.Println("自定义错误： "，err)
}
func test() (err error){

	n1 :=10
	n2 :=0
	if n2 == 0{
		//抛出自定义错误
		return erros.New("除数不能为0")
	}else{
		result :=n1/n2
		fmt.Println("reshult")
		//如果没有错误，返回零值
		return nil
	}
}
```

## 数组

### 数组的遍历

法1：普通for循环

法2：键值循环

```
//一维数组
for key,val :=range coll{
		fmt.Printf("第%d个学生成绩为：%d\n",key + 1,val)
}
//coll是数组，遍历索引用key接收，索引位置上的值用val

//二维数组
for key,value :=range arr {
	for k,v :=range value{
		fmt.Printf("arr[%v][%v]=%v\t",key,k,v)
	}
	fmt.Println()
}
```

### 数组的初始化

```
var arr1 [3]int = [3]int{1,2,3}
var arr2 = [3]int{1,2,3}
var arr3 = [...]int{1,2,3}
var arr4 = [...]int{1:66,2:99,3:98}
```

## 切片slice(特有)

### slice定义

构建在数组之上。

```
var arr [6]int = [6]int{1,2,3,4,5,6}
var slice []int = arr[1:3]
//切除一段片段，从1开始，到3结束，不包含3
slice := arr[1:3]
	fmt.Println("arr:",arr)
	fmt.Println("slice:",slice)
	fmt.Println("slice元素个数："，len(slice))
	fmt.Println("slice容量："，cap(slice))
	//容量是元素的二倍左右
```

slice底层是一个结构体，第一部分是指针，存地址；第二部分存长度；第三部分存容量

```
//用make函数定义切片,3个参数：1.切片类型；2.切片长度；3.切片容量
slice := make([]int,4,20)
```

make底层创建一个数组，对外不可见，不能直接操作，要通过slice间接访问。

```
slice := []int{1,4,7}
//此定义方式类似make
```

### slice遍历

```
slice := make([]int,4,20)
slice[0] = 11
slice[1] = 22
slice[2] = 33
slice[3] = 44
//法1：for循环
for i :=0;i < len(slice);i++{
	fmt.Printf("slice[%v] = %v \t",i slice[i])
}
//法2 for-range循环
for i,v :=range slice{
	fmt.Printf("下标： %v， 元素： %v \t",i v)
}

```

### slice追加

```
var arr [6]int = [6]int{1,2,3,4,5,6}
var slice1 []int = arr[1:4]
slice2 :=append(slice,8,6)
//对数组进行扩容，形成一个新数组，原数组不变，即slice不变
```

### slice复制

```
var arr [6]int = [6]int{1,2,3,4,5,6}
var a []int = arr[1:3]
b := make([]int,4,20map
//定义两个切片
copy(b,a)
//将a中内容复制给b
```

## 映射map

将键值对进行关联，左为键(key),右为值(value)

```
var a map[int]string
//只声明map内存没有分配空间，要通过make进行初始化
a  = make(map[int]string,10)//map可以存放10个键值对
a[1001] = "ha"
a[1002] = "he"
a[1003] = "ya"
a[1004] = "ma"
```

特点：

(1)要通过make初始化

(2)map中key-value是无序的

(3)key不可以重复，若重复value会被覆盖

```
//方式2
b :=make(map[int]string)
	b[2001] = "hey"
	b[2002] = "bey"
//方式3
c := map[int]string{
		2001 = "hey"
		2002= "bey"
}
```

```
b :=make(map[int]string)
//增加
	b[2001] = "hey"
	b[2002] = "bey"
//修改
	b[2002] = "book"
//删除
	delete(b,2001)
//查找
	value，flag := b[2001]
	fmt.Println(value)//hey
	fmt.Println(flag)//ture
//获取长度
	fmt.Println(len(b))
//遍历	
	for k,v :=range b{
	fmt.Printf("key为： %v， value为： %v \t",k v)
}
```

```
//map嵌套
a :=make(map[string]map[int]string)
//赋值
a["班1"] = make(map[int]string,3)
a["班1"][1001] = "bo"
a["班1"][1002] = "aa"
a["班1"][1003] = "li"

a["班2"] = make(map[int]string,3)
a["班2"][2001] = "lee"
a["班2"][2002] = "long"
a["班2"][2003] = "fu"

for k1,v1 :=range a{
	fmt.Printf(k1)
	for k2,v2 :=range v1{
	fmt.Printf("学号为： %v， 姓名为： %v \t",k2 v2)
	}
	fmt.Printf()
}
```

## 结构体

```
//方式1
type Teacher struct{
	Name string
	Age int
	School string
}
func main(){
	var t Teacher
	t.Name = "kihua"
	t.Age = 18
	t.School = "lishan"
}
```

```
//方式2
var t Teacher =Tercher{"kihua",18,"lishan"}
```

```
//方式3
var t *Teacher = new(Teacher)
/*	(*t).Name = "kihua"
	(*t).Age = 18
	(*t).School = "lishan" */
	t.Name = "kihua"
	t.Age = 18
	t.School = "lishan"	
```

```
//方式4
var t *Teacher = &Teacher{kihua",18,"lishan}
//	t.School = "lishan"	
```

## 方法

```
//定义结构体
type Person struct{
	Name string
}
//给结构体绑定方法test
func (p Person) test(){
	fmt.Println(p.Name)
}
func main(){
	var p Person
	p.Name = "ming"
	p.test()
	
}
```

## 封装

```
type person srtuct{
	Name striing
	age int
}
//定义工厂模式函数
func NewPerson(name string) *person{
	return &person{
	Name : name,
	}
}
//定义set和get方法，对age进行封装
func (p *person) SetAge(age int){
	if age > 0 && age < 150{
	p.age = age
	}else{
		fmt.Println("出错")
	}
}
func (p *person) GetAge() int {
	return p.age
}
```

再创建另一个包，

```
p : model.NewPerson("lee")
p.SetAge(200)
```

## 组合

```
//定义动物结构体
type Animal struct{
	Age int
	Wight float32
}
//绑定方法
func (an *Animal) ShowInfo(){
fmt.Printf("动物的年龄是： %v，动物的体重是： %v"，an.Age,an.Weight)
}
func (an *Animal) Shout(){
	fmt.Printf("喊叫“)
}
//定义结构体cat
type Cat struct{
	Animal
//提高复用性
}
func(c *Cat) scratch(){
	fmt.Printfli("挠人")
}

func main(){
	cat := &Cat{}
	cat.Animal.Age = 3
	cat.Animal.Weight = 10.6
	cat.Animal.Shout()
	cat.Animal.ShowInfo
	cat.scratch()
}
//	cat.Animal.Shout() 中"Animal"可省略
```

## 接口

```
type SayHello interface{
	sayHello()
}
type Chinese struct{

}
func(person Chinese)sayHello(){
	fmt.println("你好")
}
type American struct{

}
func(person American)sayHello(){
	fmt.println("hi")
}
//定义一个函数，专门用来各国人打招呼，接收具有SayHello能力的变量
func greet(s SayHello){
	s.sayHello()
}

func main(){
	c :=Chinese{}
	a :=American{}
	greet(a)
	greet(b)
}
```

## 断言

```
ch,flag :=s.(Chinese)//看s能否转成Chinese类型并赋给ch变量，flag判断是否转成功
if flag == ture{
	ch.niuYangGe
}else{
	fmt.println("no")	
}
```

