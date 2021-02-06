package main

import "fmt"

// NewBook is the func of create new Book instance.
func NewBook(title, author string, price float32, publish bool) *Book {
	return &Book{
		title:   title,
		author:  author,
		price:   price,
		publish: publish,
	}
}

// inputBook 是一个从终端获取输入 返回Book指针的一个函数
func inputBook() *Book {
	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Println("请依次根据提示输入相关信息。")
	fmt.Print("title:")
	fmt.Scanf("%s\n", &title)
	fmt.Print("author:")
	fmt.Scanf("%s\n", &author)
	fmt.Print("price:")
	fmt.Scanf("%f\n", &price)
	fmt.Print("publish:")
	fmt.Scanf("%t\n", &publish)
	newBook := NewBook(title, author, price, publish)
	return newBook
}

// AddBook 是添加书籍的方法
func AddBook() {
	newBook := inputBook()
	for _, v := range AllBooks {
		if v.title == newBook.title {
			fmt.Printf("书名为%s的书已经存在！\n", newBook.title)
			return
		}
	}
	AllBooks = append(AllBooks, newBook)
}

// ModifyBook is the func modify the book info.
func ModifyBook() {
	newBook := inputBook()

	for index, v := range AllBooks {
		if v.title == newBook.title {
			AllBooks[index] = newBook
			fmt.Printf("书名为%s的书籍信息更新成功！\n", newBook.title)
			return
		}
	}
	fmt.Printf("根据书名：%s无法查找到书籍信息！\n", newBook.title)
}

//showAllBook is the func of list all books.
func showAllBook() {
	for _, v := range AllBooks {
		fmt.Printf("title:《%s》 author:%s price:%.2f publish:%t\n", v.title, v.author, v.price, v.publish)
	}
}
