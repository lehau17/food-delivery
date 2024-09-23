package uploadmodel

import "github.com/lehau17/food_delivery/common"

func ErrCannotSaved(err error) *common.AppError {
	return common.NewFullErrorResponse(400, err, "Error occurred while saving image into S3", err.Error(), "ErrNotSavedImage")
}

func ErrCannotGetDimeinsion(err error) *common.AppError {
	return common.NewFullErrorResponse(400, err, "Error occurred while reading config image", err.Error(), "ErorReadConfig")
}
