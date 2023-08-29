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
				Description: "Address of the KEMP LoadMaster.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_SERVER", nil),
			},
			"api_key": &schema.Schema{
				Description: "API Key for authentication.",
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_API_KEY", ""),
			},
			"api_user": &schema.Schema{
				Description: "Username for KEMP LoadMaster API operations.",
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_API_USER", ""),
			},
			"api_pass": &schema.Schema{
				Description: "Password for KEMP LoadMaster API operations..",
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_API_PASS", ""),
			},
			"api_version": &schema.Schema{
				Description: "Use 1 for the old XML based API, 2 (default) for JSON.",
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOADMASTER_API_VERSION", 2),
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
	server := "https://" + d.Get("server").(string)
	apikey := d.Get("api_key").(string)
	apiuser := d.Get("api_user").(string)
	apipass := d.Get("api_pass").(string)
	apiversion := d.Get("api_version").(int)

	var diags diag.Diagnostics

	c := lmclient.NewClient(apikey, apiuser, apipass, server, apiversion)
	return c, diags
}
