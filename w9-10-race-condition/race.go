package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 0

	/**
	 * This is an example of race condition
	 *
	 * Race condition is two or more operations if they access the same memory location and at least one of them is a write operation
	 * In this particular example two concurrent programms write to one variable x, and depends on a concurrent order
	 * final x can be 2 or -2 which is not deterministic
	 */

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			x += 1
			fmt.Printf("go plus 1 executed, x is %v\n", x)
			if x > 1 {
				break
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			x -= 1
			fmt.Printf("go plus 2 executed, x is %v\n", x)
			if x < -1 {
				break
			}
		}
	}()
	wg.Wait()

	fmt.Printf("Completed with x = %v\n", x)
}
