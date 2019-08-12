package fortimngclient

import (
	"fmt"
)

type JSONSysAdminProfiles struct {
	ProfileId     string `json:"profileid"`
	Description   string `json:"description"`
	DeviceManager string `json:"device-manager"`
	//	DeviceOp      string `json:"device-op"`
	/*
		AdomLock           string `json:"adom-lock"`
		AdomPolicyPackages string `json:"adom-policy-packages"`
		AdomSwitch         string `json:"adom-switch"`
		AppFilter          string `json:"app-filter"`
		Assignment         string `json:"assignment"`
		ChangePassword     string `json:"change-password"`
		ConfigRetrieve     string `json:"config-retrieve"`
		ConfigRevert       string `json:"config-revert"`
		ConsistencyCheck   string `json:"consistency-check"`
				DataMask string `json:"datamask"`
			   "datamask-custom-fields": [
			     {
			       "field-category": [
			         "log"
			       ],
			       "field-name": "string",
			       "field-status": "enable",
			       "field-type": "string"
			     }
			   ],
			   "datamask-custom-priority": "disable",
			   "datamask-fields": [
			     "user"
			   ],
			   "datamask-key": [
			     "ENC MzI1Nzc3MjAyNTg1Njg0NNKOn5kCfNawE/VnDbtMpWXduJpvaREIOxBK4PNmJmqeRwgB9loHz7FqcMzTT5DrD50rb65MQrxNOiuHZ7eM/qmDuMiCMym4F4p2r819t/tQ0emIgt9MTrccrMAZN5Mv9Kmkp5KFjedrsRnbofB058Bi9VBs"
			   ],
		DeployManagement         string `json:"deploy-management"`
		DeviceAp                 string `json:"device-ap"`
		DeviceConfig             string `json:"device-config"`
		DeviceForticlient        string `json:"device-forticlient"`
		DeviceFortiswithc        string `json:"device-fortiswitch"`
		DevicePolicyPackageLock  string `json:"device-policy-package-lock"`
		DeviceProfile            string `json:"device-profile"`
		DeviceRevisionDeletion   string `json:"device-revision-deletion"`
		DeviceWanLinkLoadBalance string `json:"device-wan-link-load-balance"`
		EventManagement          string `json:"event-management"`
		FgdCenterAdvanced        string `json:"fgd-center-advanced"`
		FgdCenterFmwMgmt         string `json:"fgd-center-fmw-mgmt"`
		FgdCenterLicensing       string `json:"fgd-center-licensing"`
		FgdCenter                string `json:"fgd_center"`
		GlobalPolicyPackages     string `json:"global-policy-packages"`
		ImportPolicyPackages     string `json:"import-policy-packages"`
		IntfMapping              string `json:"intf-mapping"`
		IpsFilter                string `json:"ips-filter"`
		LogViewer                string `json:"log-viewer"`
		PolicyObjects            string `json:"policy-objects"`
		ReadPasswd               string `json:"read-passwd"`
		RealtimeMonitor          string `json:"realtime-monitor"`
		ReportViewer             string `json:"report-viewer"`
		Scope                    string `json:"scope"`
		SetInstallTargets        string `json:"set-install-targets"`
		SystemSetting            string `json:"system-setting"`
		TermAccess               string `json:"term-access"`
		Type                     string `json:"type"`
		VpnManager               string `json:"vpn-manager"`
		WebFilter                string `json:"web-filter"`
	*/
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

	out = &JSONSysAdminProfiles{}
	if data["profileid"] != nil {
		out.ProfileId = data["profileid"].(string)
	}
	if data["description"] != nil {
		out.Description = data["description"].(string)
	}
	if data["device-manager"] != nil {
		out.DeviceManager = c.AccessRight2Str(int(data["device-manager"].(float64)))
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
