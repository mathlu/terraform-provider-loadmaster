package loadmaster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lmclient "github.com/mathlu/loadmaster-go-client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_SERVER", nil),
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_API_KEY", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"loadmaster_virtual_service": dataSourceVs(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"loadmaster_virtual_service": resourceVs(),
			"loadmaster_real_server":     resourceRs(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	server := d.Get("server").(string)
	apikey := d.Get("api_key").(string)

	var diags diag.Diagnostics

	c := lmclient.NewClient(apikey, server)
	return c, diags
}
