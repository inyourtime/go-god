package s3

import (
	"context"
	"fmt"
	"gopher/src/coreplugins"
	"gopher/src/logs"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/matoous/go-nanoid/v2"
)

type s3Handler struct {
	client *s3.Client
}

func NewS3Handler() s3Handler {
	return s3Handler{
		client: initS3(),
	}
}

func initS3() *s3.Client {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", coreplugins.Config.R2.AccountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(coreplugins.Config.R2.AccessKeyID, coreplugins.Config.R2.AccessKeySecret, "")),
	)
	if err != nil {
		logs.Error(err)
	}

	return s3.NewFromConfig(cfg)
}

func (h s3Handler) Download(path string) string {
	presignClient := s3.NewPresignClient(h.client)
	presignResult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(coreplugins.Config.R2.Bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		panic("Couldn't get presigned URL for PutObject")
	}
	return presignResult.URL
}

func (h s3Handler) ListObjects() []types.Object {
	listObjectsOutput, err := h.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &coreplugins.Config.R2.Bucket,
	})
	if err != nil {
		logs.Error(err)
	}
	// for _, object := range listObjectOutput {
	// 	obj, _ := json.MarshalIndent(object, "", "\t")
	// 	fmt.Println(string(obj))
	// }
	return listObjectsOutput.Contents
}

func (h s3Handler) UploadFile(file *multipart.FileHeader) {
	ct := file.Header["Content-Type"][0]
	f, err := file.Open()
	if err != nil {
		logs.Error(err)
	}
	p := generatePath(getFileType(ct))
	_, err = h.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(coreplugins.Config.R2.Bucket),
		Key:         aws.String(p),
		Body:        f,
		ContentType: aws.String(ct),
	})
	if err != nil {
		logs.Error(err)
	}
}

func getFileType(mimeType string) string {
	parts := strings.Split(mimeType, "/")
	return parts[len(parts) - 1]
}

func generatePath(fileType string) string {
	rd, err := gonanoid.New(10)
	if err != nil {
		logs.Error(err)
	}
	path := rd + "." + fileType
	return path
}
