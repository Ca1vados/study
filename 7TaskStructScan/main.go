package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputHandler struct { // Создаем некий тип InputHandler
}

func (ih InputHandler) ReadString() string { // создаем метод для нашего типа, выводящий строку
	reader := bufio.NewReader(os.Stdin) // создаем сканер reader, считывающий данные с консоли
	fmt.Print("Введите строку: ")
	text, _ := reader.ReadString('\n') // создаем переменную text в которую считываем данные до Энтера(\n) из консоли. Методом чего является ReadString?
	return strings.TrimSpace(text)     // функция возвращает строку text, обработанную функцией strings.TrimSpace (удаляет все ведущие и конечные пробелы из строки), уточнить что именно.
}

func (ih InputHandler) ReadInt() (int, error) { // создаем метод ReadInt для InputHandler, возвращающий int и ошибку
	fmt.Print("Введите целое число: ")
	var num int              // создаем интовую переменную
	_, err := fmt.Scan(&num) // сканируем из консоли в переменную num. Возможна ошибка (очевидно ввод чего-либо кроме целого числа)
	if err != nil {          // если происходит ошибка, обрабатываем её: возвращаем ноль и созданную нами текстовую ошибку
		return 0, errors.New("не удалось считать целое число")
	}
	return num, nil // возвращаем интовое значение num и nil (почему nil?)
}

func (ih InputHandler) ReadFloat() (float64, error) { // создаем метод ReadInt для InputHandler, возвращающий float64 и ошибку
	reader := bufio.NewReader(os.Stdin) // создаем сканер reader, считывающий данные с консоли
	fmt.Print("Введите число с плавающей точкой: ")
	input, _ := reader.ReadString('\n')                             // создаем переменную input в которую считываем данные до Энтера из консоли.
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // убираем лишние пробелы (вопрос выше) и конвертируем строку input в float64. Записываем в переменную number. в err заносим возможную ошибку.
	if err != nil {                                                 // если происходит ошибка, обрабатываем её: возвращаем ноль и созданную нами текстовую ошибку
		return 0, errors.New("не удалось считать число с плавающей точкой")
	}
	return number, nil // возвращаем float64 значение number и nil (почему nil)
}

func main() {
	ih := InputHandler{} // создаем переменную ih типа InputHandler

	str := ih.ReadString()                   // в переменную str вызываем метод ReadString для переменной ih типа InputHandler
	fmt.Printf("Вы ввели строку: %s\n", str) // выводим str в консоль

	num, err := ih.ReadInt() // в переменную num вызываем метод ReadInt для переменной ih типа InputHandler. Метод также возвращает ошибку
	if err != nil {
		fmt.Println("Ошибка при вводе целого числа:", err) // т.к метод возвращает ошибку, обрабатываем её (внутри метода была ошибка на сканере, а это ошибка метода?)
	} else {
		fmt.Printf("Вы ввели целое число: %d\n", num) // если ошибки нет, выводим переменную num в консоль
	}

	floatNum, err := ih.ReadFloat() // в переменную floatNum вызываем метод ReadFloat для переменной ih типа InputHandler. Метод также возвращает ошибку
	if err != nil {                 // обрабатываем ошибку
		fmt.Println("Ошибка при вводе числа с плавающей точкой:", err)
	} else {
		fmt.Printf("Вы ввели число с плавающей точкой: %f\n", floatNum) // выводим в консоль floatNum
	}
}
