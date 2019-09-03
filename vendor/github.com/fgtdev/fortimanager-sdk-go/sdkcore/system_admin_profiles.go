package fortimngclient

import (
	"fmt"
)

type JSONSysAdminProfileGenernal struct {
	ProfileId            string `json:"profileid"`
	Description          string `json:"description"`
	SystemSetting        string `json:"system-setting"`
	AdomSwitch           string `json:"adom-switch"`
	DeployManagement     string `json:"deploy-management"`
	ImportPolicyPackages string `json:"import-policy-packages"`
	IntfMapping          string `json:"intf-mapping"`
	DeviceAp             string `json:"device-ap"`
	DeviceFortiClient    string `json:"device-forticlient"`
	DeviceFortiSwitch    string `json:"device-fortiswitch"`
	VpnManager           string `json:"vpn-manager"`
	LogViewer            string `json:"log-viewer"`
}

type JSONSysAdminProfileFortiGuard struct {
	FgdCenter          string `json:"fgd_center"`
	FgdCenterAdvanced  string `json:"fgd-center-advanced"`
	FgdCenterFmwMgmt   string `json:"fgd-center-fmw-mgmt"`
	FgdCenterLicensing string `json:"fgd-center-licensing"`
}

type JSONSysAdminProfileDeviceManager struct {
	DeviceManager            string `json:"device-manager"`
	DeviceOp                 string `json:"device-op"`
	ConfigRetrieve           string `json:"config-retrieve"`
	ConfigRevert             string `json:"config-revert"`
	DeviceRevisionDeletion   string `json:"device-revision-deletion"`
	TermAccess               string `json:"term-access"`
	DeviceConfig             string `json:"device-config"`
	DeviceProfile            string `json:"device-profile"`
	DeviceWanLinkLoadBalance string `json:"device-wan-link-load-balance"`
}

type JSONSysAdminProfilePolicyObject struct {
	PolicyObjects        string `json:"policy-objects"`
	GlobalPolicyPackages string `json:"global-policy-packages"`
	Assignment           string `json:"assignment"`
	AdomPolicyPackages   string `json:"adom-policy-packages"`
	ConsistencyCheck     string `json:"consistency-check"`
	SetInstallTargets    string `json:"set-install-targets"`
}

