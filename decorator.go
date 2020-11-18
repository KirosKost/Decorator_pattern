package main

import (
	"log"
	"math"
	"os"
	"time"
)

type piFunc func(int) float64

//logger(Pi(n))
func wraplogger(fun piFunc, logger *log.Logger) piFunc {
	return func(n int) (result float64) {
		fn := func(n int) (result float64) {
			defer func(t time.Time) {
				logger.Printf("took=%v, n=%v, result=%v, time.Since(t), n, result")
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}

}
func Pi(n int) float64 {
	ch := make(chan float64)

	for k := 0; k <= n; k++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(ch, float64(k))
	}
	result := 0.0
	for k := 0; k <= n; k++ {
		result += <-ch
	}
	return result
}

func main() {
	g := wraplogger(Pi, log.New(os.Stdout, "Test", 1))
	g(500000)
	g(30000)
}
