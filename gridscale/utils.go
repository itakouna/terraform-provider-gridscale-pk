package gridscale

import "github.com/hashicorp/terraform/helper/schema"

func updateServerName(d *schema.ResourceData, api_client *Config, serverId string) () {
	if d.HasChange("name") {
		_, name := d.GetChange("name")
		api_client.UpdateServerName(
			serverId,
			name.(string),
		)
	}
}

func updateServerCores(d *schema.ResourceData, api_client *Config, serverId string) () {
	if d.HasChange("cores") {
		_, name := d.GetChange("cores")
		api_client.UpdateServerCores(
			serverId,
			name.(int),
		)
	}
}

func updateServerMemory(d *schema.ResourceData, api_client *Config, serverId string) () {
	if d.HasChange("memory") {
		_, memory := d.GetChange("memory")
		api_client.UpdateServerMemory(
			serverId,
			memory.(int),
		)
	}
}

func updateServerLables(d *schema.ResourceData, api_client *Config, serverId string) () {
	if d.HasChange("lables") {
		api_client.UpdateServerLabels(
			serverId,
			d.Get("lables").([]string),
		)
	}
}

func updateServerPower(d *schema.ResourceData, api_client *Config, serverId string) () {

	if d.HasChange("power_on") {
		if d.Get("power_on").(bool) {
			api_client.PowerOnServer(serverId,)
		} else {
			api_client.PowerOffServer(serverId,)
		}

	}

}

func updateServerNetwork(d *schema.ResourceData, api_client *Config, serverId string) () {
	if d.HasChange("connect") {
		if d.Get("connect").(bool){
			api_client.ConnectNetwork(
				d.Get("network_id").(string),
				d.Get("ordering").(int),
				serverId,
			)
		} else{
			api_client.DisconnectNetwork(
				d.Get("network_id").(string),
				serverId,
			)
		}
	}
}
func updateServerStorage(d *schema.ResourceData, api_client *Config, serverId string) () {

	api_client.ConnectStorage(
		d.Get("storage_id").(string),
		d.Get("bootdevice").(bool),
		serverId,
	)

	api_client.DisconnectStorage(
		d.Get("storage_id").(string),
		serverId,
	)
}

func updateServerIsoImage(d *schema.ResourceData, api_client *Config, serverId string) () {
	api_client.DisconnectIsoImage(
		d.Get("iso_image_id").(string),
		serverId,
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


func updateStorageName(d *schema.ResourceData, api_client *Config, storageId string) () {
	if d.HasChange("name") {
		api_client.UpdateStorageName(
			storageId,
			d.Get("name").(string),
		)
	}
}

func updateStorageCapacity(d *schema.ResourceData, api_client *Config, storageId string) () {
	if d.HasChange("capacity") {
		api_client.UpdateStorageCapacity(
			storageId,
			d.Get("capacity").(int),
		)
	}
}
