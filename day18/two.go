package main

import (
	"fmt"
)

type User struct{
	ID int
	Name string
	Email string 
}


// userRepository

type UserRepository interface{
	 Create(u *User) error
	 FindById(Id int) (*User, error)
	 FindByEmail(email string) (*User, error)
	 Delete(Id int) error
}


// UserService
type UserService struct{
	repo UserRepository
}

// NewUserService

func NewUserService(r UserRepository) *UserService{
	return &UserService{
		repo: r,
	}
}

// Mongodb 

// mongodb repository
type MongodbRepository struct{

}

func (m MongodbRepository)Create(u *User)error{
	fmt.Println("")
}
