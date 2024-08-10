package usecase

import (
	"errors"

	"github.com/Nahom-Derese/Learning_Go/Task-7/task-manager/domain"
)

// TaskUseCase is a struct that defines the usecases for the task entity
type UserUseCase struct {
	UserRepo domain.UserRepository
	// contextTimeout time.Duration
}

// NewTaskUseCase is a function that returns a new TaskUseCase
func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
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
