package gridscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	// "github.com/rancher/go-rancher/client"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRIDSCALE_API_URL", nil),
				Description: "",
			},
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRIDSCALE_API_TOKEN", nil),
				Description: "",
			},
			"user_uuid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRIDSCALE_USER_UUID", nil),
				Description: "",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"gridscale_server":  resourceGridScaleServer(),
			"gridscale_network": resourceGridScaleNetwork(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		API_URL:   d.Get("api_url").(string),
		API_TOKEN: d.Get("api_token").(string),
		USER_UUID: d.Get("user_uuid").(string),
	}

	err := config.CreateClient()

	return config, err
}