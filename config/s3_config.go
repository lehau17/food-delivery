package config

import (
	"os"

	uploadprovider "github.com/lehau17/food_delivery/components/provider"
)

func NewS3Instance() uploadprovider.UploadProvider {
	bucket := os.Getenv("S3_BUCKET_NAME")
	region := os.Getenv("S3_REGION")
	apiKey := os.Getenv("S3_API_KEY")
	secret := os.Getenv("S3_SECRET")
	domain := os.Getenv("S3_DOMAIN")
	uploadProvider := uploadprovider.NewS3Provider(bucket, region, apiKey, secret, domain)
	return uploadProvider
}
