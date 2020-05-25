package main

/**
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
*/
import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	fmt.Println(Books{"132", "cq", "cc", 1})
	fmt.Println(Books{title: "132"})
	var book1 Books
	var book2 Books
	book1.title = "go"
	book2.title = "go2"

	fmt.Println(book1.title)
	fmt.Println(book2)

	printBook(book1)
	fmt.Println("----------------")
	printBook2(&book1)

}
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook2(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
