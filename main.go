package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// 1. Настройки подключения
	url := "http://localhost/myFirst/hs/access/"
	username := "admin"
	password := "" // Пароль пустой

	// 2. Создаем клиент с таймаутом
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 3. Формируем запрос
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// 4. Устанавливаем авторизацию
	req.SetBasicAuth(username, password)

	// 5. Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// 6. Читаем тело ответа полностью
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения тела ответа: %v", err)
	}

	// 7. УДАЛЯЕМ UTF-8 BOM (те самые символы 'ï' / EF BB BF)
	// Это критично для корректной работы с ответами из 1С
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf"))

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Printf("Сервер вернул ошибку %d: %s", resp.StatusCode, string(bodyBytes))
		return
	}

	// 8. Декодируем чистый JSON
	// Используем map, если структура JSON заранее неизвестна
	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v. \nСырые данные: %s", err, string(bodyBytes))
	}

	// Выводим результат
	fmt.Println("Данные успешно получены:")
	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}
