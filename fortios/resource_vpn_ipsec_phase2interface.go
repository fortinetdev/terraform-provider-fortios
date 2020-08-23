package fortios

import (
	"fmt"
	"log"

	"github.com/fortinetdev/forti-sdk-go/fortios/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVPNIPsecPhase2Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPNIPsecPhase2InterfaceCreate,
		Read:   resourceVPNIPsecPhase2InterfaceRead,
		Update: resourceVPNIPsecPhase2InterfaceUpdate,
		Delete: resourceVPNIPsecPhase2InterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"phase1name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"proposal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"src_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVPNIPsecPhase2InterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	phase1name := d.Get("phase1name").(string)
	proposal := d.Get("proposal").(string)
	comments := d.Get("comments").(string)
	srcAddrType := d.Get("src_addr_type").(string)
	srcStartIP := d.Get("src_start_ip").(string)
	srcEndIP := d.Get("src_end_ip").(string)
	srcSubnet := d.Get("src_subnet").(string)
	dstAddrType := d.Get("dst_addr_type").(string)
	srcName := d.Get("src_name").(string)
	dstName := d.Get("dst_name").(string)
	dstStartIP := d.Get("dst_start_ip").(string)
	dstEndIP := d.Get("dst_end_ip").(string)
	dstSubnet := d.Get("dst_subnet").(string)

	if srcStartIP == "" {
		srcStartIP = "0.0.0.0"
	}

	if srcEndIP == "" {
		srcEndIP = "0.0.0.0"
	}

	if srcSubnet == "" {
		srcSubnet = "0.0.0.0 0.0.0.0"
	}

	if dstStartIP == "" {
		dstStartIP = "0.0.0.0"
	}

	if dstEndIP == "" {
		dstEndIP = "0.0.0.0"
	}

	if dstSubnet == "" {
		dstSubnet = "0.0.0.0 0.0.0.0"
	}

	if srcAddrType == "range" {
		srcSubnet = srcStartIP + " " + srcEndIP
	}

	if dstAddrType == "range" {
		dstSubnet = dstStartIP + " " + dstEndIP
	}

	//Build input data by sdk
	i := &forticlient.JSONVPNIPsecPhase2Interface{
		Name:        name,
		Phase1name:  phase1name,
		Proposal:    proposal,
		Comments:    comments,
		SrcAddrType: srcAddrType,
		SrcStartIP:  srcStartIP,
		SrcEndIP:    srcEndIP,
		SrcSubnet:   srcSubnet,
		DstAddrType: dstAddrType,
		SrcName:     srcName,
		DstName:     dstName,
		DstStartIP:  dstStartIP,
		DstEndIP:    dstEndIP,
		DstSubnet:   dstSubnet,
	}

	//Call process by sdk
	o, err := c.CreateVPNIPsecPhase2Interface(i)
	if err != nil {
		return fmt.Errorf("Error creating VPN IPsec Phase2Interface: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceVPNIPsecPhase2InterfaceRead(d, m)
}

func resourceVPNIPsecPhase2InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	phase1name := d.Get("phase1name").(string)
	proposal := d.Get("proposal").(string)
	comments := d.Get("comments").(string)
	srcAddrType := d.Get("src_addr_type").(string)
	srcStartIP := d.Get("src_start_ip").(string)
	srcEndIP := d.Get("src_end_ip").(string)
	srcSubnet := d.Get("src_subnet").(string)
	dstAddrType := d.Get("dst_addr_type").(string)
	srcName := d.Get("src_name").(string)
	dstName := d.Get("dst_name").(string)
	dstStartIP := d.Get("dst_start_ip").(string)
	dstEndIP := d.Get("dst_end_ip").(string)
	dstSubnet := d.Get("dst_subnet").(string)

	if srcStartIP == "" {
		srcStartIP = "0.0.0.0"
	}

	if srcEndIP == "" {
		srcEndIP = "0.0.0.0"
	}

	if srcSubnet == "" {
		srcSubnet = "0.0.0.0 0.0.0.0"
	}

	if dstStartIP == "" {
		dstStartIP = "0.0.0.0"
	}

	if dstEndIP == "" {
		dstEndIP = "0.0.0.0"
	}

	if dstSubnet == "" {
		dstSubnet = "0.0.0.0 0.0.0.0"
	}

	if srcAddrType == "range" {
		srcSubnet = srcStartIP + " " + srcEndIP
	}

	if dstAddrType == "range" {
		dstSubnet = dstStartIP + " " + dstEndIP
	}

	//Build input data by sdk
	i := &forticlient.JSONVPNIPsecPhase2Interface{
		Name:        name,
		Phase1name:  phase1name,
		Proposal:    proposal,
		Comments:    comments,
		SrcAddrType: srcAddrType,
		SrcStartIP:  srcStartIP,
		SrcEndIP:    srcEndIP,
		SrcSubnet:   srcSubnet,
		DstAddrType: dstAddrType,
		SrcName:     srcName,
		DstName:     dstName,
		DstStartIP:  dstStartIP,
		DstEndIP:    dstEndIP,
		DstSubnet:   dstSubnet,
	}

	//Call process by sdk
	_, err := c.UpdateVPNIPsecPhase2Interface(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VPN IPsec Phase2Interface: %s", err)
	}

	return resourceVPNIPsecPhase2InterfaceRead(d, m)
}

func resourceVPNIPsecPhase2InterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	err := c.DeleteVPNIPsecPhase2Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VPN IPsec Phase2Interface: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceVPNIPsecPhase2InterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadVPNIPsecPhase2Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VPN IPsec Phase2Interface: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	d.Set("phase1name", o.Phase1name)
	d.Set("proposal", o.Proposal)
	d.Set("comments", o.Comments)
	d.Set("src_addr_type", o.SrcAddrType)
	d.Set("src_start_ip", o.SrcStartIP)
	d.Set("src_end_ip", o.SrcEndIP)
	// d.Set("src_subnet", o.SrcSubnet)
	d.Set("src_subnet", validateConvIPMask2CIDR(d.Get("src_subnet").(string), o.SrcSubnet))
	d.Set("dst_addr_type", o.DstAddrType)
	d.Set("src_name", o.SrcName)
	d.Set("dst_name", o.DstName)
	d.Set("dst_start_ip", o.DstStartIP)
	d.Set("dst_end_ip", o.DstEndIP)
	// d.Set("dst_subnet", o.DstSubnet)
	d.Set("dst_subnet", validateConvIPMask2CIDR(d.Get("dst_subnet").(string), o.DstSubnet))

	return nil
}
