package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/tristandostaler/shufflesoar/resources"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"shufflesoar_app_authentication": resources.ResourceAppAuthentication(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
