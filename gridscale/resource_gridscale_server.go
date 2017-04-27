package gridscale

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGridScaleServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceGridScaleServerCreate,
		Read:   resourceGridScaleServerRead,
		Update: resourceGridScaleServerUpdate,
		Delete: resourceGridScaleServerDelete,
		Schema: map[string]*schema.Schema{
			//Server parameters
			"location_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"iso_image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"labels":{
				Type: schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"power_on": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bootdevice": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ordering": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceGridScaleServerCreate(d *schema.ResourceData, meta interface{}) error {

	api_client := meta.(*Config)
	server, err := api_client.CreateServer(
		d.Get("location_uuid").(string),
		d.Get("name").(string),
		d.Get("cores").(int),
		d.Get("memory").(int),
		nil,
	)
	if err != nil {
		return err
	}
	d.SetId(server.ID)
	return resourceGridScaleServerRead(d, meta)
}

func resourceGridScaleServerRead(d *schema.ResourceData, meta interface{}) error {
	serverId := d.Id()
	api_client := meta.(*Config)
	server, err := api_client.GetServer(serverId)

	if err != nil {
		return err
	}

	d.Set("name", server.Name)

	return nil
}

func resourceGridScaleServerUpdate(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	serverId := d.Id()

	updateServerName(d, api_client, serverId)
	updateServerCores(d, api_client, serverId)
	updateServerMemory(d, api_client, serverId)
	updateServerNetwork(d,api_client, serverId)
	updateServerStorage(d,api_client, serverId)
	//updateServerPower(d,api_client, serverId)
	updateServerIsoImage(d,api_client, serverId)

	return resourceGridScaleServerRead(d, meta)
}


func resourceGridScaleServerDelete(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	serverId := d.Id()
	err := api_client.DeleteServer(serverId)

	if err != nil {
		return err
	}
	d.SetId("")

	return nil
}
