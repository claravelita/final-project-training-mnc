package dtos

import "time"

type (
	PhotoRequest struct {
		Title    string `json:"title" validate:"required"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url" validate:"required"`
	}

	PhotoResponse struct {
		ID        int           `json:"id"`
		Title     string        `json:"title"`
		Caption   string        `json:"caption"`
		PhotoURL  string        `json:"photo_url"`
		UserID    int           `json:"user_id"`
		CreatedAt *time.Time    `json:"created_at,omitempty"`
		UpdatedAt *time.Time    `json:"updated_at,omitempty"`
		Users     *UserResponse `json:"User,omitempty"`
	}
)
