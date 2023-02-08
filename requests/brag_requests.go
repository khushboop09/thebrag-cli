package requests

type AddBragRequest struct {
	Title   string
	Details string
}

type UpdateBragRequest struct {
	ID      int
	Title   string
	Details string
}
