package data_sources

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tristandostaler/terraform-provider-shufflesoar/client"
	"github.com/tristandostaler/terraform-provider-shufflesoar/utils"
)

func DataSourceAllAppAuthentication() *schema.Resource {
	r := &schema.Resource{
		ReadContext: dataSourceAllAppAuthenticationRead,
		Schema: map[string]*schema.Schema{
			"all_app_auths": {
				Type: schema.TypeList,
				Elem: client.GetDefaultAppSchema(),
			},
		},
	}

	r = utils.RecurseSetSchemaStatus(r, utils.Computed, true)

	return r
}

func dataSourceAllAppAuthenticationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*client.ShuffleClient)

	allAppAuth, err := c.GetAllAppAuth()

	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	allAppAuthMap := make([]interface{}, len(allAppAuth))

	for i, app := range allAppAuth {
		appTemp := utils.GenerateMapFromFields(app)
		allAppAuthMap[i] = appTemp
	}
	// appTemp := utils.GenerateMapFromFields(allAppAuth[0])
	// allAppAuthMap[0] = appTemp

	// log.Printf("[DEBUG] Got allAppAuthMap: %+v ", allAppAuthMap)

	if err := d.Set("all_app_auths", allAppAuthMap); err != nil {
		log.Printf("[ERROR] Got error (%+v) setting with allAppAuthMap: %+v ", err, allAppAuthMap)
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
