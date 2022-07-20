package tour

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last_fib := 0
	fib := 1
	return func() int{
		to_sum := last_fib
		last_fib = fib
		
		fib += to_sum
		return fib
	}
}

func RunFib() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}