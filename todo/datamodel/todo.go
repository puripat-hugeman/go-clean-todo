package datamodel

import (
	"time"

	"github.com/google/uuid"
)

type TodoRequestEntity struct {
	Title string
	Image string
	// Status string
}

//TodoCreateEntity has JSON and GORM tag because It's use for referenced for other struct
type TodoCreateEntity struct {
	Uuid  string `json:"todoUuid" gorm:"primaryKey;column:uuid"`
	Title string `json:"title" gorm:"column:title"`
	// Status    enums.Status `json:"status" gorm:"column:status"`
	Image     string    `json:"image" gorm:"column:image"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at"`
	// UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime;column:created_at"`
}

type TodoGetEntity struct {
	Uuid  string
	Title string
	// Status    enums.Status
	Image     string
	CreatedAt time.Time
	// UpdatedAt time.Time
}

func NewTodo(
	title string,
	// status enums.Status,
	image string,
) (*TodoCreateEntity, error) {
	// if status.IsValid() {
	// 	return nil, fmt.Errorf("invalid status: %s", status)
	// }
	return &TodoCreateEntity{
		Uuid:  uuid.NewString(),
		Title: title,
		// Status: status,
		Image: image,
	}, nil
}
