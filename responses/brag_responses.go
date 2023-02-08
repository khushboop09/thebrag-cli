package responses

import "time"

type GetAllBragResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Brag `json:"data"`
}

type GetABragResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Brag struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Details    string    `json:"details"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type PostBragResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type DeleteBragResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
