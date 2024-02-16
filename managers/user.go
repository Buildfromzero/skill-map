package managers

import (
	"errors"

	"github.com/buildfromzero/skill-map/common"
	"github.com/buildfromzero/skill-map/database"
	"github.com/buildfromzero/skill-map/models"
)

type UserManager struct {
	// dbClient
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(userData *common.UserCreationInput) (*models.User, error) {
	newUser := &models.User{FullName: userData.FullName, Email: userData.Email}
	database.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}

	return newUser, nil
}

func (userMgr *UserManager) List() ([]models.User, error) {
	users := []models.User{}
	database.DB.Find(&users)
	return users, nil
}
