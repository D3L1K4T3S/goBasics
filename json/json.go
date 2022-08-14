package json

//Докачать gjson
//Посмотреть gotify, extremecode канал на ютубу
//Скачать Как писать кодогенераторы в Go. Дмитрий Смотров
//Команда curl

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsBlocked bool   `json:"isBlocked"`
}

func main() {
	sv := User{
		Name:      "Egor",
		Age:       20,
		IsBlocked: true,
	}
	boolVar, _ := json.Marshal(sv)
	fmt.Println(string(boolVar))
	var dat map[string]interface{}
	json.Unmarshal(boolVar, &dat)
	fmt.Println(dat)
}
