---
page_title: "shufflesoar Provider"
subcategory: "provider"
description: |-
  A terraform provider for https://github.com/frikky/Shuffle
---


# Suffle SOAR Provider


A terraform provider for https://github.com/frikky/Shuffle


## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **shuffle_api_token** (String) Shuffle's API token