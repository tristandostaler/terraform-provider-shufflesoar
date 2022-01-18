terraform {
  required_providers {
    shufflesoar = {
      source = "tristandostaler/shufflesoar"
    }
  }
}

provider "shufflesoar" {
  shuffle_api_token = var.shuffle_api_token
}

