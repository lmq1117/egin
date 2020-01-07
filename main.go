package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 1000000
	fibN := fibC(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	//fibC(45)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	} else {
		return fib(x-1) + fib(x-2)
	}
}

func fibC(x int) int {
	var f []int
	for i := 0; i < x; i++ {
		if i < 2 {
			f = append(f, i+1)
		} else {
			f = append(f, f[i-1]+f[i-2])
		}
		time.Sleep(time.Millisecond * 100)
	}
	//fmt.Println(f)
	return f[x-2]
}
