package model

type Status = string

const (
	Active   Status = "A"
	Complete Status = "C"
	// Expired  Status = "E" // Could be used for setting deadlines on tasks
	Deleted Status = "D"
)

func IsValidStatus(str string) bool {
	if str == Active || str == Complete || str == Deleted {
		return true
	}
	return false
}

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

	CreateTaskParams struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		UserId      uint   `json:"userId"`
	}

	PatchTaskParams struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *Status `json:"status"`
	}
)
