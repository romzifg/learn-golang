package web

type CategoryCreateRequest struct {
	Name string `validate:"required, max=200, min=3"`
}

type CategoryResponse struct {
	Id   int
	Name string
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required, max=200, min=3"`
}
