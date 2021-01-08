// A simple illustration of how to use a closure in golang.

package main

import (
	"fmt"
)

const fibMAX = 20 // 92 is the largest value before overwhelming type int

func fibo() func() []int32 {

	nums := make([]int32, 0)
	nums = append(nums, 0, 1)

	// The payload of the fibo() function, an anonymous function with persistent state
	return (func() []int32 {
		x := nums[len(nums)-2]
		y := nums[len(nums)-1]
		nums = append(nums, x+y)
		return nums
	})

}

func main() {
	f := fibo()
	seq := make([]int32, 0)
	for i := 0; i < fibMAX; i++ {
		fibs := f()
		seq = append(seq, fibs[i])
	}
	fmt.Printf("%v\n", seq)

	for _, n := range seq {
		fmt.Println(n)
	}

	f = nil                 // dereference f
	f = fibo()              //reinitialize f
	fmt.Printf("%v\n", f()) //show that fibo() is starting fresh
	fmt.Printf("%v\n", f())
	fmt.Printf("%v\n", f())
	fmt.Printf("%v\n", f())
	fmt.Printf("%v\n", f())

}
