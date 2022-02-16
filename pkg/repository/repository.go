package repository

type Auth struct {
}

type User struct {
}

type Repository struct {
	Auth
	User
}

func NewRepository() *Repository {
	return &Repository{}
}
