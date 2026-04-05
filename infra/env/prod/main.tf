module "vpc" {
  source = "../../modules/vpc"

  env      = var.env
  vpc_cidr = var.cidr
}
