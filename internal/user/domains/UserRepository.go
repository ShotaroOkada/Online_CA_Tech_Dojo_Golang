package domains

// UserRepository is interface
type UserRepository interface {
	Create(user User) (string, error)
	Get() (User, error)
	Update(user User) error
}
