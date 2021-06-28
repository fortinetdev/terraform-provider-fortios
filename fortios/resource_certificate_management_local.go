// Author: Justin Roberts (@poroping)
// Documentation:
// Justin Roberts (@poroping),

// Description: Local certificate management.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	forticlient "github.com/fortinetdev/forti-sdk-go/fortios/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCertificateManagementLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateManagementLocalCreate,
		Read:   resourceCertificateManagementLocalRead,
		Update: resourceCertificateManagementLocalUpdate,
		Delete: resourceCertificateManagementLocalDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				ForceNew:     true,
				Optional:     true,
				Sensitive:    true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"private_key": &schema.Schema{
				Type:      schema.TypeString,
				ForceNew:  true,
				Optional:  true,
				Sensitive: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return old == strings.TrimSuffix(new, "\n")
				}, // required due to api stripping trailing \n on import
			},
			"certificate_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_key_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_ca": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_valid": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"issuer": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subject": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"not_before": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"not_after": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha1_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"ike_localid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ike_localid_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceCertificateManagementLocalCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := d.Get("name").(string)

	i := &forticlient.JSONSystemVpnCertificateLocalImport{
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Certificate: d.Get("certificate").(string),
		Password:    d.Get("password").(string),
		PrivateKey:  d.Get("private_key").(string),
		Scope:       d.Get("scope").(string),
	}

	_, err := c.CreateSystemVpnCertificateLocalImport(i, vdomparam)

	if err != nil {
		return fmt.Errorf("error creating CertificateManagementLocalImport resource: %v", err)
	}

	d.SetId(mkey)

	obj, err := getObjectCertificateManagementLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating CertificateManagementLocal resource while getting object: %v", err)
	}

	scope := d.Get("scope").(string)
	var o map[string]interface{}
	if scope == "global" {
		o, err = c.UpdateCertificateLocal(obj, mkey, vdomparam)
	} else {
		o, err = c.UpdateVpnCertificateLocal(obj, mkey, vdomparam)
	}

	if err != nil {
		return fmt.Errorf("error creating CertificateManagementLocal resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateManagementLocal")
	}

	return resourceCertificateManagementLocalRead(d, m)
}

func resourceCertificateManagementLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectCertificateManagementLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating CertificateManagementLocal resource while getting object: %v", err)
	}

	scope := d.Get("scope").(string)
	var o map[string]interface{}
	if scope == "global" {
		o, err = c.UpdateCertificateLocal(obj, mkey, vdomparam)
	} else {
		o, err = c.UpdateVpnCertificateLocal(obj, mkey, vdomparam)
	}

	if err != nil {
		return fmt.Errorf("error updating CertificateManagementLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateManagementLocal")
	}

	return resourceCertificateManagementLocalRead(d, m)
}

func resourceCertificateManagementLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	scope := d.Get("scope").(string)
	var err error
	if scope == "global" {
		err = c.DeleteCertificateLocal(mkey, vdomparam)
	} else {
		err = c.DeleteVpnCertificateLocal(mkey, vdomparam)
	}

	if err != nil {
		return fmt.Errorf("error deleting CertificateManagementLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateManagementLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	// get certificate from monitor endpoint

	i := &forticlient.JSONSystemCertificateDownload{
		Mkey: mkey,
		Type: "local-cer",
	}

	cert, err := c.ReadSystemCertificateDownload(i, vdomparam)
	if err != nil {
		cert = ""
		// return fmt.Errorf("error reading VpnCertificateLocalImport resource: %v", err)
	}

	if cert == "" {
		log.Printf("[WARN] certificate (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	// get cert attributes from cmdb endpoint

	scope := d.Get("scope").(string)

	var o map[string]interface{}

	if scope == "global" {
		o, err = c.ReadCertificateLocal(mkey, vdomparam)
	} else {
		o, err = c.ReadVpnCertificateLocal(mkey, vdomparam)
	}

	if err != nil {
		return fmt.Errorf("error reading CertificateManagementLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateManagementLocal(d, o, cert, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading CertificateManagementLocal resource from API: %v", err)
	}
	return nil
}

func flattenCertificateManagementLocalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementLocalPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementLocalComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementLocalPrivateKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementLocalCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementLocalScope(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementLocalIkeLocalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementLocalIkeLocalidType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCertificateManagementLocal(d *schema.ResourceData, o map[string]interface{}, cert string, sv string) error {
	var err error

	parsed_cert, err := parseDownloadedPemCertificate(cert)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	var certificate_details []interface{}
	certificate_details = append(certificate_details, parsed_cert)

	if err = d.Set("certificate_details", certificate_details); err != nil {
		return fmt.Errorf("error reading certificate details: %v", err)
	}

	if err = d.Set("name", flattenCertificateManagementLocalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenCertificateManagementLocalComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	if err = d.Set("certificate", cert); err != nil {
		if !fortiAPIPatch(cert) {
			return fmt.Errorf("error reading certificate: %v", err)
		}
	}

	if err = d.Set("scope", flattenCertificateManagementLocalScope(o["range"], d, "scope", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("error reading scope: %v", err)
		}
	}

	if err = d.Set("ike_localid", flattenCertificateManagementLocalIkeLocalid(o["ike-localid"], d, "ike_localid", sv)); err != nil {
		if !fortiAPIPatch(o["ike-localid"]) {
			return fmt.Errorf("error reading ike_localid: %v", err)
		}
	}

	if err = d.Set("ike_localid_type", flattenCertificateManagementLocalIkeLocalidType(o["ike-localid-type"], d, "ike_localid_type", sv)); err != nil {
		if !fortiAPIPatch(o["ike-localid-type"]) {
			return fmt.Errorf("error reading ike_localid_type: %v", err)
		}
	}

	return nil
}

func flattenCertificateManagementLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCertificateManagementLocalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementLocalPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementLocalComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementLocalPrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementLocalCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementLocalScope(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementLocalIkeLocalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementLocalIkeLocalidType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateManagementLocal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandCertificateManagementLocalComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid"); ok {

		t, err := expandCertificateManagementLocalIkeLocalid(d, v, "ike_localid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid_type"); ok {

		t, err := expandCertificateManagementLocalIkeLocalidType(d, v, "ike_localid_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid-type"] = t
		}
	}

	return &obj, nil
}
