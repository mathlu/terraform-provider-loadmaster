package loadmaster

import (
	"context"
	"strconv"

	lmclient "github.com/mathlu/loadmaster-go-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVs() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retreive information about an existing Virtual Service.",
		ReadContext: dataSourceVsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address of the Virtual Service.",
			},
			"nickname": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The \"friendly\" name of the service.",
			},
			"port": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The port for the Virtual Service.",
			},
			"protocol": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The protocol used for the Virtual Service.",
			},
		},
	}
}

func dataSourceVsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lmclient.Client)

	vcID := d.Get("id").(int)

	vc, err := c.GetVs(vcID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(vc.Index))
	_ = d.Set("address", vc.Address)
	_ = d.Set("nickname", vc.NickName)
	_ = d.Set("port", vc.Port)
	_ = d.Set("protocol", vc.Protocol)
	return nil

}
