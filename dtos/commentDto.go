package dtos

import "time"

type (
	CommentRequest struct {
		Message string `json:"message" validate:"required"`
		PhotoID int    `json:"photo_id" validate:"required"`
	}

	CommentResponse struct {
		ID        int            `json:"id"`
		UserID    int            `json:"user_id"`
		PhotoID   int            `json:"photo_id"`
		Message   string         `json:"message"`
		CreatedAt *time.Time     `json:"created_at,omitempty"`
		UpdatedAt *time.Time     `json:"updated_at,omitempty"`
		Users     *UserResponse  `json:"User,omitempty"`
		Photos    *PhotoResponse `json:"Photo,omitempty"`
	}

	CommentUpdateRequest struct {
		Message string `json:"message" validate:"required"`
	}
)
