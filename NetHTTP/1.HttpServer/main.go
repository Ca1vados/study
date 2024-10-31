package main

import (
	"FirstHttp/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Создание сервера")
	http.HandleFunc("/", handlers.Hello)
	http.HandleFunc("/sort", handlers.Sort)

	http.HandleFunc("/sort/quick_sort", handlers.QuickSortAPI) // !!

	http.HandleFunc("/sort/bubble_sort", handlers.BubbleSort)
	http.HandleFunc("/sort/insertion_sort", handlers.InsertionSort)
	http.HandleFunc("/help", handlers.Help)

	http.HandleFunc("/gen_user", handlers.GetUserApi) // !1

	http.ListenAndServe(":8080", nil)
}
