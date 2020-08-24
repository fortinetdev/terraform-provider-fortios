// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure API users.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemApiUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemApiUserCreate,
		Read:   resourceSystemApiUserRead,
		Update: resourceSystemApiUserUpdate,
		Delete: resourceSystemApiUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"api_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"cors_allow_origin": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 269),
				Optional:     true,
				Computed:     true,
			},
			"peer_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"trusthost": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_trusthost": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_trusthost": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemApiUserCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemApiUser(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemApiUser resource while getting object: %v", err)
	}

	o, err := c.CreateSystemApiUser(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemApiUser resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemApiUser")
	}

	return resourceSystemApiUserRead(d, m)
}

func resourceSystemApiUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemApiUser(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemApiUser resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemApiUser(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemApiUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemApiUser")
	}

	return resourceSystemApiUserRead(d, m)
}

func resourceSystemApiUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemApiUser(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemApiUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemApiUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemApiUser(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemApiUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemApiUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemApiUser resource from API: %v", err)
	}
	return nil
}

func flattenSystemApiUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserApiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserVdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemApiUserVdomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemApiUserVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserCorsAllowOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserPeerAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserPeerGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserTrusthost(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenSystemApiUserTrusthostId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenSystemApiUserTrusthostType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_trusthost"
		if _, ok := i["ipv4-trusthost"]; ok {
			tmp["ipv4_trusthost"] = flattenSystemApiUserTrusthostIpv4Trusthost(i["ipv4-trusthost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_trusthost"
		if _, ok := i["ipv6-trusthost"]; ok {
			tmp["ipv6_trusthost"] = flattenSystemApiUserTrusthostIpv6Trusthost(i["ipv6-trusthost"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemApiUserTrusthostId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserTrusthostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserTrusthostIpv4Trusthost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemApiUserTrusthostIpv6Trusthost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemApiUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemApiUserName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemApiUserComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("api_key", flattenSystemApiUserApiKey(o["api-key"], d, "api_key")); err != nil {
		if !fortiAPIPatch(o["api-key"]) {
			return fmt.Errorf("Error reading api_key: %v", err)
		}
	}

	if err = d.Set("accprofile", flattenSystemApiUserAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vdom", flattenSystemApiUserVdom(o["vdom"], d, "vdom")); err != nil {
			if !fortiAPIPatch(o["vdom"]) {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdom"); ok {
			if err = d.Set("vdom", flattenSystemApiUserVdom(o["vdom"], d, "vdom")); err != nil {
				if !fortiAPIPatch(o["vdom"]) {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			}
		}
	}

	if err = d.Set("schedule", flattenSystemApiUserSchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("cors_allow_origin", flattenSystemApiUserCorsAllowOrigin(o["cors-allow-origin"], d, "cors_allow_origin")); err != nil {
		if !fortiAPIPatch(o["cors-allow-origin"]) {
			return fmt.Errorf("Error reading cors_allow_origin: %v", err)
		}
	}

	if err = d.Set("peer_auth", flattenSystemApiUserPeerAuth(o["peer-auth"], d, "peer_auth")); err != nil {
		if !fortiAPIPatch(o["peer-auth"]) {
			return fmt.Errorf("Error reading peer_auth: %v", err)
		}
	}

	if err = d.Set("peer_group", flattenSystemApiUserPeerGroup(o["peer-group"], d, "peer_group")); err != nil {
		if !fortiAPIPatch(o["peer-group"]) {
			return fmt.Errorf("Error reading peer_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("trusthost", flattenSystemApiUserTrusthost(o["trusthost"], d, "trusthost")); err != nil {
			if !fortiAPIPatch(o["trusthost"]) {
				return fmt.Errorf("Error reading trusthost: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("trusthost"); ok {
			if err = d.Set("trusthost", flattenSystemApiUserTrusthost(o["trusthost"], d, "trusthost")); err != nil {
				if !fortiAPIPatch(o["trusthost"]) {
					return fmt.Errorf("Error reading trusthost: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemApiUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemApiUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemApiUserVdomName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemApiUserVdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserCorsAllowOrigin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserPeerAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserPeerGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemApiUserTrusthostId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemApiUserTrusthostType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_trusthost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipv4-trusthost"], _ = expandSystemApiUserTrusthostIpv4Trusthost(d, i["ipv4_trusthost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_trusthost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipv6-trusthost"], _ = expandSystemApiUserTrusthostIpv6Trusthost(d, i["ipv6_trusthost"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemApiUserTrusthostId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthostIpv4Trusthost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthostIpv6Trusthost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemApiUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemApiUserName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSystemApiUserComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("api_key"); ok {
		t, err := expandSystemApiUserApiKey(d, v, "api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-key"] = t
		}
	}

	if v, ok := d.GetOk("accprofile"); ok {
		t, err := expandSystemApiUserAccprofile(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemApiUserVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandSystemApiUserSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("cors_allow_origin"); ok {
		t, err := expandSystemApiUserCorsAllowOrigin(d, v, "cors_allow_origin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-allow-origin"] = t
		}
	}

	if v, ok := d.GetOk("peer_auth"); ok {
		t, err := expandSystemApiUserPeerAuth(d, v, "peer_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-auth"] = t
		}
	}

	if v, ok := d.GetOk("peer_group"); ok {
		t, err := expandSystemApiUserPeerGroup(d, v, "peer_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-group"] = t
		}
	}

	if v, ok := d.GetOk("trusthost"); ok {
		t, err := expandSystemApiUserTrusthost(d, v, "trusthost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost"] = t
		}
	}

	return &obj, nil
}
