package ginupload

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	uploadbiz "github.com/lehau17/food_delivery/modules/upload/biz"
)

func UploadImage(appCtx appcontext.AppContect) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		fmt.Printf(fileHeader.Filename)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		defer file.Close()
		bataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(bataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		biz := uploadbiz.NewUploadBiz(*appCtx.UploadProvider())
		img, err := biz.Upload(c.Request.Context(), bataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}
		c.JSON(200, img)
	}
}
