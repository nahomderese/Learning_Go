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

type UserRepository interface {
	Save(user User) (User, error)
	Update(username string, user User) (User, error)
	FindUser(id string) (User, error)
	FindByUsername(username string) (User, bool)
	FindAll() []User
	Delete(username string) error
	DeleteAll() error
}
