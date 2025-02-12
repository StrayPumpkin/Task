package main
import "fmt"

type Student struct{
	Name string
	Score []float32
}//定义Student结构体

func (s Student) Display(){
	fmt.Printf(" 姓名：%s ,成绩：%v ",s.Name,s.Score)
	fmt.Println("")
}//给结构体绑定方法

func main(){
	var students []Student //定义切片
	var a int
	for {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("1.录入成绩")
		fmt.Println("2.查看成绩")
		fmt.Println("3.退出")
		fmt.Println("请选择操作: ")
		fmt.Scanln(&a)

		switch a {
		case 1:
			var name string
			var score float32
            fmt.Printf("请输入学生姓名: ")
            fmt.Scanln(&name)
			fmt.Printf("请输入学生成绩: ")
            fmt.Scanln(&score)
			
			student :=Student{
				Name: name, 
				Score: []float32{score},
				}		
            students = append(students, student)	
			fmt.Printf("录入成功 ")
		
		case 2:
			if len(students) == 0{
				fmt.Printf("未检测到学生信息 ")
			}
			for _,student :=range students{
				student.Display() //通过遍历输出成绩
			}
		case 3:
			fmt.Printf("退出 ")
		default:
		fmt.Printf("请输入正确操作")
		}	
	}
}