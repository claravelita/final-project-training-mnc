package domain

type SocialMedia struct {
	ID             int    `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	Name           string `json:"name" gorm:"NOT NULL"`
	SocialMediaURL string `json:"social_media_url" gorm:"NOT NULL"`
	UserID         int    `json:"user_id"`

	Users User `json:"User" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
