package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/btcsuite/btcutil/base58"
)

func main() {
	type Uid struct {
		localId    uint32
		objectType int
		shardId    uint32
	}

	uid := Uid{1, 1, 1}
	val := uint64(uid.localId)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardId)
	encode := base58.Encode([]byte(fmt.Sprintf("%v", val)))
	log.Println("Encoding :>>>>", encode)
	decode := base58.Decode(encode)
	// fmt.Println("Decode :>>>>", string(decode))
	valInt, _ := strconv.Atoi(string(decode))
	// fmt.Println("Decode :>>>>", valInt)

	fmt.Printf("Binary 1: %b", 0x3FF)
	fmt.Printf("Binary 2: %b", 0x3FFFF)
	u := Uid{
		localId:    uint32(valInt >> 28),
		objectType: int(valInt >> 18 & 0x3FF),
		shardId:    uint32(valInt & 0x3FFFF),
	}
	// encode := base58.Encode([]byte(uid))
	fmt.Println("Nguoc lai : >>>>>", u.localId, u.objectType, u.shardId)

}
