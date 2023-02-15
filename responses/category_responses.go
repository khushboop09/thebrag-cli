package responses

type GetAllCategoriesResponse struct {
	Status  int
	Message string
	Data    []Category
}

type Category struct {
	ID   int
	Name string
}

type AddCategoryResponse struct {
	Status  int
	Message string
	Data    string
}
