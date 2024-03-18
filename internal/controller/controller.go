package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"vktest/internal/convert"
	"vktest/internal/models"
	"vktest/internal/service"
)

type Controller interface {
	AddActor(w http.ResponseWriter, r *http.Request)
	UpdateActor(w http.ResponseWriter, r *http.Request)
	DeleteActor(w http.ResponseWriter, r *http.Request)
	ListActor(w http.ResponseWriter, r *http.Request)
	AddMovie(w http.ResponseWriter, r *http.Request)
	UpdateMovie(w http.ResponseWriter, r *http.Request)
	DeleteMovie(w http.ResponseWriter, r *http.Request)
	ListMovie(w http.ResponseWriter, r *http.Request)
	SearchMovie(w http.ResponseWriter, r *http.Request)
}
type CinemathequeController struct {
	useCase service.Service
}

func NewController(srv service.Service) Controller {
	return &CinemathequeController{
		useCase: srv,
	}
}
func (c CinemathequeController) AddActor(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-AddActor")
	body, _ := io.ReadAll(r.Body)
	actor := models.Actor{}
	err := json.Unmarshal(body, &actor)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = actor.ValidateActor() //пример валидации по полю birth_date
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.useCase.AddActor(convert.ActorModelToDomain(actor))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (c CinemathequeController) UpdateActor(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-UpdateActor")
	body, _ := io.ReadAll(r.Body)
	actor := models.Actor{}
	err := json.Unmarshal(body, &actor)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.useCase.UpdateActor(convert.ActorModelToDomain(actor))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c CinemathequeController) DeleteActor(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-DeleteActor")
	body, _ := io.ReadAll(r.Body)
	actor := models.Actor{}
	err := json.Unmarshal(body, &actor)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.useCase.DeleteActor(actor.Id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c CinemathequeController) ListActor(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-ListActor")
	actors, err := c.useCase.ListActor()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(convert.ListActorToResponse(actors))
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (c CinemathequeController) AddMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-AddMovie")
	body, _ := io.ReadAll(r.Body)
	movie := models.CreateMovieRequest{}
	err := json.Unmarshal(body, &movie)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.useCase.AddMovie(convert.MovieModelToDomain(movie), movie.Actors)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (c CinemathequeController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-UpdateMovie")
	body, _ := io.ReadAll(r.Body)
	movie := models.CreateMovieRequest{}
	err := json.Unmarshal(body, &movie)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.useCase.UpdateMovie(convert.MovieModelToDomain(movie))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (c CinemathequeController) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-DeleteMovie")
	body, _ := io.ReadAll(r.Body)
	movie := models.CreateMovieRequest{}
	err := json.Unmarshal(body, &movie)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.useCase.DeleteMovie(movie.Id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func (c CinemathequeController) ListMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-ListMovie")

	body, _ := io.ReadAll(r.Body)
	filter := models.ListMovieFilter{}
	err := json.Unmarshal(body, &filter)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	movies, err := c.useCase.ListMovie(filter)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(movies)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
func (c CinemathequeController) SearchMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkpoint-SearchMovie")
	body, _ := io.ReadAll(r.Body)
	fragment := models.SearchMovieFilter{}
	err := json.Unmarshal(body, &fragment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	movies, err := c.useCase.SearchMovie(fragment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(movies)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
