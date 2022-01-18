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

type GetAppResponse struct {
	Data    []App
	Success bool
}

type CreateOrUpdateResponse struct {
	Success bool
	Id      string
}

func NewShuffleClient(apiToken string) (*ShuffleClient, error) {
	return &ShuffleClient{
		Url:      "https://shuffler.io/api/v1/apps/authentication",
		APIToken: apiToken,
	}, nil
}

func (c *ShuffleClient) CreateOrUpdateAppAuth(app App) (string, error) {
	// marshal User to json
	jsonData, err := json.Marshal(app)
	if err != nil {
		return "", err
	}
	body, statusCode, err := c.makeRequest(http.MethodDelete, c.Url, jsonData)
	if err != nil {
		return "", err
	}

	var responseJson CreateOrUpdateResponse
	if err := json.Unmarshal([]byte(body), &responseJson); err != nil {
		log.Printf("[WARN] Failed to unmarshal on read: %+v", body)
		return "", err
	}

	if !responseJson.Success {
		log.Printf("[WARN] Failed to add app auth: %+v", body)
		return "", fmt.Errorf("[WARN] Failed to add app auth: %+v", body)
	}

	log.Printf("[INFO] Create or Update Response: %d %s", statusCode, string(body))

	return responseJson.Id, nil
}

func (c *ShuffleClient) DeleteAppAuth(id string) error {
	body, statusCode, err := c.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s", c.Url, id), nil)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Delete Response: %d %s", statusCode, string(body))

	return nil
}

func (c *ShuffleClient) GetAppAuth(id string) (App, error) {
	body, _, err := c.makeRequest(http.MethodGet, c.Url, nil)
	if err != nil {
		return App{}, err
	}

	var responseJson GetAppResponse
	if err := json.Unmarshal([]byte(body), &responseJson); err != nil {
		log.Printf("[WARN] Failed to unmarshal on read: %+v", body)
		return App{}, err
	}

	apps := responseJson.Data

	for _, app := range apps {
		if app.Id != id {
			log.Printf("[WARN] App %s not matching %s", app.Id, id)
			continue
		}

		return app, nil
	}

	return App{}, fmt.Errorf("App (%s) not found", id)
}

func (c *ShuffleClient) makeRequest(method string, url string, body []byte) ([]byte, int, error) {
	client := &http.Client{}
	var req *http.Request
	var err error

	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		// set the HTTP method, url, and request body
		req, err = http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		if err != nil {
			defer req.Body.Close()
		}
	}

	if err != nil {
		return nil, -1, err
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.APIToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}

	return rbody, resp.StatusCode, nil
}
