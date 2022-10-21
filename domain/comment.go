package domain

type Comment struct {
	ID      int    `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	UserID  int    `json:"user_id"`
	PhotoID int    `json:"photo_id"`
	Message string `json:"message" gorm:"NOT NULL"`

	AuditTable AuditTable `json:"-" gorm:"embedded"`
	Users      User       `json:"User" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photos     Photo      `json:"Photo" gorm:"foreignKey:PhotoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
