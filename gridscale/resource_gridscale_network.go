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
		},
	}
}

func resourceGridScaleNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	gridscale_client := meta.(*Config)
	gridscale_client.CreateNetwork()
	
	return resourceGridScaleNetworkRead(d, meta)
}

func resourceGridScaleNetworkRead(d *schema.ResourceData, meta interface{}) error {
	gridscale_client := meta.(*Config)
	gridscale_client.GetNetwork()

	return nil
}

func resourceGridScaleNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	gridscale_client := meta.(*Config)
	gridscale_client.UpdateNetworkName()

	return resourceGridScaleNetworkRead(d, meta)
}

func resourceGridScaleNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	gridscale_client := meta.(*Config)
	gridscale_client.DeleteNetwork()
	return nil
}