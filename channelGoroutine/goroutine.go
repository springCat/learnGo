package channelGoroutine

import (
	"fmt"
	"time"
)

func Do() {
	var c = make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		c <- 10
		close(c)
	}()
	for {
		select {
		case x, ok := <-c:
			if ok {
				fmt.Println(x)
			} else {
				fmt.Println("closed")
				return
			}
		default:
			fmt.Println("waiting")
			time.Sleep(time.Second)
		}
	}
}
