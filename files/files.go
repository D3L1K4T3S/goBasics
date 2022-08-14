package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//ls -lah - посмотреть права доступа у файлов 0777

func files() {
	writeTiFile()
	AppendFile()
	readFile()

	c := exec.Command("top")
	c.Stdin = os.Stdin   //Канал входа данных
	c.Stdout = os.Stdout //Канал выхода данных
	c.Stderr = os.Stderr //Канал ошибок
	c.Run()

	c = exec.Command("cat", "test.txt")
}

func AppendFile() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic("don't open file")
	}
	if _, err = f.WriteString("Egor is hero\n"); err != nil {
		panic("don't write to file")
	}
}

func writeTiFile() {
	data := []byte("my name is Egor\n")
	ioutil.WriteFile("test.txt", data, 0777)
}

func readFile() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
