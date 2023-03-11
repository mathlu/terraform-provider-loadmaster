package loadmaster

import (
        "context"

        "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_SERVER", nil),
				Description: "Loadmaster server IP address.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_USERNAME", nil),
				Description: "User to authenticate with Loadmaster server.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_PASSWORD", nil),
				Description: "Password to authenticate with LOADMASTER server.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
		},
		DataSourcesMap: map[string]*schema.Resource{
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

  return nil, diag.Diagnostics{diag.Diagnostic{Summary: ""}}
}
