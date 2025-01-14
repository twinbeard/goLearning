package service

import "github.com/twinbeard/goLearning/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserByID() string {
	return us.userRepo.GetUserByID()
}
