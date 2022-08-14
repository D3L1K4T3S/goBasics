package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()                                          //Каналы связаны, если верхний канал закрывается, то нижний тоже закроется, если они разные
	ctx, cancel := context.WithCancel(ctx)                               //Возвращает канал и функцию завершения контекста
	ctx, cancel = context.WithDeadline(ctx, time.Now().Add(time.Second)) //Когда пройдет время контекст отменяется
	ctx, cancel = context.WithTimeout(ctx, time.Second)                  //То же самое что deadline только сам прибавляет к time.Now время указанно

	//Иммитация закрытик контекста. Через 100 миллисекунд закроется контекст
	// defer cancel() //Закрытие после выполнение запроса

	//Если через 100 миллисекунд не выполнится запрос, он закроется. Можно увеличить время ожидания запроса
	go func() {
		err := cancelRequest(ctx)
		if err != nil {
			cancel()
		}
	}()

	doRequest(ctx, "http://yandex.ru")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return fmt.Errorf("fail request")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response completed, status code: %d", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("request takes too long")
	}
}
