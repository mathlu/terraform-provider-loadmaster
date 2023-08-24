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
		"virtual_service_id": &schema.Schema{
			Type:        schema.ValueType(schema.TypeInt),
			Required:    true,
			Description: "Virtual Service index number.",
		},
		"address": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Required:    true,
			Description: "The Real Server IP address.",
		},
		"port": &schema.Schema{
			Type:        schema.ValueType(schema.TypeInt),
			Required:    true,
			Description: "The port on the Real Server to be used.",
		},
		"id": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Computed:    true,
			Description: "The ID of this resource.",
		},
	}
}

func resourceRs() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a resource for Real Servers. Use this to create and manage Real Servers assigned to Virtual Services in the KEMP LoadMaster.",
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
		VSIndex: d.Get("virtual_service_id").(int),
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

	vsid := d.Get("virtual_service_id").(int)
	id := d.Id()
	i, _ := strconv.Atoi(id)
	rc, err := c.GetRs(i, vsid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(rc.RsIndex))
	d.Set("address", rc.Addr)
	d.Set("port", rc.Port)
	d.Set("virtual_service_id", rc.VSIndex)
	return diags
}
func resourceRsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*lmclient.Client)

	rs := &lmclient.Rs{
		Rsi:     "!" + d.Id(),
		VSIndex: d.Get("virtual_service_id").(int),
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

	vsid := d.Get("virtual_service_id").(int)
	id := d.Id()
	i, _ := strconv.Atoi(id)
	_, err := c.DeleteRs(i, vsid)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
