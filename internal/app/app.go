package app

import (
	"fmt"
	"html/template"
	"net/http"
	"route-nn/internal/client"
	"route-nn/internal/config"
	"route-nn/internal/delivery/html"
	"time"
)

// // Вшиваем папку templates прямо в бинарник.
// // Директива должна быть ровно над переменной.
// //
// //go:embed route-nn/internal/delivery/html/templates/index.html
// var templateFS embed.FS

// Структура для передачи данных в HTML-шаблон
type PageData struct {
	Result      string
	CurrentDate string // дата по умолчанию
}

// Run #3
func Run(cfg *config.Config) error {
	apiClient := client.NewClient(cfg)

	tmpl := template.Must(template.ParseFS(html.Files, "templates/index.html"))

	// 1. Отображение главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Передаем структуру с текущей датой, чтобы input не был пустым
		data := PageData{
			CurrentDate: time.Now().Format("2006-01-02"),
		}
		tmpl.Execute(w, data)
		// tmpl.Execute(w, nil)
	})

	// 2. Обработка кнопки "Рассчитать"
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Получаем данные из формы
		code1 := r.FormValue("code1")
		code2 := r.FormValue("code2")
		date := r.FormValue("date")

		// Делаем запрос к 1C (используем существующий метод)
		var apiData map[string]interface{}
		// Здесь можно модифицировать URL или параметры под запрос к 1C
		if err := apiClient.GetJSON(cfg.Url, &apiData); err != nil {
			http.Error(w, "Ошибка получения данных из 1C", http.StatusInternalServerError)
			return
		}

		// ВАЖНО: Здесь проводим "сложные вычисления" на стороне Go
		// Для примера просто выведем полученные ключи
		calcRes := fmt.Sprintf("Обработано для %s и %s на дату %s\nДанные 1С: %v",
			code1, code2, date, apiData)

		// Передаем результат и выбранную дату обратно,
		// чтобы форма не сбрасывалась к "сегодняшнему" числу после расчета
		tmpl.Execute(w, PageData{
			Result:      calcRes,
			CurrentDate: date,
		})
		// tmpl.Execute(w, PageData{Result: calcRes})
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	return http.ListenAndServe("localhost:8080", nil)
}

// // Run #2 with *resty.Client
// func Run(cfg *config.Config) error {
// 	apiClient := client.NewClient(cfg)

// 	// call GetJSON method
// 	var result map[string]interface{}
// 	if err := apiClient.GetJSON(cfg.Url, &result); err != nil {
// 		// return err
// 		log.Fatalf("GetJSON error: %s", err)
// 	}

// 	// output
// 	// fmt.Println("Данные успешно получены:")
// 	for key, value := range result {
// 		fmt.Printf("%s: %v\n", key, value)
// 	}

// 	return nil
// }

// // Run #1 with &http.Client
// func Run(cfg *config.Config) error {

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
// }
