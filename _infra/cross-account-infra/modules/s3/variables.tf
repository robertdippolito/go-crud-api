variable "bucket_name" {
  description = "Name of the S3 bucket to create"
  type        = string
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}