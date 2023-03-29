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
		ReadContext: dataSourceVsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nickname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
	d.Set("address", vc.Address)
	d.Set("nickname", vc.NickName)
	d.Set("port", vc.Port)
	d.Set("protocol", vc.Protocol)
	return nil

}
