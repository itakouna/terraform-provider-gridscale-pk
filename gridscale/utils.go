package gridscale

import "github.com/hashicorp/terraform/helper/schema"

func updateServerName(d *schema.ResourceData, api_client *Config, server_id string) () {
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
		api_client.UpdateServerCores(
			server_id,
			name.(int),
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

func updateServerLables(d *schema.ResourceData, api_client *Config, server_id string) () {
	if d.HasChange("lables") {
		api_client.UpdateServerLabels(
			server_id,
			d.Get("lables").([]string),
		)
	}
}

func updateServerPower(d *schema.ResourceData, api_client *Config, server_id string) () {

	if d.HasChange("power_on") {
		if d.Get("power_on").(bool) {
			api_client.PowerOnServer(server_id,)
		} else {
			api_client.PowerOffServer(server_id,)
		}

	}

}

func updateServerNetwork(d *schema.ResourceData, api_client *Config, server_id string) () {
	if d.HasChange("connect") {
		if d.Get("connect").(bool){
			api_client.ConnectNetwork(
				d.Get("network_id").(string),
				d.Get("ordering").(int),
				server_id,
			)
		} else{
			api_client.DisconnectNetwork(
				d.Get("network_id").(string),
				server_id,
			)
		}
	}
}
func updateServerStorage(d *schema.ResourceData, api_client *Config, server_id string) () {

	api_client.ConnectStorage(
		d.Get("storage_id").(string),
		d.Get("bootdevice").(bool),
		server_id,
	)

	api_client.DisconnectStorage(
		d.Get("storage_id").(string),
		server_id,
	)
}

func updateServerIsoImage(d *schema.ResourceData, api_client *Config, server_id string) () {
	api_client.DisconnectIsoImage(
		d.Get("iso_image_id").(string),
		server_id,
	)
}

func updateNetworkName(d *schema.ResourceData, api_client *Config, networkId string) () {
	if d.HasChange("name") {
		api_client.UpdateNetworkName(
			networkId,
			d.Get("name").(string),
		)
	}
}

func updateNetworkLabels(d *schema.ResourceData, api_client *Config, networkId string) () {
	if d.HasChange("labels") {
		api_client.UpdateNetworkLabels(
			networkId,
			d.Get("labels").([]string),
		)
	}
}
