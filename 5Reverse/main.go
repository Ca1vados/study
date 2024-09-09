package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// функция, которая принимает строку, создает из неё слайс, переворачивает его и возвращает перевернутую строку
func reverse(enter string) string {
	strSlice := make([]string, len(enter)) // создаем слайс в который разобьем строку пользователя
	strSlice = strings.Split(enter, "")
	revStrSlice := make([]string, 0) // создаем слайс в который поместим перевернутую строку
	for i := len(strSlice) - 1; i >= 0; i-- {
		revStrSlice = append(revStrSlice, strSlice[i])
	}
	var str string // создаем строку, которую соберем из перевернутого слайса
	str = strings.Join(revStrSlice[:], "")
	return str
}

func main() {
	var line string //переменнаая в которую сохраняем результат ввода пользователя
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
	}

	fmt.Println(reverse(line))
}
