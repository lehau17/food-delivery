package orderstore

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type sqlStore struct {
	db  *gorm.DB
	rbd *redis.Client
}

func NewSqlStore(db *gorm.DB, rbd *redis.Client) *sqlStore {
	return &sqlStore{db: db, rbd: rbd}
}
