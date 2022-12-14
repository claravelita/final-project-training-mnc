package dtos

import "time"

type (
	UserRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=6"`
		Age      int    `json:"age" validate:"required"`
	}

	UserUpdateRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=6"`
	}

	UserResponse struct {
		ID              int        `json:"id,omitempty"`
		Username        string     `json:"username"`
		Email           string     `json:"email"`
		Age             int        `json:"age,omitempty"`
		UpdateAt        *time.Time `json:"update_at,omitempty"`
		ProfileImageURL *string    `json:"profile_image_url,omitempty"`
	}
)