type JSONSysAdminProfiles struct {
	*JSONSysAdminProfileGenernal
	*JSONSysAdminProfileFortiGuard
	*JSONSysAdminProfileDeviceManager
	*JSONSysAdminProfilePolicyObject
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateSystemAdminProfiles(params *JSONSysAdminProfiles, method string) (err error) {
	defer c.Trace("CreateUpdateSystemAdminProfiles")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/admin/profile",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateSystemAdminProfile failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemAdminProfiles(id string) (out *JSONSysAdminProfiles, err error) {
	defer c.Trace("ReadSystemAdminProfiles")()

	p := map[string]interface{}{
		"url": "/cli/global/system/admin/profile/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemAdminProfile failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSysAdminProfiles{
		JSONSysAdminProfileGenernal:      &JSONSysAdminProfileGenernal{},
		JSONSysAdminProfileFortiGuard:    &JSONSysAdminProfileFortiGuard{},
		JSONSysAdminProfileDeviceManager: &JSONSysAdminProfileDeviceManager{},
		JSONSysAdminProfilePolicyObject:  &JSONSysAdminProfilePolicyObject{},
	}
	if data["profileid"] != nil {
		out.ProfileId = data["profileid"].(string)
	}
	if data["description"] != nil {
		out.Description = data["description"].(string)
	}
	if data["system-setting"] != nil {
		out.SystemSetting = c.AccessRight2Str(int(data["system-setting"].(float64)))
	}
	if data["adom-switch"] != nil {
		out.AdomSwitch = c.AccessRight2Str(int(data["adom-switch"].(float64)))
	}
	if data["deploy-management"] != nil {
		out.DeployManagement = c.AccessRight2Str(int(data["deploy-management"].(float64)))
	}
	if data["import-policy-packages"] != nil {
		out.ImportPolicyPackages = c.AccessRight2Str(int(data["import-policy-packages"].(float64)))
	}
	if data["intf-mapping"] != nil {
		out.IntfMapping = c.AccessRight2Str(int(data["intf-mapping"].(float64)))
	}
	if data["device-ap"] != nil {
		out.DeviceAp = c.AccessRight2Str(int(data["device-ap"].(float64)))
	}
	if data["device-forticlient"] != nil {
		out.DeviceFortiClient = c.AccessRight2Str(int(data["device-forticlient"].(float64)))
	}
	if data["device-fortiswitch"] != nil {
		out.DeviceFortiSwitch = c.AccessRight2Str(int(data["device-fortiswitch"].(float64)))
	}
	if data["vpn-manager"] != nil {
		out.VpnManager = c.AccessRight2Str(int(data["vpn-manager"].(float64)))
	}
	if data["log-viewer"] != nil {
		out.LogViewer = c.AccessRight2Str(int(data["log-viewer"].(float64)))
	}
	if data["device-forticlient"] != nil {
		out.DeviceFortiClient = c.AccessRight2Str(int(data["device-forticlient"].(float64)))
	}

	if data["fgd_center"] != nil {
		out.FgdCenter = c.AccessRight2Str(int(data["fgd_center"].(float64)))
	}
	if data["fgd-center-advanced"] != nil {
		out.FgdCenterAdvanced = c.AccessRight2Str(int(data["fgd-center-advanced"].(float64)))
	}
	if data["fgd-center-fmw-mgmt"] != nil {
		out.FgdCenterFmwMgmt = c.AccessRight2Str(int(data["fgd-center-fmw-mgmt"].(float64)))
	}
	if data["fgd-center-licensing"] != nil {
		out.FgdCenterLicensing = c.AccessRight2Str(int(data["fgd-center-licensing"].(float64)))
	}

	if data["device-manager"] != nil {
		out.DeviceManager = c.AccessRight2Str(int(data["device-manager"].(float64)))
	}
	if data["device-op"] != nil {
		out.DeviceOp = c.AccessRight2Str(int(data["device-op"].(float64)))
	}
	if data["config-retrieve"] != nil {
		out.ConfigRetrieve = c.AccessRight2Str(int(data["config-retrieve"].(float64)))
	}
	if data["config-revert"] != nil {
		out.ConfigRevert = c.AccessRight2Str(int(data["config-revert"].(float64)))
	}
	if data["device-revision-deletion"] != nil {
		out.DeviceRevisionDeletion = c.AccessRight2Str(int(data["device-revision-deletion"].(float64)))
	}
	if data["term-access"] != nil {
		out.TermAccess = c.AccessRight2Str(int(data["term-access"].(float64)))
	}
	if data["device-config"] != nil {
		out.DeviceConfig = c.AccessRight2Str(int(data["device-config"].(float64)))
	}
	if data["device-profile"] != nil {
		out.DeviceProfile = c.AccessRight2Str(int(data["device-profile"].(float64)))
	}
	if data["device-wan-link-load-balance"] != nil {
		out.DeviceWanLinkLoadBalance = c.AccessRight2Str(int(data["device-wan-link-load-balance"].(float64)))
	}

	if data["policy-objects"] != nil {
		out.PolicyObjects = c.AccessRight2Str(int(data["policy-objects"].(float64)))
	}
	if data["global-policy-packages"] != nil {
		out.GlobalPolicyPackages = c.AccessRight2Str(int(data["global-policy-packages"].(float64)))
	}
	if data["assignment"] != nil {
		out.Assignment = c.AccessRight2Str(int(data["assignment"].(float64)))
	}
	if data["adom-policy-packages"] != nil {
		out.AdomPolicyPackages = c.AccessRight2Str(int(data["adom-policy-packages"].(float64)))
	}
	if data["consistency-check"] != nil {
		out.ConsistencyCheck = c.AccessRight2Str(int(data["consistency-check"].(float64)))
	}
	if data["set-install-targets"] != nil {
		out.SetInstallTargets = c.AccessRight2Str(int(data["set-install-targets"].(float64)))
	}

	return
}

func (c *FortiMngClient) DeleteSystemAdminProfiles(id string) (err error) {
	defer c.Trace("DeleteSystemAdminProfiles")()

	p := map[string]interface{}{
		"url": "/cli/global/system/admin/profile/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteSystemAdminProfiles failed :%s", err)
		return
	}

	return
}
