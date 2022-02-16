package repository

type User struct {
}

type Repository struct {
	User
}

func NewRepository() *Repository {
	return &Repository{}
}
