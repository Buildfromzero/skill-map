package managers

import "github.com/buildfromzero/skill-map/models"

type UserManager struct {
}

func NewUserManagerFrom() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(user *models.User) (*models.User, error) {

	return nil, nil
}
