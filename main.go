package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

var mtx = sync.Mutex{}
var money = atomic.Int32{} // usd
var bank = atomic.Int32{}  // usd

func payHandler(w http.ResponseWriter, r *http.Request) {
	// str := "Новый платёж обработан!"
	// b := []byte(str)

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Fail to read HTTP body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	pamentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("Fail to convert HTTP body to int:", err)
		return
	}

	mtx.Lock()
	if money.Load()-int32(pamentAmount) >= 0 {
		money.Add(int32(-pamentAmount))
		fmt.Println("Оплата прошла успешно:", money.Load())
	} else {
		fmt.Println("Не хватает денег на проведение оплаты!")
	}
	mtx.Unlock()

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Fail to read HTTP body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("Fail to convert HTTP body to int:", err)
		return
	}

	mtx.Lock()
	if money.Load() >= int32(saveAmount) {
		money.Add(int32(-saveAmount))

		bank.Add(int32(saveAmount))

		fmt.Println("Новое значение переменной money", money.Load())
		fmt.Println("Новое значение переменной bank", bank.Load())
	} else {
		fmt.Println("Не хватает денег, чтобы положить в копилку!")
	}
	mtx.Unlock()
}

/* func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	str := "Hello, World!)"
	b := []byte(str)

	_, err := w.Write(b)

	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err)
	} else {
		fmt.Println("Я корректно обработал HTTP запрос!")
	}
} */

func main() {
	money.Add(1000)

	// http.HandleFunc("/", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	fmt.Println("Запускаю HTTP сервер!")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error:", err)
	}

	fmt.Println("Программа закончила своё выполнение!")
}
