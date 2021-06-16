// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Justin Roberts (@poroping)
// Documentation:
// Justin Roberts (@poroping),

// Description: Download system certificates.

package fortios

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"strconv"
	"time"

	forticlient "github.com/fortinetdev/forti-sdk-go/fortios/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemCertificateDownload() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCertificateDownloadRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer_cn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject_cn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"valid_not_before": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"valid_not_after": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_valid": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemCertificateDownloadRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	i := &forticlient.JSONSystemCertificateDownload{
		Mkey: d.Get("name").(string),
		Type: d.Get("type").(string),
	}

	o, err := c.ReadSystemCertificateDownload(i)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateDownload: %v", err)
	}

	if o == "" {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemCertificateDownload(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateDownload from API: %v", err)
	}

	d.SetId(i.Mkey)

	return nil
}

func dataSourceRefreshObjectSystemCertificateDownload(d *schema.ResourceData, o string) error {
	var err error

	der, rest := pem.Decode([]byte(o))

	if der == nil {
		return fmt.Errorf("Error valid certificate not found: %v", rest)
	}

	cert, err2 := x509.ParseCertificate(der.Bytes)

	if err2 != nil {
		return fmt.Errorf("Error parsing certificate: %v", err)
	}

	expired := time.Now().Before(cert.NotAfter)
	prevalid := time.Now().After(cert.NotBefore)
	valid := expired && prevalid

	if err = d.Set("issuer_cn", cert.Issuer.CommonName); err != nil {
		if !fortiAPIPatch(o) {
			return fmt.Errorf("Error reading certificate issuer cn: %v", err)
		}
	}

	if err = d.Set("subject_cn", cert.Subject.CommonName); err != nil {
		if !fortiAPIPatch(o) {
			return fmt.Errorf("Error reading certificate subject cn: %v", err)
		}
	}

	if err = d.Set("serial", cert.SerialNumber.String()); err != nil {
		if !fortiAPIPatch(o) {
			return fmt.Errorf("Error reading certificate serial number: %v", err)
		}
	}

	if err = d.Set("version", cert.Version); err != nil {
		if !fortiAPIPatch(o) {
			return fmt.Errorf("Error reading certificate version: %v", err)
		}
	}

	if err = d.Set("valid_not_before", cert.NotBefore.Format(time.RFC3339)); err != nil {
		if !fortiAPIPatch(o) {
			return fmt.Errorf("Error reading certificate validity: %v", err)
		}
	}

	if err = d.Set("valid_not_after", cert.NotAfter.Format(time.RFC3339)); err != nil {
		if !fortiAPIPatch(o) {
			return fmt.Errorf("Error reading certificate validity: %v", err)
		}
	}

	if err = d.Set("is_valid", valid); err != nil {
		if !fortiAPIPatch(o) {
			return fmt.Errorf("Error reading certificate validity: %v", err)
		}
	}

	if err = d.Set("certificate", o); err != nil {
		if !fortiAPIPatch(o) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemCertificateDownloadFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
