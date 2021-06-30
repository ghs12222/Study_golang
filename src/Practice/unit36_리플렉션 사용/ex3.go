// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Person struct {
// 	name string `tag1:"Name1" tag2:"Name"`
// 	age  int    `tag2:"Age1" tag2:"Age"`
// }

// func main() {
// 	p := Person{}
// 	name, ok := reflect.TypeOf(p).FieldByName("name")
// 	fmt.Println(name, ok)
// 	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

// 	// age, ok := reflect.TypeOf(p).FieldByName("age")
// 	// fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))
// }
