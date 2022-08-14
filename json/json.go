package main

//Использование gjson и sjson фреймворков
import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"reflect"
	"strings"
)

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsBlocked bool   `json:"isBlocked"`
}

// Поиск в файле по значению array.#(field == value)#.field
// Выпешет первый который найдет, если >= то все которые больше или равны value
func main() {
	sv := User{
		Name:      "Egor",
		Age:       20,
		IsBlocked: true,
	}
	json, _ := json.Marshal(sv)
	fmt.Println(reflect.TypeOf(json))
	str := string(json)
	value := gjson.Get(str, "name")
	fmt.Println(value)
	tempByte := `{"name":{"first":"Egor","last":"Moisey"},"age":47, "array":[1,2,3,4]'}`
	fmt.Println(reflect.TypeOf(tempByte))
	str = string(tempByte)
	value = gjson.Get(str, "name.first")
	fmt.Println(value)
	fmt.Print("Размерность массива: ")
	fmt.Println(gjson.Get(str, "array.#"))

	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		} else {
			return strings.ToLower(json)
		}
		return json
	})

	//Использование модификатора, описанного выше @name_modifier:value
	fmt.Println(gjson.Get(str, "name.first|@case:upper"))

	//Парсинг json
	fmt.Println(gjson.Parse(str).Get("name"))

	if gjson.Valid(str) {
		panic("JSON ISN'T VALID")
	}

	result, ok := gjson.Parse(str).Value().(map[string]interface{})
	if !ok {
		panic("MAP ISN'T CORRECT")
	}

	fmt.Println(result)

	tmp := gjson.Parse(str)
	fmt.Println(reflect.TypeOf(tmp))

	str, _ = sjson.Set(str, "name.first", "Ivan")
	fmt.Println(str)
	//Добавление нового объекта
	str, _ = sjson.Set(str, "book", map[string]interface{}{"M.Bulgakov": "Master and Margarita"})
	fmt.Println(str)
}
