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

func (b Book) displayDetails() {
	fmt.Println("The book's name is "+b.title+", written by "+b.author+" in ", b.year, ".")
}
func (p Person) greet() string {
	return "Hello, " + p.name
}
func (r Rectangle) calculateArea() float64 {
	return r.length * r.width
}
func main() {
	obj1 := Rectangle{12, 6}
	result := obj1.calculateArea()
	fmt.Println(result)

	person1 := Person{"Kamal", 67}
	fmt.Println(person1.greet())

	book1 := Book{"Human Anatomy", "Dr. Ashvini Kumar Dwivedi", 2023}
	book1.displayDetails()

}
