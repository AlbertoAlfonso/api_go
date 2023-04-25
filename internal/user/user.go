package user

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	Create(user User) (User, error)
	Update(id int, user User) (User, error)
	Delete(id int) error
}
