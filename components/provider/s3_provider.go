package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lehau17/food_delivery/common"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string // cloudFront
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{bucketName: bucketName, region: region, apiKey: apiKey, secret: secret, domain: domain}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey,
			provider.secret,
			"",
		),
	})
	if err != nil {
		log.Fatalln(err)
	}
	provider.session = sess
	return provider
}

func (provider *s3Provider) SaveFileUpload(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	filebytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        filebytes,
	})
	if err != nil {
		return nil, err
	}
	img := &common.Image{
		Url: fmt.Sprintf("%s/%s", provider.domain, dst),
	}
	return img, nil
}
