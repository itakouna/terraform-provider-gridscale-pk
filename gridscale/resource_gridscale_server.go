package gridscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/parce-iot/gridscale"
	"container/list"
)

func resourceGridScaleServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceGridScaleServerCreate,
		Read:   resourceGridScaleServerRead,
		Update: resourceGridScaleServerUpdate,
		Delete: resourceGridScaleServerDelete,
		Schema: map[string]*schema.Schema{
			//Server parameters
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cores": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"location_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"labels":{
				Type: schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"power_on": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ordering": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceGridScaleServerCreate(d *schema.ResourceData, meta interface{}) error {

	api_client := meta.(*Config)
	api_client.CreateServer(
		d.Get("location_id").(string),
		d.Get("name").(string),
		d.Get("cores").(int),
		d.Get("memory").(int),
		d.Get("lables").([]string),
	)
	return resourceGridScaleServerRead(d, meta)
}

func resourceGridScaleServerRead(d *schema.ResourceData, meta interface{}) error {
	serverId := d.Id()
	api_client := meta.(*Config)
	_, server := api_client.GetServer(serverId)
	return server
}

func resourceGridScaleServerUpdate(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	server_id := d.Id()

	updateServerName(d, api_client, server_id)
	updateServerCores(d, api_client, server_id)
	updateServerMemory(d, api_client, server_id)
	if d.HasChange("lables") {
		api_client.UpdateServerLabels(
			server_id,
			d.Get("lables").([]string),
		)
	}
	if d.HasChange("cores") {
		api_client.PowerOnServer(
			server_id,
		)
	}
	api_client.PowerOffServer(
		server_id,
	)

	api_client.DisconnectIPAddress(
		d.Get("ip_address").(string),
		server_id,
	)

	api_client.ConnectNetwork(
		d.Get("network_id").(string),
		d.Get("ordering").(int),
		server_id,
	)

	api_client.DisconnectNetwork(
		d.Get("network_id").(string),
		server_id,
	)
	api_client.ConnectStorage(
		d.Get("storage_id").(string),
		d.Get("bootdevice").(bool),
		server_id,
	)

	api_client.DisconnectStorage(
		d.Get("storage_id").(string),
		server_id,
	)

	api_client.ConnectIsoImage(
		d.Get("iso_image_id").(string),
		server_id,
	)

	api_client.DisconnectIsoImage(
		d.Get("iso_image_id").(string),
		server_id,
	)
	return resourceGridScaleServerRead(d, meta)
}


func resourceGridScaleServerDelete(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	server_id := d.Id()
	api_client.GetServer(server_id)
	api_client.DeleteServer(server_id)

	return nil
}

