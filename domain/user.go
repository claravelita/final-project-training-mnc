package domain

type User struct {
	ID       int    `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	Username string `json:"username" gorm:"NOT NULL"`
	Email    string `json:"email" gorm:"NOT NULL"`
	Password string `json:"password" gorm:"NOT NULL"`
	Age      int    `json:"age"`

	AuditTable AuditTable `json:"-" gorm:"embedded"`
}
