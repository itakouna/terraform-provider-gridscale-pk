package tests

import (
	"testing"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	gridscale_config "github.com/gridscale/terraform-provider-gridscale/gridscale"
	"github.com/parce-iot/gridscale"
)



func TestAccGridScaleServer_Basic(t *testing.T) {
	var server gridscale.Server
	serverName := "webserver"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDGridScaleServerDestroyCheck,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckProfitbricksServerConfig_basic, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridScaleServerExists("gridscale_server.webserver", &server),
				),
			},
		},
	})
}

func testAccCheckDGridScaleServerDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*gridscale_config.Config)
	for _, rs := range s.RootModule().Resources {
		resp := client.DeleteServer(rs.Primary.ID)

		if resp != nil {
			return fmt.Errorf("Server still exists %s %s", rs.Primary.ID, resp)
		}
	}

	return nil
}

func testAccCheckGridScaleServerAttributes(n string, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("testAccCheckGridScaleServerAttributes: Not found: %s", n)
		}
		if rs.Primary.Attributes["name"] != name {
			return fmt.Errorf("Bad name: %s", rs.Primary.Attributes["name"])
		}
		return nil
	}
}

func testAccCheckGridScaleServerExists(n string, server *gridscale.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*gridscale_config.Config)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckGridScaleServerExists: Not found: %s", n)
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

const testAccCheckProfitbricksServerConfig_basic = `
resource "gridscale_server" "webserver" {
  name = "%s"
  cores = 1
  memory = 1024
  location_id = "ttt4t"
}`

const testAccCheckProfitbricksServerConfig_update = `
resource "gridscale_datacenter" "foobar" {
	name       = "server-test"
	location = "us/las"
}

resource "gridscale_lan" "webserver_lan" {
  datacenter_id = "${gridscale_datacenter.foobar.id}"
  public = true
  name = "public"
}

resource "gridscale_server" "webserver" {
  name = "updated"
  datacenter_id = "${gridscale_datacenter.foobar.id}"
  cores = 1
  ram = 1024
  availability_zone = "ZONE_1"
  cpu_family = "AMD_OPTERON"
  volume {
    name = "system"
    size = 5
    disk_type = "SSD"
    image_name ="ubuntu-16.04"
    image_password = "K3tTj8G14a3EgKyNeeiY"
}
  nic {
    lan = "${gridscale_lan.webserver_lan.id}"
    dhcp = true
    firewall_active = true
    firewall {
      protocol = "TCP"
      name = "SSH"
      port_range_start = 22
      port_range_end = 22
    }
  }
}`
