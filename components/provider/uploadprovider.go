package uploadprovider

import (
	"context"

	"github.com/lehau17/food_delivery/common"
)

type UploadProvider interface {
	SaveFileUpload(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
