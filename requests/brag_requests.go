package requests

type AddBragRequest struct {
	Title      string
	Details    string
	CategoryID int
}

type UpdateBragRequest struct {
	ID         int
	Title      string
	Details    string
	CategoryID int
}

type ExportBragsRequest struct {
	From       string
	To         string
	CategoryId int
}
