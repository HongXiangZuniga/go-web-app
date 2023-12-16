package auth

type Repository interface {
	AuthUser(email, password string) (*bool, error)
	GetSalt(email string) (*string, error)
}
