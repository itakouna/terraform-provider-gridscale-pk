package gridscale

import (
	"testing"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/parce-iot/gridscale"
)



func TestAccGridScaleStorage_Basic(t *testing.T) {
	var storage gridscale.Storage
	storageName := "teststorage"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {testAccPreCheck(t)},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDGridScaleStorageDestroyCheck,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckGridScaleStorageConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridScaleStorageExists("gridscale_storage.teststorage", &storage),
					testAccCheckGridScaleStorageAttributes("gridscale_storage.teststorage", storageName),
					resource.TestCheckResourceAttr("gridscale_storage.teststorage", "name", storageName),
				),
			},
			resource.TestStep{
				Config: testAccCheckGridScaleStorageConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridScaleStorageAttributes("gridscale_storage.teststorage", "updatedstorage"),
					resource.TestCheckResourceAttr("gridscale_storage.teststorage", "name", "updatedstorage"),
					resource.TestCheckResourceAttr("gridscale_storage.teststorage", "capacity", "2"),
				),
			},
		},
	})
}

func testAccCheckDGridScaleStorageDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		storage, _ := client.GetStorage(rs.Primary.ID)
		if storage == nil {
			return nil
		}
		client.DeleteStorage(rs.Primary.ID)
		/*if err != nil {
			return fmt.Errorf("Storage %s was not deleted: error to %s", rs.Primary.ID, err)
		}*/
	}

	return nil
}

func testAccCheckGridScaleStorageAttributes(n string, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckGridScaleStorageAttributes: Not found: %s", name)
		}
		if rs.Primary.Attributes["name"] != name {
			return fmt.Errorf("Bad name: %s", rs.Primary.Attributes["name"])
		}
		return nil
	}
}

func testAccCheckGridScaleStorageExists(n string, storage *gridscale.Storage) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Config)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckGridScaleStorageExists: Not found: %s", s.RootModule())
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		foundStorage, status := client.GetStorage(rs.Primary.ID)

		if status != nil {
			return fmt.Errorf("Error occured while fetching Storage: %s", rs.Primary.ID)
		}
		if foundStorage.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

const testAccCheckGridScaleStorageConfig_basic = `
resource "gridscale_storage" "teststorage" {
  name = "teststorage"
  capacity = "1"
  location_uuid = "45ed677b-3702-4b36-be2a-a2eab9827950"
}`

const testAccCheckGridScaleStorageConfig_update = `
resource "gridscale_storage" "teststorage" {
  name = "updatedstorage"
  capacity = "2"
  location_uuid = "45ed677b-3702-4b36-be2a-a2eab9827950"
}`
