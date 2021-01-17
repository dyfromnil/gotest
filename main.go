package main

import "fmt"
import "time"
import "runtime"
import "github.com/dyfromnil/gotest/compute"

func main() {
	var after <-chan time.Time
	after = time.After(time.Duration(1000) * time.Millisecond)

	var sum int
	ch := make(chan int)
	go func() {
		sum = 0
		var a int
		for i := 0; i < 4500; i++ {
			sum += i
			a = 45 * 987
		}
		sum += a
		time.Sleep(time.Duration(2000) * time.Millisecond)
		ch <- sum
	}()
	for {
		fmt.Println("wait....")
		select {
		case <-after:
			fmt.Println("n s's later")
		case a := <-ch:
			fmt.Println("result is:", a)
			goto endFor
		}
	}
endFor:
	fmt.Println("this is the end!")

	ans := make(chan int)
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	n := 3

	fmt.Println("begin computing...")

	compute.Mut(n, ans)

	for i := 0; i < n; i++ {
		r := <-ans
		fmt.Println("mul:", r)
	}

}
