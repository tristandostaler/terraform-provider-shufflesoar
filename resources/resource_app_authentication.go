package resources

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tristandostaler/terraform-provider-shufflesoar/client"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func ResourceAppAuthentication() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppAuthenticationCreate,
		Read:   resourceAppAuthenticationRead,
		Update: resourceAppAuthenticationUpdate,
		Delete: resourceAppAuthenticationDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The App name to link this authentication config to. This must match an existing App's name",
			},
			"app_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The App Id of the App to link this authentication config to",
			},
			"large_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The base64 string for the image to display. Format: data:image/png;base64,THE_BASE64",
			},
			"fields": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is a json map of all the required fields for this app authentication. The name of the fields must match the names in the authentication parameters and there must be the same number of parameters and fields.",
			},
			"label": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The text to display in the Shuffle UI",
			},
		},
	}
}

func createAppObj(d *schema.ResourceData) (client.App, error) {
	appId := d.Get("app_id").(string)
	if appId == "" {
		appId, _ = randomHex(16)
		d.Set("app_id", appId)
	}
	log.Printf("[INFO] app_id: %s", appId)

	app := client.App{
		App: client.AppAuthentication{
			Name:       d.Get("name").(string),
			Id:         appId,
			LargeImage: d.Get("large_image").(string),
		},
		Fields: []client.Field{},
		Label:  d.Get("label").(string),
		Active: true,
	}

	var fields []map[string]string
	fieldsString := d.Get("fields").(string)
	if err := json.Unmarshal([]byte(fieldsString), &fields); err != nil {
		log.Printf("[WARN] Failed to unmarshal on read: %+v", fieldsString)
		return client.App{}, err
	}

	for _, field := range fields {
		app.Fields = append(app.Fields, client.Field{Key: field["key"], Value: field["value"]})
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

	app, err := c.GetAppAuth(id)
	if err != nil {
		log.Printf("[WARN] App (%s) not found, removing from state", id)
		d.SetId("")
		return nil
	}

	jsonData, err := json.Marshal(app.Fields)
	if err != nil {
		return err
	}

	d.Set("app_id", app.App.Id)
	d.Set("name", app.App.Name)
	d.Set("label", app.Label)
	d.Set("fields", string(jsonData))
	d.Set("large_image", app.App.LargeImage)

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
