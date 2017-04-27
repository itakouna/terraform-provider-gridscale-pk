package gridscale

import (
	"testing"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/parce-iot/gridscale"
)



func TestAccGridScale_Basic(t *testing.T) {
	var network gridscale.Network
	networkName := "testnetwork"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {testAccPreCheck(t)},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDGridScaleNetworkDestroyCheck,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckGridScaleNetworkConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridScaleNetworkExists("gridscale_network.testnetwork", &network),
					testAccCheckGridScaleNetworkAttributes("gridscale_network.testnetwork", networkName),
					resource.TestCheckResourceAttr("gridscale_network.testnetwork", "name", networkName),
				),
			},
			resource.TestStep{
				Config: testAccCheckGridScaleNetworkConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridScaleNetworkAttributes("gridscale_network.testnetwork", "updatednetwork"),
					resource.TestCheckResourceAttr("gridscale_network.testnetwork", "name", "updatednetwork"),

				),
			},
		},
	})
}

func testAccCheckDGridScaleNetworkDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		sever, _ := client.GetNetwork(rs.Primary.ID)
		if sever == nil {
			return nil
		}
		err := client.DeleteNetwork(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Network %s was not deleted: error to %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckGridScaleNetworkAttributes(n string, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckGridScaleNetworkAttributes: Not found: %s", name)
		}
		if rs.Primary.Attributes["name"] != name {
			return fmt.Errorf("Bad name: %s", rs.Primary.Attributes["name"])
		}
		return nil
	}
}

func testAccCheckGridScaleNetworkExists(n string, network *gridscale.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Config)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckGridScaleNetworkExists: Not found: %s", s.RootModule())
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		foundNetwork, status := client.GetNetwork(rs.Primary.ID)

		if status != nil {
			return fmt.Errorf("Error occured while fetching Network: %s", rs.Primary.ID)
		}
		if foundNetwork.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

const testAccCheckGridScaleNetworkConfig_basic = `
resource "gridscale_network" "testnetwork" {
  name = "testnetwork"
  l2security = "true"
  location_uuid = "45ed677b-3702-4b36-be2a-a2eab9827950"
}`

const testAccCheckGridScaleNetworkConfig_update = `
resource "gridscale_network" "testnetwork" {
  name = "updatednetwork"
  location_uuid = "45ed677b-3702-4b36-be2a-a2eab9827950"
}`
