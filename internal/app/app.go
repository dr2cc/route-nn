package app

import (
	"fmt"
	"route-nn/internal/client"
	"route-nn/internal/config"
)

func Run(cfg *config.Config) error {
	// 1. Используем ваш конструктор
	apiClient := client.NewClient(cfg)

	// 2. Вызываем метод (BOM и авторизация уже внутри)
	var result map[string]interface{}
	if err := apiClient.GetJSON(cfg.Url, &result); err != nil {
		return err // или log.Fatalf
	}

	// 3. Вывод
	fmt.Println("Данные успешно получены:")
	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}

	return nil

	// // Создаем клиент с таймаутом
	// client := &http.Client{
	// 	Timeout: cfg.Timeout * time.Second, // ♊ предложил 10
	// }

	// // 3. Формируем запрос
	// req, err := http.NewRequest("GET", cfg.Url, nil)
	// if err != nil {
	// 	log.Fatalf("Ошибка создания запроса: %v", err)
	// }

	// // 4. Устанавливаем авторизацию
	// req.SetBasicAuth(cfg.Username, cfg.Password)

	// // 5. Выполняем запрос
	// resp, err := client.Do(req)

	// if err != nil {
	// 	log.Fatalf("Ошибка выполнения запроса: %v", err)
	// }
	// defer resp.Body.Close()

	// // 6. Читаем тело ответа полностью
	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Ошибка чтения тела ответа: %v", err)
	// }

	// // 7. УДАЛЯЕМ UTF-8 BOM (те самые символы 'ï' / EF BB BF)
	// // Это критично для корректной работы с ответами из 1С
	// bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf"))

	// // Проверяем статус ответа
	// if resp.StatusCode != http.StatusOK {
	// 	// Формируем ошибку, включая статус и тело ответа для контекста
	// 	return fmt.Errorf("server returned error %d: %s", resp.StatusCode, string(bodyBytes))
	// }

	// // 8. Декодируем чистый JSON
	// // Используем map, если структура JSON заранее неизвестна
	// var result map[string]interface{}
	// if err := json.Unmarshal(bodyBytes, &result); err != nil {
	// 	log.Fatalf("Ошибка парсинга JSON: %v. \nСырые данные: %s", err, string(bodyBytes))
	// }

	// // Выводим результат
	// fmt.Println("Данные успешно получены:")
	// for key, value := range result {
	// 	fmt.Printf("%s: %v\n", key, value)
	// }

	// return nil
}
