package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ShuffleClient struct {
	Url      string
	APIToken string
}

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
	Name       string
	Id         string
	LargeImage string `json:"large_image"`
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

func NewShuffleClient(apiToken string) (*ShuffleClient, error) {
	return &ShuffleClient{
		Url:      "https://shuffler.io/api/v1/apps/authentication",
		APIToken: apiToken,
	}, nil
}

func (c *ShuffleClient) CreateOrUpdateAppAuth(app App) (string, error) {
	// initialize http client
	client := &http.Client{}

	// marshal User to json
	jsonData, err := json.Marshal(app)
	if err != nil {
		return "", err
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, c.Url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.APIToken)

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

func (c *ShuffleClient) DeleteAppAuth(id string) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", c.Url, id), nil)
	if err != nil {
		return err
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.APIToken)

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

	return nil
}

func (c *ShuffleClient) GetAppAuth(id string) (App, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, c.Url, nil)
	if err != nil {
		return App{}, err
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.APIToken)

	resp, err := client.Do(req)
	if err != nil {
		return App{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return App{}, err
	}

	var responseJson map[string]interface{}
	if err := json.Unmarshal([]byte(body), &responseJson); err != nil {
		log.Printf("[WARN] Failed to unmarshal on read: %+v", body)
		return App{}, err
	}

	apps, ok := responseJson["data"].([]interface{})
	if !ok {
		log.Printf("[ERROR] expected %T to be an array of object", responseJson["data"])
		return App{}, fmt.Errorf("expected %T to be an array of object", responseJson["data"])
	}

	for _, app := range apps {
		if app.(map[string]interface{})["id"].(string) != id {
			log.Printf("[WARN] App %s not matching %s", app.(map[string]interface{})["id"].(string), id)
			continue
		}
		appObject := App{
			App: AppAuthentication{
				Name:       app.(map[string]interface{})["app"].(map[string]interface{})["name"].(string),
				Id:         app.(map[string]interface{})["app"].(map[string]interface{})["id"].(string),
				LargeImage: app.(map[string]interface{})["app"].(map[string]interface{})["large_image"].(string),
			},
			Fields: []Field{},
			Label:  app.(map[string]interface{})["label"].(string),
			Id:     id,
			Active: app.(map[string]interface{})["active"].(bool),
		}

		fields := app.(map[string]interface{})["fields"].([]interface{})
		for _, field := range fields {
			appObject.Fields = append(appObject.Fields, Field{Key: field.(map[string]interface{})["key"].(string), Value: field.(map[string]interface{})["value"].(string)})
		}

		return appObject, nil
	}

	return App{}, fmt.Errorf("App (%s) not found", id)
}
