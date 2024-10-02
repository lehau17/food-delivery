package common

import (
	"context"
	"log"
)

func AppRecover(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered:", r)
		}
	}()
}
