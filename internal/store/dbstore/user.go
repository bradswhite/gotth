package dbstore

import (
	"errors"
	"gotth/internal/store"
)

type UserStore struct {
	Users []store.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		Users: []store.User{
			// set some default users
			{
				Email:    "1@example.com",
				Password: "password",
			},
		},
	}
}

func (s *UserStore) CreateUser(email string, password string) error {

	for _, user := range s.Users {
		if user.Email == email {
			return errors.New("user already exists")
		}
	}

	s.Users = append(s.Users, store.User{Email: email, Password: password})
	return nil
}

func (s *UserStore) GetUser(email string) (*store.User, error) {
	for _, user := range s.Users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
