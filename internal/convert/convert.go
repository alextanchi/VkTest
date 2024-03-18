package convert

import (
	"vktest/internal/domain"
	"vktest/internal/models"
	"vktest/internal/repository"
)

func ActorModelToDomain(actor models.Actor) domain.Actor {
	return domain.Actor{
		Id:        actor.Id,
		Name:      actor.Name,
		Surname:   actor.Surname,
		Sex:       actor.Sex,
		BirthDate: actor.BirthDate,
	}
}
func MovieModelToDomain(movie models.CreateMovieRequest) domain.Movie {
	return domain.Movie{
		Id:          movie.Id,
		Title:       movie.Title,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Rating:      movie.Rating,
	}
}
func ListActorToDomain(listActor []*repository.ListActor) map[domain.Actor][]domain.Movie {
	collection := make(map[domain.Actor][]domain.Movie)
	for _, el := range listActor {
		var a domain.Actor
		var m domain.Movie
		if el.ActorId != nil {
			a.Id = *el.ActorId
			a.Name = *el.ActorName
			a.Surname = *el.ActorSurname
			a.Sex = *el.ActorSex
			a.BirthDate = *el.ActorBirthDate

		}
		if _, ok := collection[a]; !ok {
			collection[a] = make([]domain.Movie, 0)
		}
		if el.MovieId != nil {
			m.Id = *el.MovieId
			m.Title = *el.MovieTitle
			m.Description = *el.MovieDescription
			m.ReleaseDate = *el.MovieReleaseDate
			m.Rating = *el.MovieRating
			collection[a] = append(collection[a], m)
		}

	}
	return collection
}

func ListActorToResponse(m map[domain.Actor][]domain.Movie) []models.ListActorResponse {
	listActorResponse := make([]models.ListActorResponse, 0)
	for k, el := range m {
		var l models.ListActorResponse
		l.Id = k.Id
		l.Name = k.Name
		l.Surname = k.Surname
		l.Sex = k.Sex
		l.BirthDate = k.BirthDate
		l.Movies = make([]models.Movie, 0)
		for _, movie := range el {
			var actorMovie models.Movie
			actorMovie.Id = movie.Id
			actorMovie.Title = movie.Title
			actorMovie.Description = movie.Description
			actorMovie.ReleaseDate = movie.ReleaseDate
			actorMovie.Rating = movie.Rating
			l.Movies = append(l.Movies, actorMovie)
		}
		listActorResponse = append(listActorResponse, l)
	}

	return listActorResponse

}
