provider "shufflesoar" {

}

resource "shufflesoar_app_authentication" "example" {
  shuffle_api_token = var.shuffle_api_token
  name              = "AWS ses"
  label             = "A test app"
  fields            = <<EOF
[{
    "key": "access_key",
    "value": "1234"
}, {
    "key": "secret_key",
    "value": "1234"
}, {
    "key": "region",
    "value": "1234"
}]
EOF

}
