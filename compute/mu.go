package compute

import "fmt"

func Mut(n int, ans chan int) {
	for i := 0; i < n; i++ {
		fmt.Println("gorutine:", i, "start")
		go func() {
			var a int
			var b int
			for j := 0; j < 99; j++ {
				for i := 0; i < 9999999; i++ {
					a = i*i + i
				}
				b = a
			}
			ans <- b
		}()
	}
}
