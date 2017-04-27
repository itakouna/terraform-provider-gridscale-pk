package gridscale

import (
	"github.com/hashicorp/terraform/helper/schema"
)



func resourceGridScaleStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceGridScaleStorageCreate,
		Read:   resourceGridScaleStorageRead,
		Update: resourceGridScaleStorageUpdate,
		Delete: resourceGridScaleStorageDelete,
		Schema: map[string]*schema.Schema{
			"location_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"capacity": {
				Type:     schema.TypeInt,
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

func resourceGridScaleStorageCreate(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	storage, err := api_client.CreateStorage(
		d.Get("location_uuid").(string),
		d.Get("name").(string),
		d.Get("capacity").(int),
		nil,
		nil,
	)
	if err != nil {
		return err
	}
	d.SetId(storage.ID)
	return resourceGridScaleStorageRead(d, meta)
}

func resourceGridScaleStorageRead(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	storageId := d.Id()

	storage, err := api_client.GetStorage(storageId)
	if err != nil {
		return err
	}

	d.Set("name", storage.Name)
	return nil
}

func resourceGridScaleStorageUpdate(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	storageId := d.Id()

	updateStorageName(d, api_client, storageId)
	updateStorageCapacity(d, api_client, storageId)

	return resourceGridScaleStorageRead(d, meta)
}

func resourceGridScaleStorageDelete(d *schema.ResourceData, meta interface{}) error {
	api_client := meta.(*Config)
	storageId := d.Id()
	err := api_client.DeleteStorage(storageId)

	if err != nil {
		return err
	}
	d.SetId("")

	return nil
}
