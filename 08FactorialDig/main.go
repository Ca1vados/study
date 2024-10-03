package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Вычисляет факториал числа n
func factorial(n int) int {
	if n == 0 {
		return 1 // Факториал 0 равен 1
	} else {
		return n * factorial(n-1) // Рекурсивный вызов для вычисления факториала
	}
}

func main() {
	var ErrCustom = errors.New("Введенное значение не является целым неотрицательным числом")
	scanner := bufio.NewScanner(os.Stdin) // Создание нового сканера для ввода

	for {
		fmt.Print("Введите целое неотрицательное число для вычисления факториала или -1 для выхода из программы: ")
		scanner.Scan() // Считывание строки из ввода - происходит отчистка буфера
		input := scanner.Text()
		dig, err := strconv.Atoi(strings.TrimSpace(input)) // Преобразование введенной строки в целое число

		if err != nil {
			fmt.Println(ErrCustom.Error()) // Вывод сообщения об ошибке, если ввод некорректен
			continue
		} else if dig < -1 {
			fmt.Println("Введите число больше нуля") // Сообщение, если число меньше -1
		} else if dig == -1 {
			break // Выход из цикла, если введено -1
		} else {
			fmt.Printf("Факториал числа %d: %d\n", dig, factorial(dig)) // Вычисление и вывод факториала
		}
	}
}
