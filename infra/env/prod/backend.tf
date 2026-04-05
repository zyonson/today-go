provider "aws" {
  region = "ap-northeast-1"
}

terraform {
  backend "s3" {}
}
