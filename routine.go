package main

import (
	"fmt"
	"time"
)

//Если гланвынй поток завершается то и звершаются все остальные. При это чтение из канала блочит поток. Каналы позволяют синхронизироваться поток
//Избежать гонок каналов. Канал неьлзя использовать если поток 1. Лочиться поток. Канал - это "мостик" между двумя потоками

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
	close(ch) //Закрытие канала, чтобы канал не лочился (Выход из цикла). Нельзя ничего отпралвтья в закрытый канал
}

// Использование селектов
// Разобраться с default в select, потому что он срабатывает каждый раз
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

	//Написать калькулятор на горутинах
	//ПРимер чтобы вычитывал сразу несколько паралелльно вычслений
}
