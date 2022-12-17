package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,max=200,min=3"`
}

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=200,min=3"`
}
