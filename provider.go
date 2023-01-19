package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/tristandostaler/terraform-provider-shufflesoar/client"
	"github.com/tristandostaler/terraform-provider-shufflesoar/data_sources"
	"github.com/tristandostaler/terraform-provider-shufflesoar/resources"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown

	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		if s.Deprecated != "" {
			desc += " " + s.Deprecated
		}
		return strings.TrimSpace(desc)
	}
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"shuffle_base_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Shuffle's base URL (i.e https://shuffler.io or https://ca.shuffler.io)",
			},
			"shuffle_api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Shuffle's API token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"shufflesoar_app_authentication": resources.ResourceAppAuthentication(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"shufflesoar_all_app_authentications": data_sources.DataSourceAllAppAuthentication(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	shuffle_base_url := d.Get("shuffle_base_url").(string)
	shuffle_api_token := d.Get("shuffle_api_token").(string)

	c, _ := client.NewShuffleClient(shuffle_base_url, shuffle_api_token)

	return c, nil
}
