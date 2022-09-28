// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch location services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLocationCreate,
		Read:   resourceSwitchControllerLocationRead,
		Update: resourceSwitchControllerLocationUpdate,
		Delete: resourceSwitchControllerLocationDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"address_civic": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"additional_code": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"block": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"branch_road": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"building": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"city": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"city_division": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"country": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"country_subdivision": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"county": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"direction": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"floor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"landmark": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"language": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"number": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"number_suffix": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"place_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"post_office_box": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"postal_community": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"primary_road": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"road_section": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"room": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"script": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"seat": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"street": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"street_name_post_mod": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"street_name_pre_mod": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"street_suffix": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"sub_branch_road": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"trailing_str_suffix": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"unit": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"zip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"parent_key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"coordinates": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"altitude": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"altitude_unit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"datum": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"latitude": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"longitude": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"parent_key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"elin_number": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"elin_num": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"parent_key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSwitchControllerLocationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerLocation(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLocation resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerLocation(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLocation resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerLocation")
	}

	return resourceSwitchControllerLocationRead(d, m)
}

func resourceSwitchControllerLocationUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerLocation(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLocation resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerLocation(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLocation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerLocation")
	}

	return resourceSwitchControllerLocationRead(d, m)
}

func resourceSwitchControllerLocationDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerLocation(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLocation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLocationRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerLocation(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLocation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLocation(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLocation resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLocationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "additional"
	if _, ok := i["additional"]; ok {

		result["additional"] = flattenSwitchControllerLocationAddressCivicAdditional(i["additional"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "additional_code"
	if _, ok := i["additional-code"]; ok {

		result["additional_code"] = flattenSwitchControllerLocationAddressCivicAdditionalCode(i["additional-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block"
	if _, ok := i["block"]; ok {

		result["block"] = flattenSwitchControllerLocationAddressCivicBlock(i["block"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "branch_road"
	if _, ok := i["branch-road"]; ok {

		result["branch_road"] = flattenSwitchControllerLocationAddressCivicBranchRoad(i["branch-road"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "building"
	if _, ok := i["building"]; ok {

		result["building"] = flattenSwitchControllerLocationAddressCivicBuilding(i["building"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "city"
	if _, ok := i["city"]; ok {

		result["city"] = flattenSwitchControllerLocationAddressCivicCity(i["city"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "city_division"
	if _, ok := i["city-division"]; ok {

		result["city_division"] = flattenSwitchControllerLocationAddressCivicCityDivision(i["city-division"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "country"
	if _, ok := i["country"]; ok {

		result["country"] = flattenSwitchControllerLocationAddressCivicCountry(i["country"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "country_subdivision"
	if _, ok := i["country-subdivision"]; ok {

		result["country_subdivision"] = flattenSwitchControllerLocationAddressCivicCountrySubdivision(i["country-subdivision"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "county"
	if _, ok := i["county"]; ok {

		result["county"] = flattenSwitchControllerLocationAddressCivicCounty(i["county"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "direction"
	if _, ok := i["direction"]; ok {

		result["direction"] = flattenSwitchControllerLocationAddressCivicDirection(i["direction"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "floor"
	if _, ok := i["floor"]; ok {

		result["floor"] = flattenSwitchControllerLocationAddressCivicFloor(i["floor"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "landmark"
	if _, ok := i["landmark"]; ok {

		result["landmark"] = flattenSwitchControllerLocationAddressCivicLandmark(i["landmark"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "language"
	if _, ok := i["language"]; ok {

		result["language"] = flattenSwitchControllerLocationAddressCivicLanguage(i["language"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {

		result["name"] = flattenSwitchControllerLocationAddressCivicName(i["name"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "number"
	if _, ok := i["number"]; ok {

		result["number"] = flattenSwitchControllerLocationAddressCivicNumber(i["number"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "number_suffix"
	if _, ok := i["number-suffix"]; ok {

		result["number_suffix"] = flattenSwitchControllerLocationAddressCivicNumberSuffix(i["number-suffix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "place_type"
	if _, ok := i["place-type"]; ok {

		result["place_type"] = flattenSwitchControllerLocationAddressCivicPlaceType(i["place-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "post_office_box"
	if _, ok := i["post-office-box"]; ok {

		result["post_office_box"] = flattenSwitchControllerLocationAddressCivicPostOfficeBox(i["post-office-box"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "postal_community"
	if _, ok := i["postal-community"]; ok {

		result["postal_community"] = flattenSwitchControllerLocationAddressCivicPostalCommunity(i["postal-community"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "primary_road"
	if _, ok := i["primary-road"]; ok {

		result["primary_road"] = flattenSwitchControllerLocationAddressCivicPrimaryRoad(i["primary-road"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "road_section"
	if _, ok := i["road-section"]; ok {

		result["road_section"] = flattenSwitchControllerLocationAddressCivicRoadSection(i["road-section"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "room"
	if _, ok := i["room"]; ok {

		result["room"] = flattenSwitchControllerLocationAddressCivicRoom(i["room"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "script"
	if _, ok := i["script"]; ok {

		result["script"] = flattenSwitchControllerLocationAddressCivicScript(i["script"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "seat"
	if _, ok := i["seat"]; ok {

		result["seat"] = flattenSwitchControllerLocationAddressCivicSeat(i["seat"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "street"
	if _, ok := i["street"]; ok {

		result["street"] = flattenSwitchControllerLocationAddressCivicStreet(i["street"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "street_name_post_mod"
	if _, ok := i["street-name-post-mod"]; ok {

		result["street_name_post_mod"] = flattenSwitchControllerLocationAddressCivicStreetNamePostMod(i["street-name-post-mod"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "street_name_pre_mod"
	if _, ok := i["street-name-pre-mod"]; ok {

		result["street_name_pre_mod"] = flattenSwitchControllerLocationAddressCivicStreetNamePreMod(i["street-name-pre-mod"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "street_suffix"
	if _, ok := i["street-suffix"]; ok {

		result["street_suffix"] = flattenSwitchControllerLocationAddressCivicStreetSuffix(i["street-suffix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sub_branch_road"
	if _, ok := i["sub-branch-road"]; ok {

		result["sub_branch_road"] = flattenSwitchControllerLocationAddressCivicSubBranchRoad(i["sub-branch-road"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "trailing_str_suffix"
	if _, ok := i["trailing-str-suffix"]; ok {

		result["trailing_str_suffix"] = flattenSwitchControllerLocationAddressCivicTrailingStrSuffix(i["trailing-str-suffix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unit"
	if _, ok := i["unit"]; ok {

		result["unit"] = flattenSwitchControllerLocationAddressCivicUnit(i["unit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "zip"
	if _, ok := i["zip"]; ok {

		result["zip"] = flattenSwitchControllerLocationAddressCivicZip(i["zip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "parent_key"
	if _, ok := i["parent-key"]; ok {

		result["parent_key"] = flattenSwitchControllerLocationAddressCivicParentKey(i["parent-key"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerLocationAddressCivicAdditional(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicAdditionalCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicBlock(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicBranchRoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicBuilding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCityDivision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCountrySubdivision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCounty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicFloor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicLandmark(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicLanguage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicNumberSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicPlaceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicPostOfficeBox(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicPostalCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicPrimaryRoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicRoadSection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicRoom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicScript(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicSeat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicStreet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicStreetNamePostMod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicStreetNamePreMod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicStreetSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicSubBranchRoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicTrailingStrSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicZip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicParentKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinates(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "altitude"
	if _, ok := i["altitude"]; ok {

		result["altitude"] = flattenSwitchControllerLocationCoordinatesAltitude(i["altitude"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "altitude_unit"
	if _, ok := i["altitude-unit"]; ok {

		result["altitude_unit"] = flattenSwitchControllerLocationCoordinatesAltitudeUnit(i["altitude-unit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "datum"
	if _, ok := i["datum"]; ok {

		result["datum"] = flattenSwitchControllerLocationCoordinatesDatum(i["datum"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "latitude"
	if _, ok := i["latitude"]; ok {

		result["latitude"] = flattenSwitchControllerLocationCoordinatesLatitude(i["latitude"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "longitude"
	if _, ok := i["longitude"]; ok {

		result["longitude"] = flattenSwitchControllerLocationCoordinatesLongitude(i["longitude"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "parent_key"
	if _, ok := i["parent-key"]; ok {

		result["parent_key"] = flattenSwitchControllerLocationCoordinatesParentKey(i["parent-key"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerLocationCoordinatesAltitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesAltitudeUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesDatum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesLatitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesLongitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesParentKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationElinNumber(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "elin_num"
	if _, ok := i["elin-num"]; ok {

		result["elin_num"] = flattenSwitchControllerLocationElinNumberElinNum(i["elin-num"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "parent_key"
	if _, ok := i["parent-key"]; ok {

		result["parent_key"] = flattenSwitchControllerLocationElinNumberParentKey(i["parent-key"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerLocationElinNumberElinNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLocationElinNumberParentKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerLocation(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerLocationName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("address_civic", flattenSwitchControllerLocationAddressCivic(o["address-civic"], d, "address_civic", sv)); err != nil {
			if !fortiAPIPatch(o["address-civic"]) {
				return fmt.Errorf("Error reading address_civic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("address_civic"); ok {
			if err = d.Set("address_civic", flattenSwitchControllerLocationAddressCivic(o["address-civic"], d, "address_civic", sv)); err != nil {
				if !fortiAPIPatch(o["address-civic"]) {
					return fmt.Errorf("Error reading address_civic: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("coordinates", flattenSwitchControllerLocationCoordinates(o["coordinates"], d, "coordinates", sv)); err != nil {
			if !fortiAPIPatch(o["coordinates"]) {
				return fmt.Errorf("Error reading coordinates: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("coordinates"); ok {
			if err = d.Set("coordinates", flattenSwitchControllerLocationCoordinates(o["coordinates"], d, "coordinates", sv)); err != nil {
				if !fortiAPIPatch(o["coordinates"]) {
					return fmt.Errorf("Error reading coordinates: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("elin_number", flattenSwitchControllerLocationElinNumber(o["elin-number"], d, "elin_number", sv)); err != nil {
			if !fortiAPIPatch(o["elin-number"]) {
				return fmt.Errorf("Error reading elin_number: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("elin_number"); ok {
			if err = d.Set("elin_number", flattenSwitchControllerLocationElinNumber(o["elin-number"], d, "elin_number", sv)); err != nil {
				if !fortiAPIPatch(o["elin-number"]) {
					return fmt.Errorf("Error reading elin_number: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerLocationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerLocationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "additional"
	if _, ok := d.GetOk(pre_append); ok {

		result["additional"], _ = expandSwitchControllerLocationAddressCivicAdditional(d, i["additional"], pre_append, sv)
	}
	pre_append = pre + ".0." + "additional_code"
	if _, ok := d.GetOk(pre_append); ok {

		result["additional-code"], _ = expandSwitchControllerLocationAddressCivicAdditionalCode(d, i["additional_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block"
	if _, ok := d.GetOk(pre_append); ok {

		result["block"], _ = expandSwitchControllerLocationAddressCivicBlock(d, i["block"], pre_append, sv)
	}
	pre_append = pre + ".0." + "branch_road"
	if _, ok := d.GetOk(pre_append); ok {

		result["branch-road"], _ = expandSwitchControllerLocationAddressCivicBranchRoad(d, i["branch_road"], pre_append, sv)
	}
	pre_append = pre + ".0." + "building"
	if _, ok := d.GetOk(pre_append); ok {

		result["building"], _ = expandSwitchControllerLocationAddressCivicBuilding(d, i["building"], pre_append, sv)
	}
	pre_append = pre + ".0." + "city"
	if _, ok := d.GetOk(pre_append); ok {

		result["city"], _ = expandSwitchControllerLocationAddressCivicCity(d, i["city"], pre_append, sv)
	}
	pre_append = pre + ".0." + "city_division"
	if _, ok := d.GetOk(pre_append); ok {

		result["city-division"], _ = expandSwitchControllerLocationAddressCivicCityDivision(d, i["city_division"], pre_append, sv)
	}
	pre_append = pre + ".0." + "country"
	if _, ok := d.GetOk(pre_append); ok {

		result["country"], _ = expandSwitchControllerLocationAddressCivicCountry(d, i["country"], pre_append, sv)
	}
	pre_append = pre + ".0." + "country_subdivision"
	if _, ok := d.GetOk(pre_append); ok {

		result["country-subdivision"], _ = expandSwitchControllerLocationAddressCivicCountrySubdivision(d, i["country_subdivision"], pre_append, sv)
	}
	pre_append = pre + ".0." + "county"
	if _, ok := d.GetOk(pre_append); ok {

		result["county"], _ = expandSwitchControllerLocationAddressCivicCounty(d, i["county"], pre_append, sv)
	}
	pre_append = pre + ".0." + "direction"
	if _, ok := d.GetOk(pre_append); ok {

		result["direction"], _ = expandSwitchControllerLocationAddressCivicDirection(d, i["direction"], pre_append, sv)
	}
	pre_append = pre + ".0." + "floor"
	if _, ok := d.GetOk(pre_append); ok {

		result["floor"], _ = expandSwitchControllerLocationAddressCivicFloor(d, i["floor"], pre_append, sv)
	}
	pre_append = pre + ".0." + "landmark"
	if _, ok := d.GetOk(pre_append); ok {

		result["landmark"], _ = expandSwitchControllerLocationAddressCivicLandmark(d, i["landmark"], pre_append, sv)
	}
	pre_append = pre + ".0." + "language"
	if _, ok := d.GetOk(pre_append); ok {

		result["language"], _ = expandSwitchControllerLocationAddressCivicLanguage(d, i["language"], pre_append, sv)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok {

		result["name"], _ = expandSwitchControllerLocationAddressCivicName(d, i["name"], pre_append, sv)
	}
	pre_append = pre + ".0." + "number"
	if _, ok := d.GetOk(pre_append); ok {

		result["number"], _ = expandSwitchControllerLocationAddressCivicNumber(d, i["number"], pre_append, sv)
	}
	pre_append = pre + ".0." + "number_suffix"
	if _, ok := d.GetOk(pre_append); ok {

		result["number-suffix"], _ = expandSwitchControllerLocationAddressCivicNumberSuffix(d, i["number_suffix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "place_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["place-type"], _ = expandSwitchControllerLocationAddressCivicPlaceType(d, i["place_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "post_office_box"
	if _, ok := d.GetOk(pre_append); ok {

		result["post-office-box"], _ = expandSwitchControllerLocationAddressCivicPostOfficeBox(d, i["post_office_box"], pre_append, sv)
	}
	pre_append = pre + ".0." + "postal_community"
	if _, ok := d.GetOk(pre_append); ok {

		result["postal-community"], _ = expandSwitchControllerLocationAddressCivicPostalCommunity(d, i["postal_community"], pre_append, sv)
	}
	pre_append = pre + ".0." + "primary_road"
	if _, ok := d.GetOk(pre_append); ok {

		result["primary-road"], _ = expandSwitchControllerLocationAddressCivicPrimaryRoad(d, i["primary_road"], pre_append, sv)
	}
	pre_append = pre + ".0." + "road_section"
	if _, ok := d.GetOk(pre_append); ok {

		result["road-section"], _ = expandSwitchControllerLocationAddressCivicRoadSection(d, i["road_section"], pre_append, sv)
	}
	pre_append = pre + ".0." + "room"
	if _, ok := d.GetOk(pre_append); ok {

		result["room"], _ = expandSwitchControllerLocationAddressCivicRoom(d, i["room"], pre_append, sv)
	}
	pre_append = pre + ".0." + "script"
	if _, ok := d.GetOk(pre_append); ok {

		result["script"], _ = expandSwitchControllerLocationAddressCivicScript(d, i["script"], pre_append, sv)
	}
	pre_append = pre + ".0." + "seat"
	if _, ok := d.GetOk(pre_append); ok {

		result["seat"], _ = expandSwitchControllerLocationAddressCivicSeat(d, i["seat"], pre_append, sv)
	}
	pre_append = pre + ".0." + "street"
	if _, ok := d.GetOk(pre_append); ok {

		result["street"], _ = expandSwitchControllerLocationAddressCivicStreet(d, i["street"], pre_append, sv)
	}
	pre_append = pre + ".0." + "street_name_post_mod"
	if _, ok := d.GetOk(pre_append); ok {

		result["street-name-post-mod"], _ = expandSwitchControllerLocationAddressCivicStreetNamePostMod(d, i["street_name_post_mod"], pre_append, sv)
	}
	pre_append = pre + ".0." + "street_name_pre_mod"
	if _, ok := d.GetOk(pre_append); ok {

		result["street-name-pre-mod"], _ = expandSwitchControllerLocationAddressCivicStreetNamePreMod(d, i["street_name_pre_mod"], pre_append, sv)
	}
	pre_append = pre + ".0." + "street_suffix"
	if _, ok := d.GetOk(pre_append); ok {

		result["street-suffix"], _ = expandSwitchControllerLocationAddressCivicStreetSuffix(d, i["street_suffix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sub_branch_road"
	if _, ok := d.GetOk(pre_append); ok {

		result["sub-branch-road"], _ = expandSwitchControllerLocationAddressCivicSubBranchRoad(d, i["sub_branch_road"], pre_append, sv)
	}
	pre_append = pre + ".0." + "trailing_str_suffix"
	if _, ok := d.GetOk(pre_append); ok {

		result["trailing-str-suffix"], _ = expandSwitchControllerLocationAddressCivicTrailingStrSuffix(d, i["trailing_str_suffix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unit"
	if _, ok := d.GetOk(pre_append); ok {

		result["unit"], _ = expandSwitchControllerLocationAddressCivicUnit(d, i["unit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "zip"
	if _, ok := d.GetOk(pre_append); ok {

		result["zip"], _ = expandSwitchControllerLocationAddressCivicZip(d, i["zip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "parent_key"
	if _, ok := d.GetOk(pre_append); ok {

		result["parent-key"], _ = expandSwitchControllerLocationAddressCivicParentKey(d, i["parent_key"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerLocationAddressCivicAdditional(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicAdditionalCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicBranchRoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicBuilding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCityDivision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCountrySubdivision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCounty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicFloor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicLandmark(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicLanguage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicNumberSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicPlaceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicPostOfficeBox(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicPostalCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicPrimaryRoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicRoadSection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicRoom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicScript(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicSeat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicStreet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicStreetNamePostMod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicStreetNamePreMod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicStreetSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicSubBranchRoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicTrailingStrSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicZip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicParentKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "altitude"
	if _, ok := d.GetOk(pre_append); ok {

		result["altitude"], _ = expandSwitchControllerLocationCoordinatesAltitude(d, i["altitude"], pre_append, sv)
	}
	pre_append = pre + ".0." + "altitude_unit"
	if _, ok := d.GetOk(pre_append); ok {

		result["altitude-unit"], _ = expandSwitchControllerLocationCoordinatesAltitudeUnit(d, i["altitude_unit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "datum"
	if _, ok := d.GetOk(pre_append); ok {

		result["datum"], _ = expandSwitchControllerLocationCoordinatesDatum(d, i["datum"], pre_append, sv)
	}
	pre_append = pre + ".0." + "latitude"
	if _, ok := d.GetOk(pre_append); ok {

		result["latitude"], _ = expandSwitchControllerLocationCoordinatesLatitude(d, i["latitude"], pre_append, sv)
	}
	pre_append = pre + ".0." + "longitude"
	if _, ok := d.GetOk(pre_append); ok {

		result["longitude"], _ = expandSwitchControllerLocationCoordinatesLongitude(d, i["longitude"], pre_append, sv)
	}
	pre_append = pre + ".0." + "parent_key"
	if _, ok := d.GetOk(pre_append); ok {

		result["parent-key"], _ = expandSwitchControllerLocationCoordinatesParentKey(d, i["parent_key"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerLocationCoordinatesAltitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesAltitudeUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesDatum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesLatitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesLongitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesParentKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationElinNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "elin_num"
	if _, ok := d.GetOk(pre_append); ok {

		result["elin-num"], _ = expandSwitchControllerLocationElinNumberElinNum(d, i["elin_num"], pre_append, sv)
	}
	pre_append = pre + ".0." + "parent_key"
	if _, ok := d.GetOk(pre_append); ok {

		result["parent-key"], _ = expandSwitchControllerLocationElinNumberParentKey(d, i["parent_key"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerLocationElinNumberElinNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationElinNumberParentKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLocation(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerLocationName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("address_civic"); ok {

		t, err := expandSwitchControllerLocationAddressCivic(d, v, "address_civic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-civic"] = t
		}
	}

	if v, ok := d.GetOk("coordinates"); ok {

		t, err := expandSwitchControllerLocationCoordinates(d, v, "coordinates", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinates"] = t
		}
	}

	if v, ok := d.GetOk("elin_number"); ok {

		t, err := expandSwitchControllerLocationElinNumber(d, v, "elin_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["elin-number"] = t
		}
	}

	return &obj, nil
}
