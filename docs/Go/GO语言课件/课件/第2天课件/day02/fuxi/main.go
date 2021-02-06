package main
import "fmt"
func main(){
	s := "hello中国"
	for i:=0;i<len(s);i++{
		fmt.Println(s[i])//默认按照ASCII码去打印
		fmt.Printf("%c\n", s[i])
	}
	fmt.Println()
	for index,char:=range s{
		fmt.Println(index, char)
		fmt.Printf("%c\n", char)
	}
}