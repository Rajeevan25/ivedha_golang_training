package main

import (
	"errors"
	"fmt"
	"regexp"
)

// 1
type Calcalator struct {
	output float64
}

func enter() *Calcalator {
	return &Calcalator{}
}

func (c *Calcalator) add(num float64) *Calcalator {
	c.output += num
	return c

}
func (c *Calcalator) subtract(num float64) *Calcalator {
	c.output -= num
	return c
}
func (c *Calcalator) multiply(num float64) *Calcalator {
	c.output *= num
	return c
}
func (c *Calcalator) getOutput() float64 {
	return c.output
}

// 2
func (c *Calcalator) divide(num float64) (*Calcalator, error) {
	if num == 0 {
		c.output = 0
		return nil, errors.New("Zero Error")
	}
	c.output /= num
	return c, nil

}

//3

type User struct {
	name  string
	email string
}

func NewUser(username, email string) *User {
	return &User{
		name:  username,
		email: email,
	}
}

func (u *User) validateEmail() bool {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9.]+@[a-zA-Z0-9.]+\.[a-zA-Z]`)
	return emailRegex.MatchString(u.email)
}

func main() {
	operation := enter()
	result := operation.add(75).subtract(34).multiply(3).getOutput()
	fmt.Println(result)
	operation.divide(4)
	fmt.Println(operation.getOutput())
	operation.divide(0)
	fmt.Println(operation.getOutput())

	user1 := NewUser("Rajeevan", "rajeevan@gmai*l.com")
	fmt.Println(user1.validateEmail())

}
