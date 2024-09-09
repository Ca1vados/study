package main

import (
	"fmt"
	"strings"
)

// функция разбивающая строку на символы

func splitLine(str string) []string {
	str = strings.ToLower(str)
	symbols := strings.Split(str, "")
	return symbols
}

func countVowelConsonant(str string) {
	var vowel = make([]string, 0)     // создаем слайс для глассных
	var consonant = make([]string, 0) // создаем слайс для согласных
	for _, i := range splitLine(str) {
		switch i {
		case "а", "е", "ё", "и", "о", "у", "ы", "э", "ю", "я", "a", "e", "i", "o", "u", "y":
			vowel = append(vowel, i)
		case "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z", "б", "в", "г", "д", "ж", "з", "й", "к", "л", "м", "н", "п", "р", "с", "т", "ф", "х", "ц", "ч", "ш", "щ":
			consonant = append(consonant, i)
		}
	}
	numVowel := len(vowel)
	numConsonant := len(consonant)
	fmt.Printf("Количество гласных: %d\nКоличество согласных: %d", numVowel, numConsonant)
}

func main() {
	var line string //переменнаая в которую сохраняем результат ввода пользователя
	fmt.Print("Введите строку: ")
	fmt.Scanf("%s", &line)
	fmt.Println(splitLine(line))
	countVowelConsonant(line)
}
