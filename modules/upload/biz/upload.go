package uploadbiz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"io"
	"path/filepath"
	"time"

	"github.com/lehau17/food_delivery/common"
	uploadprovider "github.com/lehau17/food_delivery/components/provider"
	uploadmodel "github.com/lehau17/food_delivery/modules/upload/model"
)

type CreateUploadStore interface {
	CreateImage(ctx context.Context, data *common.Image) error
}
type uploadBiz struct {
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	if len(data) == 0 {
		return nil, errors.New("no data provided")
	}

	// Tạo một bytes.Reader mới từ dữ liệu
	imgReader := bytes.NewBuffer(data)
	// imgReader.Seek(0, io.SeekStart)
	// Gọi getImageDemension với imgReader
	w, h, err := getImageDemension(imgReader)
	if err != nil {
		return nil, uploadmodel.ErrCannotGetDimeinsion(err)
	}

	filExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), filExt)

	img, err := biz.provider.SaveFileUpload(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCannotSaved(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = filExt
	return img, nil
}

func getImageDemension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
