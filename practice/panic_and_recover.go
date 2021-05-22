package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/*
1. A defer statement pushes a function call onto a stack.
   The list of saved calls is executed after the surrounding function returns.
eg : func b() {
		for i := 0; i < 4; i++ {
			defer fmt.Print(i)
		}
	}
output : 3,2,1,0

2. Panic is a built-in function that stops the ordinary flow of control and begins panicking. 
   When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller
   Recover is a built-in function that regains control of a panicking goroutine. During normal execution, a call to recover will return nil and have no other effect
