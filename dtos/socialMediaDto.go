package dtos

import "time"

type (
	SocialMediaRequest struct {
		Name           string `json:"name" validate:"required"`
		SocialMediaURL string `json:"social_media_url" validate:"required"`
	}

	SocialMediaResponse struct {
		ID             int           `json:"id"`
		Name           string        `json:"name"`
		SocialMediaURL string        `json:"social_media_url"`
		UserID         int           `json:"user_id"`
		CreatedAt      *time.Time    `json:"created_at,omitempty"`
		UpdatedAt      *time.Time    `json:"updated_at,omitempty"`
		Users          *UserResponse `json:"User,omitempty"`
	}
)
