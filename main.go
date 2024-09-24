package main

import (
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	uploadprovider "github.com/lehau17/food_delivery/components/provider"
	"github.com/lehau17/food_delivery/middlewares"
	"github.com/lehau17/food_delivery/modules/restaurent/tranport/ginrestaurant"
	"github.com/lehau17/food_delivery/modules/upload/tranport/ginupload"
	usertransport "github.com/lehau17/food_delivery/modules/user/transport"
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
	var uploadProvider uploadprovider.UploadProvider
	uploadProvider = uploadprovider.NewS3Provider(bucket, region, apiKey, secret, domain)
	ctx := appcontext.NewAppContext(db, &uploadProvider)

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
	gRes := r.Group("/restaurants")
	r.Use(middlewares.Recovery(ctx))
	gRes.GET("/", ginrestaurant.GetListRestaurant(ctx))
	gRes.POST("/", ginrestaurant.CreateRestaurant(ctx))
	gRes.DELETE("/:id", ginrestaurant.DeleteRestaurant(ctx))
	gRes.POST("/upload", ginupload.UploadImage(ctx))

	gUser := r.Group("/user")
	gUser.POST("/register", usertransport.RegisterUser(ctx))

	// manager version
	// v1 := r.Group("/v1")
	// v1.POST("/product", func(c *gin.Context) {
	// 	fmt.Print(c.Request.Body)

	// 	c.JSON(201, gin.H{
	// 		"message": "Hello world",
	// 		"payload": body,
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080

}
