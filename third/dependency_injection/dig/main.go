package main

import (
	"fmt"

	"go.uber.org/dig"
)

type User struct {
	Id int
}

type Repo interface {
	Get(int) (*User, error)
}

type UserRepo struct{}

func NewUserRepo() Repo {
	return &UserRepo{}
}

func (u *UserRepo) Get(id int) (*User, error) {
	return &User{Id: id}, nil
}

type UserService struct {
	repo Repo
}

func NewUserService(repo Repo) *UserService {
	return &UserService{repo: repo}
}

func main() {
	container := dig.New()
	_ = container.Provide(NewUserRepo)
	_ = container.Provide(NewUserService)
	_ = container.Invoke(func(s *UserService) {
		user, err := s.repo.Get(100)
		fmt.Println(user, err) // &{100} <nil>
	})
}
