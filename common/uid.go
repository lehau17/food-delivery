package common

import (
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

// Uid structure
type Uid struct {
	localId    uint32
	objectType int
	shardId    uint32
}

// Implement the String method for Uid
func (uid Uid) String() string {
	val := uint64(uid.localId)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardId)
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

// Constructor for Uid
func NewUid(localId uint32, objectType int, shardId uint32) Uid {
	return Uid{localId: localId, objectType: objectType, shardId: shardId}
}

// Getter for localId
func (uid Uid) GetLocalId() uint32 {
	return uid.localId
}

// DecomposeUid to reconstruct Uid from base58 string
func DecomposeUid(s string) (Uid, error) {
	uidBytes := base58.Decode(s)
	if len(uidBytes) == 0 {
		return Uid{}, errors.New("invalid uid")
	}

	val := binary.BigEndian.Uint64(uidBytes)
	if 1<<18 > val {
		return Uid{}, errors.New("invalid uid")
	}
	u := Uid{
		localId:    uint32(val >> 28),
		objectType: int(val >> 18 & 0x3FF),
		shardId:    uint32(val & 0x3FFFF),
	}
	return u, nil
}

// FromBase58 decodes a base58 encoded string into Uid
func FromBase58(s string) (Uid, error) {
	return DecomposeUid(s)
}

// Scan implements the sql.Scanner interface
func (uid *Uid) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	s, ok := value.(string)
	if !ok {
		return errors.New("invalid uid type")
	}

	decodedUid, err := FromBase58(s)
	if err != nil {
		return err
	}

	*uid = decodedUid
	return nil
}

// Value implements the driver.Valuer interface
func (uid Uid) Value() (driver.Value, error) {
	return uid.String(), nil
}
