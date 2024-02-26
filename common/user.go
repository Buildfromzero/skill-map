package common

type UserCreationInput struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type UserUpdateInput struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type CompetenceInput struct {
	Skill int `json:"skill" binding:"required"`
	Rank  int `json:"rank" binding:"required"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

func NewUserUpdateInput() *UserUpdateInput {
	return &UserUpdateInput{}
}

func NewCompetenceInput() *CompetenceInput {
	return &CompetenceInput{}
}
