package main

import (
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub/localps"
	"github.com/lehau17/food_delivery/config"
	"github.com/lehau17/food_delivery/router"
	"github.com/lehau17/food_delivery/skio"
	"github.com/lehau17/food_delivery/subcriber"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to add environment variables")
	}
	//
	secretkey := os.Getenv("SYSTEM_SECRET")
	db := config.NewSqlInstance()
	uploadProvider := config.NewS3Instance()
	ps := localps.NewLocalPubsub()
	rdb := config.NewRedisInstance()

	ctx := appcontext.NewAppContext(db, &uploadProvider, secretkey, ps, rdb)
	consumberJob := subcriber.NewConsumerEngine(ctx)
	consumberJob.Start()

	r := router.InitRouter("debug", ctx)
	r.StaticFS("/public", http.Dir("./asset"))
	skio.NewRtEngine().Run(ctx, r)
	r.Run() // listen and serve on 0.0.0.0:8080

}
