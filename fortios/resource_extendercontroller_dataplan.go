// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: FortiExtender dataplan configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtenderControllerDataplan() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtenderControllerDataplanCreate,
		Read:   resourceExtenderControllerDataplanRead,
		Update: resourceExtenderControllerDataplanUpdate,
		Delete: resourceExtenderControllerDataplanDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"modem_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"slot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iccid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
			},
			"carrier": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
			},
			"apn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Sensitive:    true,
			},
			"pdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signal_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(50, 100),
				Optional:     true,
				Computed:     true,
			},
			"signal_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(600, 18000),
				Optional:     true,
				Computed:     true,
			},
			"capacity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 102400000),
				Optional:     true,
			},
			"monthly_fee": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000000),
				Optional:     true,
			},
			"billing_date": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),
				Optional:     true,
				Computed:     true,
			},
			"overage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferred_subnet": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"private_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExtenderControllerDataplanCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtenderControllerDataplan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerDataplan resource while getting object: %v", err)
	}

	o, err := c.CreateExtenderControllerDataplan(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerDataplan resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerDataplan")
	}

	return resourceExtenderControllerDataplanRead(d, m)
}

func resourceExtenderControllerDataplanUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtenderControllerDataplan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerDataplan resource while getting object: %v", err)
	}

	o, err := c.UpdateExtenderControllerDataplan(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerDataplan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerDataplan")
	}

	return resourceExtenderControllerDataplanRead(d, m)
}

func resourceExtenderControllerDataplanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtenderControllerDataplan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtenderControllerDataplan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerDataplanRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtenderControllerDataplan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerDataplan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtenderControllerDataplan(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerDataplan resource from API: %v", err)
	}
	return nil
}

func flattenExtenderControllerDataplanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanModemId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanSlot(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanIccid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanApn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanPdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanSignalThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtenderControllerDataplanSignalPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtenderControllerDataplanCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtenderControllerDataplanMonthlyFee(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtenderControllerDataplanBillingDate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtenderControllerDataplanOverage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerDataplanPreferredSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtenderControllerDataplanPrivateNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtenderControllerDataplan(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenExtenderControllerDataplanName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenExtenderControllerDataplanModemId(o["modem-id"], d, "modem_id", sv)); err != nil {
		if !fortiAPIPatch(o["modem-id"]) {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("type", flattenExtenderControllerDataplanType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("slot", flattenExtenderControllerDataplanSlot(o["slot"], d, "slot", sv)); err != nil {
		if !fortiAPIPatch(o["slot"]) {
			return fmt.Errorf("Error reading slot: %v", err)
		}
	}

	if err = d.Set("iccid", flattenExtenderControllerDataplanIccid(o["iccid"], d, "iccid", sv)); err != nil {
		if !fortiAPIPatch(o["iccid"]) {
			return fmt.Errorf("Error reading iccid: %v", err)
		}
	}

	if err = d.Set("carrier", flattenExtenderControllerDataplanCarrier(o["carrier"], d, "carrier", sv)); err != nil {
		if !fortiAPIPatch(o["carrier"]) {
			return fmt.Errorf("Error reading carrier: %v", err)
		}
	}

	if err = d.Set("apn", flattenExtenderControllerDataplanApn(o["APN"], d, "apn", sv)); err != nil {
		if !fortiAPIPatch(o["APN"]) {
			return fmt.Errorf("Error reading apn: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenExtenderControllerDataplanAuthType(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("username", flattenExtenderControllerDataplanUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("pdn", flattenExtenderControllerDataplanPdn(o["PDN"], d, "pdn", sv)); err != nil {
		if !fortiAPIPatch(o["PDN"]) {
			return fmt.Errorf("Error reading pdn: %v", err)
		}
	}

	if err = d.Set("signal_threshold", flattenExtenderControllerDataplanSignalThreshold(o["signal-threshold"], d, "signal_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["signal-threshold"]) {
			return fmt.Errorf("Error reading signal_threshold: %v", err)
		}
	}

	if err = d.Set("signal_period", flattenExtenderControllerDataplanSignalPeriod(o["signal-period"], d, "signal_period", sv)); err != nil {
		if !fortiAPIPatch(o["signal-period"]) {
			return fmt.Errorf("Error reading signal_period: %v", err)
		}
	}

	if err = d.Set("capacity", flattenExtenderControllerDataplanCapacity(o["capacity"], d, "capacity", sv)); err != nil {
		if !fortiAPIPatch(o["capacity"]) {
			return fmt.Errorf("Error reading capacity: %v", err)
		}
	}

	if err = d.Set("monthly_fee", flattenExtenderControllerDataplanMonthlyFee(o["monthly-fee"], d, "monthly_fee", sv)); err != nil {
		if !fortiAPIPatch(o["monthly-fee"]) {
			return fmt.Errorf("Error reading monthly_fee: %v", err)
		}
	}

	if err = d.Set("billing_date", flattenExtenderControllerDataplanBillingDate(o["billing-date"], d, "billing_date", sv)); err != nil {
		if !fortiAPIPatch(o["billing-date"]) {
			return fmt.Errorf("Error reading billing_date: %v", err)
		}
	}

	if err = d.Set("overage", flattenExtenderControllerDataplanOverage(o["overage"], d, "overage", sv)); err != nil {
		if !fortiAPIPatch(o["overage"]) {
			return fmt.Errorf("Error reading overage: %v", err)
		}
	}

	if err = d.Set("preferred_subnet", flattenExtenderControllerDataplanPreferredSubnet(o["preferred-subnet"], d, "preferred_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["preferred-subnet"]) {
			return fmt.Errorf("Error reading preferred_subnet: %v", err)
		}
	}

	if err = d.Set("private_network", flattenExtenderControllerDataplanPrivateNetwork(o["private-network"], d, "private_network", sv)); err != nil {
		if !fortiAPIPatch(o["private-network"]) {
			return fmt.Errorf("Error reading private_network: %v", err)
		}
	}

	return nil
}

func flattenExtenderControllerDataplanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtenderControllerDataplanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanModemId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanSlot(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanIccid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanApn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanPdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanSignalThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanSignalPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanMonthlyFee(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanBillingDate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanOverage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanPreferredSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerDataplanPrivateNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtenderControllerDataplan(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandExtenderControllerDataplanName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok {
		t, err := expandExtenderControllerDataplanModemId(d, v, "modem_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandExtenderControllerDataplanType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("slot"); ok {
		t, err := expandExtenderControllerDataplanSlot(d, v, "slot", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slot"] = t
		}
	} else if d.HasChange("slot") {
		obj["slot"] = nil
	}

	if v, ok := d.GetOk("iccid"); ok {
		t, err := expandExtenderControllerDataplanIccid(d, v, "iccid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iccid"] = t
		}
	} else if d.HasChange("iccid") {
		obj["iccid"] = nil
	}

	if v, ok := d.GetOk("carrier"); ok {
		t, err := expandExtenderControllerDataplanCarrier(d, v, "carrier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier"] = t
		}
	} else if d.HasChange("carrier") {
		obj["carrier"] = nil
	}

	if v, ok := d.GetOk("apn"); ok {
		t, err := expandExtenderControllerDataplanApn(d, v, "apn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["APN"] = t
		}
	} else if d.HasChange("apn") {
		obj["APN"] = nil
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandExtenderControllerDataplanAuthType(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandExtenderControllerDataplanUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	} else if d.HasChange("username") {
		obj["username"] = nil
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandExtenderControllerDataplanPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	} else if d.HasChange("password") {
		obj["password"] = nil
	}

	if v, ok := d.GetOk("pdn"); ok {
		t, err := expandExtenderControllerDataplanPdn(d, v, "pdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["PDN"] = t
		}
	}

	if v, ok := d.GetOkExists("signal_threshold"); ok {
		t, err := expandExtenderControllerDataplanSignalThreshold(d, v, "signal_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-threshold"] = t
		}
	}

	if v, ok := d.GetOk("signal_period"); ok {
		t, err := expandExtenderControllerDataplanSignalPeriod(d, v, "signal_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-period"] = t
		}
	}

	if v, ok := d.GetOkExists("capacity"); ok {
		t, err := expandExtenderControllerDataplanCapacity(d, v, "capacity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capacity"] = t
		}
	} else if d.HasChange("capacity") {
		obj["capacity"] = nil
	}

	if v, ok := d.GetOkExists("monthly_fee"); ok {
		t, err := expandExtenderControllerDataplanMonthlyFee(d, v, "monthly_fee", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monthly-fee"] = t
		}
	} else if d.HasChange("monthly_fee") {
		obj["monthly-fee"] = nil
	}

	if v, ok := d.GetOk("billing_date"); ok {
		t, err := expandExtenderControllerDataplanBillingDate(d, v, "billing_date", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-date"] = t
		}
	}

	if v, ok := d.GetOk("overage"); ok {
		t, err := expandExtenderControllerDataplanOverage(d, v, "overage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage"] = t
		}
	}

	if v, ok := d.GetOkExists("preferred_subnet"); ok {
		t, err := expandExtenderControllerDataplanPreferredSubnet(d, v, "preferred_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-subnet"] = t
		}
	}

	if v, ok := d.GetOk("private_network"); ok {
		t, err := expandExtenderControllerDataplanPrivateNetwork(d, v, "private_network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-network"] = t
		}
	}

	return &obj, nil
}
