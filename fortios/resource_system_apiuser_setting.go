package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemAPIUserSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAPIUserSettingCreate,
		Read:   resourceSystemAPIUserSettingRead,
		Update: resourceSystemAPIUserSettingUpdate,
		Delete: resourceSystemAPIUserSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"trusthost": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ipv4_trusthost": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemAPIUserSettingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	accprofile := d.Get("accprofile").(string)
	vdom := d.Get("vdom").([]interface{})
	trusthost := d.Get("trusthost").([]interface{})
	comments := d.Get("comments").(string)

	var vdoms []forticlient.MultValue

	for _, v := range vdom {
		if v == nil {
			return fmt.Errorf("null value")
		}
		vdoms = append(vdoms,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var trusthosts []forticlient.APIUserMultValue

	for _, item := range trusthost {
		detail := item.(map[string]interface{})

		tItem := forticlient.APIUserMultValue{}

		if v, ok := detail["type"]; ok {
			tItem.TypeF = v.(string)
		}
		if v, ok := detail["ipv4_trusthost"]; ok {
			tItem.Ipv4TrustHost = v.(string)
		}
		trusthosts = append(trusthosts, tItem)
	}

	//Build input data by sdk
	i := &forticlient.JSONSystemAPIUserSetting{
		Name:       name,
		Accprofile: accprofile,
		Vdom:       vdoms,
		Trusthost:  trusthosts,
		Comments:   comments,
	}

	//Call process by sdk
	o, err := c.CreateSystemAPIUserSetting(i)
	if err != nil {
		return fmt.Errorf("Error creating System APIUser Setting: %s", err)
	}

	// Set index for d
	d.SetId(o.Mkey)

	return resourceSystemAPIUserSettingRead(d, m)
}

func resourceSystemAPIUserSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	accprofile := d.Get("accprofile").(string)
	vdom := d.Get("vdom").([]interface{})
	trusthost := d.Get("trusthost").([]interface{})
	comments := d.Get("comments").(string)

	var vdoms []forticlient.MultValue

	for _, v := range vdom {
		if v == nil {
			return fmt.Errorf("null value")
		}
		vdoms = append(vdoms,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var trusthosts []forticlient.APIUserMultValue

	for _, item := range trusthost {
		detail := item.(map[string]interface{})

		tItem := forticlient.APIUserMultValue{}

		if v, ok := detail["type"]; ok {
			tItem.TypeF = v.(string)
		}
		if v, ok := detail["ipv4_trusthost"]; ok {
			tItem.Ipv4TrustHost = v.(string)
		}
		trusthosts = append(trusthosts, tItem)
	}

	//Build input data by sdk
	i := &forticlient.JSONSystemAPIUserSetting{
		Name:       name,
		Accprofile: accprofile,
		Vdom:       vdoms,
		Trusthost:  trusthosts,
		Comments:   comments,
	}

	//Call process by sdk
	_, err := c.UpdateSystemAPIUserSetting(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System APIUser Setting: %s", err)
	}

	return resourceSystemAPIUserSettingRead(d, m)
}

func resourceSystemAPIUserSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteSystemAPIUserSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting System APIUser Setting: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceSystemAPIUserSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemAPIUserSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System APIUser Setting: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	d.Set("accprofile", o.Accprofile)
	vdom := forticlient.ExtractString(o.Vdom)
	if err := d.Set("vdom", vdom); err != nil {
		log.Printf("[WARN] Error setting System APIUser Setting for (%s): %s", d.Id(), err)
	}

	var items []interface{}
	for _, z := range o.Trusthost {
		m := make(map[string]interface{})
		m["type"] = z.TypeF
		m["ipv4_trusthost"] = z.Ipv4TrustHost

		items = append(items, m)
	}

	if err := d.Set("trusthost", items); err != nil {
		log.Printf("[WARN] Error setting System APIUser Setting for (%s): %s", d.Id(), err)
	}
	d.Set("comments", o.Comments)

	return nil
}
