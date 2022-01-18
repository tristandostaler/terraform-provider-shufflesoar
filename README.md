# ShuffleSOAR
A terraform provider for https://github.com/frikky/Shuffle

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/tristandostaler)

## Using with terraform
- First, add the required provider:
```
# provider.tf
terraform {
  required_providers {
    shufflesoar = {
      source = "tristandostaler/shufflesoar"
    }
  }
}
```
- Get the token from environment variables:
```
# variables.tf
# to setup the shuffle_api_token: export TF_VAR_shuffle_api_token=YOURTOKEN
variable "shuffle_api_token" {
  type = string
}
```
- Then config the provider:
```
# main.tf
provider "shufflesoar" {
  shuffle_api_token = var.shuffle_api_token
}
```
- Now you can use it:
```
# main.tf
resource "shufflesoar_app_authentication" "example" {
  ...
}
```

## Resources
- https://www.hashicorp.com/blog/writing-custom-terraform-providers
- https://www.infracloud.io/blogs/developing-terraform-custom-provider/
- https://www.terraform.io/plugin
- https://learn.hashicorp.com/tutorials/terraform/provider-update?in=terraform/providers

## Setup
- Follow the instructions here: https://www.infracloud.io/blogs/developing-terraform-custom-provider/
- More info in https://www.terraform.io/cli/config/config-file#provider-installation
- Set the content of .terraformrc
- Then:
```
export TF_VAR_shuffle_api_token=YOURTOKEN
export TF_LOG=TRACE
sudo mkdir -p /usr/share/terraform/providers
sudo chmod -R 777 /usr/share/terraform/providers
make
make test_clean
```
