package dto

type EntityField struct {
	Field     string `json:"field" binding:"required"`
	Value     string `json:"value" binding:"required"`
	IsVisible bool   `json:"isVisible" binding:"required"`
}

type AddEntityDto struct {
	Fields []EntityField `json:"fields" binding:"required"`
	Name   string        `json:"name" binding:"required"`
}

type DeleteEntityDto struct {
	id string
}

type UpdateEntityDto struct {
	AddEntityDto
	ID string `json:"id" binding:"required"`
}
