// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure MPSK profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerMpskProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerMpskProfileCreate,
		Read:   resourceWirelessControllerMpskProfileRead,
		Update: resourceWirelessControllerMpskProfileUpdate,
		Delete: resourceWirelessControllerMpskProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"mpsk_concurrent_clients": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"mpsk_external_server_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk_external_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"mpsk_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"vlan_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
						},
						"mpsk_key": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"key_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"passphrase": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
									"sae_password": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
									"sae_pk": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sae_private_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 359),
										Optional:     true,
									},
									"concurrent_client_limit_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"concurrent_clients": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"comment": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
									"mpsk_schedules": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 35),
													Optional:     true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerMpskProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectWirelessControllerMpskProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerMpskProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerMpskProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerMpskProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerMpskProfile")
	}

	return resourceWirelessControllerMpskProfileRead(d, m)
}

func resourceWirelessControllerMpskProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectWirelessControllerMpskProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerMpskProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerMpskProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerMpskProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerMpskProfile")
	}

	return resourceWirelessControllerMpskProfileRead(d, m)
}

func resourceWirelessControllerMpskProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerMpskProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerMpskProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerMpskProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadWirelessControllerMpskProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerMpskProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerMpskProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerMpskProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerMpskProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerMpskProfileMpskExternalServerAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskExternalServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerMpskProfileMpskGroupName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_type"
		if cur_v, ok := i["vlan-type"]; ok {
			tmp["vlan_type"] = flattenWirelessControllerMpskProfileMpskGroupVlanType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if cur_v, ok := i["vlan-id"]; ok {
			tmp["vlan_id"] = flattenWirelessControllerMpskProfileMpskGroupVlanId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_key"
		if cur_v, ok := i["mpsk-key"]; ok {
			tmp["mpsk_key"] = flattenWirelessControllerMpskProfileMpskGroupMpskKey(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerMpskProfileMpskGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupVlanType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupVlanId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKey(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if cur_v, ok := i["key-type"]; ok {
			tmp["key_type"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyKeyType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if cur_v, ok := i["mac"]; ok {
			tmp["mac"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyMac(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := i["passphrase"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["passphrase"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_password"
		if _, ok := i["sae-password"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["sae_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if cur_v, ok := i["sae-pk"]; ok {
			tmp["sae_pk"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePk(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if cur_v, ok := i["sae-private-key"]; ok {
			tmp["sae_private_key"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if cur_v, ok := i["concurrent-client-limit-type"]; ok {
			tmp["concurrent_client_limit_type"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if cur_v, ok := i["concurrent-clients"]; ok {
			tmp["concurrent_clients"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if cur_v, ok := i["comment"]; ok {
			tmp["comment"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyComment(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if cur_v, ok := i["mpsk-schedules"]; ok {
			tmp["mpsk_schedules"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyKeyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedulesName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedulesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerMpskProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWirelessControllerMpskProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("mpsk_concurrent_clients", flattenWirelessControllerMpskProfileMpskConcurrentClients(o["mpsk-concurrent-clients"], d, "mpsk_concurrent_clients", sv)); err != nil {
		if !fortiAPIPatch(o["mpsk-concurrent-clients"]) {
			return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if err = d.Set("mpsk_external_server_auth", flattenWirelessControllerMpskProfileMpskExternalServerAuth(o["mpsk-external-server-auth"], d, "mpsk_external_server_auth", sv)); err != nil {
		if !fortiAPIPatch(o["mpsk-external-server-auth"]) {
			return fmt.Errorf("Error reading mpsk_external_server_auth: %v", err)
		}
	}

	if err = d.Set("mpsk_external_server", flattenWirelessControllerMpskProfileMpskExternalServer(o["mpsk-external-server"], d, "mpsk_external_server", sv)); err != nil {
		if !fortiAPIPatch(o["mpsk-external-server"]) {
			return fmt.Errorf("Error reading mpsk_external_server: %v", err)
		}
	}

	if err = d.Set("mpsk_type", flattenWirelessControllerMpskProfileMpskType(o["mpsk-type"], d, "mpsk_type", sv)); err != nil {
		if !fortiAPIPatch(o["mpsk-type"]) {
			return fmt.Errorf("Error reading mpsk_type: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("mpsk_group", flattenWirelessControllerMpskProfileMpskGroup(o["mpsk-group"], d, "mpsk_group", sv)); err != nil {
			if !fortiAPIPatch(o["mpsk-group"]) {
				return fmt.Errorf("Error reading mpsk_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_group"); ok {
			if err = d.Set("mpsk_group", flattenWirelessControllerMpskProfileMpskGroup(o["mpsk-group"], d, "mpsk_group", sv)); err != nil {
				if !fortiAPIPatch(o["mpsk-group"]) {
					return fmt.Errorf("Error reading mpsk_group: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerMpskProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerMpskProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskExternalServerAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskExternalServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWirelessControllerMpskProfileMpskGroupName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan-type"], _ = expandWirelessControllerMpskProfileMpskGroupVlanType(d, i["vlan_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan-id"], _ = expandWirelessControllerMpskProfileMpskGroupVlanId(d, i["vlan_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vlan-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mpsk-key"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKey(d, i["mpsk_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mpsk-key"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerMpskProfileMpskGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupVlanType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupVlanId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-type"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyKeyType(d, i["key_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["passphrase"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase(d, i["passphrase"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["passphrase"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sae-password"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePassword(d, i["sae_password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sae-password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sae-pk"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePk(d, i["sae_pk"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sae-private-key"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey(d, i["sae_private_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sae-private-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["concurrent-client-limit-type"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(d, i["concurrent_client_limit_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["concurrent-clients"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(d, i["concurrent_clients"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyComment(d, i["comment"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comment"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mpsk-schedules"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(d, i["mpsk_schedules"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mpsk-schedules"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyKeyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedulesName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedulesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerMpskProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerMpskProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("mpsk_concurrent_clients"); ok {
		t, err := expandWirelessControllerMpskProfileMpskConcurrentClients(d, v, "mpsk_concurrent_clients", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-concurrent-clients"] = t
		}
	} else if d.HasChange("mpsk_concurrent_clients") {
		obj["mpsk-concurrent-clients"] = nil
	}

	if v, ok := d.GetOk("mpsk_external_server_auth"); ok {
		t, err := expandWirelessControllerMpskProfileMpskExternalServerAuth(d, v, "mpsk_external_server_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-external-server-auth"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_external_server"); ok {
		t, err := expandWirelessControllerMpskProfileMpskExternalServer(d, v, "mpsk_external_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-external-server"] = t
		}
	} else if d.HasChange("mpsk_external_server") {
		obj["mpsk-external-server"] = nil
	}

	if v, ok := d.GetOk("mpsk_type"); ok {
		t, err := expandWirelessControllerMpskProfileMpskType(d, v, "mpsk_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-type"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_group"); ok || d.HasChange("mpsk_group") {
		t, err := expandWirelessControllerMpskProfileMpskGroup(d, v, "mpsk_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-group"] = t
		}
	}

	return &obj, nil
}
