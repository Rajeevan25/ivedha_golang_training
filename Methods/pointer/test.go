package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}
type Person struct {
	name string
	age  int
}
type Book struct {
	title  string
	author string
	year   int
}

func (b *Book) changeAuthor(newAuthor string) {
	b.author = newAuthor
}
func (p *Person) updateAge(newAge int) {
	p.age = newAge
}
func (r *Rectangle) increaseLengthByTwice() {
	r.length *= 2
}
func main() {
	obj1 := Rectangle{2, 6}
	obj1.increaseLengthByTwice()
	fmt.Println("The new length of rectangle is ", obj1.length)

	person1 := Person{"Kamal", 67}
	person1.updateAge(56)
	fmt.Println("New age of that person : ", person1.age)

	book1 := Book{"Human Anatomy", "Dr. Ashvini Kumar Dwivedi", 2023}
	book1.changeAuthor("Sajjan singh Dev")
	fmt.Println("Updated author's name is ", book1.author)
}
