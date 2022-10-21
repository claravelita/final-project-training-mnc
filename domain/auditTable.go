package domain

import (
	"time"
)

// AuditTable : embedded for entities with created, updated, at
type AuditTable struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
