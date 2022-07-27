package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

type T struct{}

func (t *T) Add(a, b int) string {
	fmt.Printf("a + b is %+v\n", a+b)
	return "a + b"
}

func TestRefect(t *testing.T) {
	// funcName := "Add"
	// typeT := &T{}

	// a := reflect.ValueOf(1)
	// b := reflect.ValueOf(2)

	// log.Println("reflect.ValueOf(1) \n", a)
	// log.Println("reflect.ValueOf(2) \n", b)

	// in := []reflect.Value{a, b}
	// returnValue := reflect.ValueOf(typeT).MethodByName(funcName).Call(in)

	// log.Println("returnValue \n", returnValue[0].String())

	funcName := "GetName"
	arg1 := "hello"
	arg2 := "world"

	log.Println("funcName, arg:", funcName, arg1, arg2)

	typeHandler := &Handler{}

	refectValueArg1 := reflect.ValueOf(arg1)
	refectValueArg2 := reflect.ValueOf(arg2)
	in := []reflect.Value{refectValueArg1, refectValueArg2}

	valueFunc := reflect.ValueOf(typeHandler).MethodByName(funcName)
	valueReturn := valueFunc.Call(in)

	returnValue := valueReturn[0].String()

	log.Println("returnValues:", returnValue)
}
