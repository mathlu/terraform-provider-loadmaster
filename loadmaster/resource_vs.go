package loadmaster

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lmclient "github.com/mathlu/loadmaster-go-client"
)

func resourceVs() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVsCreate,
		ReadContext:   resourceVsRead,
		UpdateContext: resourceVsUpdate,
		DeleteContext: resourceVsDelete,
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"nickname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"layer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceVsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lmclient.Client)

	vs := &lmclient.Vs{
		Address:  d.Get("address").(string),
		Protocol: d.Get("protocol").(string),
		Port:     d.Get("port").(string),
		NickName: d.Get("nickname").(string),
		Enable:   d.Get("enable").(bool),
	}
	vc, err := c.CreateVs(vs)

	if err != nil {
		return diag.FromErr(err)
	}
	i := strconv.Itoa(vc.Index)
	d.SetId(i)

	return resourceVsRead(ctx, d, m)
}
func resourceVsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*lmclient.Client)

	id := d.Id()
	i, _ := strconv.Atoi(id)
	vc, err := c.GetVs(i)
	if err != nil {
		return diag.FromErr(err)
	}

	//d.SetId(strconv.Itoa(vc.Index))
	d.Set("address", vc.Address)
	d.Set("nickname", vc.NickName)
	d.Set("port", vc.Port)
	d.Set("protocol", vc.Protocol)
	d.Set("layer", vc.Layer)
	d.Set("enable", vc.Enable)
	return diags
}
func resourceVsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*lmclient.Client)
	i, _ := strconv.Atoi(d.Id())

	vs := &lmclient.Vs{
		Index:    i,
		Address:  d.Get("address").(string),
		Protocol: d.Get("protocol").(string),
		Port:     d.Get("port").(string),
		NickName: d.Get("nickname").(string),
		Enable:   d.Get("enable").(bool),
	}

	_, err := c.ModifyVs(vs)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
func resourceVsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*lmclient.Client)

	id := d.Id()
	i, _ := strconv.Atoi(id)
	_, err := c.DeleteVs(i)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
