variable "bucket_arn" {
    description = "Bucket ARN of cross account bucket"
    type = string
}

variable "namespace" {
  description = "Kubernetes namespace in Account A"
  type        = string
}

variable "service_account" {
  description = "Kubernetes service account in Account A"
  type        = string
}

variable "issuer_url" {
  description = "OIDC issuer URL from Account A's EKS cluster"
  type        = string
}

variable "provider_arn" {
  description = "OIDC provider ARN from Account A"
  type        = string
}

variable "prefix" {
  description = "S3 key prefix pods are allowed to access"
  type        = string
  default     = "app/"
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}