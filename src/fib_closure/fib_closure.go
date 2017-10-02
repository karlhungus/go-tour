package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	//fn = fn -1 + fn - 2
	// f1, f0 = 1, 0
	fnMinus2, fnMinus1, fn := 0, 1, 1
	return func() int {
		fnMinus2, fnMinus1, fn = fnMinus1, fn, fnMinus1+fn
		return fnMinus1 - fnMinus2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
