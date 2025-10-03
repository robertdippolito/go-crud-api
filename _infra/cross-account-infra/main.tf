provider "aws" {
    region = "us-east-1"
    profile = "default"
    assume_role {
        role_arn = "arn:aws:iam::802594765618:role/OrganizationAccountAccessRole"
        session_name = "tf-deploy"
    }
}

module "s3" {
    source = "./modules/s3"

    bucket_name = "my-cross-account-bucket-10032025"
    tags = {
        project = "cross-account-demo"
        env     = "dev"
    }
}

module "iam" {
    source = "./modules/iam"

    bucket_arn = module.s3.bucket_arn
    namespace = "api"
    service_account = "api-sa"
    issuer_url       = "oidc.eks.us-east-1.amazonaws.com/id/62A59FE098E07B5F883C4B140909EC0F"
    provider_arn     = "arn:aws:iam::802594765618:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/62A59FE098E07B5F883C4B140909EC0F"
    tags = {
        project = "cross-account-demo"
        env     = "dev"
    }
}