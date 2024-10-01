package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lehau17/food_delivery/components/pubsub"
	"github.com/lehau17/food_delivery/components/pubsub/localps"
)

func main() {
	// Tạo một local pubsub
	ps := localps.NewLocalPubsub()

	// Tạo context để truyền qua các hàm
	ctx := context.Background()

	// Subscribe vào một topic (chủ đề)
	topic := pubsub.Topic("order_created")
	ch, closeFunc := ps.Subcribe(ctx, topic)
	ch2, _ := ps.Subcribe(ctx, topic)

	// Định nghĩa một goroutine lắng nghe sự kiện
	go func() {
		for msg := range ch {
			log.Printf("Received message 1")
			fmt.Println("Received message:", msg.String())
		}
	}()

	go func() {
		for msg := range ch2 {
			log.Printf("Received message 2:::::>>>>>")
			fmt.Println("Received message:", msg.String())
		}
	}()
	// Publish một vài tin nhắn
	for i := 1; i <= 100; i++ {
		message := pubsub.NewMessage(topic, fmt.Sprintf("Order #%d", i))
		ps.Publish(ctx, topic, message)
		time.Sleep(1 * time.Second) // Tạm dừng 1 giây giữa các lần publish
	}

	// Sau 6 giây đóng kết nối subscribe
	time.Sleep(6 * time.Second)
	closeFunc()

	fmt.Println("Closed subscription.")
}
