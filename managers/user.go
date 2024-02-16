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

func (userMgr *UserManager) Get(id string) (models.User, error) {
	user := models.User{}

	database.DB.First(&user, id)

	return user, nil
}

func (userMgr *UserManager) Update(userId string, userData *common.UserUpdateInput) (*models.User, error) {

	user := models.User{}

	database.DB.First(&user, userId)

	// user.FullName = userData.FullName
	// user.Email = userData.Email

	// database.DB.Save(&user)

	database.DB.Model(&user).Updates(models.User{FullName: userData.FullName, Email: userData.Email})

	return &user, nil
}

func (userMgr *UserManager) Delete(id string) error {
	user := models.User{}

	database.DB.First(&user, id)
	database.DB.Delete(&user)
	return nil
}
