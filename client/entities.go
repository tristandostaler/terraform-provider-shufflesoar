package client

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type ContactInfo struct {
	Name string
	Url  string
}

func GetDefaultContactInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
			},
			"url": {
				Type: schema.TypeString,
			},
		},
	}
}

type ReferenceInfo struct {
	DocumentationUrl string `json:"documentation_url"`
	GithubUrl        string `json:"github_url"`
}

func GetDefaultReferenceInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"documentation_url": {
				Type: schema.TypeString,
			},
			"github_url": {
				Type: schema.TypeString,
			},
		},
	}
}

type FolderMount struct {
	FolderMount       bool   `json:"folder_mount"`
	SourceFolder      string `json:"source_folder"`
	DestinationFolder string `json:"destination_folder"`
}

func GetDefaultFolderMountSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"folder_mount": {
				Type: schema.TypeBool,
			},
			"source_folder": {
				Type: schema.TypeString,
			},
			"destination_folder": {
				Type: schema.TypeString,
			},
		},
	}
}

type AuthenticationParameterSchema struct {
	Type string
}

func GetDefaultAuthenticationParameterSchemaSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
			},
		},
	}
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

func GetDefaultAuthenticationParameterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
			},
			"id": {
				Type: schema.TypeString,
			},
			"name": {
				Type: schema.TypeString,
			},
			"example": {
				Type: schema.TypeString,
			},
			"multiline": {
				Type: schema.TypeBool,
			},
			"required": {
				Type: schema.TypeBool,
			},
			"in": {
				Type: schema.TypeString,
			},
			"schema": {
				Type: schema.TypeList,
				Elem: GetDefaultAuthenticationParameterSchemaSchema(),
			},
			"scheme": {
				Type: schema.TypeString,
			},
		},
	}
}

