package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gorilla/mux"
)

func (h *Handler) GetS3Object(w http.ResponseWriter, r *http.Request) {
	if h.S3Client == nil {
		http.Error(w, "S3 client not configured", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "Missing object key", http.StatusBadRequest)
		return
	}

	output, err := h.S3Client.GetObject(r.Context(), &s3.GetObjectInput{
		Bucket: aws.String(h.Config.S3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		var nsk *types.NoSuchKey
		var re *awshttp.ResponseError
		switch {
		case errors.As(err, &nsk):
			http.Error(w, "Object not found", http.StatusNotFound)
		case errors.As(err, &re) && re.HTTPStatusCode() == http.StatusNotFound:
			http.Error(w, "Object not found", http.StatusNotFound)
		default:
			http.Error(w, fmt.Sprintf("Failed to fetch object: %v", err), http.StatusInternalServerError)
		}
		return
	}
	defer output.Body.Close()

	if output.ContentType != nil && *output.ContentType != "" {
		w.Header().Set("Content-Type", aws.ToString(output.ContentType))
	}
	if output.ETag != nil {
		w.Header().Set("ETag", aws.ToString(output.ETag))
	}
	if output.LastModified != nil {
		w.Header().Set("Last-Modified", output.LastModified.Format(http.TimeFormat))
	}
	if output.ContentLength != nil && *output.ContentLength >= 0 {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", aws.ToInt64(output.ContentLength)))
	}

	if _, err := io.Copy(w, output.Body); err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %v", err), http.StatusInternalServerError)
		return
	}
}
