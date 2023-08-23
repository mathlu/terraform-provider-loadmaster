package loadmaster

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lmclient "github.com/mathlu/loadmaster-go-client"
)

func GetRsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"status": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Computed: true,
		},
		"vsindex": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Required: true,
		},
		"address": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Required: true,
		},
		"port": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Required: true,
		},
	}
}

func resourceRs() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRsCreate,
		ReadContext:   resourceRsRead,
		UpdateContext: resourceRsUpdate,
		DeleteContext: resourceRsDelete,
		Schema:        GetRsSchema(),
	}
}

func resourceRsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lmclient.Client)

	rs := &lmclient.Rs{
		Addr:    d.Get("address").(string),
		Port:    d.Get("port").(int),
		VSIndex: d.Get("vsindex").(int),
	}
	rc, err := c.CreateRs(rs)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(rc.RsIndex))

	return resourceRsRead(ctx, d, m)
}
func resourceRsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*lmclient.Client)

	vsid := d.Get("vsindex").(int)
	id := d.Id()
	i, _ := strconv.Atoi(id)
	rc, err := c.GetRs(i, vsid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(rc.RsIndex))
	d.Set("address", rc.Addr)
	d.Set("port", rc.Port)
	d.Set("vsindex", rc.VSIndex)
	return diags
}
func resourceRsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*lmclient.Client)

	rs := &lmclient.Rs{
		Rsi:     "!" + d.Id(),
		VSIndex: d.Get("vsindex").(int),
		NewPort: strconv.Itoa(d.Get("port").(int)),
		Addr:    d.Get("address").(string),
	}

	_, err := c.ModifyRs(rs)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
func resourceRsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*lmclient.Client)

	vsid := d.Get("vsindex").(int)
	id := d.Id()
	i, _ := strconv.Atoi(id)
	_, err := c.DeleteRs(i, vsid)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
