package service

import (
	"GO-CRM-API-SHOPDEV/internal/repo"
	"GO-CRM-API-SHOPDEV/pkg/response"
)

//import "GO-CRM-API-SHOPDEV/internal/repo"
//
//type UserService struct {
//	userRepo *repo.UserRepo
//}
//
//func NewUserService() *UserService {
//	return &UserService{
//		userRepo: repo.NewUserRepo(),
//	}
//}
//
//func (us *UserService) GetInfoUseService() string {
//
//	return us.userRepo.GetInFoUser()
//}

type IUserservice interface {
	Register(email string, purpose string) int
}
type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserservice {

	return &userService{userRepo: userRepo}
}

func (us *userService) Register(email string, purpose string) int {

	//1 check mail exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}
