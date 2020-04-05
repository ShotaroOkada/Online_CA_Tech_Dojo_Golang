package domains

// UserRepository is interface
type UserRepository interface {
	Create(user *User) (token string, err error)
	Get(token string) (*User, error)
	Update(user User) error
}
