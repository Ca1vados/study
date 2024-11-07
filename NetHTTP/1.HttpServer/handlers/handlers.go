package handlers

import (
	"FirstHttp/pkg/sortfuncs"
	"FirstHttp/usecase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Вызвали обработчик hello")
	bh := []byte("Hello Creator!")
	w.Write(bh)
}

func Sort(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Вызвали обработчик hello")
	bh := []byte("Доступны следующие варианты сортировки: быстрая, вставкой, пузырьками. подробности: localhost:8080/help")
	w.Write(bh)
}

func QuickSortAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Вызвали обработчик quickSort")
	unsortedBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	unsortedString := string(unsortedBody)

	sortedArr, err := usecase.QuickSortUsecase(unsortedString)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(sortedArr)
	}
}

func BubbleSort(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Вызвали обработчик bubbleSort")
	unsortedBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	unsortedString := string(unsortedBody)
	var unsortedIntArray []int
	unsortedStrings := strings.Split(unsortedString, " ")
	for _, s := range unsortedStrings {
		num, err := strconv.Atoi(s)
		if err == nil {
			unsortedIntArray = append(unsortedIntArray, num)
		}
	}
	sortedArray := sortfuncs.BubbleSort(unsortedIntArray)
	fmt.Fprint(w, sortedArray)
}

func InsertionSort(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Вызвали обработчик insertionSort")
	unsortedBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	unsortedString := string(unsortedBody)
	var unsortedIntArray []int
	unsortedStrings := strings.Split(unsortedString, " ")
	for _, s := range unsortedStrings {
		num, err := strconv.Atoi(s)
		if err == nil {
			unsortedIntArray = append(unsortedIntArray, num)
		}
	}
	sortedArray := sortfuncs.InsertionSort(unsortedIntArray)
	fmt.Fprint(w, sortedArray)
}

func Help(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Вызвали обработчик help")
	fmt.Fprintf(w, "Доступны следующие ссылки:\n\n")
	fmt.Fprintf(w, "localhost:8080/ : приветствие\n")
	fmt.Fprintf(w, "localhost:8080/sort : варианты сортировки\n")
	fmt.Fprintf(w, "localhost:8080/sort/quick_sort - отправить в body raw числа через пробел - вернуть массив, отсортированный быстрой сортировкой в порядке возрастания\n")
	fmt.Fprintf(w, "localhost:8080/sort/bubble_sort - отправить в body raw числа через пробел - вернуть массив, отсортированный пузырьками в порядке возрастания\n")
	fmt.Fprintf(w, "localhost:8080/sort/insertion_sort - отправить в body raw числа через пробел - вернуть массив, отсортированный методом вставки в порядке возрастания\n")
	fmt.Fprintf(w, "localhost:8080/help : команды бота\n")
}

func GetUserApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Вызвали обработчик bubbleSort")
	unsortedBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	login := string(unsortedBody)

	user, err := usecase.GenerateUserPass(login)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}
