package web

type CategoryCreateRequest struct {
	Name string
}

type CategoryResponse struct {
	Id   int
	Name string
}

type CategoryUpdateRequest struct {
	Id   int
	Name string
}
