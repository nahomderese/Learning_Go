package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

type UserRes struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
}

//go:generate mockery --name UserRepository
type UserRepository interface {
	Save(user User) (User, error)
	Update(username string, user User) (User, error)
	FindUser(id string) (User, error)
	FindByUsername(username string) (User, bool)
	FindAll() []User
	Delete(username string) error
	DeleteAll() error
}

//go:generate mockery --name UserUsecase
type UserUsecase interface {
	PromoteUser(username string) (UserRes, error)
	Delete(username string) error
	DeleteAll() error
	FindAll() []UserRes
	FindByUsername(username string) (UserRes, bool)
	FindUser(id string) (UserRes, error)
	FindUserByUsername(username string) (UserRes, error)
	CreateUser(user User) (UserRes, error)
	Login(username string, password string) (string, error)
	Signup(username string, password string) (UserRes, error)
	Update(username string, user UserRes) (UserRes, error)
}
