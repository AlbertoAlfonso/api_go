package user

import (
	"errors"
	"sync"
)

type InMemoryUserRepository struct {
	mu     sync.Mutex
	users  []User
	nextID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  []User{},
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) GetAll() ([]User, error) {
	return r.users, nil
}

func (r *InMemoryUserRepository) GetByID(id int) (User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (r *InMemoryUserRepository) Create(user User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)

	return user, nil
}

func (r *InMemoryUserRepository) Update(id int, user User) (User, error) {
	for index, existingUser := range r.users {
		if existingUser.ID == id {
			user.ID = id
			r.users[index] = user
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (r *InMemoryUserRepository) Delete(id int) error {
	for index, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:index], r.users[index+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
