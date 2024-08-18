package usecase

import (
	"errors"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/infrastructure"
)

// TaskUseCase is a struct that defines the usecases for the task entity
type UserUseCase struct {
	UserRepo domain.UserRepository
	// contextTimeout time.Duration
}

// NewTaskUseCase is a function that returns a new TaskUseCase
func NewUserUseCase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUseCase{UserRepo: userRepo}
}

// PromoteUser implements UserRepository.
func (u *UserUseCase) PromoteUser(username string) (domain.UserRes, error) {
	user, exists := u.FindByUsername(username)

	if !exists {
		return domain.UserRes{}, errors.New("user not found")
	}

	if user.Role == "admin" {
		return domain.UserRes{}, errors.New("user is already an admin")
	}

	user.Role = "admin"
	updatedUser, err := u.Update(user.Username, user)

	if err != nil {
		return domain.UserRes{}, errors.New("error promoting user")
	}

	return updatedUser, nil
}

// Delete implements UserRepository.
func (u *UserUseCase) Delete(username string) error {

	err := u.UserRepo.Delete(username)

	if err != nil {
		return errors.New("user not found")
	}

	return nil

}

// DeleteAll implements UserRepository.
func (u *UserUseCase) DeleteAll() error {

	err := u.UserRepo.DeleteAll()

	if err != nil {
		return errors.New("error deleting all users")
	}

	return nil
}

// FindAll implements UserRepository.
func (u *UserUseCase) FindAll() []domain.UserRes {

	users := u.UserRepo.FindAll()
	res := make([]domain.UserRes, len(users))

	for i, user := range users {
		userRes := domain.UserRes{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		}
		res[i] = userRes
	}

	return res

}

// FindByUsername implements UserRepository.
func (u *UserUseCase) FindByUsername(username string) (domain.UserRes, bool) {

	user, exists := u.UserRepo.FindByUsername(username)

	userRes := domain.UserRes{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	return userRes, exists

}

// FindUser implements UserRepository.
func (u *UserUseCase) FindUser(id string) (domain.UserRes, error) {

	user, err := u.UserRepo.FindUser(id)

	if err != nil {
		return domain.UserRes{}, errors.New("user not found")
	}

	userRes := domain.UserRes{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	return userRes, err

}

// FindUserByUsername implements UserRepository.
func (u *UserUseCase) FindUserByUsername(username string) (domain.UserRes, error) {

	user, exist := u.UserRepo.FindByUsername(username)

	if !exist {
		return domain.UserRes{}, errors.New("user not found")
	}

	userRes := domain.UserRes{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	return userRes, nil
}

// CreateUser implements UserRepository.
func (u *UserUseCase) CreateUser(user domain.User) (domain.UserRes, error) {

	user, err := u.UserRepo.Save(user)

	if err != nil {
		return domain.UserRes{}, errors.New("user not saved")
	}

	userRes := domain.UserRes{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	return userRes, nil
}

// Save implements UserRepository.
func (u *UserUseCase) Update(username string, user domain.UserRes) (domain.UserRes, error) {

	newUser, err := u.UserRepo.Update(username, domain.User{ID: user.ID, Username: user.Username, Role: user.Role})

	if err != nil {
		return domain.UserRes{}, errors.New("user not saved")
	}

	userRes := domain.UserRes{
		ID:       newUser.ID,
		Username: newUser.Username,
		Role:     newUser.Role,
	}

	return userRes, nil

}

// Login implements UserRepository.
func (u *UserUseCase) Login(username string, password string) (string, error) {

	userInDB, exists := u.UserRepo.FindByUsername(username)

	if !exists {
		return "", errors.New("user not found")
	}

	// compare password
	err := infrastructure.ComparePasswords(userInDB.Password, password)

	if err != nil {
		return "", errors.New("invalid password " + err.Error())
	}

	// generate token
	token, err := infrastructure.GenerateToken(userInDB.Username, userInDB.ID.Hex())

	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil

}

// Login implements UserRepository.
func (u *UserUseCase) Signup(username string, password string) (domain.UserRes, error) {

	// hash password
	hash, err := infrastructure.HashPassword(password)

	if err != nil {
		return domain.UserRes{}, errors.New("error hashing password")
	}

	// if no users in the db, create an admin user
	users := u.FindAll()

	role := "regular"
	if len(users) == 0 {
		role = "admin"
	}

	// check if user already exists
	_, exists := u.FindByUsername(username)

	if exists {
		return domain.UserRes{}, errors.New("Username already exists")
	}

	userData := domain.User{Username: username, Password: string(hash), Role: role}
	newUser, error := u.CreateUser(userData)

	res := domain.UserRes{
		ID:       newUser.ID,
		Username: newUser.Username,
		Role:     newUser.Role,
	}

	if error != nil {
		return domain.UserRes{}, errors.New("error creating user")
	}

	return res, nil

}
