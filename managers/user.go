package managers

import (
	"errors"

	"github.com/buildfromzero/skill-map/common"
	"github.com/buildfromzero/skill-map/models"
	"github.com/buildfromzero/skill-map/storage"
)

type UserManager interface {
	Create(inputData *common.UserCreationInput) (*models.User, error)
	List() ([]models.User, error)
	Get(id string) (*models.User, error)
	Update(userId string, inputData *common.UserUpdateInput) (*models.User, error)
	Delete(id string) error
	AddNewSkill(userId string, inputData *common.CompetenceInput) (*models.User, error)
}

type userManager struct {
	// DatabaseDriver
	// dbClient
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (userMgr *userManager) Create(inputData *common.UserCreationInput) (*models.User, error) {
	newUser := &models.User{FullName: inputData.FullName, Email: inputData.Email}
	storage.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}

	userMgr.prefetchUser(newUser)
	return newUser, nil
}

func (userMgr *userManager) List() ([]models.User, error) {
	users := []models.User{}
	storage.DB.Preload("Competence").Preload("Competence.Skill").Find(&users)
	return users, nil
}

func (userMgr *userManager) Get(id string) (*models.User, error) {
	user := models.NewUser()

	storage.DB.First(&user, id)
	if user.ID == 0 {
		return nil, errors.New("no user found")
	}

	userMgr.prefetchUser(user)

	return user, nil
}

func (userMgr *userManager) Update(userId string, inputData *common.UserUpdateInput) (*models.User, error) {

	user := models.NewUser()

	storage.DB.First(user, userId)

	if user.ID == 0 {
		return nil, errors.New("no user found")
	}

	storage.DB.Model(&user).Updates(models.User{FullName: inputData.FullName, Email: inputData.Email})

	userMgr.prefetchUser(user)

	return user, nil
}

func (userMgr *userManager) Delete(id string) error {
	user := models.NewUser()

	storage.DB.First(user, id)

	if user.ID == 0 {
		return errors.New("no user found")
	}

	storage.DB.Delete(user)
	return nil
}

func (userMgr *userManager) AddNewSkill(userId string, inputData *common.CompetenceInput) (*models.User, error) {
	user := models.NewUser()

	storage.DB.First(user, userId)

	if user.ID == 0 {
		return nil, errors.New("no user found")
	}

	skill := models.NewSkill()

	storage.DB.First(skill, inputData.Skill)

	competenceObj := models.NewCompetence()
	competenceObj.User = *user
	competenceObj.Skill = *skill
	competenceObj.Rank = inputData.Rank
	storage.DB.Create(competenceObj)
	userMgr.prefetchUser(user)

	return user, nil
}

func (userMgr *userManager) prefetchUser(user *models.User) {
	storage.DB.Model(user).Preload("Competence").Preload("Competence.Skill").Find(user)
}
