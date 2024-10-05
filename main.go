package main

import (
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	uploadprovider "github.com/lehau17/food_delivery/components/provider"
	"github.com/lehau17/food_delivery/components/pubsub/localps"
	"github.com/lehau17/food_delivery/middlewares"
	"github.com/lehau17/food_delivery/modules/category/transport/gincategory"
	"github.com/lehau17/food_delivery/modules/food/transport/ginfood"
	"github.com/lehau17/food_delivery/modules/likeuser/transport/ginlikerestaurant"
	"github.com/lehau17/food_delivery/modules/restaurent/tranport/ginrestaurant"
	"github.com/lehau17/food_delivery/modules/upload/tranport/ginupload"
	usertransport "github.com/lehau17/food_delivery/modules/user/transport/gintransport"
	"github.com/lehau17/food_delivery/subcriber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// nạp biến môi trường
	err := godotenv.Load()
	if err != nil {
		panic("failed to add environment variables")
	}
	//
	dsn := os.Getenv("MYSQL_CONNECT_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	if err != nil {
		panic("failed to connect database")
	}
	bucket := os.Getenv("S3_BUCKET_NAME")
	region := os.Getenv("S3_REGION")
	apiKey := os.Getenv("S3_API_KEY")
	secret := os.Getenv("S3_SECRET")
	domain := os.Getenv("S3_DOMAIN")
	secretkey := os.Getenv("SYSTEM_SECRET")
	var uploadProvider uploadprovider.UploadProvider
	uploadProvider = uploadprovider.NewS3Provider(bucket, region, apiKey, secret, domain)
	ps := localps.NewLocalPubsub()
	ctx := appcontext.NewAppContext(db, &uploadProvider, secretkey, ps)
	consumberJob := subcriber.NewConsumerEngine(ctx)
	consumberJob.Start()
	// Test UID

	// type Uid struct {
	// 	localId    uint32
	// 	objectType int
	// 	shardId    uint32
	// }

	// func(uid Uid) String(){
	// 	return "hehe"
	// }

	// Implement phương thức String cho kiểu Uid
	// func (uid Uid) String() string {
	// 	// Tạo giá trị từ các trường của Uid
	// 	val := uint64(uid.localId)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardId)<<0
	// 	// Chuyển đổi giá trị thành chuỗi base58
	// 	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
	// }
	// func NewUidhewhw(localId uint32, objectType int, shardId uint32) Uid {
	// 	return Uid{localId: localId, objectType: objectType, shardId: shardId}
	// }

	//api using gorn
	gin.SetMode("debug")
	r := gin.Default()
	r.Use(middlewares.Recovery(ctx))
	{
		gRes := r.Group("/restaurants", middlewares.CheckAuth(ctx))
		gRes.GET("/", ginrestaurant.GetListRestaurant(ctx))
		gRes.POST("/", ginrestaurant.CreateRestaurant(ctx))
		gRes.DELETE("/:id", ginrestaurant.DeleteRestaurant(ctx))
		gRes.POST("/upload", ginupload.UploadImage(ctx))
		gRes.POST("/:id/like", ginlikerestaurant.LikeRestaurant(ctx))
		gRes.POST("/:id/unlike", ginlikerestaurant.UnlikeRestaurant(ctx))
	}
	{
		gUser := r.Group("/user")
		gUser.POST("/register", usertransport.RegisterUser(ctx))
		gUser.POST("/login", usertransport.Login(ctx))
		gUser.GET("/profile", middlewares.CheckAuth(ctx), usertransport.Profile(ctx))
	}
	{
		gCate := r.Group("/categories", middlewares.CheckAuth(ctx))
		gCate.POST("/", gincategory.CreateCategory(ctx))
		gCate.GET("/", gincategory.GetListCategory(ctx))
		gCate.PATCH("/:id", gincategory.UpdateCategory(ctx))
		gCate.DELETE("/:id", gincategory.DeleteCategory(ctx))
	}
	{
		gFood := r.Group("/foods", middlewares.CheckAuth(ctx))
		gFood.POST("/", ginfood.CreateFood(ctx))
		// gFood.GET("/", gincategory.GetListCategory(ctx))
		// gFood.PATCH("/:id", gincategory.UpdateCategory(ctx))
		gFood.DELETE("/:id", ginfood.DeleteFood(ctx))
	}

	r.Run() // listen and serve on 0.0.0.0:8080

}
