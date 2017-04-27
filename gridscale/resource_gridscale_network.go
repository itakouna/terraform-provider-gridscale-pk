package gridscale

import (
	"github.com/hashicorp/terraform/helper/schema"
)



func resourceGridScaleNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceGridScaleNetworkCreate,
		Read:   resourceGridScaleNetworkRead,
		Update: resourceGridScaleNetworkUpdate,
		Delete: resourceGridScaleNetworkDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"l2security": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"labels":{
				Type: schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func resourceGridScaleNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	network, err := api_client.CreateNetwork(
		d.Get("location_uuid").(string),
		d.Get("name").(string),
		d.Get("l2security").(bool),
		nil,
	)
	if err != nil {
		return err
	}
	d.SetId(network.ID)
	return resourceGridScaleNetworkRead(d, meta)
}

func resourceGridScaleNetworkRead(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	networkId := d.Id()

	network, err := api_client.GetNetwork(networkId)
	if err != nil {
		return err
	}

	d.Set("name", network.Name)
	return nil
}

func resourceGridScaleNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	networkId := d.Id()
	updateNetworkName(d, api_client, networkId)
	updateNetworkLabels(d, api_client, networkId)

	return resourceGridScaleNetworkRead(d, meta)
}

func resourceGridScaleNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	networkId := d.Id()
	err := api_client.DeleteNetwork(networkId)

	if err != nil {
		return err
	}
	d.SetId("")

	return nil
}