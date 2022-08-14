package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	mu sync.Mutex //Объявлять переменную до которой она должна "защищать". Блокировать только эксклюзивный доступ.
	// Остальной код лучше не включать в критическую секцию
	rw sync.RWMutex //Можно запрашивать блокировку на чтение или запись отдельно. Пока идет запись нельзя читать
	//Использование тогда, когда код изменяет охраняемые данные. Если код критический секции изменяет данные, то ставить блокировку на запись
	//а где чтение, ставить блокировку на чтение и запись.
	c map[string]int
}

func (c *Counter) CountMe() map[string]int {
	c.rw.Lock()
	defer c.rw.Unlock()
	return c.c
}

func (c *Counter) CountMeAgain() map[string]int {
	c.rw.RLock() //Не заблокируют map если не было блокировки на запись
	defer c.rw.RUnlock()
	return c.c
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++ //Критическая секция, которая помогает избежать гонки потоков. Потоки выстроятся в очередь
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

func main() {

	var wg sync.WaitGroup
	var counter uint64

	for i := 0; i < 10; i++ {
		wg.Add(1)
		k := i
		go func() {
			defer wg.Done()
			fmt.Printf("%d goroutine wirking...\n", k)
			time.Sleep(time.Second * 10)
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)
			}

		}()
	}

	wg.Wait()
	fmt.Printf("all done, counter = %d\n", counter)

	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(c.Value(key))
}
