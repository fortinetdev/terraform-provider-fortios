// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure GRE tunnel.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemGreTunnel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemGreTunnelRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sequence_number_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sequence_number_reception": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"checksum_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"checksum_reception": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_outbound": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"key_inbound": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dscp_copying": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"keepalive_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"keepalive_failtimes": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemGreTunnelRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemGreTunnel: type error")
	}

	o, err := c.ReadSystemGreTunnel(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemGreTunnel: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemGreTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemGreTunnel from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemGreTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelRemoteGw6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelLocalGw6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelSequenceNumberTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelSequenceNumberReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelChecksumTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelChecksumReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelKeyOutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelKeyInbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelDscpCopying(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelKeepaliveInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGreTunnelKeepaliveFailtimes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemGreTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemGreTunnelName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemGreTunnelInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_version", dataSourceFlattenSystemGreTunnelIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("remote_gw6", dataSourceFlattenSystemGreTunnelRemoteGw6(o["remote-gw6"], d, "remote_gw6")); err != nil {
		if !fortiAPIPatch(o["remote-gw6"]) {
			return fmt.Errorf("Error reading remote_gw6: %v", err)
		}
	}

	if err = d.Set("local_gw6", dataSourceFlattenSystemGreTunnelLocalGw6(o["local-gw6"], d, "local_gw6")); err != nil {
		if !fortiAPIPatch(o["local-gw6"]) {
			return fmt.Errorf("Error reading local_gw6: %v", err)
		}
	}

	if err = d.Set("remote_gw", dataSourceFlattenSystemGreTunnelRemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("local_gw", dataSourceFlattenSystemGreTunnelLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("sequence_number_transmission", dataSourceFlattenSystemGreTunnelSequenceNumberTransmission(o["sequence-number-transmission"], d, "sequence_number_transmission")); err != nil {
		if !fortiAPIPatch(o["sequence-number-transmission"]) {
			return fmt.Errorf("Error reading sequence_number_transmission: %v", err)
		}
	}

	if err = d.Set("sequence_number_reception", dataSourceFlattenSystemGreTunnelSequenceNumberReception(o["sequence-number-reception"], d, "sequence_number_reception")); err != nil {
		if !fortiAPIPatch(o["sequence-number-reception"]) {
			return fmt.Errorf("Error reading sequence_number_reception: %v", err)
		}
	}

	if err = d.Set("checksum_transmission", dataSourceFlattenSystemGreTunnelChecksumTransmission(o["checksum-transmission"], d, "checksum_transmission")); err != nil {
		if !fortiAPIPatch(o["checksum-transmission"]) {
			return fmt.Errorf("Error reading checksum_transmission: %v", err)
		}
	}

	if err = d.Set("checksum_reception", dataSourceFlattenSystemGreTunnelChecksumReception(o["checksum-reception"], d, "checksum_reception")); err != nil {
		if !fortiAPIPatch(o["checksum-reception"]) {
			return fmt.Errorf("Error reading checksum_reception: %v", err)
		}
	}

	if err = d.Set("key_outbound", dataSourceFlattenSystemGreTunnelKeyOutbound(o["key-outbound"], d, "key_outbound")); err != nil {
		if !fortiAPIPatch(o["key-outbound"]) {
			return fmt.Errorf("Error reading key_outbound: %v", err)
		}
	}

	if err = d.Set("key_inbound", dataSourceFlattenSystemGreTunnelKeyInbound(o["key-inbound"], d, "key_inbound")); err != nil {
		if !fortiAPIPatch(o["key-inbound"]) {
			return fmt.Errorf("Error reading key_inbound: %v", err)
		}
	}

	if err = d.Set("dscp_copying", dataSourceFlattenSystemGreTunnelDscpCopying(o["dscp-copying"], d, "dscp_copying")); err != nil {
		if !fortiAPIPatch(o["dscp-copying"]) {
			return fmt.Errorf("Error reading dscp_copying: %v", err)
		}
	}

	if err = d.Set("diffservcode", dataSourceFlattenSystemGreTunnelDiffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("keepalive_interval", dataSourceFlattenSystemGreTunnelKeepaliveInterval(o["keepalive-interval"], d, "keepalive_interval")); err != nil {
		if !fortiAPIPatch(o["keepalive-interval"]) {
			return fmt.Errorf("Error reading keepalive_interval: %v", err)
		}
	}

	if err = d.Set("keepalive_failtimes", dataSourceFlattenSystemGreTunnelKeepaliveFailtimes(o["keepalive-failtimes"], d, "keepalive_failtimes")); err != nil {
		if !fortiAPIPatch(o["keepalive-failtimes"]) {
			return fmt.Errorf("Error reading keepalive_failtimes: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemGreTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
