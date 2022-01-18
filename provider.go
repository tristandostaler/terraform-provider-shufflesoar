package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/tristandostaler/terraform-provider-shufflesoar/client"
	"github.com/tristandostaler/terraform-provider-shufflesoar/resources"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"shuffle_api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Shuffle's API token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"shufflesoar_app_authentication": resources.ResourceAppAuthentication(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	shuffle_api_token := d.Get("shuffle_api_token").(string)

	c, _ := client.NewShuffleClient(shuffle_api_token)

	return c, nil
}
