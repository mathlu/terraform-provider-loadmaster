package loadmaster

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lmclient "github.com/mathlu/loadmaster-go-client"
)

func GetVsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"address": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Required:    true,
			Description: "The IP address of the Virtual Service.",
		},
		"enable": &schema.Schema{
			Type:        schema.ValueType(schema.TypeBool),
			Optional:    true,
			Default:     true,
			Description: "Enable or disable the virtual server.",
		},
		"nickname": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Optional:    true,
			Description: "Specifies the \"friendly\" name of the service.",
		},
		"port": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Required:    true,
			Description: "The port for the Virtual Service.",
		},
		"protocol": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Required:    true,
			Description: "The protocol to be used for the Virtual Service.",
		},
		"layer": &schema.Schema{
			Type:        schema.ValueType(schema.TypeInt),
			Optional:    true,
			Description: "Network Layer for the service to run at (7 or 4).",
			Default:     7,
		},
		"type": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Optional:    true,
			Default:     "gen",
			Description: "Specifies the type of service being load balanced (gen, http, http2, ts, tls, or log).",
		},
		"defaultgw": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Optional:    true,
			Description: "Specify the Virtual Service-specific default gateway to be used and to send responses back to clients. If not set, the global default gateway will be used",
		},
		"id": &schema.Schema{
			Type:        schema.ValueType(schema.TypeString),
			Computed:    true,
			Description: "The ID of this resource.",
		},
	}
}

func resourceVs() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a resource for Virtual Services. Use this to create and manage Virtual Services in the KEMP LoadMaster.",
		CreateContext: resourceVsCreate,
		ReadContext:   resourceVsRead,
		UpdateContext: resourceVsUpdate,
		DeleteContext: resourceVsDelete,
		Schema:        GetVsSchema(),
	}
}

func resourceVsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lmclient.Client)

	vs := &lmclient.Vs{
		Address:   d.Get("address").(string),
		Port:      d.Get("port").(string),
		NickName:  d.Get("nickname").(string),
		Type:      d.Get("type").(string),
		Protocol:  d.Get("protocol").(string),
		Enable:    d.Get("enable").(bool),
		Layer:     d.Get("layer").(int),
		DefaultGW: d.Get("defaultgw").(string),
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

	d.SetId(strconv.Itoa(vc.Index))
	_ = d.Set("address", vc.Address)
	_ = d.Set("port", vc.VSPort)
	_ = d.Set("nickname", vc.NickName)
	_ = d.Set("type", vc.Type)
	_ = d.Set("protocol", vc.Protocol)
	_ = d.Set("enable", vc.Enable)
	_ = d.Set("layer", vc.Layer)
	_ = d.Set("defaultgw", vc.DefaultGW)
	return diags
}
func resourceVsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*lmclient.Client)
	i, _ := strconv.Atoi(d.Id())

	oldport, _ := d.GetChange("port")
	vs := &lmclient.Vs{
		Index:     i,
		Port:      oldport.(string),
		Address:   d.Get("address").(string),
		VSPort:    d.Get("port").(string),
		NickName:  d.Get("nickname").(string),
		Type:      d.Get("type").(string),
		Protocol:  d.Get("protocol").(string),
		Enable:    d.Get("enable").(bool),
		Layer:     d.Get("layer").(int),
		DefaultGW: d.Get("defaultgw").(string),
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
