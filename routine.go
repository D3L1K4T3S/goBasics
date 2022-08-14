package main

import (
	"fmt"
	"time"
)

//Если главный поток завершается, то и завершаются все остальные. При этом чтение из канала блокирует поток. Каналы позволяют синхронизироваться поток.
//Предотвратить гонки каналов. Канал нельзя использовать, если поток нет рутин. Канал - это "мостик" между двумя потоками

func main() {
	ch := make(chan int)      //Создание канала
	go Say("Hello world", ch) //СОздание потока
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("==========================")
	data := make(chan int)
	exit := make(chan int)

	go func() {
		exit <- 0
	}()
	selectOne(data, exit)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)
}

func Say(word string, ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println(word)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Microsecond * 5)
		ch <- i
	}
	close(ch) //Закрытие канала, чтобы канал не заблокировался(Выход из цикла). Нельзя ничего отправлять в закрытый канал
}

// Использование select

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
			/*default:
			fmt.Println("wait...")
			time.Sleep(3 * time.Second)

			*/
		}
	}
}
