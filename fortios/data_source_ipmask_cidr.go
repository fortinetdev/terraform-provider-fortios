package fortios

import (
	"crypto/md5"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIPMaskCIDR() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPMaskCIDRRead,
		Schema: map[string]*schema.Schema{
			"ipmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipmasklist": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cidrlist": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func convIpmask2CIDR(v string) string {
	line := strings.Split(v, " ")
	if len(line) >= 2 {
		ip := line[0]
		mask := line[1]
		prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()

		return ip + "/" + strconv.Itoa(prefixSize)
	}

	return ""
}

func dataSourceIPMaskCIDRRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	t := d.Get("ipmask").(string)

	tid := ""

	if t != "" {
		d.Set("cidr", convIpmask2CIDR(t))

		tid += t
	}

	v := d.Get("ipmasklist")
	if v != nil {
		l := v.([]interface{})
		if len(l) != 0 {
			result := make([](string), 0, len(l))

			for _, r := range l {
				result = append(result, convIpmask2CIDR(r.(string)))
				tid += r.(string)
			}
			d.Set("cidrlist", result)
		}
	}

	if tid != "" {
		h := md5.New()
		io.WriteString(h, tid)

		v1 := fmt.Sprintf("%x", h.Sum(nil))
		d.SetId(v1)
	} else {
		d.SetId("")
	}

	return nil
}
