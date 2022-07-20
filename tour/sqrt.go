package tour

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
    var guess = 1.0
    for math.Abs(guess*guess-x) > 0.001 {
            fmt.Println(guess, guess*guess)
            guess -= (guess*guess - x) / (2*guess)
        } 
    return guess
}

func RunSqrt() {
	var res = Sqrt(25)
	fmt.Println(res, res*res)
}