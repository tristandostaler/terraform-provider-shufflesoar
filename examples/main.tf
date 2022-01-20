module "shufflesoar_resources" {
  source = "./resources"

  shuffle_api_token = var.shuffle_api_token
}

module "shufflesoar_data_sources" {
  source = "./data_sources"

  shuffle_api_token = var.shuffle_api_token
}
