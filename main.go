package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost/myFirst/hs/access/"

	// 1. Создаем клиент (рекомендуется задавать таймаут)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 2. Создаем объект запроса
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Устанавливаем базовую авторизацию (логин, пароль)
	// Если пароля нет, передаем пустую строку ""
	req.SetBasicAuth("admin", "")

	// 4. Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// 5. Проверка статуса (например, 401 Unauthorized, если логин неверен)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка сервера: %s", resp.Status)
	}

	// 6. Читаем JSON (замените map на свою структуру для удобства)
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	fmt.Printf("Данные получены: %+v\n", result)
}
