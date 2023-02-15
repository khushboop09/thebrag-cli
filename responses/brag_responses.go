package responses

type GetAllBragResponse struct {
	Status  int
	Message string
	Data    []Brag
}

type GetABragResponse struct {
	Status  int
	Message string
	Data    interface{}
}

type Brag struct {
	ID           int
	Title        string
	Details      string
	CategoryID   int
	CategoryName string
}

type PostBragResponse struct {
	Status  int
	Message string
	Data    string
}

type PutBragResponse struct {
	Status  int
	Message string
	Data    string
}

type DeleteBragResponse struct {
	Status  int
	Message string
	Data    string
}
