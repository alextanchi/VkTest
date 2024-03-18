package domain

type Movie struct {
	Id          string
	Title       string
	Description string
	ReleaseDate string
	Rating      float64
}

type Actor struct {
	Id        string
	Name      string
	Surname   string
	Sex       string
	BirthDate string
}

type MovieActor struct {
	MovieId string
	ActorId string
}

type Role struct {
	Id   string
	Role string
}

type Account struct {
	Id       string
	Login    string
	Password string
}
