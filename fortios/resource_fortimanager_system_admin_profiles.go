package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemAdminProfiles() *schema.Resource {
	return &schema.Resource{
		Create: createFMGSystemAdminProfiles,
		Read:   readFMGSystemAdminProfiles,
		Update: updateFMGSystemAdminProfiles,
		Delete: deleteFMGSystemAdminProfiles,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"profileid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_setting": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"adom_switch": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"deploy_management": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"import_policy_packages": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"intf_mapping": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_ap": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_forticlient": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_fortiswitch": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"vpn_manager": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"log_viewer": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"fortiguard_center": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"fortiguard_center_advanced": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"fortiguard_center_firmware_managerment": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"fortiguard_center_licensing": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_manager": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_operation": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"config_retrieve": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"config_revert": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_revision_deletion": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"terminal_access": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_config": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_profile": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"device_wan_link_load_balance": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"policy_objects": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"global_policy_packages": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"assignment": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"adom_policy_packages": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"consistency_check": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"set_install_targets": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
		},
	}
}

func createFMGSystemAdminProfiles(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGSystemAdminProfiles")()

	i := &fmgclient.JSONSysAdminProfiles{
		JSONSysAdminProfileGenernal: &fmgclient.JSONSysAdminProfileGenernal{
			ProfileId:            d.Get("profileid").(string),
			Description:          d.Get("description").(string),
			SystemSetting:        d.Get("system_setting").(string),
			AdomSwitch:           d.Get("adom_switch").(string),
			DeployManagement:     d.Get("deploy_management").(string),
			ImportPolicyPackages: d.Get("import_policy_packages").(string),
			IntfMapping:          d.Get("intf_mapping").(string),
			DeviceAp:             d.Get("device_ap").(string),
			DeviceFortiClient:    d.Get("device_forticlient").(string),
			DeviceFortiSwitch:    d.Get("device_fortiswitch").(string),
			VpnManager:           d.Get("vpn_manager").(string),
			LogViewer:            d.Get("log_viewer").(string),
		},
		JSONSysAdminProfileFortiGuard: &fmgclient.JSONSysAdminProfileFortiGuard{
			FgdCenter:          d.Get("fortiguard_center").(string),
			FgdCenterAdvanced:  d.Get("fortiguard_center_advanced").(string),
			FgdCenterFmwMgmt:   d.Get("fortiguard_center_firmware_managerment").(string),
			FgdCenterLicensing: d.Get("fortiguard_center_licensing").(string),
		},
		JSONSysAdminProfileDeviceManager: &fmgclient.JSONSysAdminProfileDeviceManager{
			DeviceManager:            d.Get("device_manager").(string),
			DeviceOp:                 d.Get("device_operation").(string),
			ConfigRetrieve:           d.Get("config_retrieve").(string),
			ConfigRevert:             d.Get("config_revert").(string),
			DeviceRevisionDeletion:   d.Get("device_revision_deletion").(string),
			TermAccess:               d.Get("terminal_access").(string),
			DeviceConfig:             d.Get("device_config").(string),
			DeviceProfile:            d.Get("device_profile").(string),
			DeviceWanLinkLoadBalance: d.Get("device_wan_link_load_balance").(string),
		},
		JSONSysAdminProfilePolicyObject: &fmgclient.JSONSysAdminProfilePolicyObject{
			PolicyObjects:        d.Get("policy_objects").(string),
			GlobalPolicyPackages: d.Get("global_policy_packages").(string),
			Assignment:           d.Get("assignment").(string),
			AdomPolicyPackages:   d.Get("adom_policy_packages").(string),
			ConsistencyCheck:     d.Get("consistency_check").(string),
			SetInstallTargets:    d.Get("set_install_targets").(string),
		},
	}

	err := c.CreateUpdateSystemAdminProfiles(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Admin Profiles: %s", err)
	}

	d.SetId(i.ProfileId)

	return readFMGSystemAdminProfiles(d, m)
}

