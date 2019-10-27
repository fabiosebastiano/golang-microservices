package mutexIncapsulated

import (
	"fmt"
	"sync"
)

var (
	atomicCounter = AtomicCounter{}
)

type AtomicCounter struct {
	value int
	lock  sync.Mutex
}

func (a *AtomicCounter) Increase() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}
func (a *AtomicCounter) Decrease() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value--
}
func (a *AtomicCounter) Value() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("valore finale contatore: %d", atomicCounter.value))
}

func updateCounter(wg *sync.WaitGroup) {
	atomicCounter.Increase()
	wg.Done()
}
