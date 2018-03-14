package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	m := map[string]interface{}{
		"foo":    foo,
		"bar":    bar,
		"fooErr": fooErr,
	}

	//assuming nil err in "foo" and "bar" calls
	res, err := CallFn(m, "foo")
	fmt.Println("res,err calling 'foo':", res, err)
	res, err = CallFn(m, "bar", 1, 2, 3)
	fmt.Println("res,err calling 'bar':", res, err, " Usable value:", res[0].Int())

	//returning error forcefully in "fooErr" call
	res, err = CallFn(m, "fooErr")
	fmt.Println("res,err calling 'fooErr':", res, err)
	//handle error obj
	errVal, ok := res[0].Interface().(error)
	if !ok {
		fmt.Println("Error returned by fooErr function was nil.")
	} else {
		fmt.Println("Err returned by fooErr:", errVal)
	}
	
	//Call "bar" with err in func call
res, err = CallFn(m, "bar")
	fmt.Println("res,err calling 'bar' with intentional err:", res, err)
	
	fmt.Println("Done!")
}
func foo() error {
	fmt.Println("in foo")
	return errors.New("test err")
}

func fooErr() error {
	fmt.Println("in fooErr")
	return errors.New("test err")
}

func bar(a, b, c int) (int, error) {
	fmt.Println("in bar")
	return a + b + c, nil
}

func CallFn(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("Invalid number of params passed.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
