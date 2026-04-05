provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_vpc" "today_go_vpc" {
  cidr_block = var.vpc_cidr
  tags = {
    Name = "${var.env}-today-go-vpc"
  }
}
