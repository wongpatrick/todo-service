package model

type Status = string

const (
	Active   Status = "A"
	Complete Status = "C"
	Expired  Status = "E"
	Deleted  Status = "D"
)

type (
	Task struct {
		Id          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      Status `json:"status"`
		UserId      uint   `json:"userId"`
		CreatedAt   string `json:"createdAt"`
		ModifiedAt  string `json:"modifiedAt"` 
	}
)
