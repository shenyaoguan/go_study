package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	User
	Class string
}

func (u User) SayName(name string) {

	fmt.Println("My name is", name)
}

func main() {
	u := User{
		Name: "Alice",
		Age:  18,
		Sex:  true,
	}
	//s := Student{
	//	u,
	//	"Class 1",
	//}
	check(u)
	//fmt.Println(s)
}

func check(inter interface{}) {
	//t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	m := v.Method(0)
	m.Call([]reflect.Value{reflect.ValueOf("Alice")})
}
