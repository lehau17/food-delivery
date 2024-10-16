package citystore

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type sqlStore struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewSqlStore(db *gorm.DB, rdb *redis.Client) *sqlStore {
	return &sqlStore{db: db, rdb: rdb}
}
