package loadmaster

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceVs(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"loadmaster_virtual_service.foo", "nickname", "bar"),
					resource.TestCheckResourceAttr(
						"loadmaster_virtual_service.foo", "defaultgw", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"loadmaster_virtual_service.foo", "port", "8080"),
					resource.TestCheckResourceAttr(
						"loadmaster_virtual_service.foo", "type", "gen"),
					resource.TestCheckResourceAttr(
						"loadmaster_virtual_service.foo", "address", "192.168.1.10"),
					resource.TestCheckResourceAttr(
						"loadmaster_virtual_service.foo", "protocol", "tcp"),
				),
			},
		},
	})
}

const testAccResourceVs = `
resource "loadmaster_virtual_service" "foo" {
  address   = "192.168.1.10"
  protocol  = "tcp"
  port      = "8080"
  nickname  = "bar"
  type      = "gen"
  defaultgw = "192.168.1.1"
  checktype = "tcp"
}
`
