package repo

import (
	"fmt"

	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/internal/database"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetUserByID() string {
// 	return " Truong dep trai"
// }

// dùng INTERFACE
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userReppository struct {
	sqlc *database.Queries
}

// GetUserByEmail implements IUserRepository.
func (up *userReppository) GetUserByEmail(email string) bool {
	// row := global.Mdb.Table("go_crm_user").Where("usr = ?", email).First(&models.GoCrmUserV2{}).RowsAffected
	fmt.Println("SQL C")
	user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}
	return user.UsrID != 0

}

func NewUserRepository() IUserRepository {
	return &userReppository{
		sqlc: database.New(global.Mdbc),
	}
}
