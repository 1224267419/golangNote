package main

import "fmt"

type book struct {
	name  string
	athor string
	price int
}

func main() {
	var book1 book
	var book2 book

	book1.name = "书1"
	book1.athor = "作者1"
	book1.price = 100
	book2.name = "书2"
	book2.athor = "作者2"
	book2.price = 30
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book2.price)
	fmt.Println(book1.name)

	var struct_pointer *book
	struct_pointer = &book1
	fmt.Println(struct_pointer.name)

	var number []int
	fmt.Println((number == nil))

}
