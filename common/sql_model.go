package common

import (
	"time"
)

type SqlModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	Fake_id   string     `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (s *SqlModel) GenUid(dbType int) {
	uid := NewUid(uint32(s.Id), dbType, 1)
	s.Fake_id = uid.String()
}
