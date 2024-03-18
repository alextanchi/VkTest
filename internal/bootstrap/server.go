package bootstrap

import (
	"net/http"
	"time"
	"vktest/internal/controller"
)

type Server struct {
	httpServer *http.Server
	cnt        controller.Controller
}

// NewServer объединили  пакет http и контроллер
func NewServer(cnt controller.Controller) Server {
	return Server{
		httpServer: &http.Server{
			Addr:           ":8080",
			MaxHeaderBytes: 1 << 20,          //1MB
			ReadTimeout:    10 * time.Second, //10 сек
			WriteTimeout:   10 * time.Second,
		},
		cnt: cnt,
	}
}

// InitRoutes инициализируем все наши эндпоинты
func (s Server) InitRoutes() {

	http.HandleFunc("/actor", s.cnt.AddActor)
	http.HandleFunc("/actor/update", s.cnt.UpdateActor)
	http.HandleFunc("/actor/delete", s.cnt.DeleteActor)
	http.HandleFunc("/actor/list", s.cnt.ListActor)
	http.HandleFunc("/movie", s.cnt.AddMovie)
	http.HandleFunc("/movie/update", s.cnt.UpdateMovie)
	http.HandleFunc("/movie/delete", s.cnt.DeleteMovie)
	http.HandleFunc("/movie/list", s.cnt.ListMovie)
	http.HandleFunc("/movie/search", s.cnt.SearchMovie)

}
func (s Server) Run() error {
	return s.httpServer.ListenAndServe()
}
