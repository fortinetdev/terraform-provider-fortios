package fmgclient

import (
	"fmt"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

// JSONSysAdminProfileGenernal contains the params for creating system admin profiles with genernal part
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

// JSONSysAdminProfileFortiGuard contains the params for creating system admin profiles with fortiguard part
type JSONSysAdminProfileFortiGuard struct {
	FgdCenter          string `json:"fgd_center"`
	FgdCenterAdvanced  string `json:"fgd-center-advanced"`
	FgdCenterFmwMgmt   string `json:"fgd-center-fmw-mgmt"`
	FgdCenterLicensing string `json:"fgd-center-licensing"`
}

// JSONSysAdminProfileDeviceManager contains the params for creating system admin profiles with devicemanager part
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

// JSONSysAdminProfilePolicyObject contains the params for creating system admin profiles with policy object part
type JSONSysAdminProfilePolicyObject struct {
	PolicyObjects        string `json:"policy-objects"`
	GlobalPolicyPackages string `json:"global-policy-packages"`
	Assignment           string `json:"assignment"`
	AdomPolicyPackages   string `json:"adom-policy-packages"`
	ConsistencyCheck     string `json:"consistency-check"`
	SetInstallTargets    string `json:"set-install-targets"`
}

// JSONSysAdminProfiles contains the params for creating system admin profiles
type JSONSysAdminProfiles struct {
	*JSONSysAdminProfileGenernal
	*JSONSysAdminProfileFortiGuard
	*JSONSysAdminProfileDeviceManager
	*JSONSysAdminProfilePolicyObject
}

// CreateUpdateSystemAdminProfiles is for creating or updating the system admin profiles
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateSystemAdminProfiles(params *JSONSysAdminProfiles, method string) (err error) {
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

// ReadSystemAdminProfiles is for reading the system admin profiles
// Input:
//   @id: admin profile id
// Output:
//   @out: admin profile infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadSystemAdminProfiles(id string) (out *JSONSysAdminProfiles, err error) {
	defer c.Trace("ReadSystemAdminProfiles")()

	p := map[string]interface{}{
		"url": "/cli/global/system/admin/profile/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemAdminProfile failed: %s", err)
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
		out.SystemSetting = util.AccessRight2Str(int(data["system-setting"].(float64)))
	}
	if data["adom-switch"] != nil {
		out.AdomSwitch = util.AccessRight2Str(int(data["adom-switch"].(float64)))
	}
	if data["deploy-management"] != nil {
		out.DeployManagement = util.AccessRight2Str(int(data["deploy-management"].(float64)))
	}
	if data["import-policy-packages"] != nil {
		out.ImportPolicyPackages = util.AccessRight2Str(int(data["import-policy-packages"].(float64)))
	}
	if data["intf-mapping"] != nil {
		out.IntfMapping = util.AccessRight2Str(int(data["intf-mapping"].(float64)))
	}
	if data["device-ap"] != nil {
		out.DeviceAp = util.AccessRight2Str(int(data["device-ap"].(float64)))
	}
	if data["device-forticlient"] != nil {
		out.DeviceFortiClient = util.AccessRight2Str(int(data["device-forticlient"].(float64)))
	}
	if data["device-fortiswitch"] != nil {
		out.DeviceFortiSwitch = util.AccessRight2Str(int(data["device-fortiswitch"].(float64)))
	}
	if data["vpn-manager"] != nil {
		out.VpnManager = util.AccessRight2Str(int(data["vpn-manager"].(float64)))
	}
	if data["log-viewer"] != nil {
		out.LogViewer = util.AccessRight2Str(int(data["log-viewer"].(float64)))
	}
	if data["device-forticlient"] != nil {
		out.DeviceFortiClient = util.AccessRight2Str(int(data["device-forticlient"].(float64)))
	}

	if data["fgd_center"] != nil {
		out.FgdCenter = util.AccessRight2Str(int(data["fgd_center"].(float64)))
	}
	if data["fgd-center-advanced"] != nil {
		out.FgdCenterAdvanced = util.AccessRight2Str(int(data["fgd-center-advanced"].(float64)))
	}
	if data["fgd-center-fmw-mgmt"] != nil {
		out.FgdCenterFmwMgmt = util.AccessRight2Str(int(data["fgd-center-fmw-mgmt"].(float64)))
	}
	if data["fgd-center-licensing"] != nil {
		out.FgdCenterLicensing = util.AccessRight2Str(int(data["fgd-center-licensing"].(float64)))
	}

	if data["device-manager"] != nil {
		out.DeviceManager = util.AccessRight2Str(int(data["device-manager"].(float64)))
	}
	if data["device-op"] != nil {
		out.DeviceOp = util.AccessRight2Str(int(data["device-op"].(float64)))
	}
	if data["config-retrieve"] != nil {
		out.ConfigRetrieve = util.AccessRight2Str(int(data["config-retrieve"].(float64)))
	}
	if data["config-revert"] != nil {
		out.ConfigRevert = util.AccessRight2Str(int(data["config-revert"].(float64)))
	}
	if data["device-revision-deletion"] != nil {
		out.DeviceRevisionDeletion = util.AccessRight2Str(int(data["device-revision-deletion"].(float64)))
	}
	if data["term-access"] != nil {
		out.TermAccess = util.AccessRight2Str(int(data["term-access"].(float64)))
	}
	if data["device-config"] != nil {
		out.DeviceConfig = util.AccessRight2Str(int(data["device-config"].(float64)))
	}
	if data["device-profile"] != nil {
		out.DeviceProfile = util.AccessRight2Str(int(data["device-profile"].(float64)))
	}
	if data["device-wan-link-load-balance"] != nil {
		out.DeviceWanLinkLoadBalance = util.AccessRight2Str(int(data["device-wan-link-load-balance"].(float64)))
	}

	if data["policy-objects"] != nil {
		out.PolicyObjects = util.AccessRight2Str(int(data["policy-objects"].(float64)))
	}
	if data["global-policy-packages"] != nil {
		out.GlobalPolicyPackages = util.AccessRight2Str(int(data["global-policy-packages"].(float64)))
	}
	if data["assignment"] != nil {
		out.Assignment = util.AccessRight2Str(int(data["assignment"].(float64)))
	}
	if data["adom-policy-packages"] != nil {
		out.AdomPolicyPackages = util.AccessRight2Str(int(data["adom-policy-packages"].(float64)))
	}
	if data["consistency-check"] != nil {
		out.ConsistencyCheck = util.AccessRight2Str(int(data["consistency-check"].(float64)))
	}
	if data["set-install-targets"] != nil {
		out.SetInstallTargets = util.AccessRight2Str(int(data["set-install-targets"].(float64)))
	}

	return
}

// DeleteSystemAdminProfiles is for deleting the specific admin profile
// Input:
//   @id: admin profile id
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteSystemAdminProfiles(id string) (err error) {
	defer c.Trace("DeleteSystemAdminProfiles")()

	p := map[string]interface{}{
		"url": "/cli/global/system/admin/profile/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteSystemAdminProfiles failed: %s", err)
		return
	}

	return
}
