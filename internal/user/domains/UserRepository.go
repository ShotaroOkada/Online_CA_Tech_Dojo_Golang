package domains

// UserRepository is interface
type UserRepository interface {
	Create(name string) (token string, err error)
	Get(token string) (*User, error)
	Update(user User) error
}
