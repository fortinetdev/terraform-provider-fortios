// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure local FortiGuard Web Filter local ratings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebfilterFtgdLocalRating() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFtgdLocalRatingCreate,
		Read:   resourceWebfilterFtgdLocalRatingRead,
		Update: resourceWebfilterFtgdLocalRatingUpdate,
		Delete: resourceWebfilterFtgdLocalRatingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required: true,
			},
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rating": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceWebfilterFtgdLocalRatingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterFtgdLocalRating(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalRating resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterFtgdLocalRating(obj)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalRating resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdLocalRating")
	}

	return resourceWebfilterFtgdLocalRatingRead(d, m)
}

func resourceWebfilterFtgdLocalRatingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterFtgdLocalRating(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalRating resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterFtgdLocalRating(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalRating resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdLocalRating")
	}

	return resourceWebfilterFtgdLocalRatingRead(d, m)
}

func resourceWebfilterFtgdLocalRatingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWebfilterFtgdLocalRating(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFtgdLocalRating resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFtgdLocalRatingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebfilterFtgdLocalRating(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalRating resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFtgdLocalRating(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalRating resource from API: %v", err)
	}
	return nil
}


func flattenWebfilterFtgdLocalRatingUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalRatingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalRatingRating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectWebfilterFtgdLocalRating(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("url", flattenWebfilterFtgdLocalRatingUrl(o["url"], d, "url")); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("status", flattenWebfilterFtgdLocalRatingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("rating", flattenWebfilterFtgdLocalRatingRating(o["rating"], d, "rating")); err != nil {
		if !fortiAPIPatch(o["rating"]) {
			return fmt.Errorf("Error reading rating: %v", err)
		}
	}


	return nil
}

func flattenWebfilterFtgdLocalRatingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandWebfilterFtgdLocalRatingUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalRatingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalRatingRating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectWebfilterFtgdLocalRating(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("url"); ok {
		t, err := expandWebfilterFtgdLocalRatingUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebfilterFtgdLocalRatingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("rating"); ok {
		t, err := expandWebfilterFtgdLocalRatingRating(d, v, "rating")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rating"] = t
		}
	}


	return &obj, nil
}

