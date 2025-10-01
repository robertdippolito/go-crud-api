package handlers

import (
	"k8s-api/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Handler struct {
	Config   *config.AppConfig
	S3Client *s3.Client
}
