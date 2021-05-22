// error values indicate an abnormal state
// error is an interface with Error method internally
// error = nil, means everything is fine
// https://github.com/golang/go/wiki/Errors
// https://gobyexample.com/errors
// https://blog.golang.org/defer-panic-and-recover

package main

import (
	"errors"
	"fmt"
)

// custom error type
type CustomError struct {
	msg    string
	Offset int64
}

func (e *CustomError) Error() string {
	return e.msg + " offset is " + fmt.Sprint(e.Offset)
}

func main() {
	x := 8
	err := testError(x)

	if err != nil {
		fmt.Println(err)
	}

	if ae, ok := err.(*CustomError); ok {
		fmt.Println(ae.msg)
		fmt.Println(ae.Offset)
	}
}

func testError(x int) error {

	if x > 10 { // return internal error
		return errors.New("x is greater than 10")
	}
	return &CustomError{"x is less than 10", int64(x)} // returing custom error
}
