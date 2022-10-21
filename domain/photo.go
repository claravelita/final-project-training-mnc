package domain

type Photo struct {
	ID       int    `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	Title    string `json:"title" gorm:"NOT NULL"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" gorm:"NOT NULL"`
	UserID   int    `json:"user_id"`

	AuditTable AuditTable `json:"-" gorm:"embedded"`
	Users      User       `json:"User" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
