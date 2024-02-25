package managers

import (
	"errors"

	"github.com/buildfromzero/skill-map/common"
	"github.com/buildfromzero/skill-map/models"
	"github.com/buildfromzero/skill-map/storage"
)

type SkillManager interface {
	// Skill
	Create(inputData *common.SkillCreationInput) (*models.Skill, error)
	List() ([]models.Skill, error)
	Get(id string) (*models.Skill, error)
	Update(id string, inputData *common.SkillUpdateInput) (*models.Skill, error)
	Delete(id string) error
	// Skill Group
	CreateGroup(inputData *common.SkillGroupCreationInput) (*models.SkillGroup, error)
	ListGroup() ([]models.SkillGroup, error)
	GetGroup(id string) (*models.SkillGroup, error)
	UpdateGroup(id string, inputData *common.SkillGroupUpdateInput) (*models.SkillGroup, error)
	DeleteGroup(id string) error
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

	skillGroup := models.NewSkillGroup()

	storage.DB.First(skillGroup, inputData.Group)
	newSkillObj.SkillGroup = *skillGroup

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

func (skillMgr *skillManager) Get(id string) (*models.Skill, error) {
	skillObj := models.NewSkill()

	storage.DB.First(skillObj, id)

	if skillObj.ID == 0 {
		return nil, errors.New("no skill found")
	}

	return skillObj, nil
}

func (skillMgr *skillManager) Update(id string, inputData *common.SkillUpdateInput) (*models.Skill, error) {
	skillObj := &models.Skill{}

	storage.DB.First(skillObj, id)

	if skillObj.ID == 0 {
		return nil, errors.New("item does not exist")
	}

	storage.DB.Model(skillObj).Updates(models.Skill{Name: inputData.Name})
	// TODO: handle errors
	return skillObj, nil
}

func (skillMgr *skillManager) Delete(id string) error {
	skillObj := models.Skill{}

	storage.DB.First(&skillObj, id)
	if skillObj.ID == 0 {
		return errors.New("item does not exist")
	}
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
	skillGroups := []models.SkillGroup{}

	storage.DB.Model(&skillGroups).Preload("Skills").Find(&skillGroups)
	// TODO: handle errors
	return skillGroups, nil
}

func (skillMgr *skillManager) GetGroup(id string) (*models.SkillGroup, error) {
	skillGroupObj := models.NewSkillGroup()

	storage.DB.First(skillGroupObj, id)
	// storage.DB.Model(skillGroupObj).Preload("Skills").Find(skillGroupObj)

	if skillGroupObj.ID == 0 {
		return nil, errors.New("item does not exist")
	}

	return skillGroupObj, nil
}

func (skillMgr *skillManager) UpdateGroup(id string, inputData *common.SkillGroupUpdateInput) (*models.SkillGroup, error) {

	skillGroupObj := models.NewSkillGroup()

	storage.DB.First(skillGroupObj, id)

	if skillGroupObj.ID == 0 {
		return nil, errors.New("item does not exist")
	}

	// skillMapping, err := skillMgr.getSkillList(inputData.Skills)

	// if err != nil {
	// 	return nil, err
	// }

	storage.DB.Model(skillGroupObj).Updates(models.Skill{Name: inputData.Name})
	// storage.DB.Model(skillGroupObj).Association("Skills").Replace(skillMapping)
	// TODO: handle errors
	return skillGroupObj, nil
}

// func (skillMgr *skillManager) getSkillList(inputList []int) ([]*models.Skill, error) {

// 	skills := []*models.Skill{}

// 	for _, id := range inputList {
// 		skill := models.NewSkill()
// 		storage.DB.First(skill, id)
// 		if skill.ID == 0 {
// 			return nil, fmt.Errorf("skill with id %v not exists", id)
// 		}
// 		skills = append(skills, skill)
// 	}

// 	return skills, nil
// }

func (skillMgr *skillManager) DeleteGroup(id string) error {
	skillGroupObj := models.NewSkillGroup()

	storage.DB.First(skillGroupObj, id)

	if skillGroupObj.ID == 0 {
		return errors.New("item does not exist")
	}
	storage.DB.Delete(skillGroupObj)
	// TODO: handle errors
	return nil
}
