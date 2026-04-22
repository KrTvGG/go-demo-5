package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("Новый проект")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды от 1 до 4")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)

	r := strings.NewReader("Привет, я поток данных")
	block := make([]byte, 4)
	for {
		_, err := r.Read(block)
		fmt.Printf("%q\n", block)
		if err == io.EOF {
			break
		}
	}
}