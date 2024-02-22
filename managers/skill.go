package managers

import (
	"errors"

	"github.com/buildfromzero/skill-map/common"
	"github.com/buildfromzero/skill-map/models"
	"github.com/buildfromzero/skill-map/storage"
)

type SkillManager interface {
	Create(inputData *common.SkillCreationInput) (*models.Skill, error)
	List() ([]models.Skill, error)
	Get(id string) (models.Skill, error)
	Update(id string, inputData *common.SkillUpdateInput) (*models.Skill, error)
	Delete(id string) error
}

type skillManager struct {
	// DatabaseDriver
	// dbClient
}

func NewSkillManager() SkillManager {
	return &skillManager{}
}

func (skillMgr *skillManager) Create(inputData *common.SkillCreationInput) (*models.Skill, error) {
	newSkillObj := &models.Skill{Name: inputData.Name}
	storage.DB.Create(newSkillObj)

	if newSkillObj.ID == 0 {
		return nil, errors.New("skill creation failed")
	}

	return newSkillObj, nil
}

func (skillMgr *skillManager) List() ([]models.Skill, error) {
	skillObj := []models.Skill{}

	storage.DB.Find(&skillObj)
	// TODO: handle errors
	return skillObj, nil
}

func (skillMgr *skillManager) Get(id string) (models.Skill, error) {
	skillObj := models.Skill{}

	storage.DB.First(&skillObj, id)

	return skillObj, nil
}

func (skillMgr *skillManager) Update(id string, inputData *common.SkillUpdateInput) (*models.Skill, error) {
	skillObj := models.Skill{}

	storage.DB.First(&skillObj, id)
	storage.DB.Model(&skillObj).Updates(models.Skill{Name: inputData.Name})
	// TODO: handle errors
	return &skillObj, nil
}

func (skillMgr *skillManager) Delete(id string) error {
	skillObj := models.Skill{}

	storage.DB.First(&skillObj, id)
	storage.DB.Delete(&skillObj)
	// TODO: handle errors
	return nil
}

func (skillMgr *skillManager) CreateGroup(inputData *common.SkillGroupCreationInput) (*models.SkillGroup, error) {
	newSkillGroupObj := &models.SkillGroup{Name: inputData.Name}
	storage.DB.Create(newSkillGroupObj)

	if newSkillGroupObj.ID == 0 {
		return nil, errors.New("skill group creation failed")
	}

	return newSkillGroupObj, nil
}

func (skillMgr *skillManager) ListGroup() ([]models.SkillGroup, error) {
	skillGroupObj := []models.SkillGroup{}

	storage.DB.Find(&skillGroupObj)
	// TODO: handle errors
	return skillGroupObj, nil
}

func (skillMgr *skillManager) GetGroup(id string) (models.SkillGroup, error) {
	skillGroupObj := models.SkillGroup{}

	storage.DB.First(&skillGroupObj, id)

	return skillGroupObj, nil
}

func (skillMgr *skillManager) UpdateGroup(id string, inputData *common.SkillGroupUpdateInput) (*models.SkillGroup, error) {
	skillGroupObj := models.SkillGroup{}

	storage.DB.First(&skillGroupObj, id)
	storage.DB.Model(&skillGroupObj).Updates(models.Skill{Name: inputData.Name})
	// TODO: handle errors
	return &skillGroupObj, nil
}

func (skillMgr *skillManager) DeleteGroup(id string) error {
	skillGroupObj := models.SkillGroup{}

	storage.DB.First(&skillGroupObj, id)
	storage.DB.Delete(&skillGroupObj)
	// TODO: handle errors
	return nil
}
