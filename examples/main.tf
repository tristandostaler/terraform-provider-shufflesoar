module "shufflesoar_resources" {
  source = "./resources"

  shuffle_base_url  = var.shuffle_base_url
  shuffle_api_token = var.shuffle_api_token
}

module "shufflesoar_data_sources" {
  source = "./data_sources"

  shuffle_base_url  = var.shuffle_base_url
  shuffle_api_token = var.shuffle_api_token
}
