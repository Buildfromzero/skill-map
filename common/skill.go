package common

type SkillCreationInput struct {
	Name string `json:"name"`
}

type SkillUpdateInput struct {
	Name string `json:"name"`
}

type SkillGroupCreationInput struct {
	Name string `json:"name"`
}

type SkillGroupUpdateInput struct {
	Name   string `json:"name"`
	Skills []int  `json:"skills"`
}

func NewSkillCreationInput() *SkillCreationInput {
	return &SkillCreationInput{}
}

func NewSkillUpdateInput() *SkillUpdateInput {
	return &SkillUpdateInput{}
}

func NewSkillGroupCreationInput() *SkillGroupCreationInput {
	return &SkillGroupCreationInput{}
}

func NewSkillGroupUpdateInput() *SkillGroupUpdateInput {
	return &SkillGroupUpdateInput{}
}
