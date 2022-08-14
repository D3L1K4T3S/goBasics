package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"math/rand"
	"time"
)

type Figure struct {
	age  int8
	name string
}

type Person struct {
	Name string
	Age  int
}

func main() {
	number := "vasya"
	fmt.Println(number)

	var a complex64 = 10 + 10i + 10 + 10i
	var b complex64 = 12 + 12i

	num := rand.Intn(100 - 50)

	var point *int = &num

	fmt.Println(point, *point)
	fmt.Println(&num, num)
	square(point)
	fmt.Println(&num, num)

	fmt.Println()
	fmt.Println(a + b)
	fmt.Println(time.Now().String())

	number = "Hello world!"
	fmt.Println(number)

	var input string = "privet"
	_ = input

	fmt.Println(SumMult(3199, 100000))

	fmt.Println(calculate(10, 5, divide))

	if compare(5, 10) {
		fmt.Println("True")
	}

	ob := Figure{20, "Egor"}
	ob.Square()

	//Метка
Label:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 20; j++ {
			//fmt.Println(i, " - ", j)
			if i >= 10 {
				continue Label //Метка к внешнему циклу
			}
		}
	}

	array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array)

	people := [3]Person{Person{"Egor", 20}, Person{"Rostislav", 19}, Person{"Slava", 18}}

	//Итерация по массиву 1-ое значение индекс элемента, 2-ое значение массива
	for ind, person := range people {
		fmt.Println("[", ind, "] - Person: ", person, "\n")
	}

	//Слайс увеличивает в 2 раза объем массива если не хватает объема изначального массива при добавлении нового элемента

	Person := struct {
		Name, LastName, BirthDate string
	}{
		Name:      "NoName",
		LastName:  "NoName",
		BirthDate: time.Now().String(),
	}

	Person.Name = "Egor"
	Person.LastName = "Zhelagin"

	fmt.Println(Person)

	firstSlice := []int{5, 6, 7, 8}
	showAllElemnts(firstSlice...) //Передача всгео слайса - равнозначно передаче сразу всех элементов по отдельности
	//... показаывает передачу всей структуры
	//Чтобы преобразовать слайс в указатель на массив необходимо указаить правлиьное значение объема слайса

	intarray := (*[4]int)(firstSlice)
	fmt.Println(intarray)

	//Когда передает слайсы в функции то это равнозначно передачи массива через указатель. Если в функции происходит выделение
	//под новый массив памяти то меняется у слайса поинтер

	intArr := [...]int{1, 2, 3, 4, 5}
	intSlice := intArr[1:3]
	fmt.Println(intSlice)

	//Капасити он выделяет = количество элементов в массиве - кол-во элементов пропущенных выбором. Все значения [:]

	pointsMap := map[string]int{
		"xx": 123,
		"yy": 456,
	}
	p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)
	fmt.Println(&p1)

}

func Multiple(x, y int) int {
	return x * y
}

// Использование тегов
type Point struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func (p Point) method() {
	fmt.Println("call Point method")
}

// Методы
func (object Figure) Square() {
	fmt.Println(object.age * 2)
}
func compare(x, y int) bool {
	return x < y
}

func square(x *int) {
	*x *= *x
}

func SumMult(first, second int) (int, int) {
	return first + second, first * second
}

func divide(x, y int) int {
	return x / y
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

func createDivider(divider int) func(x int) int {
	divideFunc := func(x int) int { return x / divider }
	return divideFunc
}

func showAllElemnts(values ...int) {
	for _, val := range values {
		fmt.Println("Value: ", val)
	}
	fmt.Println()
}
