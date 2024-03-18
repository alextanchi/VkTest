package service

import (
	"vktest/internal/convert"
	"vktest/internal/domain"
	"vktest/internal/models"
	"vktest/internal/repository"
)

type Service interface {
	AddActor(actor domain.Actor) error
	UpdateActor(actor domain.Actor) error
	DeleteActor(id string) error
	ListActor() (map[domain.Actor][]domain.Movie, error)
	AddMovie(movie domain.Movie, actorIds []string) error
	UpdateMovie(movie domain.Movie) error
	DeleteMovie(id string) error
	ListMovie(filter models.ListMovieFilter) ([]domain.Movie, error)       //как реализовать сортировку? Несколькими методами?
	SearchMovie(fragment models.SearchMovieFilter) ([]domain.Movie, error) //здесь два разных метода нужно?
}
type CinemathequeService struct {
	storage repository.Repository
}

func NewService(repos repository.Repository) Service {

	return CinemathequeService{
		storage: repos,
	}
}

func (s CinemathequeService) AddActor(actor domain.Actor) error {

	return s.storage.AddActor(actor)
}

func (s CinemathequeService) UpdateActor(actor domain.Actor) error {

	return s.storage.UpdateActor(actor)
}

func (s CinemathequeService) DeleteActor(id string) error {

	return s.storage.DeleteActor(id)
}

func (s CinemathequeService) ListActor() (map[domain.Actor][]domain.Movie, error) {
	listActor, err := s.storage.ListActor()
	if err != nil {
		return map[domain.Actor][]domain.Movie{}, err
	}

	return convert.ListActorToDomain(listActor), err

}

func (s CinemathequeService) AddMovie(movie domain.Movie, actorsIds []string) error {

	id, err := s.storage.AddMovie(movie)
	if err != nil {
		return err
	}

	return s.storage.AddMovieActor(id, actorsIds)
}

func (s CinemathequeService) UpdateMovie(movie domain.Movie) error {

	return s.storage.UpdateMovie(movie)
}

func (s CinemathequeService) DeleteMovie(id string) error {

	return s.storage.DeleteMovie(id)
}

func (s CinemathequeService) ListMovie(filter models.ListMovieFilter) ([]domain.Movie, error) {
	return s.storage.ListMovie(filter)
}
func (s CinemathequeService) SearchMovie(fragment models.SearchMovieFilter) ([]domain.Movie, error) {
	return s.storage.SearchMovie(fragment)
}
