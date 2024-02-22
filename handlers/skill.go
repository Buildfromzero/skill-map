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

func (handler *SkillHandler) RegisterApis(r *gin.Engine) {
	skillGroup := r.Group(handler.groupName)

	skillGroup.GET("skills/", handler.List)
	skillGroup.POST("skills/", handler.Create)
	skillGroup.GET("skills/:skillid/", handler.Detail)
	skillGroup.DELETE("skills/:skillid/", handler.Delete)
	skillGroup.PATCH("skills/:skillid/", handler.Update)
}

func (handler *SkillHandler) Create(ctx *gin.Context) {

	inputData := common.NewSkillCreationInput()

	err := ctx.BindJSON(&inputData)

	if err != nil {
		common.BadResponse(ctx, "Failed to bind data")
		return
	}

	newSkill, err := handler.skillManager.Create(inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to create user")
		return
	}

	ctx.JSON(http.StatusOK, newSkill)
}

func (handler *SkillHandler) List(ctx *gin.Context) {

	allSkills, err := handler.skillManager.List()

	if err != nil {
		common.BadResponse(ctx, "failed to get all skills")
		return
	}

	ctx.JSON(http.StatusOK, allSkills)
}

func (handler *SkillHandler) Detail(ctx *gin.Context) {

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

func (handler *SkillHandler) Delete(ctx *gin.Context) {

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

func (handler *SkillHandler) Update(ctx *gin.Context) {

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

	user, err := handler.skillManager.Update(skillId, inputData)

	if err != nil {
		common.BadResponse(ctx, "failed to update user")
		return
	}

	ctx.JSON(http.StatusOK, user)
}
