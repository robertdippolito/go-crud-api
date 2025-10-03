data "aws_iam_policy_document" "trust" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRoleWithWebIdentity"]
    principals {
      type        = "Federated"
      identifiers = [var.provider_arn]
    }
    condition {
      test     = "StringEquals"
      variable = "${local.issuer_hostpath}:sub"
      values   = [local.sa_sub]
    }
    condition {
      test     = "StringEquals"
      variable = "${local.issuer_hostpath}:aud"
      values   = ["sts.amazonaws.com"]
    }
  }
}

locals {
  issuer_hostpath = replace(var.issuer_url, "https://", "")
  sa_sub          = "system:serviceaccount:${var.namespace}:${var.service_account}"
}

resource "aws_iam_role" "irsa_role" {
  name               = "EKS-Cross-Account-S3-${var.service_account}"
  assume_role_policy = data.aws_iam_policy_document.trust.json
  tags               = var.tags
}

data "aws_iam_policy_document" "s3" {
  statement {
    sid       = "ListBucketPrefix"
    actions   = ["s3:ListBucket"]
    resources = [var.bucket_arn]
  }
  statement {
    sid       = "RWObjectsInPrefix"
    actions   = ["s3:GetObject"]
    resources = ["${var.bucket_arn}/*"]
  }
}

resource "aws_iam_policy" "s3" {
  name   = "EKS-${var.namespace}-${var.service_account}-S3"
  policy = data.aws_iam_policy_document.s3.json
}

resource "aws_iam_role_policy_attachment" "attach" {
  role       = aws_iam_role.irsa_role.name
  policy_arn = aws_iam_policy.s3.arn
}