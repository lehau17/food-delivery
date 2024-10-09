package foodbiz

import (
	"context"
	"fmt"
	"log"
	"time"
)

type FoodUpdateAvgPointStore interface {
	UpdateAvgPoint(ctx context.Context) error
}

type FoodUpdateAvgPointBiz struct {
	store FoodUpdateAvgPointStore
}

func NewFoodUpdateAvgPointBiz(store FoodUpdateAvgPointStore) *FoodUpdateAvgPointBiz {
	return &FoodUpdateAvgPointBiz{store: store}
}

func (b *FoodUpdateAvgPointBiz) FoodUpdateAvgPointBiz(ctx context.Context) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	errchan := make(chan error)

	go func() {
		for {
			fmt.Println("Vao for vo tan")

			<-ticker.C
			// Gọi hàm tính lại điểm đánh giá trung bình
			fmt.Println("Nhan duoc data channel ")

			if err := b.store.UpdateAvgPoint(ctx); err != nil {
				log.Println(err)
				errchan <- err
			}
		}
		// return nil
	}()
	return nil
}
