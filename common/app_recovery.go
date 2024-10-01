package common

import (
	"context"
	"fmt"
)

func AppRecover(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
}
