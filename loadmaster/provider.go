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
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_API_TOKEN", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"loadmaster_vs": dataSourceVs(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"loadmaster_vs": resourceVs(),
			"loadmaster_rs": resourceRs(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	server := d.Get("server").(string)
	api_token := d.Get("api_token").(string)

	var diags diag.Diagnostics

	c := lmclient.NewClient(api_token, server)
	return c, diags
}
