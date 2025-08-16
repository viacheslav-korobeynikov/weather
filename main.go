package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)

	// Создали Reader
	data := strings.NewReader("Привет! Я поток данных")
	// Блок потока данных из Reader
	byteData := make([]byte, 4)
	for {
		_, err := data.Read(byteData)
		fmt.Printf("%q", byteData)
		// Проверка на то, что файл закончился
		if err == io.EOF {
			break
		}
	}
}
