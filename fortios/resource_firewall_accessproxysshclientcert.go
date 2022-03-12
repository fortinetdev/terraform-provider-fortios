// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Access Proxy SSH client certificate.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallAccessProxySshClientCert() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxySshClientCertCreate,
		Read:   resourceFirewallAccessProxySshClientCertRead,
		Update: resourceFirewallAccessProxySshClientCertUpdate,
		Delete: resourceFirewallAccessProxySshClientCertDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_x11_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_agent_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_port_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_pty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_user_rc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"critical": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"data": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"auth_ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallAccessProxySshClientCertCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallAccessProxySshClientCert(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxySshClientCert resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallAccessProxySshClientCert(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxySshClientCert resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxySshClientCert")
	}

	return resourceFirewallAccessProxySshClientCertRead(d, m)
}

func resourceFirewallAccessProxySshClientCertUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallAccessProxySshClientCert(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxySshClientCert resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallAccessProxySshClientCert(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxySshClientCert resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxySshClientCert")
	}

	return resourceFirewallAccessProxySshClientCertRead(d, m)
}

func resourceFirewallAccessProxySshClientCertDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallAccessProxySshClientCert(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxySshClientCert resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxySshClientCertRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallAccessProxySshClientCert(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxySshClientCert resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxySshClientCert(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxySshClientCert resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxySshClientCertName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertSourceAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitX11Forwarding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitAgentForwarding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitPortForwarding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitPty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitUserRc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertCertExtension(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenFirewallAccessProxySshClientCertCertExtensionName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := i["critical"]; ok {

			tmp["critical"] = flattenFirewallAccessProxySshClientCertCertExtensionCritical(i["critical"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenFirewallAccessProxySshClientCertCertExtensionType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := i["data"]; ok {

			tmp["data"] = flattenFirewallAccessProxySshClientCertCertExtensionData(i["data"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallAccessProxySshClientCertCertExtensionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertCertExtensionCritical(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertCertExtensionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertCertExtensionData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertAuthCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxySshClientCert(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallAccessProxySshClientCertName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("source_address", flattenFirewallAccessProxySshClientCertSourceAddress(o["source-address"], d, "source_address", sv)); err != nil {
		if !fortiAPIPatch(o["source-address"]) {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	if err = d.Set("permit_x11_forwarding", flattenFirewallAccessProxySshClientCertPermitX11Forwarding(o["permit-x11-forwarding"], d, "permit_x11_forwarding", sv)); err != nil {
		if !fortiAPIPatch(o["permit-x11-forwarding"]) {
			return fmt.Errorf("Error reading permit_x11_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_agent_forwarding", flattenFirewallAccessProxySshClientCertPermitAgentForwarding(o["permit-agent-forwarding"], d, "permit_agent_forwarding", sv)); err != nil {
		if !fortiAPIPatch(o["permit-agent-forwarding"]) {
			return fmt.Errorf("Error reading permit_agent_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_port_forwarding", flattenFirewallAccessProxySshClientCertPermitPortForwarding(o["permit-port-forwarding"], d, "permit_port_forwarding", sv)); err != nil {
		if !fortiAPIPatch(o["permit-port-forwarding"]) {
			return fmt.Errorf("Error reading permit_port_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_pty", flattenFirewallAccessProxySshClientCertPermitPty(o["permit-pty"], d, "permit_pty", sv)); err != nil {
		if !fortiAPIPatch(o["permit-pty"]) {
			return fmt.Errorf("Error reading permit_pty: %v", err)
		}
	}

	if err = d.Set("permit_user_rc", flattenFirewallAccessProxySshClientCertPermitUserRc(o["permit-user-rc"], d, "permit_user_rc", sv)); err != nil {
		if !fortiAPIPatch(o["permit-user-rc"]) {
			return fmt.Errorf("Error reading permit_user_rc: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cert_extension", flattenFirewallAccessProxySshClientCertCertExtension(o["cert-extension"], d, "cert_extension", sv)); err != nil {
			if !fortiAPIPatch(o["cert-extension"]) {
				return fmt.Errorf("Error reading cert_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cert_extension"); ok {
			if err = d.Set("cert_extension", flattenFirewallAccessProxySshClientCertCertExtension(o["cert-extension"], d, "cert_extension", sv)); err != nil {
				if !fortiAPIPatch(o["cert-extension"]) {
					return fmt.Errorf("Error reading cert_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_ca", flattenFirewallAccessProxySshClientCertAuthCa(o["auth-ca"], d, "auth_ca", sv)); err != nil {
		if !fortiAPIPatch(o["auth-ca"]) {
			return fmt.Errorf("Error reading auth_ca: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxySshClientCertFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallAccessProxySshClientCertName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitX11Forwarding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitAgentForwarding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitPortForwarding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitPty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitUserRc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertCertExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandFirewallAccessProxySshClientCertCertExtensionName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["critical"], _ = expandFirewallAccessProxySshClientCertCertExtensionCritical(d, i["critical"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandFirewallAccessProxySshClientCertCertExtensionType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["data"], _ = expandFirewallAccessProxySshClientCertCertExtensionData(d, i["data"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxySshClientCertCertExtensionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertCertExtensionCritical(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertCertExtensionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertCertExtensionData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertAuthCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxySshClientCert(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallAccessProxySshClientCertName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok {

		t, err := expandFirewallAccessProxySshClientCertSourceAddress(d, v, "source_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	if v, ok := d.GetOk("permit_x11_forwarding"); ok {

		t, err := expandFirewallAccessProxySshClientCertPermitX11Forwarding(d, v, "permit_x11_forwarding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-x11-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_agent_forwarding"); ok {

		t, err := expandFirewallAccessProxySshClientCertPermitAgentForwarding(d, v, "permit_agent_forwarding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-agent-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_port_forwarding"); ok {

		t, err := expandFirewallAccessProxySshClientCertPermitPortForwarding(d, v, "permit_port_forwarding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-port-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_pty"); ok {

		t, err := expandFirewallAccessProxySshClientCertPermitPty(d, v, "permit_pty", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-pty"] = t
		}
	}

	if v, ok := d.GetOk("permit_user_rc"); ok {

		t, err := expandFirewallAccessProxySshClientCertPermitUserRc(d, v, "permit_user_rc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-user-rc"] = t
		}
	}

	if v, ok := d.GetOk("cert_extension"); ok {

		t, err := expandFirewallAccessProxySshClientCertCertExtension(d, v, "cert_extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-extension"] = t
		}
	}

	if v, ok := d.GetOk("auth_ca"); ok {

		t, err := expandFirewallAccessProxySshClientCertAuthCa(d, v, "auth_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ca"] = t
		}
	}

	return &obj, nil
}
