package models

import (
	"errors"
	"time"
)

type Movie struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
}
type CreateMovieRequest struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ReleaseDate string   `json:"release_date"`
	Rating      float64  `json:"rating"`
	Actors      []string `json:"actors"`
}

type Actor struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birth_date"`
}
type ListActorResponse struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Surname   string  `json:"surname"`
	Sex       string  `json:"sex"`
	BirthDate string  `json:"birth_date"`
	Movies    []Movie `json:"movies"`
}

type ListMovieFilter struct {
	Sorted   *string `json:"sorted"`
	SortType *string `json:"sortType"`
}
type SearchMovieFilter struct {
	MovieName *string `json:"movieName"`
	ActorName *string `json:"actorName"`
}

func (a Actor) ValidateActor() error {
	_, err := time.Parse("2006-01-02", a.BirthDate)
	if err != nil {
		return errors.New("Неверный формат поля Birth_date")
	}
	return nil
}
