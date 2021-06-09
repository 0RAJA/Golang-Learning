package main

import "fmt"

//匿名结构体涉及提升字段,看OOP
func main() {
	//1
	student1 := Student{
		studentName: "张三",
		favoriteBook: Book{
			bookName: "红楼梦",
			price:    25.5,
		},
	}
	fmt.Println(student1)
	//2
	boo1 := Book{
		bookName: "西游记",
		price:    99,
	}
	student2 := Student{
		studentName:  "李四",
		favoriteBook: boo1,
	}
	fmt.Println(student2)
	fmt.Println(student2.favoriteBook.bookName)
	//3
	book2 := Book{
		bookName: "三国演义",
		price:    98,
	}
	student3 := Student{}
	student3.studentName = "王麻子"
	student3.favoriteBook = book2
	fmt.Println(student3)
	//4
	student4 := Student2{
		studentName:  "Raja",
		favoriteBook: &book2,
	}
	fmt.Println("{", student4.studentName, *student4.favoriteBook, "}")
}

type Book struct {
	bookName string
	price    float64
}

type Student struct {
	studentName  string
	favoriteBook Book
}

type Student2 struct {
	studentName  string
	favoriteBook *Book //对于不需要更改或者需要统一的结构,最好采用指针进行引用传递
}