type Authentication struct {
	Type         string
	Required     bool
	Parameters   []AuthenticationParameter
	RedirectUri  string `json:"redirect_uri"`
	TokenUri     string `json:"token_uri"`
	RefreshUri   string `json:"refresh_uri"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func GetDefaultAuthenticationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
			},
			"required": {
				Type: schema.TypeBool,
			},
			"parameters": {
				Type: schema.TypeList,
				Elem: GetDefaultAuthenticationParameterSchema(),
			},
			"redirect_uri": {
				Type: schema.TypeString,
			},
			"token_uri": {
				Type: schema.TypeString,
			},
			"refresh_uri": {
				Type: schema.TypeString,
			},
			"client_id": {
				Type: schema.TypeString,
			},
			"client_secret": {
				Type: schema.TypeString,
			},
		},
	}
}

type Version struct {
	Version string
	Id      string
}

func GetDefaultVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version": {
				Type: schema.TypeString,
			},
			"id": {
				Type: schema.TypeString,
			},
		},
	}
}

type AppAuthentication struct {
	Name           string
	IsValid        bool `json:"is_valid"`
	Id             string
	Link           string
	AppVersion     string `json:"app_version"`
	SharingConfig  string `json:"sharing_config"`
	Generated      bool
	Downloaded     bool
	Sharing        bool
	Verified       bool
	Invalid        bool
	Activated      bool
	Tested         bool
	Hash           string
	PrivateId      string `json:"private_id"`
	Description    string
	Environment    string
	SmallImage     string        `json:"small_image"`
	LargeImage     string        `json:"large_image"`
	ContactInfo    ContactInfo   `json:"contact_info"`
	ReferenceInfo  ReferenceInfo `json:"reference_info"`
	FolderMount    FolderMount   `json:"folder_mount"`
	Actions        interface{}
	Authentication Authentication
	Tags           []string
	Categories     []string
	Created        int
	Edited         int
	LastRuntime    int `json:"last_runtime"`
	Versions       []Version
	LoopVersions   []string `json:"loop_versions"`
	Owner          string
	Public         bool
	ReferenceOrg   string `json:"reference_org"`
	ReferenceUrl   string `json:"reference_url"`
	ActionFilePath string `json:"action_file_path"`
	Documentation  string
}

func GetDefaultAppAuthenticationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				// Computed: true,
				// Required:    true,
				Description: "The App name to link this authentication config to. This must match an existing App's name",
			},
			"id": {
				Type: schema.TypeString,
				// Computed: true,
				// Optional:    true,
				Description: "The App Id of the App to link this authentication config to",
			},
			"large_image": {
				Type: schema.TypeString,
				// Computed: true,
				// Optional:    true,
				Description: "The base64 string for the image to display. Format: data:image/png;base64,THE_BASE64",
			},
			"is_valid": {
				Type: schema.TypeBool,
			},
			"link": {
				Type: schema.TypeString,
			},
			"app_version": {
				Type: schema.TypeString,
			},
			"sharing_config": {
				Type: schema.TypeString,
			},
			"generated": {
				Type: schema.TypeBool,
			},
			"downloaded": {
				Type: schema.TypeBool,
			},
			"sharing": {
				Type: schema.TypeBool,
			},
			"verified": {
				Type: schema.TypeBool,
			},
			"invalid": {
				Type: schema.TypeBool,
			},
			"activated": {
				Type: schema.TypeBool,
			},
			"tested": {
				Type: schema.TypeBool,
			},
			"hash": {
				Type: schema.TypeString,
			},
			"private_id": {
				Type: schema.TypeString,
			},
			"description": {
				Type: schema.TypeString,
			},
			"environment": {
				Type: schema.TypeString,
			},
			"small_image": {
				Type: schema.TypeString,
			},
			"tags": {
				Type: schema.TypeString,
			},
			"categories": {
				Type: schema.TypeString,
			},
			"created": {
				Type: schema.TypeInt,
			},
			"edited": {
				Type: schema.TypeInt,
			},
			"last_runtime": {
				Type: schema.TypeInt,
			},
			"loop_versions": {
				Type: schema.TypeString,
			},
			"owner": {
				Type: schema.TypeString,
			},
			"public": {
				Type: schema.TypeBool,
			},
			"reference_org": {
				Type: schema.TypeString,
			},
			"reference_url": {
				Type: schema.TypeString,
			},
			"action_file_path": {
				Type: schema.TypeString,
			},
			"documentation": {
				Type: schema.TypeString,
			},
			"versions": {
				Type: schema.TypeList,
				Elem: GetDefaultVersionSchema(),
			},
			"contact_info": {
				Type: schema.TypeList,
				Elem: GetDefaultContactInfoSchema(),
			},
			"reference_info": {
				Type: schema.TypeList,
				Elem: GetDefaultReferenceInfoSchema(),
			},
			"folder_mount": {
				Type: schema.TypeList,
				Elem: GetDefaultFolderMountSchema(),
			},
			"authentication": {
				Type: schema.TypeList,
				Elem: GetDefaultAuthenticationSchema(),
			},
		},
	}
}

type Field struct {
	Key   string
	Value string
}

func GetDefaultFieldSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type: schema.TypeString,
				// Computed: true,
			},
			"value": {
				Type: schema.TypeString,
				// Computed: true,
			},
		},
	}
}

type AppUsage struct {
	WorkfflowId string `json:"workflow_id"`
	Nodes       []string
}

func GetDefaultAppUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"workflow_id": {
				Type: schema.TypeString,
			},
			"nodes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

type App struct {
	Active            bool
	Label             string
	Id                string
	App               AppAuthentication
	Fields            []Field
	Usage             []AppUsage
	WorkflowCount     int    `json:"workflow_count"`
	NodeCount         int    `json:"node_count"`
	OrgId             string `json:"org_id"`
	Created           int
	Edited            int
	Defined           bool
	Type              string
	Encrypted         bool
	ReferenceWorkflow string `json:"reference_workflow"`
}

func GetDefaultAppSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": {
				Type: schema.TypeBool,
			},
			"label": {
				Type:        schema.TypeString,
				Description: "The text to display in the Shuffle UI",
			},
			"id": {
				Type: schema.TypeString,
			},
			"app": {
				Type:        schema.TypeList,
				Description: "A block for the app authentication settings",
				Elem:        GetDefaultAppAuthenticationSchema(),
			},
			"fields": {
				Type:        schema.TypeList,
				Description: "This is a list of all the required fields for this app authentication. The name of the fields must match the names in the authentication parameters and there must be the same number of parameters and fields.",
				Elem:        GetDefaultFieldSchema(),
			},
			"usage": {
				Type: schema.TypeList,
				Elem: GetDefaultAppUsageSchema(),
			},
			"workflow_count": {
				Type: schema.TypeInt,
			},
			"node_count": {
				Type: schema.TypeInt,
			},
			"org_id": {
				Type: schema.TypeString,
			},
			"created": {
				Type: schema.TypeInt,
			},
			"edited": {
				Type: schema.TypeInt,
			},
			"defined": {
				Type: schema.TypeBool,
			},
			"type": {
				Type: schema.TypeString,
			},
			"encrypted": {
				Type: schema.TypeBool,
			},
			"referenceworkflow": {
				Type: schema.TypeString,
			},
		},
	}
}

type GetAppResponse struct {
	Data    []App
	Success bool
}

type CreateOrUpdateResponse struct {
	Success bool
	Id      string
}
