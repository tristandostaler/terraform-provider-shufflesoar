package resources

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tristandostaler/terraform-provider-shufflesoar/client"
	"github.com/tristandostaler/terraform-provider-shufflesoar/utils"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func ResourceAppAuthentication() *schema.Resource {
	r := &schema.Resource{
		Create: resourceAppAuthenticationCreate,
		Read:   resourceAppAuthenticationRead,
		Update: resourceAppAuthenticationUpdate,
		Delete: resourceAppAuthenticationDelete,

		Schema: client.GetDefaultAppSchema().Schema,
	}

	r = utils.RecurseSetSchemaStatus(r, utils.Optional, true)

	r = utils.RecurseSetSchemaStatusByKey(r, "id", utils.Computed, true)
	r = utils.RecurseSetSchemaStatusByKey(r, "label", utils.Required, true)
	r = utils.RecurseSetSchemaStatusByKey(r, "app.name", utils.Required, true)
	r = utils.RecurseSetSchemaStatusByKey(r, "fields.key", utils.Required, true)
	r = utils.RecurseSetSchemaStatusByKey(r, "fields.value", utils.Required, true)

	r.Schema["app"].MinItems = 1
	r.Schema["app"].MaxItems = 1

	return r
}

func getAppAuthFromResourceData(d *schema.ResourceData) map[string]interface{} {
	return d.Get("app").([]interface{})[0].(map[string]interface{})
}

func createAppObj(d *schema.ResourceData) (client.App, error) {
	appAuth := getAppAuthFromResourceData(d)
	appAuthId := appAuth["id"].(string)
	if appAuthId == "" {
		appAuthId, _ = randomHex(16)
		appAuth["id"] = appAuthId
		temp := make([]interface{}, 1)
		temp[0] = appAuth
		d.Set("app", temp)
	}
	log.Printf("[INFO] appAuthId: %s", appAuthId)

	app := client.App{
		App: client.AppAuthentication{
			Name:       appAuth["name"].(string),
			Id:         appAuthId,
			LargeImage: appAuth["large_image"].(string),
		},
		Fields: []client.Field{},
		Label:  d.Get("label").(string),
		Active: true,
	}

	fieldsResources := d.Get("fields").([]interface{})
	for _, fr := range fieldsResources {
		app.Fields = append(app.Fields, client.Field{
			Key:   fr.(map[string]interface{})["key"].(string),
			Value: fr.(map[string]interface{})["value"].(string),
		})
	}

	return app, nil
}

func resourceAppAuthenticationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.ShuffleClient)

	app, err := createAppObj(d)
	if err != nil {
		return err
	}

	id, err := c.CreateOrUpdateAppAuth(app)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceAppAuthenticationRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	c := m.(*client.ShuffleClient)

	app, err := c.GetAppAuthById(id)
	if err != nil {
		log.Printf("[WARN] App (%s) not found, removing from state", id)
		d.SetId("")
		return nil
	}

	appAuth := make([]map[string]interface{}, 1)
	appAuth[0] = make(map[string]interface{})
	appAuth[0]["id"] = app.App.Id
	appAuth[0]["name"] = app.App.Name
	appAuth[0]["large_image"] = app.App.LargeImage

	fields := make([]map[string]interface{}, len(app.Fields))
	for _, f := range app.Fields {
		f1 := make(map[string]interface{})
		f1["key"] = f.Key
		f1["value"] = f.Value
		fields = append(fields, f1)
	}

	d.Set("app", appAuth)
	d.Set("name", app.App.Name)
	d.Set("label", app.Label)
	d.Set("fields", fields)

	return nil
}

func resourceAppAuthenticationUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.ShuffleClient)

	app, err := createAppObj(d)
	if err != nil {
		return err
	}
	app.Id = d.Id()

	_, err = c.CreateOrUpdateAppAuth(app)
	if err != nil {
		return err
	}

	return nil
}

func resourceAppAuthenticationDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.ShuffleClient)

	id := d.Id()

	c.DeleteAppAuth(id)

	d.SetId("")
	return nil
}
