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

func resourceExtensionControllerDataplan() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerDataplanCreate,
		Read:   resourceExtensionControllerDataplanRead,
		Update: resourceExtensionControllerDataplanUpdate,
		Delete: resourceExtensionControllerDataplanDelete,

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
				Computed: true,
			},
			"iccid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"carrier": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"apn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
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
				Computed:     true,
			},
			"monthly_fee": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000000),
				Optional:     true,
				Computed:     true,
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

func resourceExtensionControllerDataplanCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerDataplan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerDataplan resource while getting object: %v", err)
	}

	o, err := c.CreateExtensionControllerDataplan(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerDataplan resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerDataplan")
	}

	return resourceExtensionControllerDataplanRead(d, m)
}

func resourceExtensionControllerDataplanUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerDataplan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerDataplan resource while getting object: %v", err)
	}

	o, err := c.UpdateExtensionControllerDataplan(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerDataplan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerDataplan")
	}

	return resourceExtensionControllerDataplanRead(d, m)
}

func resourceExtensionControllerDataplanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtensionControllerDataplan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerDataplan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerDataplanRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerDataplan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerDataplan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerDataplan(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerDataplan resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerDataplanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanModemId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanSlot(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanIccid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanApn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanPdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanSignalThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanSignalPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanMonthlyFee(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanBillingDate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanOverage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanPreferredSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerDataplanPrivateNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtensionControllerDataplan(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenExtensionControllerDataplanName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenExtensionControllerDataplanModemId(o["modem-id"], d, "modem_id", sv)); err != nil {
		if !fortiAPIPatch(o["modem-id"]) {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("type", flattenExtensionControllerDataplanType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("slot", flattenExtensionControllerDataplanSlot(o["slot"], d, "slot", sv)); err != nil {
		if !fortiAPIPatch(o["slot"]) {
			return fmt.Errorf("Error reading slot: %v", err)
		}
	}

	if err = d.Set("iccid", flattenExtensionControllerDataplanIccid(o["iccid"], d, "iccid", sv)); err != nil {
		if !fortiAPIPatch(o["iccid"]) {
			return fmt.Errorf("Error reading iccid: %v", err)
		}
	}

	if err = d.Set("carrier", flattenExtensionControllerDataplanCarrier(o["carrier"], d, "carrier", sv)); err != nil {
		if !fortiAPIPatch(o["carrier"]) {
			return fmt.Errorf("Error reading carrier: %v", err)
		}
	}

	if err = d.Set("apn", flattenExtensionControllerDataplanApn(o["apn"], d, "apn", sv)); err != nil {
		if !fortiAPIPatch(o["apn"]) {
			return fmt.Errorf("Error reading apn: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenExtensionControllerDataplanAuthType(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("username", flattenExtensionControllerDataplanUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("password", flattenExtensionControllerDataplanPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("pdn", flattenExtensionControllerDataplanPdn(o["pdn"], d, "pdn", sv)); err != nil {
		if !fortiAPIPatch(o["pdn"]) {
			return fmt.Errorf("Error reading pdn: %v", err)
		}
	}

	if err = d.Set("signal_threshold", flattenExtensionControllerDataplanSignalThreshold(o["signal-threshold"], d, "signal_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["signal-threshold"]) {
			return fmt.Errorf("Error reading signal_threshold: %v", err)
		}
	}

	if err = d.Set("signal_period", flattenExtensionControllerDataplanSignalPeriod(o["signal-period"], d, "signal_period", sv)); err != nil {
		if !fortiAPIPatch(o["signal-period"]) {
			return fmt.Errorf("Error reading signal_period: %v", err)
		}
	}

	if err = d.Set("capacity", flattenExtensionControllerDataplanCapacity(o["capacity"], d, "capacity", sv)); err != nil {
		if !fortiAPIPatch(o["capacity"]) {
			return fmt.Errorf("Error reading capacity: %v", err)
		}
	}

	if err = d.Set("monthly_fee", flattenExtensionControllerDataplanMonthlyFee(o["monthly-fee"], d, "monthly_fee", sv)); err != nil {
		if !fortiAPIPatch(o["monthly-fee"]) {
			return fmt.Errorf("Error reading monthly_fee: %v", err)
		}
	}

	if err = d.Set("billing_date", flattenExtensionControllerDataplanBillingDate(o["billing-date"], d, "billing_date", sv)); err != nil {
		if !fortiAPIPatch(o["billing-date"]) {
			return fmt.Errorf("Error reading billing_date: %v", err)
		}
	}

	if err = d.Set("overage", flattenExtensionControllerDataplanOverage(o["overage"], d, "overage", sv)); err != nil {
		if !fortiAPIPatch(o["overage"]) {
			return fmt.Errorf("Error reading overage: %v", err)
		}
	}

	if err = d.Set("preferred_subnet", flattenExtensionControllerDataplanPreferredSubnet(o["preferred-subnet"], d, "preferred_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["preferred-subnet"]) {
			return fmt.Errorf("Error reading preferred_subnet: %v", err)
		}
	}

	if err = d.Set("private_network", flattenExtensionControllerDataplanPrivateNetwork(o["private-network"], d, "private_network", sv)); err != nil {
		if !fortiAPIPatch(o["private-network"]) {
			return fmt.Errorf("Error reading private_network: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerDataplanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtensionControllerDataplanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanModemId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanSlot(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanIccid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanApn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanPdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanSignalThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanSignalPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanMonthlyFee(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanBillingDate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanOverage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanPreferredSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanPrivateNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerDataplan(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandExtensionControllerDataplanName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok {
		t, err := expandExtensionControllerDataplanModemId(d, v, "modem_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandExtensionControllerDataplanType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("slot"); ok {
		t, err := expandExtensionControllerDataplanSlot(d, v, "slot", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slot"] = t
		}
	}

	if v, ok := d.GetOk("iccid"); ok {
		t, err := expandExtensionControllerDataplanIccid(d, v, "iccid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iccid"] = t
		}
	}

	if v, ok := d.GetOk("carrier"); ok {
		t, err := expandExtensionControllerDataplanCarrier(d, v, "carrier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier"] = t
		}
	}

	if v, ok := d.GetOk("apn"); ok {
		t, err := expandExtensionControllerDataplanApn(d, v, "apn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandExtensionControllerDataplanAuthType(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandExtensionControllerDataplanUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandExtensionControllerDataplanPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("pdn"); ok {
		t, err := expandExtensionControllerDataplanPdn(d, v, "pdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdn"] = t
		}
	}

	if v, ok := d.GetOkExists("signal_threshold"); ok {
		t, err := expandExtensionControllerDataplanSignalThreshold(d, v, "signal_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-threshold"] = t
		}
	}

	if v, ok := d.GetOk("signal_period"); ok {
		t, err := expandExtensionControllerDataplanSignalPeriod(d, v, "signal_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-period"] = t
		}
	}

	if v, ok := d.GetOkExists("capacity"); ok {
		t, err := expandExtensionControllerDataplanCapacity(d, v, "capacity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capacity"] = t
		}
	}

	if v, ok := d.GetOkExists("monthly_fee"); ok {
		t, err := expandExtensionControllerDataplanMonthlyFee(d, v, "monthly_fee", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monthly-fee"] = t
		}
	}

	if v, ok := d.GetOk("billing_date"); ok {
		t, err := expandExtensionControllerDataplanBillingDate(d, v, "billing_date", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-date"] = t
		}
	}

	if v, ok := d.GetOk("overage"); ok {
		t, err := expandExtensionControllerDataplanOverage(d, v, "overage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage"] = t
		}
	}

	if v, ok := d.GetOkExists("preferred_subnet"); ok {
		t, err := expandExtensionControllerDataplanPreferredSubnet(d, v, "preferred_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-subnet"] = t
		}
	}

	if v, ok := d.GetOk("private_network"); ok {
		t, err := expandExtensionControllerDataplanPrivateNetwork(d, v, "private_network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-network"] = t
		}
	}

	return &obj, nil
}
