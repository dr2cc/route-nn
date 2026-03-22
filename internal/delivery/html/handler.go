// internal/delivery/html/handler.go
package html

// import (
// 	"html/template"
// 	"net/http"
// )

// type RnnUseCase interface {
// 	// FetchAndCalculate(param1, param2 string) (domain.Result, error)
// }

// type Handler struct {
// 	useCase  RnnUseCase
// 	template *template.Template
// }

// func NewHandler(uc RnnUseCase) *Handler {
// 	return &Handler{
// 		useCase:  uc,
// 		template: template.Must(template.ParseGlob("templates/*.html")),
// 	}
// }

// func (h *Handler) RenderMainPage(w http.ResponseWriter, r *http.Request) {
// 	// Вызывает UseCase, получает данные, делает ExecuteTemplate
// }
