package gridscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"testing"
	"os"
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
		nil,
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
	updateServerNetwork(d,api_client, server_id)
	updateServerStorage(d,api_client, server_id)
	updateServerPower(d,api_client, server_id)
	updateServerIsoImage(d,api_client, server_id)

	return resourceGridScaleServerRead(d, meta)
}


func resourceGridScaleServerDelete(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	server_id := d.Id()
	api_client.GetServer(server_id)
	api_client.DeleteServer(server_id)

	return nil
}

func testAccPreCheck(t *testing.T) {

	if v := os.Getenv("GRIDSCALE_API_URL"); v == "" {
		t.Fatal("GRIDSCALE_API_URL must be set for acceptance tests")
	}

	if v := os.Getenv("GRIDSCALE_API_TOKEN"); v == "" {
		t.Fatal("GRIDSCALE_API_TOKEN must be set for acceptance tests")
	}

	if v := os.Getenv("GRIDSCALE_USER_UUID"); v == "" {
		t.Fatal("GRIDSCALE_USER_UUID must be set for acceptance tests")
	}
}
