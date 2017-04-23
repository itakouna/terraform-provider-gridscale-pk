package gridscale

import "github.com/hashicorp/terraform/helper/schema"

func updateServerName(d *schema.ResourceData, name, api_client *Config, server_id string) () {
	if d.HasChange("name") {
		_, name := d.GetChange("name")
		api_client.UpdateServerName(
			server_id,
			name.(string),
		)
	}
}

func updateServerCores(d *schema.ResourceData, api_client *Config, server_id string) () {
	if d.HasChange("cores") {
		_, name := d.GetChange("cores")
		api_client.UpdateServerName(
			server_id,
			name.(string),
		)
	}
}

func updateServerMemory(d *schema.ResourceData, api_client *Config, server_id string) () {
	if d.HasChange("memory") {
		_, memory := d.GetChange("memory")
		api_client.UpdateServerMemory(
			server_id,
			memory.(int),
		)
	}
}
