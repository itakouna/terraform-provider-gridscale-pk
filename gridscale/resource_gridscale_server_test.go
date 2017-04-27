package gridscale

import (
	"testing"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/parce-iot/gridscale"
)



func TestAccGridScaleServer_Basic(t *testing.T) {
	var server gridscale.Server
	serverName := "testserver"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {testAccPreCheck(t)},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDGridScaleServerDestroyCheck,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckGridScaleServerConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridScaleServerExists("gridscale_server.testserver", &server),
					testAccCheckGridScaleServerAttributes("gridscale_server.testserver", serverName),
					resource.TestCheckResourceAttr("gridscale_server.testserver", "name", serverName),
				),
			},
			resource.TestStep{
				Config: testAccCheckGridScaleServerConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridScaleServerAttributes("gridscale_server.testserver", "updatedserver"),
					resource.TestCheckResourceAttr("gridscale_server.testserver", "name", "updatedserver"),
					resource.TestCheckResourceAttr("gridscale_server.testserver", "cores", "2"),
					resource.TestCheckResourceAttr("gridscale_server.testserver", "memory", "2"),
					resource.TestCheckResourceAttr("gridscale_server.testserver", "power_on",  "true"),
					resource.TestCheckResourceAttr("gridscale_server.testserver", "ordering",  "1"),
					resource.TestCheckResourceAttr("gridscale_server.testserver", "bootdevice",  "true"),


				),
			},
		},
	})
}

func testAccCheckDGridScaleServerDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		sever, _ := client.GetServer(rs.Primary.ID)
		if sever == nil {
			return nil
		}
		err := client.DeleteServer(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Server %s was not deleted: error to %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckGridScaleServerAttributes(n string, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckGridScaleServerAttributes: Not found: %s", name)
		}
		if rs.Primary.Attributes["name"] != name {
			return fmt.Errorf("Bad name: %s", rs.Primary.Attributes["name"])
		}
		return nil
	}
}

func testAccCheckGridScaleServerExists(n string, server *gridscale.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Config)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckGridScaleServerExists: Not found: %s", s.RootModule())
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		foundServer, status := client.GetServer(rs.Primary.ID)

		if status != nil {
			return fmt.Errorf("Error occured while fetching Server: %s", rs.Primary.ID)
		}
		if foundServer.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

const testAccCheckGridScaleServerConfig_basic = `
resource "gridscale_server" "testserver" {
  name = "testserver"
  cores = 1
  memory = 1
  location_uuid = "45ed677b-3702-4b36-be2a-a2eab9827950"
}
`

const testAccCheckGridScaleServerConfig_update = `
resource "gridscale_storage" "serverstorage" {
  name = "servertorage"
  capacity = "1"
  location_uuid = "45ed677b-3702-4b36-be2a-a2eab9827950"
}

resource "gridscale_network" "servernetwork" {
  name = "servernetwork"
  l2security = "true"
  location_uuid = "45ed677b-3702-4b36-be2a-a2eab9827950"
}

resource "gridscale_server" "testserver" {
  location_uuid = "45ed677b-3702-4b36-be2a-a2eab9827950"
  name = "updatedserver"
  cores = 2
  memory = 2
  power_on = true
  storage_id = "${gridscale_storage.serverstorage.id}"
  bootdevice = true
  network_id = "${gridscale_network.servernetwork.id}"
  ordering = 1
}`
