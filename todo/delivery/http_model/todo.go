package httpmodel

import "time"

type TodoRequestDelivery struct {
	Title string `json:"title"`
	// Status string `json:"status"`
}

type TodoDelivery struct {
	Uuid  string `json:"todoUuid"`
	Title string `json:"title" `
	// Status    enums.Status `json:"status" "`
	Image     string    `json:"image" `
	CreatedAt time.Time `json:"createdAt" `
	UpdatedAt time.Time `json:"updatedAt" `
}
