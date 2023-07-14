package playground_test

import (
	"fmt"
	"sync"
)

type Animal interface {
	voice() string
}

type Dog struct {

}

func (d Dog) voice() string {
	return "gav"
}

func main() {
	wg := sync.WaitGroup{}
	var mutex = sync.RWMutex{}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock()
		go hello(i, &wg, &mutex)
		mutex.Lock()
		go increment(i, &wg, &mutex)
	}
	wg.Wait()
}

func hello(counter int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	fmt.Printf("Hello #%v\n", counter)
	mutex.RUnlock()
	wg.Done()
}

func increment(counter int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	counter++
	mutex.Unlock()
	wg.Done()
}
