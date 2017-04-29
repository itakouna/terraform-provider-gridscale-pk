package gridscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRIDSCALE_API_URL", nil),
				Description: "",
			},
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRIDSCALE_API_TOKEN", nil),
				Description: "",
			},
			"user_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRIDSCALE_USER_UUID", nil),
				Description: "",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"gridscale_server": resourceGridScaleServer(),
			"gridscale_network": resourceGridScaleNetwork(),
			"gridscale_storage": resourceGridScaleStorage(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Endpoint:  d.Get("api_url").(string),
		AuthToken: d.Get("api_token").(string),
		UserId:    d.Get("user_uuid").(string),
	}

	err := config.CreateClient()

	return config, err
}
