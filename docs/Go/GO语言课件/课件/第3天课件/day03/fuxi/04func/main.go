package main

//函数
import "fmt"

func f1(){
	fmt.Println("hello 沙河！")
}
//参数类型简写
func f2(name1, name2 string){
	fmt.Println(name1)
	fmt.Println(name2)
}
//可变参数 0个或多个
func f3(names ...string){
	fmt.Println(names)//[]string
}

//无返回值
func f4(){
	fmt.Println("MJJ真烦")
}
func f5(a, b int) int{
	return a+b
}
//多返回值,必须用括号括起来，用英文逗号分隔
func f6(a, b int) (int, int){
	return a+b, a-b
}

//命名的返回值
func f7(a, b int) (sum int, sub int){
	sum = a + b
	sub = a - b
	return
}

//匿名函数
// func (name string){
// 	fmt.Println("Hello", name)
// }

//闭包： 函数调用了它外层的变量
func closure(key string) func(name string){
	// key := "沙河有沙还有河！"
	return func(name string){
		fmt.Println("Hello", name)
		fmt.Println(key)
	}
}


func main(){
	//通过变量调用匿名函数
	// f("豪杰")
	//声明并直接调用函数
	// func (name string){
	// 	fmt.Println("Hello", name)
	// }("沙河")

	// 闭包
	f :=closure("沙河有沙还有河！")// 得到闭包函数
	fmt.Printf("%T\n", f)
	f("豪杰")//调用闭包函数

	f2 := closure("西二旗有两个旗")// 得到闭包函数
	f2("豪杰")//调用闭包函数
}