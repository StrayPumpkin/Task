# Go是oop

1.可以通过结构体和方法实现封装

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

另一个包中

```
p : model.NewPerson("lee")
p.SetAge(200)
```



2.通过组合实现代码复用

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



3.通过接口实现多态

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

