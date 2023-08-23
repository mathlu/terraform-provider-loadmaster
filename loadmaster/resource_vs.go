package loadmaster

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lmclient "github.com/mathlu/loadmaster-go-client"
)

func GetVsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"address": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Required: true,
		},
		"addvia": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"alertthreshold": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"allowhttp2": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"anomalyscoringthreshold": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
			Default:  100,
		},
		"bandwidth": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"blockingparanoia": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
			Default:  1,
		},
		"bodylimit": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
			Default:  1048576,
		},
		"cache": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"checkport": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
			Default:  "0",
		},
		"checktype": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
			Default:  "tcp",
		},
		"checkuse11": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"checkuseget": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"chkinterval": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"chkretrycount": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"chktimeout": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"clientcert": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"compress": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"connsperseclimit": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"enable": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
			Default:  true,
		},
		"enhancedhealthchecks": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"errorcode": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
			Default:  "0",
		},
		"espenabled": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"executingparanoia": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
			Default:  1,
		},
		"followvsid": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"forcel4": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"forcel7": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
			Default:  true,
		},
		"httpreschedule": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"idletime": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
			Default:  660,
		},
		"index": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Computed: true,
		},
		"inputauthmode": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"intercept": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"interceptmode": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"interceptopts": &schema.Schema{
			Type:     schema.ValueType(schema.TypeList),
			Computed: true,
			Elem:     &schema.Schema{Type: schema.ValueType(schema.TypeString)},
		},
		"ipreputationblocking": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"istransparent": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Computed: true,
		},
		"jsondlimit": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
			Default:  10000,
		},
		"layer": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Computed: true,
		},
		"mastervs": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Computed: true,
		},
		"mastervsid": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Computed: true,
		},
		"matchlen": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"maxconnslimit": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"multiconnect": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"needhostname": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"nickname": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
		},
		"nmatchbodyrules": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"npreprocessrules": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"nrequestrules": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"nresponserules": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"nrules": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"numberofrss": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Computed: true,
		},
		"ocspverify": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"outputauthmode": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"owaspopts": &schema.Schema{
			Type:     schema.ValueType(schema.TypeList),
			Computed: true,
			Elem:     &schema.Schema{Type: schema.ValueType(schema.TypeString)},
		},
		"passcipher": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"passsni": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"pcrelimit": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
			Default:  10000,
		},
		"persisttimeout": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
			Default:  "0",
		},
		"port": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Required: true,
		},
		"protocol": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Required: true,
		},
		"qos": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"refreshpersist": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"requestsperseclimit": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"rsminimum": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"samesite": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"schedule": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
			Default:  "rr",
		},
		"securityheaderoptions": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"serverinit": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"sslreencrypt": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"sslreverse": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		/*"sslrewrite": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
			Default:  "0",
		},*/
		"starttlsmode": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"status": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Computed: true,
		},
		"subnetoriginating": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
			Default:  true,
		},
		/*"tlstype": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
		},*/
		"transactionlimit": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"transparent": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.ValueType(schema.TypeString),
			Optional: true,
			Default:  "http",
		},
		"useforsnat": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
		"verify": &schema.Schema{
			Type:     schema.ValueType(schema.TypeInt),
			Optional: true,
		},
		"verifybearer": &schema.Schema{
			Type:     schema.ValueType(schema.TypeBool),
			Optional: true,
		},
	}
}

func resourceVs() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVsCreate,
		ReadContext:   resourceVsRead,
		UpdateContext: resourceVsUpdate,
		DeleteContext: resourceVsDelete,
		Schema:        GetVsSchema(),
	}
}

func resourceVsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.Get("forcel4").(bool) && d.Get("forcel7").(bool) {
		return diag.FromErr(fmt.Errorf("forcel4 and forcel7 must not both be true"))
	}
	c := m.(*lmclient.Client)

	interceptoptsRaw := d.Get("interceptopts").([]interface{})
	interceptopts := make([]string, len(interceptoptsRaw))
	for io, raw := range interceptoptsRaw {
		interceptopts[io] = raw.(string)
	}

	owaspoptsRaw := d.Get("owaspopts").([]interface{})
	owaspopts := make([]string, len(owaspoptsRaw))
	for oo, raw := range owaspoptsRaw {
		owaspopts[oo] = raw.(string)
	}

	vs := &lmclient.Vs{
		Address:       d.Get("address").(string),
		Port:          d.Get("port").(string),
		NickName:      d.Get("nickname").(string),
		Enable:        d.Get("enable").(bool),
		SSLReverse:    d.Get("sslreverse").(bool),
		SSLReencrypt:  d.Get("sslreencrypt").(bool),
		InterceptMode: d.Get("interceptmode").(int),
		Intercept:     d.Get("intercept").(bool),
		//InterceptOpts:           interceptopts,
		AlertThreshold: d.Get("alertthreshold").(int),
		//OwaspOpts:               owaspopts,
		BlockingParanoia:        d.Get("blockingparanoia").(int),
		IPReputationBlocking:    d.Get("ipreputationblocking").(bool),
		ExecutingParanoia:       d.Get("executingparanoia").(int),
		AnomalyScoringThreshold: d.Get("anomalyscoringthreshold").(int),
		PCRELimit:               d.Get("pcrelimit").(int),
		JSONDLimit:              d.Get("jsondlimit").(int),
		BodyLimit:               d.Get("bodylimit").(int),
		Transactionlimit:        d.Get("transactionlimit").(int),
		Transparent:             d.Get("transparent").(bool),
		SubnetOriginating:       d.Get("subnetoriginating").(bool),
		ServerInit:              d.Get("serverinit").(int),
		StartTLSMode:            d.Get("starttlsmode").(int),
		Idletime:                d.Get("idletime").(int),
		Cache:                   d.Get("cache").(bool),
		Compress:                d.Get("compress").(bool),
		Verify:                  d.Get("verify").(int),
		UseforSnat:              d.Get("useforsnat").(bool),
		ForceL4:                 d.Get("forcel4").(bool),
		ForceL7:                 d.Get("forcel7").(bool),
		MultiConnect:            d.Get("multiconnect").(bool),
		ClientCert:              d.Get("clientcert").(int),
		SecurityHeaderOptions:   d.Get("securityheaderoptions").(int),
		SameSite:                d.Get("samesite").(int),
		VerifyBearer:            d.Get("verifybearer").(bool),
		ErrorCode:               d.Get("errorcode").(string),
		CheckUse11:              d.Get("checkuse11").(bool),
		MatchLen:                d.Get("matchlen").(int),
		CheckUseGet:             d.Get("checkuseget").(int),
		//SSLRewrite:              d.Get("sslrewrite").(string),
		Type:             d.Get("type").(string),
		FollowVSID:       d.Get("followvsid").(int),
		Protocol:         d.Get("protocol").(string),
		Schedule:         d.Get("schedule").(string),
		CheckType:        d.Get("checktype").(string),
		PersistTimeout:   d.Get("persisttimeout").(string),
		CheckPort:        d.Get("checkport").(string),
		HTTPReschedule:   d.Get("httpreschedule").(bool),
		NRules:           d.Get("nrules").(int),
		NRequestRules:    d.Get("nrequestrules").(int),
		NResponseRules:   d.Get("nresponserules").(int),
		NMatchBodyRules:  d.Get("nmatchbodyrules").(int),
		NPreProcessRules: d.Get("npreprocessrules").(int),
		EspEnabled:       d.Get("espenabled").(bool),
		InputAuthMode:    d.Get("inputauthmode").(int),
		OutputAuthMode:   d.Get("outputauthmode").(int),
		AddVia:           d.Get("addvia").(int),
		QoS:              d.Get("qos").(int),
		//TLSType:                 d.Get("tlstype").(string),
		NeedHostName:         d.Get("needhostname").(bool),
		OCSPVerify:           d.Get("ocspverify").(bool),
		AllowHTTP2:           d.Get("allowhttp2").(bool),
		PassCipher:           d.Get("passcipher").(bool),
		PassSni:              d.Get("passsni").(bool),
		ChkInterval:          d.Get("chkinterval").(int),
		ChkTimeout:           d.Get("chktimeout").(int),
		ChkRetryCount:        d.Get("chkretrycount").(int),
		Bandwidth:            d.Get("bandwidth").(int),
		ConnsPerSecLimit:     d.Get("connsperseclimit").(int),
		RequestsPerSecLimit:  d.Get("requestsperseclimit").(int),
		MaxConnsLimit:        d.Get("maxconnslimit").(int),
		RefreshPersist:       d.Get("refreshpersist").(bool),
		EnhancedHealthChecks: d.Get("enhancedhealthchecks").(bool),
		RsMinimum:            d.Get("rsminimum").(int),
		NumberOfRSs:          d.Get("numberofrss").(int),
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
	d.Set("address", vc.Address)
	d.Set("vsport", vc.VSPort)
	d.Set("nickname", vc.NickName)
	d.Set("enable", vc.Enable)
	d.Set("sslreverse", vc.SSLReverse)
	d.Set("sslreencrypt", vc.SSLReencrypt)
	d.Set("interceptmode", vc.InterceptMode)
	d.Set("intercept", vc.Intercept)
	d.Set("interceptopts", vc.InterceptOpts)
	d.Set("alertthreshold", vc.AlertThreshold)
	d.Set("owaspopts", vc.OwaspOpts)
	d.Set("blockingparanoia", vc.BlockingParanoia)
	d.Set("ipreputationblocking", vc.IPReputationBlocking)
	d.Set("executingparanoia", vc.ExecutingParanoia)
	d.Set("anomalyscoringthreshold", vc.AnomalyScoringThreshold)
	d.Set("pcrelimit", vc.PCRELimit)
	d.Set("jsondlimit", vc.JSONDLimit)
	d.Set("bodylimit", vc.BodyLimit)
	d.Set("transactionlimit", vc.Transactionlimit)
	d.Set("transparent", vc.Transparent)
	d.Set("subnetoriginating", vc.SubnetOriginating)
	d.Set("serverinit", vc.ServerInit)
	d.Set("starttlsmode", vc.StartTLSMode)
	d.Set("idletime", vc.Idletime)
	d.Set("cache", vc.Cache)
	d.Set("compress", vc.Compress)
	d.Set("verify", vc.Verify)
	d.Set("useforsnat", vc.UseforSnat)
	d.Set("forcel4", vc.ForceL4)
	d.Set("forcel7", vc.ForceL7)
	d.Set("multiconnect", vc.MultiConnect)
	d.Set("clientcert", vc.ClientCert)
	d.Set("securityheaderoptions", vc.SecurityHeaderOptions)
	d.Set("samesite", vc.SameSite)
	d.Set("verifybearer", vc.VerifyBearer)
	d.Set("errorcode", vc.ErrorCode)
	d.Set("checkuse11", vc.CheckUse11)
	d.Set("matchlen", vc.MatchLen)
	d.Set("checkuseget", vc.CheckUseGet)
	//d.Set("sslrewrite", vc.SSLRewrite)
	d.Set("type", vc.Type)
	d.Set("followvsid", vc.FollowVSID)
	d.Set("protocol", vc.Protocol)
	d.Set("schedule", vc.Schedule)
	d.Set("checktype", vc.CheckType)
	d.Set("persisttimeout", vc.PersistTimeout)
	d.Set("checkport", vc.CheckPort)
	d.Set("httpreschedule", vc.HTTPReschedule)
	d.Set("nrules", vc.NRules)
	d.Set("nrequestrules", vc.NRequestRules)
	d.Set("nresponserules", vc.NResponseRules)
	d.Set("nmatchbodyrules", vc.NMatchBodyRules)
	d.Set("npreprocessrules", vc.NPreProcessRules)
	d.Set("espenabled", vc.EspEnabled)
	d.Set("inputauthmode", vc.InputAuthMode)
	d.Set("outputauthmode", vc.OutputAuthMode)
	d.Set("mastervs", vc.MasterVS)
	d.Set("mastervsid", vc.MasterVSID)
	d.Set("istransparent", vc.IsTransparent)
	d.Set("addvia", vc.AddVia)
	d.Set("qos", vc.QoS)
	//d.Set("tlstype", vc.TLSType)
	d.Set("needhostname", vc.NeedHostName)
	d.Set("ocspverify", vc.OCSPVerify)
	d.Set("allowhttp2", vc.AllowHTTP2)
	d.Set("passcipher", vc.PassCipher)
	d.Set("passsni", vc.PassSni)
	d.Set("chkinterval", vc.ChkInterval)
	d.Set("chktimeout", vc.ChkTimeout)
	d.Set("chkretrycount", vc.ChkRetryCount)
	d.Set("bandwidth", vc.Bandwidth)
	d.Set("connsperseclimit", vc.ConnsPerSecLimit)
	d.Set("requestsperseclimit", vc.RequestsPerSecLimit)
	d.Set("maxconnslimit", vc.MaxConnsLimit)
	d.Set("refreshpersist", vc.RefreshPersist)
	d.Set("enhancedhealthchecks", vc.EnhancedHealthChecks)
	d.Set("rsminimum", vc.RsMinimum)
	d.Set("numberofrss", vc.NumberOfRSs)
	return diags
}
func resourceVsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if d.Get("forcel4").(bool) && d.Get("forcel7").(bool) {
		return diag.FromErr(fmt.Errorf("forcel4 and forcel7 must not both be true"))
	}
	c := m.(*lmclient.Client)
	i, _ := strconv.Atoi(d.Id())

	interceptoptsRaw := d.Get("interceptopts").([]interface{})
	interceptopts := make([]string, len(interceptoptsRaw))
	for io, raw := range interceptoptsRaw {
		interceptopts[io] = raw.(string)
	}

	owaspoptsRaw := d.Get("owaspopts").([]interface{})
	owaspopts := make([]string, len(owaspoptsRaw))
	for oo, raw := range owaspoptsRaw {
		owaspopts[oo] = raw.(string)
	}

	oldport, _ := d.GetChange("port")
	vs := &lmclient.Vs{
		Index:                   i,
		Port:                    oldport.(string),
		Address:                 d.Get("address").(string),
		VSPort:                  d.Get("port").(string),
		NickName:                d.Get("nickname").(string),
		Enable:                  d.Get("enable").(bool),
		SSLReverse:              d.Get("sslreverse").(bool),
		SSLReencrypt:            d.Get("sslreencrypt").(bool),
		InterceptMode:           d.Get("interceptmode").(int),
		Intercept:               d.Get("intercept").(bool),
		InterceptOpts:           interceptopts,
		AlertThreshold:          d.Get("alertthreshold").(int),
		OwaspOpts:               owaspopts,
		BlockingParanoia:        d.Get("blockingparanoia").(int),
		IPReputationBlocking:    d.Get("ipreputationblocking").(bool),
		ExecutingParanoia:       d.Get("executingparanoia").(int),
		AnomalyScoringThreshold: d.Get("anomalyscoringthreshold").(int),
		PCRELimit:               d.Get("pcrelimit").(int),
		JSONDLimit:              d.Get("jsondlimit").(int),
		BodyLimit:               d.Get("bodylimit").(int),
		Transactionlimit:        d.Get("transactionlimit").(int),
		Transparent:             d.Get("transparent").(bool),
		SubnetOriginating:       d.Get("subnetoriginating").(bool),
		ServerInit:              d.Get("serverinit").(int),
		StartTLSMode:            d.Get("starttlsmode").(int),
		Idletime:                d.Get("idletime").(int),
		Cache:                   d.Get("cache").(bool),
		Compress:                d.Get("compress").(bool),
		Verify:                  d.Get("verify").(int),
		UseforSnat:              d.Get("useforsnat").(bool),
		ForceL4:                 d.Get("forcel4").(bool),
		ForceL7:                 d.Get("forcel7").(bool),
		MultiConnect:            d.Get("multiconnect").(bool),
		ClientCert:              d.Get("clientcert").(int),
		SecurityHeaderOptions:   d.Get("securityheaderoptions").(int),
		SameSite:                d.Get("samesite").(int),
		VerifyBearer:            d.Get("verifybearer").(bool),
		ErrorCode:               d.Get("errorcode").(string),
		CheckUse11:              d.Get("checkuse11").(bool),
		MatchLen:                d.Get("matchlen").(int),
		CheckUseGet:             d.Get("checkuseget").(int),
		//SSLRewrite:              d.Get("sslrewrite").(string),
		Type:             d.Get("type").(string),
		FollowVSID:       d.Get("followvsid").(int),
		Protocol:         d.Get("protocol").(string),
		Schedule:         d.Get("schedule").(string),
		CheckType:        d.Get("checktype").(string),
		PersistTimeout:   d.Get("persisttimeout").(string),
		CheckPort:        d.Get("checkport").(string),
		HTTPReschedule:   d.Get("httpreschedule").(bool),
		NRules:           d.Get("nrules").(int),
		NRequestRules:    d.Get("nrequestrules").(int),
		NResponseRules:   d.Get("nresponserules").(int),
		NMatchBodyRules:  d.Get("nmatchbodyrules").(int),
		NPreProcessRules: d.Get("npreprocessrules").(int),
		EspEnabled:       d.Get("espenabled").(bool),
		InputAuthMode:    d.Get("inputauthmode").(int),
		OutputAuthMode:   d.Get("outputauthmode").(int),
		AddVia:           d.Get("addvia").(int),
		QoS:              d.Get("qos").(int),
		//TLSType:                 d.Get("tlstype").(string),
		NeedHostName:         d.Get("needhostname").(bool),
		OCSPVerify:           d.Get("ocspverify").(bool),
		AllowHTTP2:           d.Get("allowhttp2").(bool),
		PassCipher:           d.Get("passcipher").(bool),
		PassSni:              d.Get("passsni").(bool),
		ChkInterval:          d.Get("chkinterval").(int),
		ChkTimeout:           d.Get("chktimeout").(int),
		ChkRetryCount:        d.Get("chkretrycount").(int),
		Bandwidth:            d.Get("bandwidth").(int),
		ConnsPerSecLimit:     d.Get("connsperseclimit").(int),
		RequestsPerSecLimit:  d.Get("requestsperseclimit").(int),
		MaxConnsLimit:        d.Get("maxconnslimit").(int),
		RefreshPersist:       d.Get("refreshpersist").(bool),
		EnhancedHealthChecks: d.Get("enhancedhealthchecks").(bool),
		RsMinimum:            d.Get("rsminimum").(int),
		NumberOfRSs:          d.Get("numberofrss").(int),
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
