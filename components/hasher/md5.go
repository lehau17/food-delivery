package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hasher struct{}

func NewMd5Hasher() *md5Hasher {
	return &md5Hasher{}
}

func (h *md5Hasher) Hash(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	return hex.EncodeToString(md5.Sum(nil))
}
