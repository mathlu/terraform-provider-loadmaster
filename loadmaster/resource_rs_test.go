package loadmaster

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceRs(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"loadmaster_real_server.baz", "address", "192.168.1.11"),
				),
			},
		},
	})
}

const testAccResourceRs = `
resource "loadmaster_virtual_service" "foo" {
  address   = "192.168.1.10"
  protocol  = "tcp"
  port      = "8080"
  nickname  = "bar"
  type      = "gen"
  checktype = "tcp"
}


resource "loadmaster_real_server" "baz" {
  virtual_service_id = loadmaster_virtual_service.foo.id
  address            = "192.168.1.11"
  port               = "8080"
}
`
