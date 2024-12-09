package repo

//type UserRepo struct {
//}
//
//func NewUserRepo() *UserRepo {
//
//	return &UserRepo{}
//}
//
//func (ur *UserRepo) GetInFoUser() string {
//
//	return "TanTD"
//}

// Interface version

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

func (u userRepository) GetUserByEmail(email string) bool {

	return true
}

func NewUserRepository() IUserRepository {

	return &userRepository{}
}
