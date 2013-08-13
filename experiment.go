package main

import (
    "fmt"
)

/* STRUCTS */
type Base struct {
    MyValue int
}

type Example1 struct {
    Base
    MyValue1 int
}

type Example2 struct {
    Base
    MyValue2 int
}

/* INTERFACES */
type IExample1 interface {
    IExample1()
}

type IExample2 interface {
    IExample2()
}

/* FAKE METHODS */
func (e Example1) IExample1() {}
func (e Example2) IExample2() {}


/* Guess the type of the struct */
func typeguess(i interface{}) {

    switch i.(type) {
        case IExample1:
            v := i.(*Example1)
            fmt.Printf("I'm an IExample1!\n%d\n", v.MyValue)
            break
        case IExample2:
            v := i.(*Example2)
            fmt.Printf("I'm an IExample2!\n%d\n", v.MyValue)
            break
        default:
            fmt.Printf("Unknown type\n")
    }
}

func main() {

    a := &Example1{Base{42}, 10}
    b := &Example2{Base{42}, 10}

    typeguess(b)

    _ = a
    _ = b
}
