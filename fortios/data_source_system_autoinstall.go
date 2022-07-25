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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemAutoInstall() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAutoInstallRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"auto_install_config": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_install_image": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_config_file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_image_file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemAutoInstallRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemAutoInstall"

	o, err := c.ReadSystemAutoInstall(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoInstall: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAutoInstall(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoInstall from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAutoInstallAutoInstallConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoInstallAutoInstallImage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoInstallDefaultConfigFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoInstallDefaultImageFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAutoInstall(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_install_config", dataSourceFlattenSystemAutoInstallAutoInstallConfig(o["auto-install-config"], d, "auto_install_config")); err != nil {
		if !fortiAPIPatch(o["auto-install-config"]) {
			return fmt.Errorf("Error reading auto_install_config: %v", err)
		}
	}

	if err = d.Set("auto_install_image", dataSourceFlattenSystemAutoInstallAutoInstallImage(o["auto-install-image"], d, "auto_install_image")); err != nil {
		if !fortiAPIPatch(o["auto-install-image"]) {
			return fmt.Errorf("Error reading auto_install_image: %v", err)
		}
	}

	if err = d.Set("default_config_file", dataSourceFlattenSystemAutoInstallDefaultConfigFile(o["default-config-file"], d, "default_config_file")); err != nil {
		if !fortiAPIPatch(o["default-config-file"]) {
			return fmt.Errorf("Error reading default_config_file: %v", err)
		}
	}

	if err = d.Set("default_image_file", dataSourceFlattenSystemAutoInstallDefaultImageFile(o["default-image-file"], d, "default_image_file")); err != nil {
		if !fortiAPIPatch(o["default-image-file"]) {
			return fmt.Errorf("Error reading default_image_file: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAutoInstallFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
