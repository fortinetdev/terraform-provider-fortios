// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure USB auto installation.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoInstall() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoInstallUpdate,
		Read:   resourceSystemAutoInstallRead,
		Update: resourceSystemAutoInstallUpdate,
		Delete: resourceSystemAutoInstallDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"auto_install_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_install_image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_config_file": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"default_image_file": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemAutoInstallUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAutoInstall(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoInstall resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutoInstall(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoInstall resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutoInstall")
	}

	return resourceSystemAutoInstallRead(d, m)
}

func resourceSystemAutoInstallDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAutoInstall(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemAutoInstall resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoInstall(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemAutoInstall resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoInstallRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemAutoInstall(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoInstall resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoInstall(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoInstall resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoInstallAutoInstallConfig(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoInstallAutoInstallImage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoInstallDefaultConfigFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoInstallDefaultImageFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutoInstall(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("auto_install_config", flattenSystemAutoInstallAutoInstallConfig(o["auto-install-config"], d, "auto_install_config", sv)); err != nil {
		if !fortiAPIPatch(o["auto-install-config"]) {
			return fmt.Errorf("Error reading auto_install_config: %v", err)
		}
	}

	if err = d.Set("auto_install_image", flattenSystemAutoInstallAutoInstallImage(o["auto-install-image"], d, "auto_install_image", sv)); err != nil {
		if !fortiAPIPatch(o["auto-install-image"]) {
			return fmt.Errorf("Error reading auto_install_image: %v", err)
		}
	}

	if err = d.Set("default_config_file", flattenSystemAutoInstallDefaultConfigFile(o["default-config-file"], d, "default_config_file", sv)); err != nil {
		if !fortiAPIPatch(o["default-config-file"]) {
			return fmt.Errorf("Error reading default_config_file: %v", err)
		}
	}

	if err = d.Set("default_image_file", flattenSystemAutoInstallDefaultImageFile(o["default-image-file"], d, "default_image_file", sv)); err != nil {
		if !fortiAPIPatch(o["default-image-file"]) {
			return fmt.Errorf("Error reading default_image_file: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoInstallFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAutoInstallAutoInstallConfig(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoInstallAutoInstallImage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoInstallDefaultConfigFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoInstallDefaultImageFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoInstall(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_install_config"); ok {
		if setArgNil {
			obj["auto-install-config"] = nil
		} else {
			t, err := expandSystemAutoInstallAutoInstallConfig(d, v, "auto_install_config", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-install-config"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_install_image"); ok {
		if setArgNil {
			obj["auto-install-image"] = nil
		} else {
			t, err := expandSystemAutoInstallAutoInstallImage(d, v, "auto_install_image", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-install-image"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_config_file"); ok {
		if setArgNil {
			obj["default-config-file"] = nil
		} else {
			t, err := expandSystemAutoInstallDefaultConfigFile(d, v, "default_config_file", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-config-file"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_image_file"); ok {
		if setArgNil {
			obj["default-image-file"] = nil
		} else {
			t, err := expandSystemAutoInstallDefaultImageFile(d, v, "default_image_file", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-image-file"] = t
			}
		}
	}

	return &obj, nil
}
