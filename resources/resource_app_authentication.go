package resources

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type ContactInfo struct {
	Name string
	Url  string
}

type ReferenceInfo struct {
	Documentation_url string
	Github_url        string
}

type FolderMount struct {
	Folder_mount       bool
	Source_folder      string
	Destination_folder string
}

type AuthenticationParameterSchema struct {
	AuthenticationParameterSchema_type string `json:"type"`
}

type AuthenticationParameter struct {
	Description string
	Id          string
	Name        string
	Example     string
	Multiline   bool
	Required    bool
	In          string
	Schema      AuthenticationParameterSchema
	Scheme      string
}

type Authentication struct {
	Auth_type     string `json:"type"`
	Required      bool
	Parameters    []AuthenticationParameter
	Redirect_uri  string
	Token_uri     string
	Refresh_uri   string
	Client_id     string
	Client_secret string
}

type AppAuthenticationVersion struct {
	Version string
	Id      string
}

type AppAuthentication struct {
	Name string
	Id   string
}

type Field struct {
	Key   string
	Value string
}

type App struct {
	App    AppAuthentication
	Fields []Field
	Label  string
	Id     string
	Active bool
}

var url string = "https://shuffler.io/api/v1/apps/authentication"

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
			"shuffle_api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Shuffle's API token",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The App name to link this authentication config to. This must match an existing App's name",
			},
			"app_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A unique ID for the app",
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

func createAppObj(d *schema.ResourceData) (App, error) {
	appId := d.Get("app_id").(string)
	if appId == "" {
		appId, _ = randomHex(16)
		d.Set("app_id", appId)
	}
	log.Printf("[INFO] app_id: %s", appId)

	app := App{
		App: AppAuthentication{
			Name: d.Get("name").(string),
			Id:   appId,
		},
		Fields: []Field{},
		Label:  d.Get("label").(string),
		Active: true,
	}

	var fields []map[string]string
	fieldsString := d.Get("fields").(string)
	if err := json.Unmarshal([]byte(fieldsString), &fields); err != nil {
		log.Printf("[WARN] Failed to unmarshal on read: %+v", fieldsString)
		return App{}, err
	}

	for _, field := range fields {
		app.Fields = append(app.Fields, Field{Key: field["key"], Value: field["value"]})
	}

	return app, nil
}

func getShuffleApiToken(d *schema.ResourceData) (string, error) {
	shuffle_api_token := d.Get("shuffle_api_token").(string)
	if shuffle_api_token == "" {
		return "", fmt.Errorf("shuffle_api_token is mandatory")
	}
	return shuffle_api_token, nil
}

func resourceAppAuthenticationCreate(d *schema.ResourceData, m interface{}) error {
	shuffle_api_token, err := getShuffleApiToken(d)
	if err != nil {
		return err
	}

	app, err := createAppObj(d)
	if err != nil {
		return err
	}

	id, err := createOrUpdate(app, shuffle_api_token)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func createOrUpdate(app App, shuffle_api_token string) (string, error) {
	// initialize http client
	client := &http.Client{}

	// marshal User to json
	jsonData, err := json.Marshal(app)
	if err != nil {
		return "", err
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+shuffle_api_token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseJson map[string]interface{}
	if err := json.Unmarshal([]byte(body), &responseJson); err != nil {
		log.Printf("[WARN] Failed to unmarshal on read: %+v", body)
		return "", err
	}

	log.Printf("[INFO] Create or Update Response: %d %s", resp.StatusCode, string(body))

	return responseJson["id"].(string), nil
}

func resourceAppAuthenticationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAppAuthenticationUpdate(d *schema.ResourceData, m interface{}) error {
	shuffle_api_token, err := getShuffleApiToken(d)
	if err != nil {
		return err
	}

	app, err := createAppObj(d)
	if err != nil {
		return err
	}
	app.Id = d.Id()

	_, err = createOrUpdate(app, shuffle_api_token)
	if err != nil {
		return err
	}

	return nil
}

func resourceAppAuthenticationDelete(d *schema.ResourceData, m interface{}) error {
	shuffle_api_token, err := getShuffleApiToken(d)
	if err != nil {
		return err
	}

	id := d.Id()

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", url, id), nil)
	if err != nil {
		return err
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+shuffle_api_token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Delete Response: %d %s", resp.StatusCode, string(body))

	d.SetId("")
	return nil
}
