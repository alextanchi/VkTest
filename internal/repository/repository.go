package repository

import (
	"strings"
	"vktest/internal/domain"
	"vktest/internal/models"

	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type Repository interface {
	AddActor(actor domain.Actor) error
	UpdateActor(actor domain.Actor) error
	DeleteActor(id string) error
	ListActor() ([]*ListActor, error)
	AddMovie(movie domain.Movie) (string, error)
	UpdateMovie(movie domain.Movie) error
	DeleteMovie(id string) error
	ListMovie(filter models.ListMovieFilter) ([]domain.Movie, error)
	SearchMovie(fragment models.SearchMovieFilter) ([]domain.Movie, error)

	AddMovieActor(movieId string, actorIds []string) error
}

type Cinematheque struct {
	db *sql.DB
}

func NewCinematheque(db *sql.DB) Repository { //конструктор

	return &Cinematheque{
		db: db,
	}
}

// ConnectDb добавляем подключение к базе
func ConnectDb() (*sql.DB, error) {
	//перенести креды в конфиги(env)
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres", "password", "host.docker.internal", 5432, "cinematheque")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	log.Println("подключились к БД")

	return db, nil
}

func (c Cinematheque) AddActor(actor domain.Actor) error {

	id := uuid.NewString()

	querySql := `INSERT INTO actor
    (id,
     name, 
     surname,
     sex,
     birth_date
     ) VALUES ($1, $2, $3, $4, $5)`

	_, err := c.db.Exec(querySql,
		id,
		actor.Name,
		actor.Surname,
		actor.Sex,
		actor.BirthDate,
	)

	if err != nil {
		return err
	}
	return nil
}

func (c Cinematheque) UpdateActor(actor domain.Actor) error {

	//избежать sql инъекций посредством плейсхолдеров
	querySql := "UPDATE actor SET "
	if actor.Name != "" {
		querySql += "name = '" + actor.Name + "',"
	}
	if actor.Surname != "" {
		querySql += "surname = '" + actor.Surname + "',"
	}
	if actor.Sex != "" {
		querySql += "sex = '" + actor.Sex + "',"
	}
	if actor.BirthDate != "" {
		querySql += " birth_date = '" + actor.BirthDate + "',"
	}
	querySql = querySql[:len(querySql)-1] + " WHERE id = $1"
	log.Println("Запрос: ", querySql)
	_, err := c.db.Exec(querySql, actor.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c Cinematheque) DeleteActor(id string) error {
	_, err := c.db.Exec("DELETE FROM actor WHERE id = $1", id)
	return err
}

type ListActor struct {
	ActorId          *string
	ActorName        *string
	ActorSurname     *string
	ActorSex         *string
	ActorBirthDate   *string //можно передлать в формат даты
	MovieId          *string
	MovieTitle       *string
	MovieDescription *string
	MovieReleaseDate *string //можно передлать в формат даты
	MovieRating      *float64
}

func (c Cinematheque) ListActor() ([]*ListActor, error) {

	querySql := `SELECT
     a.id,
     a.name,
     a.surname,
     a.sex,
     a.birth_date,
     m.id,
     m.title,
     m.description,
     m.release_date,
     m.rating
 FROM actor a
     LEFT JOIN movie_actor ma on a.id = ma.actor_id
     LEFT JOIN movie m on m.id = ma.movie_id`

	rows, err := c.db.Query(querySql)
	if err != nil {

		return nil, err
	}

	defer rows.Close()

	listactor := make([]*ListActor, 0)
	for rows.Next() {

		a := &ListActor{}

		err = rows.Scan(
			&a.ActorId,
			&a.ActorName,
			&a.ActorSurname,
			&a.ActorSex,
			&a.ActorBirthDate,
			&a.MovieId,
			&a.MovieTitle,
			&a.MovieDescription,
			&a.MovieReleaseDate,
			&a.MovieRating,
		)
		if err == sql.ErrNoRows {
			return []*ListActor{}, nil
		}

		if err != nil {
			return nil, err
		}
		listactor = append(listactor, a)

	}

	return listactor, nil
}

func (c Cinematheque) AddMovie(movie domain.Movie) (string, error) {

	id := uuid.NewString()

	querySql := `INSERT INTO movie
    (id,
     title, 
     description,
     release_date,
     rating
     ) VALUES ($1, $2, $3, $4, $5)`

	_, err := c.db.Exec(querySql,
		id,
		movie.Title,
		movie.Description,
		movie.ReleaseDate,
		movie.Rating,
	)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (c Cinematheque) AddMovieActor(movieId string, actorIds []string) error {
	querySql := `INSERT INTO
		    movie_actor (movie_id, actor_id)
			VALUES ($1, $2)`

	for _, el := range actorIds {
		_, err := c.db.Exec(querySql, movieId, el)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c Cinematheque) UpdateMovie(movie domain.Movie) error {
	//избежать sql инъекций посредством плейсхолдеров
	querySql := "UPDATE movie SET "
	if movie.Title != "" {
		querySql += "title = '" + movie.Title + "',"
	}
	if movie.Description != "" {
		querySql += "description = '" + movie.Description + "',"
	}
	if movie.ReleaseDate != "" {
		querySql += "release_date = '" + movie.ReleaseDate + "',"
	}

	//предобразуем тип (будем учитывать только одну цифру после запятой)
	ratingString := strconv.FormatFloat(movie.Rating, 'f', 1, 64)
	if movie.Rating != 0 {
		querySql += "rating = " + ratingString + ","
	}

	querySql = querySql[:len(querySql)-1] + " WHERE id = $1"
	log.Println("Запрос: ", querySql)

	_, err := c.db.Exec(querySql, movie.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c Cinematheque) DeleteMovie(id string) error {
	_, err := c.db.Exec("DELETE FROM movie WHERE id = $1", id)
	return err
}

func (c Cinematheque) ListMovie(filter models.ListMovieFilter) ([]domain.Movie, error) {

	querySql := `SELECT 
    	id,
     	title, 
     	description,
     	release_date,
     	rating
		FROM movie`

	//избежать sql инъекций посредством плейсхолдеров
	if filter.Sorted != nil && filter.SortType != nil {
		querySql += " order by " + *filter.Sorted + " " + *filter.SortType
	} else {
		querySql += " order by rating desc"
	}

	rows, err := c.db.Query(querySql)
	if err != nil {

		return nil, err
	}

	defer rows.Close()

	movies := make([]domain.Movie, 0)
	for rows.Next() {

		m := domain.Movie{}
		err = rows.Scan(
			&m.Id,
			&m.Title,
			&m.Description,
			&m.ReleaseDate,
			&m.Rating,
		)
		if err == sql.ErrNoRows {
			return []domain.Movie{}, nil
		}

		if err != nil {

			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}
func (c Cinematheque) SearchMovie(fragment models.SearchMovieFilter) ([]domain.Movie, error) {

	querySql := `SELECT 
    	m.id,
     	title, 
     	description,
     	release_date,
     	rating
		FROM movie m
		LEFT JOIN movie_actor mv on m.id = mv.movie_id
		LEFT JOIN actor a on a.id = mv.actor_id `

	//избежать sql инъекций посредством плейсхолдеров
	if fragment.MovieName != nil {
		querySql += "where title ilike '%" + strings.ToLower(*fragment.MovieName) + "%'"
	}
	if fragment.ActorName != nil {
		querySql += "where a.name ilike '%" + strings.ToLower(*fragment.ActorName) + "%'"
	}

	rows, err := c.db.Query(querySql)
	if err != nil {

		return nil, err
	}

	defer rows.Close()

	movies := make([]domain.Movie, 0)
	for rows.Next() {

		m := domain.Movie{}
		err = rows.Scan(
			&m.Id,
			&m.Title,
			&m.Description,
			&m.ReleaseDate,
			&m.Rating,
		)
		if err == sql.ErrNoRows {
			return []domain.Movie{}, nil
		}

		if err != nil {

			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}
