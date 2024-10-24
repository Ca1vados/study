
## Содержание

1. [Ошибки в Go](#ошибки-в-go)
   - Заведение собственных ошибок
   - Обработка ошибок
   - Использование `panic` и `recover`
   - Пример обработки ошибок
2. [Структуры в Go](#структуры-в-go)
   - Определение и объявление
   - Методы структур: значение и указатель
3. [Считывание с клавиатуры в Go](#считывание-с-клавиатуры-в-go)
   - Считывание до нажатия Enter
   - Считывание с разделителем
   - Считывание до специального символа
4. [Пример программы](#пример-программы)

## Ошибки в Go

### Примеры обработки ошибок на языке Python и Golang
```python
data = ""
try:
	data = curl.ReadFromServer("8.8.8.8")
	n_data = atoi(data)
catch:
	print("Err")
	return

print(data)
```

```go
data, err := curl.ReadFromServer("8.8.8.8")
if err != nil {
	fnt.Println(err.Error())
	return
}

n_data, err := atoi(data)
if err != nil {
	fnt.Println(err.Error())
	return
}

fmt.Println(data)
```

### Заведение собственных ошибок

Создайте собственные ошибки для более детальной обработки:

```go
import "errors"

var ErrCustom = errors.New("custom error occurred")
```

### Обработка ошибок

Обработка ошибок в Go обычно выполняется с помощью проверки возвращаемого значения `error`:

```go
func doSomething() error {
    // Код, который может вызвать ошибку
    return nil // или ошибка
}

err := doSomething()
if err != nil {
    fmt.Println("Ошибка:", err.Error())
} else {
    fmt.Println("Успешное выполнение")
}
```

### Использование `panic` и `recover`

`panic` используется для критических ошибок, когда программа не может продолжать выполнение. `recover` позволяет перехватить панику:

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Восстановление после паники:", r)
        }
    }()
    mightPanic()
    fmt.Println("Этот код выполнится, если не произойдет паники")
}

func mightPanic() {
    panic("произошла критическая ошибка")
}
```

### Пример обработки ошибок

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println("Результат:", result)
}
```

## Структуры в Go

### Определение и объявление

Структуры используются для объединения полей:

```go
type Person struct {
    Name string
    Age  int
    pass string
}

func New(name string, age int) (Person, error) {
	p := Person{
		Name: name,
		Age: age, 
		pass: name + "123456"
	}
	if age < 0 {
		return p, errors.New("Возраст не може быть отрицательным!")
	}
	return p, nil
}
```

```go
p, err := New()
if err != nil {
	fmt.Println(err.Error())
	return
}

fmt.Println("Имя:", p.Name)
fmt.Println("Имя:", p.pass)
```
### Методы структур: значение и указатель

Методы можно объявлять как для копий структур, так и для указателей:

- Метод для копии структуры:

```go
func (p Person) GetPass() {
    return p.pass
}
```
- Вызов метода

```go
p.GetPass()
```
- Метод для указателя на структуру:

```go
func (p *Person) SetAge(age int) {
    p.Age = age
}
```

Использование указателей позволяет изменять поля структуры.

## Считывание с клавиатуры в Go

### Считывание до нажатия Enter

Используйте `bufio.NewReader` для считывания до нажатия Enter:

```go
import (
    "bufio"
    "fmt"
    "os"
)

reader := bufio.NewReader(os.Stdin)
fmt.Print("Введите строку: ")
input, _ := reader.ReadString('\n')
fmt.Println("Вы ввели:", input)
```

**Пример ввода:** `Hello, Go!`  
**Выход:** `Вы ввели: Hello, Go!`

### Считывание с разделителем

Вы можете использовать `Scan` с разделителем:

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Split(bufio.ScanWords)
fmt.Println("Введите слова, разделенные пробелами:")
for scanner.Scan() {
    fmt.Println("Вы ввели слово:", scanner.Text())
}
```

**Пример ввода:** `Go is awesome`  
**Выход:**  
```
Вы ввели слово: Go
Вы ввели слово: is
Вы ввели слово: awesome
```

### Считывание до специального символа

Считывание до специального символа возможно с `bufio.Reader`:

```go
reader := bufio.NewReader(os.Stdin)
fmt.Print("Введите до символа '!': ")
input, _ := reader.ReadString('!')
fmt.Println("Вы ввели:", input)
```

**Пример ввода:** `Hello, Go!`  
**Выход:** `Вы ввели: Hello, Go!`

## Пример программы

```go
package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type InputHandler struct {
}

func (ih InputHandler) ReadString() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите строку: ")
    text, _ := reader.ReadString('\n')
    return strings.TrimSpace(text)
}

func (ih InputHandler) ReadInt() (int, error) {
    fmt.Print("Введите целое число: ")
    var num int
    _, err := fmt.Scan(&num)
    if err != nil {
        return 0, errors.New("не удалось считать целое число")
    }
    return num, nil
}

func (ih InputHandler) ReadFloat() (float64, error) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите число с плавающей точкой: ")
    input, _ := reader.ReadString('\n')
    number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
    if err != nil {
        return 0, errors.New("не удалось считать число с плавающей точкой")
    }
    return number, nil
}

func main() {
    ih := InputHandler{}

    str := ih.ReadString()
    fmt.Printf("Вы ввели строку: %s\n", str)

    num, err := ih.ReadInt()
    if err != nil {
        fmt.Println("Ошибка при вводе целого числа:", err)
    } else {
        fmt.Printf("Вы ввели целое число: %d\n", num)
    }

    floatNum, err := ih.ReadFloat()
    if err != nil {
        fmt.Println("Ошибка при вводе числа с плавающей точкой:", err)
    } else {
        fmt.Printf("Вы ввели число с плавающей точкой: %f\n", floatNum)
    }
}
```

Этот пример показывает, как можно использовать структуры для обработки ввода и обработки ошибок, демонстрируя гибкость и модульность кода в Go.