package loadmaster
import (
  "os"
  "testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
        testAccProviders = map[string]*schema.Provider{
		"loadmaster": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
  	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("LOADMASTER_SERVER"); v == "" {
		t.Fatal("LOADMASTER_SERVER must be set for acceptance tests")
	}

	if v := os.Getenv("LOADMASTER_USERNAME"); v == "" {
		t.Fatal("LOADMASTER_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("LOADMASTER_PASSWORD"); v == "" {
		t.Fatal("LOADMASTER_PASSWORD must be set for acceptance tests")
	}
}
