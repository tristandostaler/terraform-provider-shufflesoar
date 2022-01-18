# ShuffleSOAR
A terraform provider for https://github.com/frikky/Shuffle

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
sudo mkdir -p /usr/share/terraform/providers
sudo chmod -R 777 /usr/share/terraform/providers
make
make test_clean
```