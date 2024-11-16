package service

import "GO-CRM-API-SHOPDEV/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserSerivce() *UserService {

	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUseService() string {

	return us.userRepo.GetInFoUser();
}
