package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func count(str string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		count("Sheep")
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		count("Dog")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
