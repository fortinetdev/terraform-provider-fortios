package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLogFortiAnalyzerSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortiAnalyzerSettingCreate,
		Read:   resourceLogFortiAnalyzerSettingRead,
		Update: resourceLogFortiAnalyzerSettingUpdate,
		Delete: resourceLogFortiAnalyzerSettingDelete,

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reliable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hmac_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceLogFortiAnalyzerSettingCreate(d *schema.ResourceData, m interface{}) error {
	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Get Params from d
	// status := d.Get("status").(string)
	// server := d.Get("server").(string)
	// sourceIP := d.Get("source_ip").(string)
	// uploadOption := d.Get("upload_option").(string)
	// reliable := d.Get("reliable").(string)
	// hmacAlgorithm := d.Get("hmac_algorithm").(string)
	// encAlgorithm := d.Get("enc_algorithm").(string)

	// //Build input data by sdk
	// i := &forticlient.JSONLogFortiAnalyzerSetting{
	// 	Status:       status,
	// 	Server:       server,
	// 	SourceIP:       sourceIP,
	// 	UploadOption:       uploadOption,
	// 	Reliable:       reliable,
	// 	HmacAlgorithm:       hmacAlgorithm,
	// 	EncAlgorithm:       encAlgorithm,
	// }

	// //Call process by sdk
	// o, err := c.CreateLogFortiAnalyzerSetting(i)
	// if err != nil {
	// 	return fmt.Errorf("Error creating Log FortiAnalyzer Setting: %s", err)
	// }

	// // Set index for d
	// // d.SetId(strconv.Itoa(int(o.Mkey)))
	//     d.SetId(o.Mkey)

	err := resourceLogFortiAnalyzerSettingUpdate(d, m)
	if err != nil {
		return fmt.Errorf("Error updating Log FortiAnalyzer Setting: %s", err)
	}

	return nil
}

func resourceLogFortiAnalyzerSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d

	status := d.Get("status").(string)
	server := d.Get("server").(string)
	sourceIP := d.Get("source_ip").(string)
	uploadOption := d.Get("upload_option").(string)
	reliable := d.Get("reliable").(string)
	hmacAlgorithm := d.Get("hmac_algorithm").(string)
	encAlgorithm := d.Get("enc_algorithm").(string)

	//Build input data by sdk
	i := &forticlient.JSONLogFortiAnalyzerSetting{
		Status:        status,
		Server:        server,
		SourceIP:      sourceIP,
		UploadOption:  uploadOption,
		Reliable:      reliable,
		HmacAlgorithm: hmacAlgorithm,
		EncAlgorithm:  encAlgorithm,
	}

	//Call process by sdk
	_, err := c.UpdateLogFortiAnalyzerSetting(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Log FortiAnalyzer Setting: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)
	d.SetId("1")

	return nil
}

func resourceLogFortiAnalyzerSettingDelete(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// err := c.DeleteLogFortiAnalyzerSetting(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error deleting Log FortiAnalyzer Setting: %s", err)
	// }

	// //Set index for d
	d.SetId("")

	return nil
}

func resourceLogFortiAnalyzerSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadLogFortiAnalyzerSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Log FortiAnalyzer Setting: %s", err)
	}

	//Refresh property
	d.Set("status", o.Status)
	d.Set("server", o.Server)
	d.Set("source_ip", o.SourceIP)
	d.Set("upload_option", o.UploadOption)
	d.Set("reliable", o.Reliable)
	d.Set("hmac_algorithm", o.HmacAlgorithm)
	d.Set("enc_algorithm", o.EncAlgorithm)

	return nil
}
