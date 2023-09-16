package model

type Status = string

const (
	Active   Status = "A"
	Complete Status = "C"
	// Expired  Status = "E" // Could be used for setting deadlines on tasks
	Deleted Status = "D"
)

type (
	Task struct {
		Id          uint    `json:"id"`
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *Status `json:"status"`
		UserId      uint    `json:"userId"`
		CreatedAt   string  `json:"createdAt"`
		ModifiedAt  string  `json:"modifiedAt"`
	}

	TaskParams struct {
		Id     *uint `json:"id"`
		UserId *uint `json:"userId"`
	}
)
