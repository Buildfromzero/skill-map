package handlers

import (
	"fmt"
	"net/http"

	"github.com/buildfromzero/skill-map/common"
	"github.com/buildfromzero/skill-map/managers"
	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	groupName    string
	skillManager managers.SkillManager
}

func NewSkillHandlerFrom(skillManager managers.SkillManager) *SkillHandler {
	return &SkillHandler{
		"api/",
		skillManager,
	}
}

func (handler *SkillHandler) RegisterEndpoints(r *gin.Engine) {
	skillGroup := r.Group(handler.groupName)
	// skill apis
	skillGroup.GET("skills/", handler.ListSkills)
	skillGroup.POST("skills/", handler.CreateSkill)
	skillGroup.GET("skills/:skillid/", handler.SkillDetail)
	skillGroup.DELETE("skills/:skillid/", handler.DeleteSkill)
	skillGroup.PATCH("skills/:skillid/", handler.UpdateSkill)
	// skill group apis
	skillGroup.GET("skill-groups/", handler.ListSkillGroups)
	skillGroup.POST("skill-groups/", handler.CreateSkillGroup)
	skillGroup.GET("skill-groups/:groupid/", handler.SkillGroupDetail)
	skillGroup.DELETE("skill-groups/:groupid/", handler.DeleteSkillGroup)
	skillGroup.PATCH("skill-groups/:groupid/", handler.UpdateSkillGroup)
}

func (handler *SkillHandler) CreateSkill(ctx *gin.Context) {

	inputData := common.NewSkillCreationInput()

	err := ctx.BindJSON(&inputData)

	if err != nil {
		common.BadResponse(ctx, "Failed to bind data")
		return
	}

	newSkill, err := handler.skillManager.Create(inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to create skill")
		return
	}

	ctx.JSON(http.StatusOK, newSkill)
}

func (handler *SkillHandler) ListSkills(ctx *gin.Context) {

	allSkills, err := handler.skillManager.List()

	if err != nil {
		common.BadResponse(ctx, "failed to get all skills")
		return
	}

	ctx.JSON(http.StatusOK, allSkills)
}

func (handler *SkillHandler) SkillDetail(ctx *gin.Context) {

	skillId, ok := ctx.Params.Get("skillid")

	if !ok {
		fmt.Println("invalid skill id")
	}
	skill, err := handler.skillManager.Get(skillId)

	if skill.ID == 0 {
		common.BadResponse(ctx, "no skill present")
		return
	}

	if err != nil {
		common.BadResponse(ctx, "failed to get skill")
	}

	ctx.JSON(http.StatusOK, skill)
}

func (handler *SkillHandler) DeleteSkill(ctx *gin.Context) {

	skillId, ok := ctx.Params.Get("skillid")

	if !ok {
		common.BadResponse(ctx, "invalid userid")
	}
	err := handler.skillManager.Delete(skillId)

	if err != nil {
		common.BadResponse(ctx, "failed to delete skill")
	}

	common.SuccessResponse(ctx, "deleted Skill")
}

func (handler *SkillHandler) UpdateSkill(ctx *gin.Context) {

	skillId, ok := ctx.Params.Get("skillid")

	if !ok {
		common.BadResponse(ctx, "failed to delete skill")
	}

	inputData := common.NewSkillUpdateInput()

	err := ctx.BindJSON(&inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to bind data")
		return
	}

	skill, err := handler.skillManager.Update(skillId, inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to update skill")
		return
	}

	ctx.JSON(http.StatusOK, skill)
}

func (handler *SkillHandler) CreateSkillGroup(ctx *gin.Context) {

	inputData := common.NewSkillGroupCreationInput()

	err := ctx.BindJSON(&inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to bind data")
		return
	}

	newSkillGroup, err := handler.skillManager.CreateGroup(inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to create Skill Group")
		return
	}

	ctx.JSON(http.StatusOK, newSkillGroup)
}

func (handler *SkillHandler) ListSkillGroups(ctx *gin.Context) {

	skillGroups, err := handler.skillManager.ListGroup()

	if err != nil {
		common.BadResponse(ctx, "failed to get all Skill Groups")
		return
	}

	ctx.JSON(http.StatusOK, skillGroups)
}

func (handler *SkillHandler) SkillGroupDetail(ctx *gin.Context) {

	groupId, ok := ctx.Params.Get("groupid")

	if !ok {
		fmt.Println("invalid group id")
	}
	skillGroup, err := handler.skillManager.GetGroup(groupId)

	if skillGroup.ID == 0 {
		common.BadResponse(ctx, "no group present with given id")
		return
	}

	if err != nil {
		common.BadResponse(ctx, "failed to get Skill Group")
	}

	ctx.JSON(http.StatusOK, skillGroup)
}

func (handler *SkillHandler) DeleteSkillGroup(ctx *gin.Context) {

	groupId, ok := ctx.Params.Get("groupid")

	if !ok {
		common.BadResponse(ctx, "invalid Group id")
	}
	err := handler.skillManager.DeleteGroup(groupId)

	if err != nil {
		common.BadResponse(ctx, "failed to delete Skill Group")
	}

	common.SuccessResponse(ctx, "deleted Skill Group")
}

func (handler *SkillHandler) UpdateSkillGroup(ctx *gin.Context) {

	groupId, ok := ctx.Params.Get("groupid")

	if !ok {
		common.BadResponse(ctx, "failed to get group id")
	}

	inputData := common.NewSkillGroupUpdateInput()

	err := ctx.BindJSON(&inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to bind data")
		return
	}

	skillGroup, err := handler.skillManager.UpdateGroup(groupId, inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to update Skill Group")
		return
	}

	ctx.JSON(http.StatusOK, skillGroup)
}
