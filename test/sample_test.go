package test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup")
	res := m.Run()
	fmt.Println("tear-down")
	os.Exit(res)
}

func TestAdd(t *testing.T) {

	fmt.Println("SETUP")

	//Чистка хвостов, в случае если удаление данных не произошло из-за паники в каком-то тесте (Досмотреть)
	t.Cleanup(func() {
		fmt.Println("TEARDOWN ON CLEANUP")
	})
	//setup
	//Здесь вставка данных из базы данных
	t.Run("easy", func(t *testing.T) {
		t.Parallel() //Параллельный запуск тестов при независимости тестов друг от друга
		t.Log("easy")
		var x, y, result = 2, 2, 4
		realResult := Add(x, y)
		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}
	})
	t.Run("normal", func(t *testing.T) {
		t.Parallel()
		t.Log("normal")
		var x, y, result = -2, 2, 0
		realResult := Add(x, y)
		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}
	})
	//tearDown
	//Удаление тестовых данных в базе данных

	//Можно вкалдывать t.run друг в друга - 1 легкий тест 2 легкий тест и тд. И выйдет тройная иерархия тестов
	//Запуск одного определенного теста go test -v -run TestMultiple/easy или /easy

}