func readFMGSystemAdminProfiles(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemAdminProfiles")()

	profileId := d.Id()
	o, err := c.ReadSystemAdminProfiles(profileId)
	if err != nil {
		return fmt.Errorf("Error reading System Admin Profiles: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("profileid", o.ProfileId)
	d.Set("description", o.Description)
	d.Set("system_setting", o.SystemSetting)
	d.Set("adom_switch", o.AdomSwitch)
	d.Set("deploy_management", o.DeployManagement)
	d.Set("import_policy_packages", o.ImportPolicyPackages)
	d.Set("intf_mapping", o.IntfMapping)
	d.Set("device_ap", o.DeviceAp)
	d.Set("device_forticlient", o.DeviceFortiClient)
	d.Set("device_fortiswitch", o.DeviceFortiSwitch)
	d.Set("vpn_manager", o.VpnManager)
	d.Set("log_viewer", o.LogViewer)

	d.Set("fortiguard_center", o.FgdCenter)
	d.Set("fortiguard_center_advanced", o.FgdCenterAdvanced)
	d.Set("fortiguard_center_firmware_managerment", o.FgdCenterFmwMgmt)
	d.Set("fortiguard_center_licensing", o.FgdCenterLicensing)

	d.Set("device_manager", o.DeviceManager)
	d.Set("device_operation", o.DeviceOp)
	d.Set("config_retrieve", o.ConfigRetrieve)
	d.Set("config_revert", o.ConfigRevert)
	d.Set("device_revision_deletion", o.DeviceRevisionDeletion)
	d.Set("terminal_access", o.TermAccess)
	d.Set("device_config", o.DeviceConfig)
	d.Set("device_profile", o.DeviceProfile)
	d.Set("device_wan_link_load_balance", o.DeviceWanLinkLoadBalance)

	d.Set("policy_objects", o.PolicyObjects)
	d.Set("global_policy_packages", o.GlobalPolicyPackages)
	d.Set("assignment", o.Assignment)
	d.Set("adom_policy_packages", o.AdomPolicyPackages)
	d.Set("consistency_check", o.ConsistencyCheck)
	d.Set("set_install_targets", o.SetInstallTargets)

	return nil
}

func updateFMGSystemAdminProfiles(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGSystemAdminProfiles")()

	if d.HasChange("profileid") {
		return fmt.Errorf("the profileid argument is the key and should not be modified here")
	}

	i := &fmgclient.JSONSysAdminProfiles{
		JSONSysAdminProfileGenernal: &fmgclient.JSONSysAdminProfileGenernal{
			ProfileId:            d.Get("profileid").(string),
			Description:          d.Get("description").(string),
			SystemSetting:        d.Get("system_setting").(string),
			AdomSwitch:           d.Get("adom_switch").(string),
			DeployManagement:     d.Get("deploy_management").(string),
			ImportPolicyPackages: d.Get("import_policy_packages").(string),
			IntfMapping:          d.Get("intf_mapping").(string),
			DeviceAp:             d.Get("device_ap").(string),
			DeviceFortiClient:    d.Get("device_forticlient").(string),
			DeviceFortiSwitch:    d.Get("device_fortiswitch").(string),
			VpnManager:           d.Get("vpn_manager").(string),
			LogViewer:            d.Get("log_viewer").(string),
		},
		JSONSysAdminProfileFortiGuard: &fmgclient.JSONSysAdminProfileFortiGuard{
			FgdCenter:          d.Get("fortiguard_center").(string),
			FgdCenterAdvanced:  d.Get("fortiguard_center_advanced").(string),
			FgdCenterFmwMgmt:   d.Get("fortiguard_center_firmware_managerment").(string),
			FgdCenterLicensing: d.Get("fortiguard_center_licensing").(string),
		},
		JSONSysAdminProfileDeviceManager: &fmgclient.JSONSysAdminProfileDeviceManager{
			DeviceManager:            d.Get("device_manager").(string),
			DeviceOp:                 d.Get("device_operation").(string),
			ConfigRetrieve:           d.Get("config_retrieve").(string),
			ConfigRevert:             d.Get("config_revert").(string),
			DeviceRevisionDeletion:   d.Get("device_revision_deletion").(string),
			TermAccess:               d.Get("terminal_access").(string),
			DeviceConfig:             d.Get("device_config").(string),
			DeviceProfile:            d.Get("device_profile").(string),
			DeviceWanLinkLoadBalance: d.Get("device_wan_link_load_balance").(string),
		},
		JSONSysAdminProfilePolicyObject: &fmgclient.JSONSysAdminProfilePolicyObject{
			PolicyObjects:        d.Get("policy_objects").(string),
			GlobalPolicyPackages: d.Get("global_policy_packages").(string),
			Assignment:           d.Get("assignment").(string),
			AdomPolicyPackages:   d.Get("adom_policy_packages").(string),
			ConsistencyCheck:     d.Get("consistency_check").(string),
			SetInstallTargets:    d.Get("set_install_targets").(string),
		},
	}

	err := c.CreateUpdateSystemAdminProfiles(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating System Admin Profiles: %s", err)
	}

	return readFMGSystemAdminProfiles(d, m)
}

func deleteFMGSystemAdminProfiles(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGSystemAdminProfiles")()

	profileId := d.Id()

	err := c.DeleteSystemAdminProfiles(profileId)
	if err != nil {
		return fmt.Errorf("Error deleting System Admin Profiles: %s", err)
	}

	d.SetId("")

	return nil
}
