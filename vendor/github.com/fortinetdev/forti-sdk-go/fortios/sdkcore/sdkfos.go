// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Description: SDK for FortiOS Provider

package forticlient

type creatUpdateOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}


// CreateSystemVdom API operation for FortiOS creates a new Vdom.
// Returns the index value of the Vdom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVdom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vdom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVdom API operation for FortiOS updates the specified Vdom.
// Returns the index value of the Vdom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdom API operation for FortiOS deletes the specified Vdom.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVdom API operation for FortiOS gets the Vdom
// with the specified index value.
// Returns the requested Vdom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemSaml API operation for FortiOS updates the specified Saml.
// Returns the index value of the Saml and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSaml(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/saml"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSaml API operation for FortiOS deletes the specified Saml.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSaml(mkey string) (err error) {

	//No unset API for system - saml
	return
}

// ReadSystemSaml API operation for FortiOS gets the Saml
// with the specified index value.
// Returns the requested Saml value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSaml(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/saml"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemGeneve API operation for FortiOS creates a new Geneve.
// Returns the index value of the Geneve and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - geneve chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemGeneve(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/geneve"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemGeneve API operation for FortiOS updates the specified Geneve.
// Returns the index value of the Geneve and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - geneve chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGeneve(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/geneve"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemGeneve API operation for FortiOS deletes the specified Geneve.
// Returns error for service API and SDK errors.
// See the system - geneve chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGeneve(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/geneve"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemGeneve API operation for FortiOS gets the Geneve
// with the specified index value.
// Returns the requested Geneve value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - geneve chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGeneve(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/geneve"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemGlobal API operation for FortiOS updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemGlobal API operation for FortiOS deletes the specified Global.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGlobal(mkey string) (err error) {

	//No unset API for system - global
	return
}

// ReadSystemGlobal API operation for FortiOS gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemAccprofile API operation for FortiOS creates a new Accprofile.
// Returns the index value of the Accprofile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAccprofile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/accprofile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAccprofile API operation for FortiOS updates the specified Accprofile.
// Returns the index value of the Accprofile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAccprofile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAccprofile API operation for FortiOS deletes the specified Accprofile.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAccprofile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAccprofile API operation for FortiOS gets the Accprofile
// with the specified index value.
// Returns the requested Accprofile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAccprofile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemVdomLink API operation for FortiOS creates a new Vdom Link.
// Returns the index value of the Vdom Link and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVdomLink(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vdom-link"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVdomLink API operation for FortiOS updates the specified Vdom Link.
// Returns the index value of the Vdom Link and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomLink(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-link"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomLink API operation for FortiOS deletes the specified Vdom Link.
// Returns error for service API and SDK errors.
// See the system - vdom-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomLink(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vdom-link"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVdomLink API operation for FortiOS gets the Vdom Link
// with the specified index value.
// Returns the requested Vdom Link value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomLink(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-link"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSwitchInterface API operation for FortiOS creates a new Switch Interface.
// Returns the index value of the Switch Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - switch-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSwitchInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/switch-interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSwitchInterface API operation for FortiOS updates the specified Switch Interface.
// Returns the index value of the Switch Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - switch-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSwitchInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/switch-interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSwitchInterface API operation for FortiOS deletes the specified Switch Interface.
// Returns error for service API and SDK errors.
// See the system - switch-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSwitchInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/switch-interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSwitchInterface API operation for FortiOS gets the Switch Interface
// with the specified index value.
// Returns the requested Switch Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - switch-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSwitchInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/switch-interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemObjectTagging API operation for FortiOS creates a new Object Tagging.
// Returns the index value of the Object Tagging and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tagging chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemObjectTagging(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/object-tagging"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemObjectTagging API operation for FortiOS updates the specified Object Tagging.
// Returns the index value of the Object Tagging and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tagging chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemObjectTagging(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/object-tagging"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemObjectTagging API operation for FortiOS deletes the specified Object Tagging.
// Returns error for service API and SDK errors.
// See the system - object-tagging chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemObjectTagging(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/object-tagging"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemObjectTagging API operation for FortiOS gets the Object Tagging
// with the specified index value.
// Returns the requested Object Tagging value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tagging chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemObjectTagging(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/object-tagging"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemInterface API operation for FortiOS creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemInterface API operation for FortiOS updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemInterface API operation for FortiOS deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemInterface API operation for FortiOS gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemPasswordPolicy API operation for FortiOS updates the specified Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPasswordPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/password-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPasswordPolicy API operation for FortiOS deletes the specified Password Policy.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPasswordPolicy(mkey string) (err error) {

	//No unset API for system - password-policy
	return
}

// ReadSystemPasswordPolicy API operation for FortiOS gets the Password Policy
// with the specified index value.
// Returns the requested Password Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPasswordPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/password-policy"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemPasswordPolicyGuestAdmin API operation for FortiOS updates the specified Password Policy Guest Admin.
// Returns the index value of the Password Policy Guest Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy-guest-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPasswordPolicyGuestAdmin(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/password-policy-guest-admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPasswordPolicyGuestAdmin API operation for FortiOS deletes the specified Password Policy Guest Admin.
// Returns error for service API and SDK errors.
// See the system - password-policy-guest-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPasswordPolicyGuestAdmin(mkey string) (err error) {

	//No unset API for system - password-policy-guest-admin
	return
}

// ReadSystemPasswordPolicyGuestAdmin API operation for FortiOS gets the Password Policy Guest Admin
// with the specified index value.
// Returns the requested Password Policy Guest Admin value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy-guest-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPasswordPolicyGuestAdmin(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/password-policy-guest-admin"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemSmsServer API operation for FortiOS creates a new Sms Server.
// Returns the index value of the Sms Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sms-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSmsServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/sms-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSmsServer API operation for FortiOS updates the specified Sms Server.
// Returns the index value of the Sms Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sms-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSmsServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sms-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSmsServer API operation for FortiOS deletes the specified Sms Server.
// Returns error for service API and SDK errors.
// See the system - sms-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSmsServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/sms-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSmsServer API operation for FortiOS gets the Sms Server
// with the specified index value.
// Returns the requested Sms Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sms-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSmsServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sms-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemCustomLanguage API operation for FortiOS creates a new Custom Language.
// Returns the index value of the Custom Language and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - custom-language chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCustomLanguage(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/custom-language"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemCustomLanguage API operation for FortiOS updates the specified Custom Language.
// Returns the index value of the Custom Language and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - custom-language chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCustomLanguage(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/custom-language"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCustomLanguage API operation for FortiOS deletes the specified Custom Language.
// Returns error for service API and SDK errors.
// See the system - custom-language chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCustomLanguage(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/custom-language"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemCustomLanguage API operation for FortiOS gets the Custom Language
// with the specified index value.
// Returns the requested Custom Language value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - custom-language chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCustomLanguage(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/custom-language"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAdmin API operation for FortiOS creates a new Admin.
// Returns the index value of the Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdmin(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAdmin API operation for FortiOS updates the specified Admin.
// Returns the index value of the Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdmin(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAdmin API operation for FortiOS deletes the specified Admin.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdmin(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAdmin API operation for FortiOS gets the Admin
// with the specified index value.
// Returns the requested Admin value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdmin(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemApiUser API operation for FortiOS creates a new Api User.
// Returns the index value of the Api User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemApiUser(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/api-user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemApiUser API operation for FortiOS updates the specified Api User.
// Returns the index value of the Api User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemApiUser(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemApiUser API operation for FortiOS deletes the specified Api User.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemApiUser(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemApiUser API operation for FortiOS gets the Api User
// with the specified index value.
// Returns the requested Api User value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemApiUser(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSsoAdmin API operation for FortiOS creates a new Sso Admin.
// Returns the index value of the Sso Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sso-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSsoAdmin(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/sso-admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSsoAdmin API operation for FortiOS updates the specified Sso Admin.
// Returns the index value of the Sso Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sso-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSsoAdmin(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sso-admin"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSsoAdmin API operation for FortiOS deletes the specified Sso Admin.
// Returns error for service API and SDK errors.
// See the system - sso-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSsoAdmin(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/sso-admin"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSsoAdmin API operation for FortiOS gets the Sso Admin
// with the specified index value.
// Returns the requested Sso Admin value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sso-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSsoAdmin(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sso-admin"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemSettings API operation for FortiOS updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSettings API operation for FortiOS deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the system - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSettings(mkey string) (err error) {

	//No unset API for system - settings
	return
}

// ReadSystemSettings API operation for FortiOS gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemSitTunnel API operation for FortiOS creates a new Sit Tunnel.
// Returns the index value of the Sit Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sit-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSitTunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/sit-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSitTunnel API operation for FortiOS updates the specified Sit Tunnel.
// Returns the index value of the Sit Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sit-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSitTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sit-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSitTunnel API operation for FortiOS deletes the specified Sit Tunnel.
// Returns error for service API and SDK errors.
// See the system - sit-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSitTunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/sit-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSitTunnel API operation for FortiOS gets the Sit Tunnel
// with the specified index value.
// Returns the requested Sit Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sit-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSitTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sit-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemFssoPolling API operation for FortiOS updates the specified Fsso Polling.
// Returns the index value of the Fsso Polling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFssoPolling(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fsso-polling"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFssoPolling API operation for FortiOS deletes the specified Fsso Polling.
// Returns error for service API and SDK errors.
// See the system - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFssoPolling(mkey string) (err error) {

	//No unset API for system - fsso-polling
	return
}

// ReadSystemFssoPolling API operation for FortiOS gets the Fsso Polling
// with the specified index value.
// Returns the requested Fsso Polling value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFssoPolling(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fsso-polling"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemHa API operation for FortiOS updates the specified Ha.
// Returns the index value of the Ha and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHa(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ha"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemHa API operation for FortiOS deletes the specified Ha.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHa(mkey string) (err error) {

	//No unset API for system - ha
	return
}

// ReadSystemHa API operation for FortiOS gets the Ha
// with the specified index value.
// Returns the requested Ha value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHa(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ha"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemHaMonitor API operation for FortiOS updates the specified Ha Monitor.
// Returns the index value of the Ha Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHaMonitor(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ha-monitor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemHaMonitor API operation for FortiOS deletes the specified Ha Monitor.
// Returns error for service API and SDK errors.
// See the system - ha-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHaMonitor(mkey string) (err error) {

	//No unset API for system - ha-monitor
	return
}

// ReadSystemHaMonitor API operation for FortiOS gets the Ha Monitor
// with the specified index value.
// Returns the requested Ha Monitor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHaMonitor(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ha-monitor"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemStorage API operation for FortiOS creates a new Storage.
// Returns the index value of the Storage and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - storage chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemStorage(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/storage"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemStorage API operation for FortiOS updates the specified Storage.
// Returns the index value of the Storage and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - storage chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemStorage(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/storage"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemStorage API operation for FortiOS deletes the specified Storage.
// Returns error for service API and SDK errors.
// See the system - storage chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemStorage(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/storage"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemStorage API operation for FortiOS gets the Storage
// with the specified index value.
// Returns the requested Storage value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - storage chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemStorage(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/storage"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemDedicatedMgmt API operation for FortiOS updates the specified Dedicated Mgmt.
// Returns the index value of the Dedicated Mgmt and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dedicated-mgmt chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDedicatedMgmt(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dedicated-mgmt"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDedicatedMgmt API operation for FortiOS deletes the specified Dedicated Mgmt.
// Returns error for service API and SDK errors.
// See the system - dedicated-mgmt chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDedicatedMgmt(mkey string) (err error) {

	//No unset API for system - dedicated-mgmt
	return
}

// ReadSystemDedicatedMgmt API operation for FortiOS gets the Dedicated Mgmt
// with the specified index value.
// Returns the requested Dedicated Mgmt value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dedicated-mgmt chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDedicatedMgmt(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dedicated-mgmt"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemArpTable API operation for FortiOS creates a new Arp Table.
// Returns the index value of the Arp Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemArpTable(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/arp-table"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemArpTable API operation for FortiOS updates the specified Arp Table.
// Returns the index value of the Arp Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemArpTable(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemArpTable API operation for FortiOS deletes the specified Arp Table.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemArpTable(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemArpTable API operation for FortiOS gets the Arp Table
// with the specified index value.
// Returns the requested Arp Table value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemArpTable(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpv6NeighborCache API operation for FortiOS creates a new Ipv6 Neighbor Cache.
// Returns the index value of the Ipv6 Neighbor Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpv6NeighborCache(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpv6NeighborCache API operation for FortiOS updates the specified Ipv6 Neighbor Cache.
// Returns the index value of the Ipv6 Neighbor Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpv6NeighborCache(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpv6NeighborCache API operation for FortiOS deletes the specified Ipv6 Neighbor Cache.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpv6NeighborCache(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpv6NeighborCache API operation for FortiOS gets the Ipv6 Neighbor Cache
// with the specified index value.
// Returns the requested Ipv6 Neighbor Cache value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpv6NeighborCache(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemDns API operation for FortiOS updates the specified Dns.
// Returns the index value of the Dns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDns(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDns API operation for FortiOS deletes the specified Dns.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDns(mkey string) (err error) {

	//No unset API for system - dns
	return
}

// ReadSystemDns API operation for FortiOS gets the Dns
// with the specified index value.
// Returns the requested Dns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDns(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemDdns API operation for FortiOS creates a new Ddns.
// Returns the index value of the Ddns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ddns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDdns(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ddns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDdns API operation for FortiOS updates the specified Ddns.
// Returns the index value of the Ddns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ddns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDdns(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ddns"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDdns API operation for FortiOS deletes the specified Ddns.
// Returns error for service API and SDK errors.
// See the system - ddns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDdns(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ddns"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDdns API operation for FortiOS gets the Ddns
// with the specified index value.
// Returns the requested Ddns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ddns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDdns(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ddns"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemSflow API operation for FortiOS updates the specified Sflow.
// Returns the index value of the Sflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSflow API operation for FortiOS deletes the specified Sflow.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSflow(mkey string) (err error) {

	//No unset API for system - sflow
	return
}

// ReadSystemSflow API operation for FortiOS gets the Sflow
// with the specified index value.
// Returns the requested Sflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemVdomSflow API operation for FortiOS updates the specified Vdom Sflow.
// Returns the index value of the Vdom Sflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomSflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-sflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomSflow API operation for FortiOS deletes the specified Vdom Sflow.
// Returns error for service API and SDK errors.
// See the system - vdom-sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomSflow(mkey string) (err error) {

	//No unset API for system - vdom-sflow
	return
}

// ReadSystemVdomSflow API operation for FortiOS gets the Vdom Sflow
// with the specified index value.
// Returns the requested Vdom Sflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomSflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-sflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemNetflow API operation for FortiOS updates the specified Netflow.
// Returns the index value of the Netflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNetflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/netflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemNetflow API operation for FortiOS deletes the specified Netflow.
// Returns error for service API and SDK errors.
// See the system - netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNetflow(mkey string) (err error) {

	//No unset API for system - netflow
	return
}

// ReadSystemNetflow API operation for FortiOS gets the Netflow
// with the specified index value.
// Returns the requested Netflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNetflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/netflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemVdomNetflow API operation for FortiOS updates the specified Vdom Netflow.
// Returns the index value of the Vdom Netflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomNetflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-netflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomNetflow API operation for FortiOS deletes the specified Vdom Netflow.
// Returns error for service API and SDK errors.
// See the system - vdom-netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomNetflow(mkey string) (err error) {

	//No unset API for system - vdom-netflow
	return
}

// ReadSystemVdomNetflow API operation for FortiOS gets the Vdom Netflow
// with the specified index value.
// Returns the requested Vdom Netflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomNetflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-netflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemVdomDns API operation for FortiOS updates the specified Vdom Dns.
// Returns the index value of the Vdom Dns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomDns(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-dns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomDns API operation for FortiOS deletes the specified Vdom Dns.
// Returns error for service API and SDK errors.
// See the system - vdom-dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomDns(mkey string) (err error) {

	//No unset API for system - vdom-dns
	return
}

// ReadSystemVdomDns API operation for FortiOS gets the Vdom Dns
// with the specified index value.
// Returns the requested Vdom Dns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomDns(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-dns"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemReplacemsgImage API operation for FortiOS creates a new Replacemsg Image.
// Returns the index value of the Replacemsg Image and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-image chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgImage(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/replacemsg-image"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgImage API operation for FortiOS updates the specified Replacemsg Image.
// Returns the index value of the Replacemsg Image and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-image chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgImage(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/replacemsg-image"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgImage API operation for FortiOS deletes the specified Replacemsg Image.
// Returns error for service API and SDK errors.
// See the system - replacemsg-image chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgImage(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/replacemsg-image"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgImage API operation for FortiOS gets the Replacemsg Image
// with the specified index value.
// Returns the requested Replacemsg Image value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-image chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgImage(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/replacemsg-image"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgMail API operation for FortiOS creates a new Mail.
// Returns the index value of the Mail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - mail chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgMail(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/mail"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgMail API operation for FortiOS updates the specified Mail.
// Returns the index value of the Mail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - mail chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgMail(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/mail"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgMail API operation for FortiOS deletes the specified Mail.
// Returns error for service API and SDK errors.
// See the system.replacemsg - mail chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgMail(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/mail"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgMail API operation for FortiOS gets the Mail
// with the specified index value.
// Returns the requested Mail value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - mail chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgMail(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/mail"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgHttp API operation for FortiOS creates a new Http.
// Returns the index value of the Http and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - http chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgHttp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/http"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgHttp API operation for FortiOS updates the specified Http.
// Returns the index value of the Http and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - http chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgHttp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/http"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgHttp API operation for FortiOS deletes the specified Http.
// Returns error for service API and SDK errors.
// See the system.replacemsg - http chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgHttp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/http"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgHttp API operation for FortiOS gets the Http
// with the specified index value.
// Returns the requested Http value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - http chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgHttp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/http"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgWebproxy API operation for FortiOS creates a new Webproxy.
// Returns the index value of the Webproxy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - webproxy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgWebproxy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/webproxy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgWebproxy API operation for FortiOS updates the specified Webproxy.
// Returns the index value of the Webproxy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - webproxy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgWebproxy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/webproxy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgWebproxy API operation for FortiOS deletes the specified Webproxy.
// Returns error for service API and SDK errors.
// See the system.replacemsg - webproxy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgWebproxy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/webproxy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgWebproxy API operation for FortiOS gets the Webproxy
// with the specified index value.
// Returns the requested Webproxy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - webproxy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgWebproxy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/webproxy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgFtp API operation for FortiOS creates a new Ftp.
// Returns the index value of the Ftp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - ftp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgFtp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/ftp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgFtp API operation for FortiOS updates the specified Ftp.
// Returns the index value of the Ftp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - ftp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgFtp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/ftp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgFtp API operation for FortiOS deletes the specified Ftp.
// Returns error for service API and SDK errors.
// See the system.replacemsg - ftp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgFtp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/ftp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgFtp API operation for FortiOS gets the Ftp
// with the specified index value.
// Returns the requested Ftp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - ftp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgFtp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/ftp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgNntp API operation for FortiOS creates a new Nntp.
// Returns the index value of the Nntp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - nntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgNntp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/nntp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgNntp API operation for FortiOS updates the specified Nntp.
// Returns the index value of the Nntp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - nntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgNntp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/nntp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgNntp API operation for FortiOS deletes the specified Nntp.
// Returns error for service API and SDK errors.
// See the system.replacemsg - nntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgNntp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/nntp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgNntp API operation for FortiOS gets the Nntp
// with the specified index value.
// Returns the requested Nntp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - nntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgNntp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/nntp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgFortiguardWf API operation for FortiOS creates a new Fortiguard Wf.
// Returns the index value of the Fortiguard Wf and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - fortiguard-wf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgFortiguardWf(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/fortiguard-wf"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgFortiguardWf API operation for FortiOS updates the specified Fortiguard Wf.
// Returns the index value of the Fortiguard Wf and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - fortiguard-wf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgFortiguardWf(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/fortiguard-wf"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgFortiguardWf API operation for FortiOS deletes the specified Fortiguard Wf.
// Returns error for service API and SDK errors.
// See the system.replacemsg - fortiguard-wf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgFortiguardWf(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/fortiguard-wf"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgFortiguardWf API operation for FortiOS gets the Fortiguard Wf
// with the specified index value.
// Returns the requested Fortiguard Wf value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - fortiguard-wf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgFortiguardWf(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/fortiguard-wf"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgSpam API operation for FortiOS creates a new Spam.
// Returns the index value of the Spam and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - spam chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgSpam(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/spam"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgSpam API operation for FortiOS updates the specified Spam.
// Returns the index value of the Spam and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - spam chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgSpam(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/spam"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgSpam API operation for FortiOS deletes the specified Spam.
// Returns error for service API and SDK errors.
// See the system.replacemsg - spam chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgSpam(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/spam"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgSpam API operation for FortiOS gets the Spam
// with the specified index value.
// Returns the requested Spam value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - spam chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgSpam(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/spam"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgAlertmail API operation for FortiOS creates a new Alertmail.
// Returns the index value of the Alertmail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - alertmail chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgAlertmail(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/alertmail"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgAlertmail API operation for FortiOS updates the specified Alertmail.
// Returns the index value of the Alertmail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - alertmail chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgAlertmail(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/alertmail"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgAlertmail API operation for FortiOS deletes the specified Alertmail.
// Returns error for service API and SDK errors.
// See the system.replacemsg - alertmail chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgAlertmail(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/alertmail"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgAlertmail API operation for FortiOS gets the Alertmail
// with the specified index value.
// Returns the requested Alertmail value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - alertmail chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgAlertmail(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/alertmail"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgAdmin API operation for FortiOS creates a new Admin.
// Returns the index value of the Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgAdmin(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgAdmin API operation for FortiOS updates the specified Admin.
// Returns the index value of the Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgAdmin(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/admin"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgAdmin API operation for FortiOS deletes the specified Admin.
// Returns error for service API and SDK errors.
// See the system.replacemsg - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgAdmin(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/admin"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgAdmin API operation for FortiOS gets the Admin
// with the specified index value.
// Returns the requested Admin value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgAdmin(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/admin"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgAuth API operation for FortiOS creates a new Auth.
// Returns the index value of the Auth and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - auth chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgAuth(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/auth"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgAuth API operation for FortiOS updates the specified Auth.
// Returns the index value of the Auth and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - auth chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgAuth(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/auth"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgAuth API operation for FortiOS deletes the specified Auth.
// Returns error for service API and SDK errors.
// See the system.replacemsg - auth chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgAuth(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/auth"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgAuth API operation for FortiOS gets the Auth
// with the specified index value.
// Returns the requested Auth value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - auth chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgAuth(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/auth"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgSslvpn API operation for FortiOS creates a new Sslvpn.
// Returns the index value of the Sslvpn and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - sslvpn chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgSslvpn(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/sslvpn"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgSslvpn API operation for FortiOS updates the specified Sslvpn.
// Returns the index value of the Sslvpn and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - sslvpn chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgSslvpn(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/sslvpn"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgSslvpn API operation for FortiOS deletes the specified Sslvpn.
// Returns error for service API and SDK errors.
// See the system.replacemsg - sslvpn chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgSslvpn(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/sslvpn"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgSslvpn API operation for FortiOS gets the Sslvpn
// with the specified index value.
// Returns the requested Sslvpn value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - sslvpn chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgSslvpn(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/sslvpn"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgEc API operation for FortiOS creates a new Ec.
// Returns the index value of the Ec and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - ec chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgEc(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/ec"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgEc API operation for FortiOS updates the specified Ec.
// Returns the index value of the Ec and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - ec chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgEc(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/ec"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgEc API operation for FortiOS deletes the specified Ec.
// Returns error for service API and SDK errors.
// See the system.replacemsg - ec chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgEc(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/ec"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgEc API operation for FortiOS gets the Ec
// with the specified index value.
// Returns the requested Ec value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - ec chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgEc(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/ec"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgDeviceDetectionPortal API operation for FortiOS creates a new Device Detection Portal.
// Returns the index value of the Device Detection Portal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - device-detection-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgDeviceDetectionPortal(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/device-detection-portal"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgDeviceDetectionPortal API operation for FortiOS updates the specified Device Detection Portal.
// Returns the index value of the Device Detection Portal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - device-detection-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgDeviceDetectionPortal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/device-detection-portal"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgDeviceDetectionPortal API operation for FortiOS deletes the specified Device Detection Portal.
// Returns error for service API and SDK errors.
// See the system.replacemsg - device-detection-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgDeviceDetectionPortal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/device-detection-portal"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgDeviceDetectionPortal API operation for FortiOS gets the Device Detection Portal
// with the specified index value.
// Returns the requested Device Detection Portal value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - device-detection-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgDeviceDetectionPortal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/device-detection-portal"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgNacQuar API operation for FortiOS creates a new Nac Quar.
// Returns the index value of the Nac Quar and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - nac-quar chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgNacQuar(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/nac-quar"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgNacQuar API operation for FortiOS updates the specified Nac Quar.
// Returns the index value of the Nac Quar and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - nac-quar chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgNacQuar(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/nac-quar"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgNacQuar API operation for FortiOS deletes the specified Nac Quar.
// Returns error for service API and SDK errors.
// See the system.replacemsg - nac-quar chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgNacQuar(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/nac-quar"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgNacQuar API operation for FortiOS gets the Nac Quar
// with the specified index value.
// Returns the requested Nac Quar value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - nac-quar chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgNacQuar(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/nac-quar"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgTrafficQuota API operation for FortiOS creates a new Traffic Quota.
// Returns the index value of the Traffic Quota and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - traffic-quota chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgTrafficQuota(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/traffic-quota"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgTrafficQuota API operation for FortiOS updates the specified Traffic Quota.
// Returns the index value of the Traffic Quota and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - traffic-quota chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgTrafficQuota(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/traffic-quota"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgTrafficQuota API operation for FortiOS deletes the specified Traffic Quota.
// Returns error for service API and SDK errors.
// See the system.replacemsg - traffic-quota chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgTrafficQuota(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/traffic-quota"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgTrafficQuota API operation for FortiOS gets the Traffic Quota
// with the specified index value.
// Returns the requested Traffic Quota value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - traffic-quota chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgTrafficQuota(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/traffic-quota"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgUtm API operation for FortiOS creates a new Utm.
// Returns the index value of the Utm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - utm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgUtm(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/utm"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgUtm API operation for FortiOS updates the specified Utm.
// Returns the index value of the Utm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - utm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgUtm(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/utm"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgUtm API operation for FortiOS deletes the specified Utm.
// Returns error for service API and SDK errors.
// See the system.replacemsg - utm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgUtm(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/utm"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgUtm API operation for FortiOS gets the Utm
// with the specified index value.
// Returns the requested Utm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - utm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgUtm(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/utm"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgIcap API operation for FortiOS creates a new Icap.
// Returns the index value of the Icap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - icap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgIcap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.replacemsg/icap"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgIcap API operation for FortiOS updates the specified Icap.
// Returns the index value of the Icap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - icap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgIcap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.replacemsg/icap"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgIcap API operation for FortiOS deletes the specified Icap.
// Returns error for service API and SDK errors.
// See the system.replacemsg - icap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgIcap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.replacemsg/icap"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgIcap API operation for FortiOS gets the Icap
// with the specified index value.
// Returns the requested Icap value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.replacemsg - icap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgIcap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.replacemsg/icap"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgGroup API operation for FortiOS creates a new Replacemsg Group.
// Returns the index value of the Replacemsg Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/replacemsg-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgGroup API operation for FortiOS updates the specified Replacemsg Group.
// Returns the index value of the Replacemsg Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/replacemsg-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgGroup API operation for FortiOS deletes the specified Replacemsg Group.
// Returns error for service API and SDK errors.
// See the system - replacemsg-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/replacemsg-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgGroup API operation for FortiOS gets the Replacemsg Group
// with the specified index value.
// Returns the requested Replacemsg Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/replacemsg-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemSnmpSysinfo API operation for FortiOS updates the specified Sysinfo.
// Returns the index value of the Sysinfo and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpSysinfo(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/sysinfo"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpSysinfo API operation for FortiOS deletes the specified Sysinfo.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpSysinfo(mkey string) (err error) {

	//No unset API for system.snmp - sysinfo
	return
}

// ReadSystemSnmpSysinfo API operation for FortiOS gets the Sysinfo
// with the specified index value.
// Returns the requested Sysinfo value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpSysinfo(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/sysinfo"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemSnmpCommunity API operation for FortiOS creates a new Community.
// Returns the index value of the Community and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpCommunity(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.snmp/community"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSnmpCommunity API operation for FortiOS updates the specified Community.
// Returns the index value of the Community and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpCommunity(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpCommunity API operation for FortiOS deletes the specified Community.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpCommunity(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSnmpCommunity API operation for FortiOS gets the Community
// with the specified index value.
// Returns the requested Community value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpCommunity(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSnmpUser API operation for FortiOS creates a new User.
// Returns the index value of the User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpUser(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.snmp/user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSnmpUser API operation for FortiOS updates the specified User.
// Returns the index value of the User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpUser(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpUser API operation for FortiOS deletes the specified User.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpUser(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSnmpUser API operation for FortiOS gets the User
// with the specified index value.
// Returns the requested User value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpUser(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemAutoupdatePushUpdate API operation for FortiOS updates the specified Push Update.
// Returns the index value of the Push Update and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdatePushUpdate(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/push-update"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdatePushUpdate API operation for FortiOS deletes the specified Push Update.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdatePushUpdate(mkey string) (err error) {

	//No unset API for system.autoupdate - push-update
	return
}

// ReadSystemAutoupdatePushUpdate API operation for FortiOS gets the Push Update
// with the specified index value.
// Returns the requested Push Update value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdatePushUpdate(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/push-update"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemAutoupdateSchedule API operation for FortiOS updates the specified Schedule.
// Returns the index value of the Schedule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdateSchedule(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/schedule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdateSchedule API operation for FortiOS deletes the specified Schedule.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdateSchedule(mkey string) (err error) {

	//No unset API for system.autoupdate - schedule
	return
}

// ReadSystemAutoupdateSchedule API operation for FortiOS gets the Schedule
// with the specified index value.
// Returns the requested Schedule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdateSchedule(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/schedule"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemAutoupdateTunneling API operation for FortiOS updates the specified Tunneling.
// Returns the index value of the Tunneling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdateTunneling(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/tunneling"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdateTunneling API operation for FortiOS deletes the specified Tunneling.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdateTunneling(mkey string) (err error) {

	//No unset API for system.autoupdate - tunneling
	return
}

// ReadSystemAutoupdateTunneling API operation for FortiOS gets the Tunneling
// with the specified index value.
// Returns the requested Tunneling value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdateTunneling(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/tunneling"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemSessionTtl API operation for FortiOS updates the specified Session Ttl.
// Returns the index value of the Session Ttl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSessionTtl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/session-ttl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSessionTtl API operation for FortiOS deletes the specified Session Ttl.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSessionTtl(mkey string) (err error) {

	//No unset API for system - session-ttl
	return
}

// ReadSystemSessionTtl API operation for FortiOS gets the Session Ttl
// with the specified index value.
// Returns the requested Session Ttl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSessionTtl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/session-ttl"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemDhcpServer API operation for FortiOS creates a new Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDhcpServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.dhcp/server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDhcpServer API operation for FortiOS updates the specified Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDhcpServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDhcpServer API operation for FortiOS deletes the specified Server.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDhcpServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDhcpServer API operation for FortiOS gets the Server
// with the specified index value.
// Returns the requested Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDhcpServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemDhcp6Server API operation for FortiOS creates a new Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp6 - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDhcp6Server(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.dhcp6/server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDhcp6Server API operation for FortiOS updates the specified Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp6 - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDhcp6Server(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.dhcp6/server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDhcp6Server API operation for FortiOS deletes the specified Server.
// Returns error for service API and SDK errors.
// See the system.dhcp6 - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDhcp6Server(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.dhcp6/server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDhcp6Server API operation for FortiOS gets the Server
// with the specified index value.
// Returns the requested Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp6 - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDhcp6Server(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.dhcp6/server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAlias API operation for FortiOS creates a new Alias.
// Returns the index value of the Alias and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alias chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAlias(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/alias"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAlias API operation for FortiOS updates the specified Alias.
// Returns the index value of the Alias and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alias chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlias(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/alias"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAlias API operation for FortiOS deletes the specified Alias.
// Returns error for service API and SDK errors.
// See the system - alias chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlias(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/alias"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAlias API operation for FortiOS gets the Alias
// with the specified index value.
// Returns the requested Alias value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alias chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlias(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/alias"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutoScript API operation for FortiOS creates a new Auto Script.
// Returns the index value of the Auto Script and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutoScript(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/auto-script"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutoScript API operation for FortiOS updates the specified Auto Script.
// Returns the index value of the Auto Script and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoScript(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoScript API operation for FortiOS deletes the specified Auto Script.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoScript(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutoScript API operation for FortiOS gets the Auto Script
// with the specified index value.
// Returns the requested Auto Script value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoScript(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemManagementTunnel API operation for FortiOS updates the specified Management Tunnel.
// Returns the index value of the Management Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemManagementTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/management-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemManagementTunnel API operation for FortiOS deletes the specified Management Tunnel.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemManagementTunnel(mkey string) (err error) {

	//No unset API for system - management-tunnel
	return
}

// ReadSystemManagementTunnel API operation for FortiOS gets the Management Tunnel
// with the specified index value.
// Returns the requested Management Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemManagementTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/management-tunnel"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemFortimanager API operation for FortiOS updates the specified Fortimanager.
// Returns the index value of the Fortimanager and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortimanager(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortimanager"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortimanager API operation for FortiOS deletes the specified Fortimanager.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortimanager(mkey string) (err error) {

	//No unset API for system - fortimanager
	return
}

// ReadSystemFortimanager API operation for FortiOS gets the Fortimanager
// with the specified index value.
// Returns the requested Fortimanager value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortimanager(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortimanager"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemFm API operation for FortiOS updates the specified Fm.
// Returns the index value of the Fm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFm(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fm"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFm API operation for FortiOS deletes the specified Fm.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFm(mkey string) (err error) {

	//No unset API for system - fm
	return
}

// ReadSystemFm API operation for FortiOS gets the Fm
// with the specified index value.
// Returns the requested Fm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFm(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fm"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemCentralManagement API operation for FortiOS updates the specified Central Management.
// Returns the index value of the Central Management and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCentralManagement(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/central-management"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCentralManagement API operation for FortiOS deletes the specified Central Management.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCentralManagement(mkey string) (err error) {

	//No unset API for system - central-management
	return
}

// ReadSystemCentralManagement API operation for FortiOS gets the Central Management
// with the specified index value.
// Returns the requested Central Management value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCentralManagement(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/central-management"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemZone API operation for FortiOS creates a new Zone.
// Returns the index value of the Zone and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemZone(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/zone"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemZone API operation for FortiOS updates the specified Zone.
// Returns the index value of the Zone and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemZone(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemZone API operation for FortiOS deletes the specified Zone.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemZone(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemZone API operation for FortiOS gets the Zone
// with the specified index value.
// Returns the requested Zone value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemZone(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSdnConnector API operation for FortiOS creates a new Sdn Connector.
// Returns the index value of the Sdn Connector and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sdn-connector chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSdnConnector(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/sdn-connector"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSdnConnector API operation for FortiOS updates the specified Sdn Connector.
// Returns the index value of the Sdn Connector and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sdn-connector chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSdnConnector(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sdn-connector"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSdnConnector API operation for FortiOS deletes the specified Sdn Connector.
// Returns error for service API and SDK errors.
// See the system - sdn-connector chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSdnConnector(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/sdn-connector"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSdnConnector API operation for FortiOS gets the Sdn Connector
// with the specified index value.
// Returns the requested Sdn Connector value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sdn-connector chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSdnConnector(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sdn-connector"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpv6Tunnel API operation for FortiOS creates a new Ipv6 Tunnel.
// Returns the index value of the Ipv6 Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpv6Tunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipv6-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpv6Tunnel API operation for FortiOS updates the specified Ipv6 Tunnel.
// Returns the index value of the Ipv6 Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpv6Tunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipv6-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpv6Tunnel API operation for FortiOS deletes the specified Ipv6 Tunnel.
// Returns error for service API and SDK errors.
// See the system - ipv6-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpv6Tunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipv6-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpv6Tunnel API operation for FortiOS gets the Ipv6 Tunnel
// with the specified index value.
// Returns the requested Ipv6 Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpv6Tunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipv6-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemExternalResource API operation for FortiOS creates a new External Resource.
// Returns the index value of the External Resource and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - external-resource chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemExternalResource(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/external-resource"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemExternalResource API operation for FortiOS updates the specified External Resource.
// Returns the index value of the External Resource and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - external-resource chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemExternalResource(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/external-resource"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemExternalResource API operation for FortiOS deletes the specified External Resource.
// Returns error for service API and SDK errors.
// See the system - external-resource chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemExternalResource(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/external-resource"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemExternalResource API operation for FortiOS gets the External Resource
// with the specified index value.
// Returns the requested External Resource value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - external-resource chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemExternalResource(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/external-resource"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemNetworkVisibility API operation for FortiOS updates the specified Network Visibility.
// Returns the index value of the Network Visibility and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - network-visibility chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNetworkVisibility(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/network-visibility"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemNetworkVisibility API operation for FortiOS deletes the specified Network Visibility.
// Returns error for service API and SDK errors.
// See the system - network-visibility chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNetworkVisibility(mkey string) (err error) {

	//No unset API for system - network-visibility
	return
}

// ReadSystemNetworkVisibility API operation for FortiOS gets the Network Visibility
// with the specified index value.
// Returns the requested Network Visibility value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - network-visibility chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNetworkVisibility(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/network-visibility"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemGreTunnel API operation for FortiOS creates a new Gre Tunnel.
// Returns the index value of the Gre Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - gre-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemGreTunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/gre-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemGreTunnel API operation for FortiOS updates the specified Gre Tunnel.
// Returns the index value of the Gre Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - gre-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGreTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/gre-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemGreTunnel API operation for FortiOS deletes the specified Gre Tunnel.
// Returns error for service API and SDK errors.
// See the system - gre-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGreTunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/gre-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemGreTunnel API operation for FortiOS gets the Gre Tunnel
// with the specified index value.
// Returns the requested Gre Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - gre-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGreTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/gre-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpsecAggregate API operation for FortiOS creates a new Ipsec Aggregate.
// Returns the index value of the Ipsec Aggregate and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipsec-aggregate chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpsecAggregate(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipsec-aggregate"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpsecAggregate API operation for FortiOS updates the specified Ipsec Aggregate.
// Returns the index value of the Ipsec Aggregate and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipsec-aggregate chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpsecAggregate(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipsec-aggregate"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpsecAggregate API operation for FortiOS deletes the specified Ipsec Aggregate.
// Returns error for service API and SDK errors.
// See the system - ipsec-aggregate chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpsecAggregate(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipsec-aggregate"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpsecAggregate API operation for FortiOS gets the Ipsec Aggregate
// with the specified index value.
// Returns the requested Ipsec Aggregate value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipsec-aggregate chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpsecAggregate(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipsec-aggregate"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpipTunnel API operation for FortiOS creates a new Ipip Tunnel.
// Returns the index value of the Ipip Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipip-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpipTunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipip-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpipTunnel API operation for FortiOS updates the specified Ipip Tunnel.
// Returns the index value of the Ipip Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipip-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpipTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipip-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpipTunnel API operation for FortiOS deletes the specified Ipip Tunnel.
// Returns error for service API and SDK errors.
// See the system - ipip-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpipTunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipip-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpipTunnel API operation for FortiOS gets the Ipip Tunnel
// with the specified index value.
// Returns the requested Ipip Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipip-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpipTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipip-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemMobileTunnel API operation for FortiOS creates a new Mobile Tunnel.
// Returns the index value of the Mobile Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mobile-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemMobileTunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/mobile-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemMobileTunnel API operation for FortiOS updates the specified Mobile Tunnel.
// Returns the index value of the Mobile Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mobile-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemMobileTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/mobile-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemMobileTunnel API operation for FortiOS deletes the specified Mobile Tunnel.
// Returns error for service API and SDK errors.
// See the system - mobile-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemMobileTunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/mobile-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemMobileTunnel API operation for FortiOS gets the Mobile Tunnel
// with the specified index value.
// Returns the requested Mobile Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mobile-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemMobileTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/mobile-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemPppoeInterface API operation for FortiOS creates a new Pppoe Interface.
// Returns the index value of the Pppoe Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - pppoe-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemPppoeInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/pppoe-interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemPppoeInterface API operation for FortiOS updates the specified Pppoe Interface.
// Returns the index value of the Pppoe Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - pppoe-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPppoeInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/pppoe-interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPppoeInterface API operation for FortiOS deletes the specified Pppoe Interface.
// Returns error for service API and SDK errors.
// See the system - pppoe-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPppoeInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/pppoe-interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemPppoeInterface API operation for FortiOS gets the Pppoe Interface
// with the specified index value.
// Returns the requested Pppoe Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - pppoe-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPppoeInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/pppoe-interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemVxlan API operation for FortiOS creates a new Vxlan.
// Returns the index value of the Vxlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVxlan(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vxlan"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVxlan API operation for FortiOS updates the specified Vxlan.
// Returns the index value of the Vxlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVxlan(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVxlan API operation for FortiOS deletes the specified Vxlan.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVxlan(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVxlan API operation for FortiOS gets the Vxlan
// with the specified index value.
// Returns the requested Vxlan value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVxlan(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemVirtualWirePair API operation for FortiOS creates a new Virtual Wire Pair.
// Returns the index value of the Virtual Wire Pair and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - virtual-wire-pair chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVirtualWirePair(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/virtual-wire-pair"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVirtualWirePair API operation for FortiOS updates the specified Virtual Wire Pair.
// Returns the index value of the Virtual Wire Pair and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - virtual-wire-pair chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVirtualWirePair(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/virtual-wire-pair"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVirtualWirePair API operation for FortiOS deletes the specified Virtual Wire Pair.
// Returns error for service API and SDK errors.
// See the system - virtual-wire-pair chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVirtualWirePair(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/virtual-wire-pair"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVirtualWirePair API operation for FortiOS gets the Virtual Wire Pair
// with the specified index value.
// Returns the requested Virtual Wire Pair value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - virtual-wire-pair chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVirtualWirePair(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/virtual-wire-pair"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemDnsDatabase API operation for FortiOS creates a new Dns Database.
// Returns the index value of the Dns Database and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDnsDatabase(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/dns-database"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDnsDatabase API operation for FortiOS updates the specified Dns Database.
// Returns the index value of the Dns Database and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDnsDatabase(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDnsDatabase API operation for FortiOS deletes the specified Dns Database.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDnsDatabase(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDnsDatabase API operation for FortiOS gets the Dns Database
// with the specified index value.
// Returns the requested Dns Database value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDnsDatabase(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemDnsServer API operation for FortiOS creates a new Dns Server.
// Returns the index value of the Dns Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDnsServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/dns-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDnsServer API operation for FortiOS updates the specified Dns Server.
// Returns the index value of the Dns Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDnsServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDnsServer API operation for FortiOS deletes the specified Dns Server.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDnsServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDnsServer API operation for FortiOS gets the Dns Server
// with the specified index value.
// Returns the requested Dns Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDnsServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemResourceLimits API operation for FortiOS updates the specified Resource Limits.
// Returns the index value of the Resource Limits and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemResourceLimits(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/resource-limits"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemResourceLimits API operation for FortiOS deletes the specified Resource Limits.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemResourceLimits(mkey string) (err error) {

	//No unset API for system - resource-limits
	return
}

// ReadSystemResourceLimits API operation for FortiOS gets the Resource Limits
// with the specified index value.
// Returns the requested Resource Limits value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemResourceLimits(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/resource-limits"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemVdomProperty API operation for FortiOS creates a new Vdom Property.
// Returns the index value of the Vdom Property and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-property chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVdomProperty(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vdom-property"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVdomProperty API operation for FortiOS updates the specified Vdom Property.
// Returns the index value of the Vdom Property and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-property chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomProperty(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-property"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomProperty API operation for FortiOS deletes the specified Vdom Property.
// Returns error for service API and SDK errors.
// See the system - vdom-property chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomProperty(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vdom-property"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVdomProperty API operation for FortiOS gets the Vdom Property
// with the specified index value.
// Returns the requested Vdom Property value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-property chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomProperty(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-property"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemVirtualWanLink API operation for FortiOS updates the specified Virtual Wan Link.
// Returns the index value of the Virtual Wan Link and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - virtual-wan-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVirtualWanLink(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/virtual-wan-link"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVirtualWanLink API operation for FortiOS deletes the specified Virtual Wan Link.
// Returns error for service API and SDK errors.
// See the system - virtual-wan-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVirtualWanLink(mkey string) (err error) {

	//No unset API for system - virtual-wan-link
	return
}

// ReadSystemVirtualWanLink API operation for FortiOS gets the Virtual Wan Link
// with the specified index value.
// Returns the requested Virtual Wan Link value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - virtual-wan-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVirtualWanLink(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/virtual-wan-link"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemLldpNetworkPolicy API operation for FortiOS creates a new Network Policy.
// Returns the index value of the Network Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.lldp - network-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLldpNetworkPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.lldp/network-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemLldpNetworkPolicy API operation for FortiOS updates the specified Network Policy.
// Returns the index value of the Network Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.lldp - network-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLldpNetworkPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.lldp/network-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemLldpNetworkPolicy API operation for FortiOS deletes the specified Network Policy.
// Returns error for service API and SDK errors.
// See the system.lldp - network-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLldpNetworkPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.lldp/network-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemLldpNetworkPolicy API operation for FortiOS gets the Network Policy
// with the specified index value.
// Returns the requested Network Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.lldp - network-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLldpNetworkPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.lldp/network-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemClusterSync API operation for FortiOS creates a new Cluster Sync.
// Returns the index value of the Cluster Sync and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - cluster-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemClusterSync(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/cluster-sync"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemClusterSync API operation for FortiOS updates the specified Cluster Sync.
// Returns the index value of the Cluster Sync and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - cluster-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemClusterSync(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/cluster-sync"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemClusterSync API operation for FortiOS deletes the specified Cluster Sync.
// Returns error for service API and SDK errors.
// See the system - cluster-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemClusterSync(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/cluster-sync"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemClusterSync API operation for FortiOS gets the Cluster Sync
// with the specified index value.
// Returns the requested Cluster Sync value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - cluster-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemClusterSync(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/cluster-sync"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemFortiguard API operation for FortiOS updates the specified Fortiguard.
// Returns the index value of the Fortiguard and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortiguard(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortiguard"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortiguard API operation for FortiOS deletes the specified Fortiguard.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortiguard(mkey string) (err error) {

	//No unset API for system - fortiguard
	return
}

// ReadSystemFortiguard API operation for FortiOS gets the Fortiguard
// with the specified index value.
// Returns the requested Fortiguard value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortiguard(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortiguard"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemEmailServer API operation for FortiOS updates the specified Email Server.
// Returns the index value of the Email Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemEmailServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/email-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemEmailServer API operation for FortiOS deletes the specified Email Server.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemEmailServer(mkey string) (err error) {

	//No unset API for system - email-server
	return
}

// ReadSystemEmailServer API operation for FortiOS gets the Email Server
// with the specified index value.
// Returns the requested Email Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemEmailServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/email-server"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemAlarm API operation for FortiOS updates the specified Alarm.
// Returns the index value of the Alarm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alarm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlarm(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/alarm"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAlarm API operation for FortiOS deletes the specified Alarm.
// Returns error for service API and SDK errors.
// See the system - alarm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlarm(mkey string) (err error) {

	//No unset API for system - alarm
	return
}

// ReadSystemAlarm API operation for FortiOS gets the Alarm
// with the specified index value.
// Returns the requested Alarm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alarm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlarm(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/alarm"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemMacAddressTable API operation for FortiOS creates a new Mac Address Table.
// Returns the index value of the Mac Address Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mac-address-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemMacAddressTable(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/mac-address-table"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemMacAddressTable API operation for FortiOS updates the specified Mac Address Table.
// Returns the index value of the Mac Address Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mac-address-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemMacAddressTable(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/mac-address-table"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemMacAddressTable API operation for FortiOS deletes the specified Mac Address Table.
// Returns error for service API and SDK errors.
// See the system - mac-address-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemMacAddressTable(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/mac-address-table"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemMacAddressTable API operation for FortiOS gets the Mac Address Table
// with the specified index value.
// Returns the requested Mac Address Table value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mac-address-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemMacAddressTable(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/mac-address-table"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSessionHelper API operation for FortiOS creates a new Session Helper.
// Returns the index value of the Session Helper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-helper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSessionHelper(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/session-helper"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSessionHelper API operation for FortiOS updates the specified Session Helper.
// Returns the index value of the Session Helper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-helper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSessionHelper(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/session-helper"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSessionHelper API operation for FortiOS deletes the specified Session Helper.
// Returns error for service API and SDK errors.
// See the system - session-helper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSessionHelper(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/session-helper"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSessionHelper API operation for FortiOS gets the Session Helper
// with the specified index value.
// Returns the requested Session Helper value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-helper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSessionHelper(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/session-helper"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemProxyArp API operation for FortiOS creates a new Proxy Arp.
// Returns the index value of the Proxy Arp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - proxy-arp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemProxyArp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/proxy-arp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemProxyArp API operation for FortiOS updates the specified Proxy Arp.
// Returns the index value of the Proxy Arp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - proxy-arp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemProxyArp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/proxy-arp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemProxyArp API operation for FortiOS deletes the specified Proxy Arp.
// Returns error for service API and SDK errors.
// See the system - proxy-arp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemProxyArp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/proxy-arp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemProxyArp API operation for FortiOS gets the Proxy Arp
// with the specified index value.
// Returns the requested Proxy Arp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - proxy-arp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemProxyArp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/proxy-arp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemFipsCc API operation for FortiOS updates the specified Fips Cc.
// Returns the index value of the Fips Cc and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fips-cc chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFipsCc(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fips-cc"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFipsCc API operation for FortiOS deletes the specified Fips Cc.
// Returns error for service API and SDK errors.
// See the system - fips-cc chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFipsCc(mkey string) (err error) {

	//No unset API for system - fips-cc
	return
}

// ReadSystemFipsCc API operation for FortiOS gets the Fips Cc
// with the specified index value.
// Returns the requested Fips Cc value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fips-cc chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFipsCc(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fips-cc"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemTosBasedPriority API operation for FortiOS creates a new Tos Based Priority.
// Returns the index value of the Tos Based Priority and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - tos-based-priority chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemTosBasedPriority(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/tos-based-priority"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemTosBasedPriority API operation for FortiOS updates the specified Tos Based Priority.
// Returns the index value of the Tos Based Priority and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - tos-based-priority chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemTosBasedPriority(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/tos-based-priority"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemTosBasedPriority API operation for FortiOS deletes the specified Tos Based Priority.
// Returns error for service API and SDK errors.
// See the system - tos-based-priority chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemTosBasedPriority(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/tos-based-priority"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemTosBasedPriority API operation for FortiOS gets the Tos Based Priority
// with the specified index value.
// Returns the requested Tos Based Priority value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - tos-based-priority chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemTosBasedPriority(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/tos-based-priority"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemDscpBasedPriority API operation for FortiOS creates a new Dscp Based Priority.
// Returns the index value of the Dscp Based Priority and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dscp-based-priority chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDscpBasedPriority(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/dscp-based-priority"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDscpBasedPriority API operation for FortiOS updates the specified Dscp Based Priority.
// Returns the index value of the Dscp Based Priority and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dscp-based-priority chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDscpBasedPriority(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dscp-based-priority"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDscpBasedPriority API operation for FortiOS deletes the specified Dscp Based Priority.
// Returns error for service API and SDK errors.
// See the system - dscp-based-priority chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDscpBasedPriority(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/dscp-based-priority"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDscpBasedPriority API operation for FortiOS gets the Dscp Based Priority
// with the specified index value.
// Returns the requested Dscp Based Priority value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dscp-based-priority chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDscpBasedPriority(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dscp-based-priority"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemProbeResponse API operation for FortiOS updates the specified Probe Response.
// Returns the index value of the Probe Response and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - probe-response chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemProbeResponse(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/probe-response"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemProbeResponse API operation for FortiOS deletes the specified Probe Response.
// Returns error for service API and SDK errors.
// See the system - probe-response chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemProbeResponse(mkey string) (err error) {

	//No unset API for system - probe-response
	return
}

// ReadSystemProbeResponse API operation for FortiOS gets the Probe Response
// with the specified index value.
// Returns the requested Probe Response value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - probe-response chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemProbeResponse(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/probe-response"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemLinkMonitor API operation for FortiOS creates a new Link Monitor.
// Returns the index value of the Link Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - link-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLinkMonitor(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/link-monitor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemLinkMonitor API operation for FortiOS updates the specified Link Monitor.
// Returns the index value of the Link Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - link-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLinkMonitor(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/link-monitor"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemLinkMonitor API operation for FortiOS deletes the specified Link Monitor.
// Returns error for service API and SDK errors.
// See the system - link-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLinkMonitor(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/link-monitor"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemLinkMonitor API operation for FortiOS gets the Link Monitor
// with the specified index value.
// Returns the requested Link Monitor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - link-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLinkMonitor(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/link-monitor"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemAutoInstall API operation for FortiOS updates the specified Auto Install.
// Returns the index value of the Auto Install and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-install chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoInstall(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/auto-install"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoInstall API operation for FortiOS deletes the specified Auto Install.
// Returns error for service API and SDK errors.
// See the system - auto-install chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoInstall(mkey string) (err error) {

	//No unset API for system - auto-install
	return
}

// ReadSystemAutoInstall API operation for FortiOS gets the Auto Install
// with the specified index value.
// Returns the requested Auto Install value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-install chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoInstall(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/auto-install"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemConsole API operation for FortiOS updates the specified Console.
// Returns the index value of the Console and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - console chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemConsole(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/console"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemConsole API operation for FortiOS deletes the specified Console.
// Returns error for service API and SDK errors.
// See the system - console chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemConsole(mkey string) (err error) {

	//No unset API for system - console
	return
}

// ReadSystemConsole API operation for FortiOS gets the Console
// with the specified index value.
// Returns the requested Console value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - console chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemConsole(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/console"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemNtp API operation for FortiOS updates the specified Ntp.
// Returns the index value of the Ntp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNtp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ntp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemNtp API operation for FortiOS deletes the specified Ntp.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNtp(mkey string) (err error) {

	//No unset API for system - ntp
	return
}

// ReadSystemNtp API operation for FortiOS gets the Ntp
// with the specified index value.
// Returns the requested Ntp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNtp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ntp"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemPtp API operation for FortiOS updates the specified Ptp.
// Returns the index value of the Ptp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ptp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPtp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ptp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPtp API operation for FortiOS deletes the specified Ptp.
// Returns error for service API and SDK errors.
// See the system - ptp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPtp(mkey string) (err error) {

	//No unset API for system - ptp
	return
}

// ReadSystemPtp API operation for FortiOS gets the Ptp
// with the specified index value.
// Returns the requested Ptp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ptp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPtp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ptp"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemWccp API operation for FortiOS creates a new Wccp.
// Returns the index value of the Wccp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - wccp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemWccp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/wccp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemWccp API operation for FortiOS updates the specified Wccp.
// Returns the index value of the Wccp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - wccp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemWccp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/wccp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemWccp API operation for FortiOS deletes the specified Wccp.
// Returns error for service API and SDK errors.
// See the system - wccp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemWccp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/wccp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemWccp API operation for FortiOS gets the Wccp
// with the specified index value.
// Returns the requested Wccp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - wccp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemWccp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/wccp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemNat64 API operation for FortiOS updates the specified Nat64.
// Returns the index value of the Nat64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - nat64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNat64(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/nat64"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemNat64 API operation for FortiOS deletes the specified Nat64.
// Returns error for service API and SDK errors.
// See the system - nat64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNat64(mkey string) (err error) {

	//No unset API for system - nat64
	return
}

// ReadSystemNat64 API operation for FortiOS gets the Nat64
// with the specified index value.
// Returns the requested Nat64 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - nat64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNat64(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/nat64"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemVdomRadiusServer API operation for FortiOS creates a new Vdom Radius Server.
// Returns the index value of the Vdom Radius Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-radius-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVdomRadiusServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vdom-radius-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVdomRadiusServer API operation for FortiOS updates the specified Vdom Radius Server.
// Returns the index value of the Vdom Radius Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-radius-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomRadiusServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-radius-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomRadiusServer API operation for FortiOS deletes the specified Vdom Radius Server.
// Returns error for service API and SDK errors.
// See the system - vdom-radius-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomRadiusServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vdom-radius-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVdomRadiusServer API operation for FortiOS gets the Vdom Radius Server
// with the specified index value.
// Returns the requested Vdom Radius Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-radius-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomRadiusServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-radius-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemFtmPush API operation for FortiOS updates the specified Ftm Push.
// Returns the index value of the Ftm Push and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ftm-push chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFtmPush(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ftm-push"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFtmPush API operation for FortiOS deletes the specified Ftm Push.
// Returns error for service API and SDK errors.
// See the system - ftm-push chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFtmPush(mkey string) (err error) {

	//No unset API for system - ftm-push
	return
}

// ReadSystemFtmPush API operation for FortiOS gets the Ftm Push
// with the specified index value.
// Returns the requested Ftm Push value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ftm-push chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFtmPush(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ftm-push"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemGeoipOverride API operation for FortiOS creates a new Geoip Override.
// Returns the index value of the Geoip Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - geoip-override chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemGeoipOverride(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/geoip-override"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemGeoipOverride API operation for FortiOS updates the specified Geoip Override.
// Returns the index value of the Geoip Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - geoip-override chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGeoipOverride(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/geoip-override"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemGeoipOverride API operation for FortiOS deletes the specified Geoip Override.
// Returns error for service API and SDK errors.
// See the system - geoip-override chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGeoipOverride(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/geoip-override"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemGeoipOverride API operation for FortiOS gets the Geoip Override
// with the specified index value.
// Returns the requested Geoip Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - geoip-override chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGeoipOverride(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/geoip-override"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemFortisandbox API operation for FortiOS updates the specified Fortisandbox.
// Returns the index value of the Fortisandbox and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortisandbox chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortisandbox(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortisandbox"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortisandbox API operation for FortiOS deletes the specified Fortisandbox.
// Returns error for service API and SDK errors.
// See the system - fortisandbox chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortisandbox(mkey string) (err error) {

	//No unset API for system - fortisandbox
	return
}

// ReadSystemFortisandbox API operation for FortiOS gets the Fortisandbox
// with the specified index value.
// Returns the requested Fortisandbox value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortisandbox chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortisandbox(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortisandbox"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemAffinityInterrupt API operation for FortiOS creates a new Affinity Interrupt.
// Returns the index value of the Affinity Interrupt and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - affinity-interrupt chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAffinityInterrupt(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/affinity-interrupt"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAffinityInterrupt API operation for FortiOS updates the specified Affinity Interrupt.
// Returns the index value of the Affinity Interrupt and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - affinity-interrupt chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAffinityInterrupt(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/affinity-interrupt"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAffinityInterrupt API operation for FortiOS deletes the specified Affinity Interrupt.
// Returns error for service API and SDK errors.
// See the system - affinity-interrupt chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAffinityInterrupt(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/affinity-interrupt"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAffinityInterrupt API operation for FortiOS gets the Affinity Interrupt
// with the specified index value.
// Returns the requested Affinity Interrupt value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - affinity-interrupt chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAffinityInterrupt(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/affinity-interrupt"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAffinityPacketRedistribution API operation for FortiOS creates a new Affinity Packet Redistribution.
// Returns the index value of the Affinity Packet Redistribution and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - affinity-packet-redistribution chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAffinityPacketRedistribution(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/affinity-packet-redistribution"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAffinityPacketRedistribution API operation for FortiOS updates the specified Affinity Packet Redistribution.
// Returns the index value of the Affinity Packet Redistribution and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - affinity-packet-redistribution chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAffinityPacketRedistribution(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/affinity-packet-redistribution"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAffinityPacketRedistribution API operation for FortiOS deletes the specified Affinity Packet Redistribution.
// Returns error for service API and SDK errors.
// See the system - affinity-packet-redistribution chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAffinityPacketRedistribution(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/affinity-packet-redistribution"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAffinityPacketRedistribution API operation for FortiOS gets the Affinity Packet Redistribution
// with the specified index value.
// Returns the requested Affinity Packet Redistribution value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - affinity-packet-redistribution chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAffinityPacketRedistribution(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/affinity-packet-redistribution"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemVdomException API operation for FortiOS creates a new Vdom Exception.
// Returns the index value of the Vdom Exception and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-exception chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVdomException(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vdom-exception"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVdomException API operation for FortiOS updates the specified Vdom Exception.
// Returns the index value of the Vdom Exception and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-exception chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomException(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-exception"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomException API operation for FortiOS deletes the specified Vdom Exception.
// Returns error for service API and SDK errors.
// See the system - vdom-exception chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomException(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vdom-exception"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVdomException API operation for FortiOS gets the Vdom Exception
// with the specified index value.
// Returns the requested Vdom Exception value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-exception chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomException(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-exception"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemCsf API operation for FortiOS updates the specified Csf.
// Returns the index value of the Csf and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - csf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCsf(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/csf"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCsf API operation for FortiOS deletes the specified Csf.
// Returns error for service API and SDK errors.
// See the system - csf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCsf(mkey string) (err error) {

	//No unset API for system - csf
	return
}

// ReadSystemCsf API operation for FortiOS gets the Csf
// with the specified index value.
// Returns the requested Csf value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - csf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCsf(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/csf"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemAutomationTrigger API operation for FortiOS creates a new Automation Trigger.
// Returns the index value of the Automation Trigger and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-trigger chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutomationTrigger(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/automation-trigger"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutomationTrigger API operation for FortiOS updates the specified Automation Trigger.
// Returns the index value of the Automation Trigger and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-trigger chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutomationTrigger(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/automation-trigger"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutomationTrigger API operation for FortiOS deletes the specified Automation Trigger.
// Returns error for service API and SDK errors.
// See the system - automation-trigger chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutomationTrigger(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/automation-trigger"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutomationTrigger API operation for FortiOS gets the Automation Trigger
// with the specified index value.
// Returns the requested Automation Trigger value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-trigger chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutomationTrigger(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/automation-trigger"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutomationAction API operation for FortiOS creates a new Automation Action.
// Returns the index value of the Automation Action and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-action chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutomationAction(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/automation-action"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutomationAction API operation for FortiOS updates the specified Automation Action.
// Returns the index value of the Automation Action and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-action chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutomationAction(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/automation-action"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutomationAction API operation for FortiOS deletes the specified Automation Action.
// Returns error for service API and SDK errors.
// See the system - automation-action chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutomationAction(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/automation-action"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutomationAction API operation for FortiOS gets the Automation Action
// with the specified index value.
// Returns the requested Automation Action value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-action chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutomationAction(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/automation-action"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutomationDestination API operation for FortiOS creates a new Automation Destination.
// Returns the index value of the Automation Destination and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-destination chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutomationDestination(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/automation-destination"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutomationDestination API operation for FortiOS updates the specified Automation Destination.
// Returns the index value of the Automation Destination and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-destination chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutomationDestination(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/automation-destination"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutomationDestination API operation for FortiOS deletes the specified Automation Destination.
// Returns error for service API and SDK errors.
// See the system - automation-destination chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutomationDestination(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/automation-destination"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutomationDestination API operation for FortiOS gets the Automation Destination
// with the specified index value.
// Returns the requested Automation Destination value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-destination chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutomationDestination(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/automation-destination"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutomationStitch API operation for FortiOS creates a new Automation Stitch.
// Returns the index value of the Automation Stitch and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-stitch chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutomationStitch(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/automation-stitch"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutomationStitch API operation for FortiOS updates the specified Automation Stitch.
// Returns the index value of the Automation Stitch and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-stitch chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutomationStitch(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/automation-stitch"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutomationStitch API operation for FortiOS deletes the specified Automation Stitch.
// Returns error for service API and SDK errors.
// See the system - automation-stitch chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutomationStitch(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/automation-stitch"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutomationStitch API operation for FortiOS gets the Automation Stitch
// with the specified index value.
// Returns the requested Automation Stitch value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-stitch chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutomationStitch(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/automation-stitch"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemNdProxy API operation for FortiOS updates the specified Nd Proxy.
// Returns the index value of the Nd Proxy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - nd-proxy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNdProxy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/nd-proxy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemNdProxy API operation for FortiOS deletes the specified Nd Proxy.
// Returns error for service API and SDK errors.
// See the system - nd-proxy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNdProxy(mkey string) (err error) {

	//No unset API for system - nd-proxy
	return
}

// ReadSystemNdProxy API operation for FortiOS gets the Nd Proxy
// with the specified index value.
// Returns the requested Nd Proxy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - nd-proxy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNdProxy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/nd-proxy"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateWirelessControllerInterController API operation for FortiOS updates the specified Inter Controller.
// Returns the index value of the Inter Controller and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - inter-controller chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerInterController(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/inter-controller"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerInterController API operation for FortiOS deletes the specified Inter Controller.
// Returns error for service API and SDK errors.
// See the wireless-controller - inter-controller chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerInterController(mkey string) (err error) {

	//No unset API for wireless-controller - inter-controller
	return
}

// ReadWirelessControllerInterController API operation for FortiOS gets the Inter Controller
// with the specified index value.
// Returns the requested Inter Controller value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - inter-controller chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerInterController(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/inter-controller"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateWirelessControllerGlobal API operation for FortiOS updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerGlobal API operation for FortiOS deletes the specified Global.
// Returns error for service API and SDK errors.
// See the wireless-controller - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerGlobal(mkey string) (err error) {

	//No unset API for wireless-controller - global
	return
}

// ReadWirelessControllerGlobal API operation for FortiOS gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateWirelessControllerHotspot20AnqpVenueName API operation for FortiOS creates a new Anqp Venue Name.
// Returns the index value of the Anqp Venue Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-venue-name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20AnqpVenueName(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-venue-name"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20AnqpVenueName API operation for FortiOS updates the specified Anqp Venue Name.
// Returns the index value of the Anqp Venue Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-venue-name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20AnqpVenueName(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-venue-name"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20AnqpVenueName API operation for FortiOS deletes the specified Anqp Venue Name.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-venue-name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20AnqpVenueName(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-venue-name"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20AnqpVenueName API operation for FortiOS gets the Anqp Venue Name
// with the specified index value.
// Returns the requested Anqp Venue Name value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-venue-name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20AnqpVenueName(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-venue-name"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20AnqpNetworkAuthType API operation for FortiOS creates a new Anqp Network Auth Type.
// Returns the index value of the Anqp Network Auth Type and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-network-auth-type chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20AnqpNetworkAuthType(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-network-auth-type"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20AnqpNetworkAuthType API operation for FortiOS updates the specified Anqp Network Auth Type.
// Returns the index value of the Anqp Network Auth Type and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-network-auth-type chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20AnqpNetworkAuthType(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-network-auth-type"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20AnqpNetworkAuthType API operation for FortiOS deletes the specified Anqp Network Auth Type.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-network-auth-type chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20AnqpNetworkAuthType(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-network-auth-type"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20AnqpNetworkAuthType API operation for FortiOS gets the Anqp Network Auth Type
// with the specified index value.
// Returns the requested Anqp Network Auth Type value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-network-auth-type chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20AnqpNetworkAuthType(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-network-auth-type"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20AnqpRoamingConsortium API operation for FortiOS creates a new Anqp Roaming Consortium.
// Returns the index value of the Anqp Roaming Consortium and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-roaming-consortium chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20AnqpRoamingConsortium(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-roaming-consortium"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20AnqpRoamingConsortium API operation for FortiOS updates the specified Anqp Roaming Consortium.
// Returns the index value of the Anqp Roaming Consortium and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-roaming-consortium chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20AnqpRoamingConsortium(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-roaming-consortium"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20AnqpRoamingConsortium API operation for FortiOS deletes the specified Anqp Roaming Consortium.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-roaming-consortium chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20AnqpRoamingConsortium(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-roaming-consortium"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20AnqpRoamingConsortium API operation for FortiOS gets the Anqp Roaming Consortium
// with the specified index value.
// Returns the requested Anqp Roaming Consortium value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-roaming-consortium chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20AnqpRoamingConsortium(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-roaming-consortium"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20AnqpNaiRealm API operation for FortiOS creates a new Anqp Nai Realm.
// Returns the index value of the Anqp Nai Realm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-nai-realm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20AnqpNaiRealm(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-nai-realm"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20AnqpNaiRealm API operation for FortiOS updates the specified Anqp Nai Realm.
// Returns the index value of the Anqp Nai Realm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-nai-realm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20AnqpNaiRealm(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-nai-realm"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20AnqpNaiRealm API operation for FortiOS deletes the specified Anqp Nai Realm.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-nai-realm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20AnqpNaiRealm(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-nai-realm"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20AnqpNaiRealm API operation for FortiOS gets the Anqp Nai Realm
// with the specified index value.
// Returns the requested Anqp Nai Realm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-nai-realm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20AnqpNaiRealm(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-nai-realm"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20Anqp3GppCellular API operation for FortiOS creates a new Anqp 3Gpp Cellular.
// Returns the index value of the Anqp 3Gpp Cellular and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-3gpp-cellular chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20Anqp3GppCellular(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-3gpp-cellular"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20Anqp3GppCellular API operation for FortiOS updates the specified Anqp 3Gpp Cellular.
// Returns the index value of the Anqp 3Gpp Cellular and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-3gpp-cellular chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20Anqp3GppCellular(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-3gpp-cellular"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20Anqp3GppCellular API operation for FortiOS deletes the specified Anqp 3Gpp Cellular.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-3gpp-cellular chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20Anqp3GppCellular(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-3gpp-cellular"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20Anqp3GppCellular API operation for FortiOS gets the Anqp 3Gpp Cellular
// with the specified index value.
// Returns the requested Anqp 3Gpp Cellular value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-3gpp-cellular chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20Anqp3GppCellular(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-3gpp-cellular"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20AnqpIpAddressType API operation for FortiOS creates a new Anqp Ip Address Type.
// Returns the index value of the Anqp Ip Address Type and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-ip-address-type chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20AnqpIpAddressType(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-ip-address-type"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20AnqpIpAddressType API operation for FortiOS updates the specified Anqp Ip Address Type.
// Returns the index value of the Anqp Ip Address Type and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-ip-address-type chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20AnqpIpAddressType(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-ip-address-type"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20AnqpIpAddressType API operation for FortiOS deletes the specified Anqp Ip Address Type.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-ip-address-type chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20AnqpIpAddressType(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-ip-address-type"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20AnqpIpAddressType API operation for FortiOS gets the Anqp Ip Address Type
// with the specified index value.
// Returns the requested Anqp Ip Address Type value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - anqp-ip-address-type chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20AnqpIpAddressType(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/anqp-ip-address-type"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20H2QpOperatorName API operation for FortiOS creates a new H2Qp Operator Name.
// Returns the index value of the H2Qp Operator Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-operator-name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20H2QpOperatorName(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-operator-name"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20H2QpOperatorName API operation for FortiOS updates the specified H2Qp Operator Name.
// Returns the index value of the H2Qp Operator Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-operator-name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20H2QpOperatorName(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-operator-name"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20H2QpOperatorName API operation for FortiOS deletes the specified H2Qp Operator Name.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-operator-name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20H2QpOperatorName(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-operator-name"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20H2QpOperatorName API operation for FortiOS gets the H2Qp Operator Name
// with the specified index value.
// Returns the requested H2Qp Operator Name value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-operator-name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20H2QpOperatorName(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-operator-name"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20H2QpWanMetric API operation for FortiOS creates a new H2Qp Wan Metric.
// Returns the index value of the H2Qp Wan Metric and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-wan-metric chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20H2QpWanMetric(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-wan-metric"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20H2QpWanMetric API operation for FortiOS updates the specified H2Qp Wan Metric.
// Returns the index value of the H2Qp Wan Metric and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-wan-metric chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20H2QpWanMetric(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-wan-metric"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20H2QpWanMetric API operation for FortiOS deletes the specified H2Qp Wan Metric.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-wan-metric chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20H2QpWanMetric(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-wan-metric"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20H2QpWanMetric API operation for FortiOS gets the H2Qp Wan Metric
// with the specified index value.
// Returns the requested H2Qp Wan Metric value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-wan-metric chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20H2QpWanMetric(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-wan-metric"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20H2QpConnCapability API operation for FortiOS creates a new H2Qp Conn Capability.
// Returns the index value of the H2Qp Conn Capability and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-conn-capability chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20H2QpConnCapability(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-conn-capability"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20H2QpConnCapability API operation for FortiOS updates the specified H2Qp Conn Capability.
// Returns the index value of the H2Qp Conn Capability and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-conn-capability chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20H2QpConnCapability(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-conn-capability"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20H2QpConnCapability API operation for FortiOS deletes the specified H2Qp Conn Capability.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-conn-capability chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20H2QpConnCapability(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-conn-capability"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20H2QpConnCapability API operation for FortiOS gets the H2Qp Conn Capability
// with the specified index value.
// Returns the requested H2Qp Conn Capability value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-conn-capability chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20H2QpConnCapability(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-conn-capability"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20Icon API operation for FortiOS creates a new Icon.
// Returns the index value of the Icon and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - icon chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20Icon(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/icon"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20Icon API operation for FortiOS updates the specified Icon.
// Returns the index value of the Icon and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - icon chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20Icon(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/icon"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20Icon API operation for FortiOS deletes the specified Icon.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - icon chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20Icon(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/icon"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20Icon API operation for FortiOS gets the Icon
// with the specified index value.
// Returns the requested Icon value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - icon chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20Icon(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/icon"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20H2QpOsuProvider API operation for FortiOS creates a new H2Qp Osu Provider.
// Returns the index value of the H2Qp Osu Provider and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-osu-provider chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20H2QpOsuProvider(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-osu-provider"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20H2QpOsuProvider API operation for FortiOS updates the specified H2Qp Osu Provider.
// Returns the index value of the H2Qp Osu Provider and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-osu-provider chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20H2QpOsuProvider(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-osu-provider"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20H2QpOsuProvider API operation for FortiOS deletes the specified H2Qp Osu Provider.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-osu-provider chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20H2QpOsuProvider(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-osu-provider"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20H2QpOsuProvider API operation for FortiOS gets the H2Qp Osu Provider
// with the specified index value.
// Returns the requested H2Qp Osu Provider value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - h2qp-osu-provider chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20H2QpOsuProvider(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/h2qp-osu-provider"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20QosMap API operation for FortiOS creates a new Qos Map.
// Returns the index value of the Qos Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - qos-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20QosMap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/qos-map"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20QosMap API operation for FortiOS updates the specified Qos Map.
// Returns the index value of the Qos Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - qos-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20QosMap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/qos-map"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20QosMap API operation for FortiOS deletes the specified Qos Map.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - qos-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20QosMap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/qos-map"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20QosMap API operation for FortiOS gets the Qos Map
// with the specified index value.
// Returns the requested Qos Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - qos-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20QosMap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/qos-map"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerHotspot20HsProfile API operation for FortiOS creates a new Hs Profile.
// Returns the index value of the Hs Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - hs-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerHotspot20HsProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/hs-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerHotspot20HsProfile API operation for FortiOS updates the specified Hs Profile.
// Returns the index value of the Hs Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - hs-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerHotspot20HsProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/hs-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerHotspot20HsProfile API operation for FortiOS deletes the specified Hs Profile.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - hs-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerHotspot20HsProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/hs-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerHotspot20HsProfile API operation for FortiOS gets the Hs Profile
// with the specified index value.
// Returns the requested Hs Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller.hotspot20 - hs-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerHotspot20HsProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller.hotspot20/hs-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerVap API operation for FortiOS creates a new Vap.
// Returns the index value of the Vap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - vap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerVap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/vap"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerVap API operation for FortiOS updates the specified Vap.
// Returns the index value of the Vap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - vap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerVap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/vap"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerVap API operation for FortiOS deletes the specified Vap.
// Returns error for service API and SDK errors.
// See the wireless-controller - vap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerVap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/vap"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerVap API operation for FortiOS gets the Vap
// with the specified index value.
// Returns the requested Vap value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - vap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerVap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/vap"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateWirelessControllerTimers API operation for FortiOS updates the specified Timers.
// Returns the index value of the Timers and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - timers chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerTimers(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/timers"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerTimers API operation for FortiOS deletes the specified Timers.
// Returns error for service API and SDK errors.
// See the wireless-controller - timers chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerTimers(mkey string) (err error) {

	//No unset API for wireless-controller - timers
	return
}

// ReadWirelessControllerTimers API operation for FortiOS gets the Timers
// with the specified index value.
// Returns the requested Timers value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - timers chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerTimers(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/timers"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateWirelessControllerSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the wireless-controller - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerSetting(mkey string) (err error) {

	//No unset API for wireless-controller - setting
	return
}

// ReadWirelessControllerSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateWirelessControllerBonjourProfile API operation for FortiOS creates a new Bonjour Profile.
// Returns the index value of the Bonjour Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - bonjour-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerBonjourProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/bonjour-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerBonjourProfile API operation for FortiOS updates the specified Bonjour Profile.
// Returns the index value of the Bonjour Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - bonjour-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerBonjourProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/bonjour-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerBonjourProfile API operation for FortiOS deletes the specified Bonjour Profile.
// Returns error for service API and SDK errors.
// See the wireless-controller - bonjour-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerBonjourProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/bonjour-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerBonjourProfile API operation for FortiOS gets the Bonjour Profile
// with the specified index value.
// Returns the requested Bonjour Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - bonjour-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerBonjourProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/bonjour-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerRegion API operation for FortiOS creates a new Region.
// Returns the index value of the Region and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - region chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerRegion(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/region"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerRegion API operation for FortiOS updates the specified Region.
// Returns the index value of the Region and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - region chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerRegion(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/region"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerRegion API operation for FortiOS deletes the specified Region.
// Returns error for service API and SDK errors.
// See the wireless-controller - region chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerRegion(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/region"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerRegion API operation for FortiOS gets the Region
// with the specified index value.
// Returns the requested Region value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - region chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerRegion(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/region"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerVapGroup API operation for FortiOS creates a new Vap Group.
// Returns the index value of the Vap Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - vap-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerVapGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/vap-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerVapGroup API operation for FortiOS updates the specified Vap Group.
// Returns the index value of the Vap Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - vap-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerVapGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/vap-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerVapGroup API operation for FortiOS deletes the specified Vap Group.
// Returns error for service API and SDK errors.
// See the wireless-controller - vap-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerVapGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/vap-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerVapGroup API operation for FortiOS gets the Vap Group
// with the specified index value.
// Returns the requested Vap Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - vap-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerVapGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/vap-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerWidsProfile API operation for FortiOS creates a new Wids Profile.
// Returns the index value of the Wids Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wids-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerWidsProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/wids-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerWidsProfile API operation for FortiOS updates the specified Wids Profile.
// Returns the index value of the Wids Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wids-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerWidsProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/wids-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerWidsProfile API operation for FortiOS deletes the specified Wids Profile.
// Returns error for service API and SDK errors.
// See the wireless-controller - wids-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerWidsProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/wids-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerWidsProfile API operation for FortiOS gets the Wids Profile
// with the specified index value.
// Returns the requested Wids Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wids-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerWidsProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/wids-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerBleProfile API operation for FortiOS creates a new Ble Profile.
// Returns the index value of the Ble Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - ble-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerBleProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/ble-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerBleProfile API operation for FortiOS updates the specified Ble Profile.
// Returns the index value of the Ble Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - ble-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerBleProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/ble-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerBleProfile API operation for FortiOS deletes the specified Ble Profile.
// Returns error for service API and SDK errors.
// See the wireless-controller - ble-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerBleProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/ble-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerBleProfile API operation for FortiOS gets the Ble Profile
// with the specified index value.
// Returns the requested Ble Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - ble-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerBleProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/ble-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerWtpProfile API operation for FortiOS creates a new Wtp Profile.
// Returns the index value of the Wtp Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerWtpProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/wtp-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerWtpProfile API operation for FortiOS updates the specified Wtp Profile.
// Returns the index value of the Wtp Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerWtpProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/wtp-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerWtpProfile API operation for FortiOS deletes the specified Wtp Profile.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerWtpProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/wtp-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerWtpProfile API operation for FortiOS gets the Wtp Profile
// with the specified index value.
// Returns the requested Wtp Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerWtpProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/wtp-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerWtp API operation for FortiOS creates a new Wtp.
// Returns the index value of the Wtp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerWtp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/wtp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerWtp API operation for FortiOS updates the specified Wtp.
// Returns the index value of the Wtp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerWtp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/wtp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerWtp API operation for FortiOS deletes the specified Wtp.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerWtp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/wtp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerWtp API operation for FortiOS gets the Wtp
// with the specified index value.
// Returns the requested Wtp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerWtp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/wtp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerWtpGroup API operation for FortiOS creates a new Wtp Group.
// Returns the index value of the Wtp Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerWtpGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/wtp-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerWtpGroup API operation for FortiOS updates the specified Wtp Group.
// Returns the index value of the Wtp Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerWtpGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/wtp-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerWtpGroup API operation for FortiOS deletes the specified Wtp Group.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerWtpGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/wtp-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerWtpGroup API operation for FortiOS gets the Wtp Group
// with the specified index value.
// Returns the requested Wtp Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - wtp-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerWtpGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/wtp-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerQosProfile API operation for FortiOS creates a new Qos Profile.
// Returns the index value of the Qos Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - qos-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerQosProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/qos-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerQosProfile API operation for FortiOS updates the specified Qos Profile.
// Returns the index value of the Qos Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - qos-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerQosProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/qos-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerQosProfile API operation for FortiOS deletes the specified Qos Profile.
// Returns error for service API and SDK errors.
// See the wireless-controller - qos-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerQosProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/qos-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerQosProfile API operation for FortiOS gets the Qos Profile
// with the specified index value.
// Returns the requested Qos Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - qos-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerQosProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/qos-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerUtmProfile API operation for FortiOS creates a new Utm Profile.
// Returns the index value of the Utm Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - utm-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerUtmProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/utm-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerUtmProfile API operation for FortiOS updates the specified Utm Profile.
// Returns the index value of the Utm Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - utm-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerUtmProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/utm-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerUtmProfile API operation for FortiOS deletes the specified Utm Profile.
// Returns error for service API and SDK errors.
// See the wireless-controller - utm-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerUtmProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/utm-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerUtmProfile API operation for FortiOS gets the Utm Profile
// with the specified index value.
// Returns the requested Utm Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - utm-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerUtmProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/utm-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWirelessControllerApStatus API operation for FortiOS creates a new Ap Status.
// Returns the index value of the Ap Status and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - ap-status chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWirelessControllerApStatus(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wireless-controller/ap-status"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWirelessControllerApStatus API operation for FortiOS updates the specified Ap Status.
// Returns the index value of the Ap Status and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - ap-status chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWirelessControllerApStatus(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wireless-controller/ap-status"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWirelessControllerApStatus API operation for FortiOS deletes the specified Ap Status.
// Returns error for service API and SDK errors.
// See the wireless-controller - ap-status chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWirelessControllerApStatus(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wireless-controller/ap-status"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWirelessControllerApStatus API operation for FortiOS gets the Ap Status
// with the specified index value.
// Returns the requested Ap Status value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wireless-controller - ap-status chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWirelessControllerApStatus(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wireless-controller/ap-status"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerTrafficPolicy API operation for FortiOS creates a new Traffic Policy.
// Returns the index value of the Traffic Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - traffic-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerTrafficPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/traffic-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerTrafficPolicy API operation for FortiOS updates the specified Traffic Policy.
// Returns the index value of the Traffic Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - traffic-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerTrafficPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/traffic-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerTrafficPolicy API operation for FortiOS deletes the specified Traffic Policy.
// Returns error for service API and SDK errors.
// See the switch-controller - traffic-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerTrafficPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/traffic-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerTrafficPolicy API operation for FortiOS gets the Traffic Policy
// with the specified index value.
// Returns the requested Traffic Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - traffic-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerTrafficPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/traffic-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerSwitchInterfaceTag API operation for FortiOS creates a new Switch Interface Tag.
// Returns the index value of the Switch Interface Tag and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-interface-tag chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerSwitchInterfaceTag(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/switch-interface-tag"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerSwitchInterfaceTag API operation for FortiOS updates the specified Switch Interface Tag.
// Returns the index value of the Switch Interface Tag and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-interface-tag chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerSwitchInterfaceTag(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/switch-interface-tag"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerSwitchInterfaceTag API operation for FortiOS deletes the specified Switch Interface Tag.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-interface-tag chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerSwitchInterfaceTag(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/switch-interface-tag"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerSwitchInterfaceTag API operation for FortiOS gets the Switch Interface Tag
// with the specified index value.
// Returns the requested Switch Interface Tag value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-interface-tag chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerSwitchInterfaceTag(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/switch-interface-tag"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerVlan API operation for FortiOS creates a new Vlan.
// Returns the index value of the Vlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - vlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerVlan(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/vlan"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerVlan API operation for FortiOS updates the specified Vlan.
// Returns the index value of the Vlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - vlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerVlan(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/vlan"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerVlan API operation for FortiOS deletes the specified Vlan.
// Returns error for service API and SDK errors.
// See the switch-controller - vlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerVlan(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/vlan"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerVlan API operation for FortiOS gets the Vlan
// with the specified index value.
// Returns the requested Vlan value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - vlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerVlan(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/vlan"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSwitchController8021XSettings API operation for FortiOS updates the specified 802 1X Settings.
// Returns the index value of the 802 1X Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - 802-1X-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchController8021XSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/802-1X-settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchController8021XSettings API operation for FortiOS deletes the specified 802 1X Settings.
// Returns error for service API and SDK errors.
// See the switch-controller - 802-1X-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchController8021XSettings(mkey string) (err error) {

	//No unset API for switch-controller - 802-1X-settings
	return
}

// ReadSwitchController8021XSettings API operation for FortiOS gets the 802 1X Settings
// with the specified index value.
// Returns the requested 802 1X Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - 802-1X-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchController8021XSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/802-1X-settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchControllerSecurityPolicy8021X API operation for FortiOS creates a new 802 1X.
// Returns the index value of the 802 1X and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.security-policy - 802-1X chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerSecurityPolicy8021X(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller.security-policy/802-1X"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerSecurityPolicy8021X API operation for FortiOS updates the specified 802 1X.
// Returns the index value of the 802 1X and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.security-policy - 802-1X chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerSecurityPolicy8021X(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.security-policy/802-1X"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerSecurityPolicy8021X API operation for FortiOS deletes the specified 802 1X.
// Returns error for service API and SDK errors.
// See the switch-controller.security-policy - 802-1X chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerSecurityPolicy8021X(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller.security-policy/802-1X"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerSecurityPolicy8021X API operation for FortiOS gets the 802 1X
// with the specified index value.
// Returns the requested 802 1X value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.security-policy - 802-1X chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerSecurityPolicy8021X(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.security-policy/802-1X"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerSecurityPolicyCaptivePortal API operation for FortiOS creates a new Captive Portal.
// Returns the index value of the Captive Portal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.security-policy - captive-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerSecurityPolicyCaptivePortal(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller.security-policy/captive-portal"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerSecurityPolicyCaptivePortal API operation for FortiOS updates the specified Captive Portal.
// Returns the index value of the Captive Portal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.security-policy - captive-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerSecurityPolicyCaptivePortal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.security-policy/captive-portal"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerSecurityPolicyCaptivePortal API operation for FortiOS deletes the specified Captive Portal.
// Returns error for service API and SDK errors.
// See the switch-controller.security-policy - captive-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerSecurityPolicyCaptivePortal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller.security-policy/captive-portal"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerSecurityPolicyCaptivePortal API operation for FortiOS gets the Captive Portal
// with the specified index value.
// Returns the requested Captive Portal value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.security-policy - captive-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerSecurityPolicyCaptivePortal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.security-policy/captive-portal"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSwitchControllerLldpSettings API operation for FortiOS updates the specified Lldp Settings.
// Returns the index value of the Lldp Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - lldp-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerLldpSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/lldp-settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerLldpSettings API operation for FortiOS deletes the specified Lldp Settings.
// Returns error for service API and SDK errors.
// See the switch-controller - lldp-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerLldpSettings(mkey string) (err error) {

	//No unset API for switch-controller - lldp-settings
	return
}

// ReadSwitchControllerLldpSettings API operation for FortiOS gets the Lldp Settings
// with the specified index value.
// Returns the requested Lldp Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - lldp-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerLldpSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/lldp-settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchControllerLldpProfile API operation for FortiOS creates a new Lldp Profile.
// Returns the index value of the Lldp Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - lldp-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerLldpProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/lldp-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerLldpProfile API operation for FortiOS updates the specified Lldp Profile.
// Returns the index value of the Lldp Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - lldp-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerLldpProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/lldp-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerLldpProfile API operation for FortiOS deletes the specified Lldp Profile.
// Returns error for service API and SDK errors.
// See the switch-controller - lldp-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerLldpProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/lldp-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerLldpProfile API operation for FortiOS gets the Lldp Profile
// with the specified index value.
// Returns the requested Lldp Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - lldp-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerLldpProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/lldp-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerQosDot1PMap API operation for FortiOS creates a new Dot1P Map.
// Returns the index value of the Dot1P Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - dot1p-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerQosDot1PMap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller.qos/dot1p-map"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerQosDot1PMap API operation for FortiOS updates the specified Dot1P Map.
// Returns the index value of the Dot1P Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - dot1p-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerQosDot1PMap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.qos/dot1p-map"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerQosDot1PMap API operation for FortiOS deletes the specified Dot1P Map.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - dot1p-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerQosDot1PMap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller.qos/dot1p-map"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerQosDot1PMap API operation for FortiOS gets the Dot1P Map
// with the specified index value.
// Returns the requested Dot1P Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - dot1p-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerQosDot1PMap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.qos/dot1p-map"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerQosIpDscpMap API operation for FortiOS creates a new Ip Dscp Map.
// Returns the index value of the Ip Dscp Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - ip-dscp-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerQosIpDscpMap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller.qos/ip-dscp-map"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerQosIpDscpMap API operation for FortiOS updates the specified Ip Dscp Map.
// Returns the index value of the Ip Dscp Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - ip-dscp-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerQosIpDscpMap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.qos/ip-dscp-map"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerQosIpDscpMap API operation for FortiOS deletes the specified Ip Dscp Map.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - ip-dscp-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerQosIpDscpMap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller.qos/ip-dscp-map"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerQosIpDscpMap API operation for FortiOS gets the Ip Dscp Map
// with the specified index value.
// Returns the requested Ip Dscp Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - ip-dscp-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerQosIpDscpMap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.qos/ip-dscp-map"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerQosQueuePolicy API operation for FortiOS creates a new Queue Policy.
// Returns the index value of the Queue Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - queue-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerQosQueuePolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller.qos/queue-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerQosQueuePolicy API operation for FortiOS updates the specified Queue Policy.
// Returns the index value of the Queue Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - queue-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerQosQueuePolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.qos/queue-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerQosQueuePolicy API operation for FortiOS deletes the specified Queue Policy.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - queue-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerQosQueuePolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller.qos/queue-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerQosQueuePolicy API operation for FortiOS gets the Queue Policy
// with the specified index value.
// Returns the requested Queue Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - queue-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerQosQueuePolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.qos/queue-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerQosQosPolicy API operation for FortiOS creates a new Qos Policy.
// Returns the index value of the Qos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - qos-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerQosQosPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller.qos/qos-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerQosQosPolicy API operation for FortiOS updates the specified Qos Policy.
// Returns the index value of the Qos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - qos-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerQosQosPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.qos/qos-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerQosQosPolicy API operation for FortiOS deletes the specified Qos Policy.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - qos-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerQosQosPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller.qos/qos-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerQosQosPolicy API operation for FortiOS gets the Qos Policy
// with the specified index value.
// Returns the requested Qos Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.qos - qos-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerQosQosPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.qos/qos-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerAutoConfigPolicy API operation for FortiOS creates a new Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerAutoConfigPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller.auto-config/policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerAutoConfigPolicy API operation for FortiOS updates the specified Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerAutoConfigPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.auto-config/policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerAutoConfigPolicy API operation for FortiOS deletes the specified Policy.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerAutoConfigPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller.auto-config/policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerAutoConfigPolicy API operation for FortiOS gets the Policy
// with the specified index value.
// Returns the requested Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerAutoConfigPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.auto-config/policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSwitchControllerAutoConfigDefault API operation for FortiOS updates the specified Default.
// Returns the index value of the Default and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - default chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerAutoConfigDefault(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.auto-config/default"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerAutoConfigDefault API operation for FortiOS deletes the specified Default.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - default chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerAutoConfigDefault(mkey string) (err error) {

	//No unset API for switch-controller.auto-config - default
	return
}

// ReadSwitchControllerAutoConfigDefault API operation for FortiOS gets the Default
// with the specified index value.
// Returns the requested Default value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - default chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerAutoConfigDefault(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.auto-config/default"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchControllerAutoConfigCustom API operation for FortiOS creates a new Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerAutoConfigCustom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller.auto-config/custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerAutoConfigCustom API operation for FortiOS updates the specified Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerAutoConfigCustom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller.auto-config/custom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerAutoConfigCustom API operation for FortiOS deletes the specified Custom.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerAutoConfigCustom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller.auto-config/custom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerAutoConfigCustom API operation for FortiOS gets the Custom
// with the specified index value.
// Returns the requested Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller.auto-config - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerAutoConfigCustom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller.auto-config/custom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerSwitchProfile API operation for FortiOS creates a new Switch Profile.
// Returns the index value of the Switch Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerSwitchProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/switch-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerSwitchProfile API operation for FortiOS updates the specified Switch Profile.
// Returns the index value of the Switch Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerSwitchProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/switch-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerSwitchProfile API operation for FortiOS deletes the specified Switch Profile.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerSwitchProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/switch-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerSwitchProfile API operation for FortiOS gets the Switch Profile
// with the specified index value.
// Returns the requested Switch Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerSwitchProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/switch-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerCustomCommand API operation for FortiOS creates a new Custom Command.
// Returns the index value of the Custom Command and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - custom-command chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerCustomCommand(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/custom-command"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerCustomCommand API operation for FortiOS updates the specified Custom Command.
// Returns the index value of the Custom Command and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - custom-command chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerCustomCommand(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/custom-command"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerCustomCommand API operation for FortiOS deletes the specified Custom Command.
// Returns error for service API and SDK errors.
// See the switch-controller - custom-command chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerCustomCommand(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/custom-command"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerCustomCommand API operation for FortiOS gets the Custom Command
// with the specified index value.
// Returns the requested Custom Command value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - custom-command chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerCustomCommand(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/custom-command"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerVirtualPortPool API operation for FortiOS creates a new Virtual Port Pool.
// Returns the index value of the Virtual Port Pool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - virtual-port-pool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerVirtualPortPool(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/virtual-port-pool"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerVirtualPortPool API operation for FortiOS updates the specified Virtual Port Pool.
// Returns the index value of the Virtual Port Pool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - virtual-port-pool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerVirtualPortPool(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/virtual-port-pool"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerVirtualPortPool API operation for FortiOS deletes the specified Virtual Port Pool.
// Returns error for service API and SDK errors.
// See the switch-controller - virtual-port-pool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerVirtualPortPool(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/virtual-port-pool"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerVirtualPortPool API operation for FortiOS gets the Virtual Port Pool
// with the specified index value.
// Returns the requested Virtual Port Pool value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - virtual-port-pool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerVirtualPortPool(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/virtual-port-pool"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerManagedSwitch API operation for FortiOS creates a new Managed Switch.
// Returns the index value of the Managed Switch and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - managed-switch chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerManagedSwitch(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/managed-switch"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerManagedSwitch API operation for FortiOS updates the specified Managed Switch.
// Returns the index value of the Managed Switch and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - managed-switch chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerManagedSwitch(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/managed-switch"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerManagedSwitch API operation for FortiOS deletes the specified Managed Switch.
// Returns error for service API and SDK errors.
// See the switch-controller - managed-switch chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerManagedSwitch(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/managed-switch"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerManagedSwitch API operation for FortiOS gets the Managed Switch
// with the specified index value.
// Returns the requested Managed Switch value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - managed-switch chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerManagedSwitch(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/managed-switch"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchControllerSwitchGroup API operation for FortiOS creates a new Switch Group.
// Returns the index value of the Switch Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchControllerSwitchGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch-controller/switch-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchControllerSwitchGroup API operation for FortiOS updates the specified Switch Group.
// Returns the index value of the Switch Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerSwitchGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/switch-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerSwitchGroup API operation for FortiOS deletes the specified Switch Group.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerSwitchGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch-controller/switch-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchControllerSwitchGroup API operation for FortiOS gets the Switch Group
// with the specified index value.
// Returns the requested Switch Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerSwitchGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/switch-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSwitchControllerStpSettings API operation for FortiOS updates the specified Stp Settings.
// Returns the index value of the Stp Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - stp-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerStpSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/stp-settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerStpSettings API operation for FortiOS deletes the specified Stp Settings.
// Returns error for service API and SDK errors.
// See the switch-controller - stp-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerStpSettings(mkey string) (err error) {

	//No unset API for switch-controller - stp-settings
	return
}

// ReadSwitchControllerStpSettings API operation for FortiOS gets the Stp Settings
// with the specified index value.
// Returns the requested Stp Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - stp-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerStpSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/stp-settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerStormControl API operation for FortiOS updates the specified Storm Control.
// Returns the index value of the Storm Control and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - storm-control chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerStormControl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/storm-control"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerStormControl API operation for FortiOS deletes the specified Storm Control.
// Returns error for service API and SDK errors.
// See the switch-controller - storm-control chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerStormControl(mkey string) (err error) {

	//No unset API for switch-controller - storm-control
	return
}

// ReadSwitchControllerStormControl API operation for FortiOS gets the Storm Control
// with the specified index value.
// Returns the requested Storm Control value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - storm-control chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerStormControl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/storm-control"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerGlobal API operation for FortiOS updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerGlobal API operation for FortiOS deletes the specified Global.
// Returns error for service API and SDK errors.
// See the switch-controller - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerGlobal(mkey string) (err error) {

	//No unset API for switch-controller - global
	return
}

// ReadSwitchControllerGlobal API operation for FortiOS gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerSystem API operation for FortiOS updates the specified System.
// Returns the index value of the System and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - system chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerSystem(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/system"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerSystem API operation for FortiOS deletes the specified System.
// Returns error for service API and SDK errors.
// See the switch-controller - system chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerSystem(mkey string) (err error) {

	//No unset API for switch-controller - system
	return
}

// ReadSwitchControllerSystem API operation for FortiOS gets the System
// with the specified index value.
// Returns the requested System value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - system chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerSystem(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/system"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerMacSyncSettings API operation for FortiOS updates the specified Mac Sync Settings.
// Returns the index value of the Mac Sync Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - mac-sync-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerMacSyncSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/mac-sync-settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerMacSyncSettings API operation for FortiOS deletes the specified Mac Sync Settings.
// Returns error for service API and SDK errors.
// See the switch-controller - mac-sync-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerMacSyncSettings(mkey string) (err error) {

	//No unset API for switch-controller - mac-sync-settings
	return
}

// ReadSwitchControllerMacSyncSettings API operation for FortiOS gets the Mac Sync Settings
// with the specified index value.
// Returns the requested Mac Sync Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - mac-sync-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerMacSyncSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/mac-sync-settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerSwitchLog API operation for FortiOS updates the specified Switch Log.
// Returns the index value of the Switch Log and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-log chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerSwitchLog(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/switch-log"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerSwitchLog API operation for FortiOS deletes the specified Switch Log.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-log chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerSwitchLog(mkey string) (err error) {

	//No unset API for switch-controller - switch-log
	return
}

// ReadSwitchControllerSwitchLog API operation for FortiOS gets the Switch Log
// with the specified index value.
// Returns the requested Switch Log value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - switch-log chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerSwitchLog(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/switch-log"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerIgmpSnooping API operation for FortiOS updates the specified Igmp Snooping.
// Returns the index value of the Igmp Snooping and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - igmp-snooping chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerIgmpSnooping(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/igmp-snooping"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerIgmpSnooping API operation for FortiOS deletes the specified Igmp Snooping.
// Returns error for service API and SDK errors.
// See the switch-controller - igmp-snooping chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerIgmpSnooping(mkey string) (err error) {

	//No unset API for switch-controller - igmp-snooping
	return
}

// ReadSwitchControllerIgmpSnooping API operation for FortiOS gets the Igmp Snooping
// with the specified index value.
// Returns the requested Igmp Snooping value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - igmp-snooping chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerIgmpSnooping(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/igmp-snooping"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerSflow API operation for FortiOS updates the specified Sflow.
// Returns the index value of the Sflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerSflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/sflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerSflow API operation for FortiOS deletes the specified Sflow.
// Returns error for service API and SDK errors.
// See the switch-controller - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerSflow(mkey string) (err error) {

	//No unset API for switch-controller - sflow
	return
}

// ReadSwitchControllerSflow API operation for FortiOS gets the Sflow
// with the specified index value.
// Returns the requested Sflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerSflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/sflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerQuarantine API operation for FortiOS updates the specified Quarantine.
// Returns the index value of the Quarantine and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerQuarantine(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/quarantine"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerQuarantine API operation for FortiOS deletes the specified Quarantine.
// Returns error for service API and SDK errors.
// See the switch-controller - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerQuarantine(mkey string) (err error) {

	//No unset API for switch-controller - quarantine
	return
}

// ReadSwitchControllerQuarantine API operation for FortiOS gets the Quarantine
// with the specified index value.
// Returns the requested Quarantine value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerQuarantine(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/quarantine"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSwitchControllerNetworkMonitorSettings API operation for FortiOS updates the specified Network Monitor Settings.
// Returns the index value of the Network Monitor Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - network-monitor-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerNetworkMonitorSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/network-monitor-settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerNetworkMonitorSettings API operation for FortiOS deletes the specified Network Monitor Settings.
// Returns error for service API and SDK errors.
// See the switch-controller - network-monitor-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerNetworkMonitorSettings(mkey string) (err error) {

	//No unset API for switch-controller - network-monitor-settings
	return
}

// ReadSwitchControllerNetworkMonitorSettings API operation for FortiOS gets the Network Monitor Settings
// with the specified index value.
// Returns the requested Network Monitor Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - network-monitor-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerNetworkMonitorSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/network-monitor-settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateExtenderControllerExtender API operation for FortiOS creates a new Extender.
// Returns the index value of the Extender and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the extender-controller - extender chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateExtenderControllerExtender(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/extender-controller/extender"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateExtenderControllerExtender API operation for FortiOS updates the specified Extender.
// Returns the index value of the Extender and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the extender-controller - extender chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateExtenderControllerExtender(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/extender-controller/extender"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteExtenderControllerExtender API operation for FortiOS deletes the specified Extender.
// Returns error for service API and SDK errors.
// See the extender-controller - extender chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteExtenderControllerExtender(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/extender-controller/extender"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadExtenderControllerExtender API operation for FortiOS gets the Extender
// with the specified index value.
// Returns the requested Extender value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the extender-controller - extender chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadExtenderControllerExtender(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/extender-controller/extender"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallAddress API operation for FortiOS creates a new Address.
// Returns the index value of the Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallAddress(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/address"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallAddress API operation for FortiOS updates the specified Address.
// Returns the index value of the Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallAddress(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/address"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallAddress API operation for FortiOS deletes the specified Address.
// Returns error for service API and SDK errors.
// See the firewall - address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallAddress(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/address"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallAddress API operation for FortiOS gets the Address
// with the specified index value.
// Returns the requested Address value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallAddress(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/address"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallMulticastAddress API operation for FortiOS creates a new Multicast Address.
// Returns the index value of the Multicast Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallMulticastAddress(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/multicast-address"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallMulticastAddress API operation for FortiOS updates the specified Multicast Address.
// Returns the index value of the Multicast Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallMulticastAddress(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/multicast-address"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallMulticastAddress API operation for FortiOS deletes the specified Multicast Address.
// Returns error for service API and SDK errors.
// See the firewall - multicast-address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallMulticastAddress(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/multicast-address"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallMulticastAddress API operation for FortiOS gets the Multicast Address
// with the specified index value.
// Returns the requested Multicast Address value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallMulticastAddress(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/multicast-address"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallAddress6Template API operation for FortiOS creates a new Address6 Template.
// Returns the index value of the Address6 Template and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address6-template chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallAddress6Template(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/address6-template"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallAddress6Template API operation for FortiOS updates the specified Address6 Template.
// Returns the index value of the Address6 Template and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address6-template chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallAddress6Template(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/address6-template"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallAddress6Template API operation for FortiOS deletes the specified Address6 Template.
// Returns error for service API and SDK errors.
// See the firewall - address6-template chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallAddress6Template(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/address6-template"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallAddress6Template API operation for FortiOS gets the Address6 Template
// with the specified index value.
// Returns the requested Address6 Template value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address6-template chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallAddress6Template(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/address6-template"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallAddress6 API operation for FortiOS creates a new Address6.
// Returns the index value of the Address6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallAddress6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/address6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallAddress6 API operation for FortiOS updates the specified Address6.
// Returns the index value of the Address6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallAddress6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/address6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallAddress6 API operation for FortiOS deletes the specified Address6.
// Returns error for service API and SDK errors.
// See the firewall - address6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallAddress6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/address6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallAddress6 API operation for FortiOS gets the Address6
// with the specified index value.
// Returns the requested Address6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallAddress6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/address6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallMulticastAddress6 API operation for FortiOS creates a new Multicast Address6.
// Returns the index value of the Multicast Address6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-address6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallMulticastAddress6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/multicast-address6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallMulticastAddress6 API operation for FortiOS updates the specified Multicast Address6.
// Returns the index value of the Multicast Address6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-address6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallMulticastAddress6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/multicast-address6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallMulticastAddress6 API operation for FortiOS deletes the specified Multicast Address6.
// Returns error for service API and SDK errors.
// See the firewall - multicast-address6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallMulticastAddress6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/multicast-address6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallMulticastAddress6 API operation for FortiOS gets the Multicast Address6
// with the specified index value.
// Returns the requested Multicast Address6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-address6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallMulticastAddress6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/multicast-address6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallAddrgrp API operation for FortiOS creates a new Addrgrp.
// Returns the index value of the Addrgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallAddrgrp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/addrgrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallAddrgrp API operation for FortiOS updates the specified Addrgrp.
// Returns the index value of the Addrgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallAddrgrp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/addrgrp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallAddrgrp API operation for FortiOS deletes the specified Addrgrp.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallAddrgrp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/addrgrp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallAddrgrp API operation for FortiOS gets the Addrgrp
// with the specified index value.
// Returns the requested Addrgrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallAddrgrp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/addrgrp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallAddrgrp6 API operation for FortiOS creates a new Addrgrp6.
// Returns the index value of the Addrgrp6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallAddrgrp6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/addrgrp6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallAddrgrp6 API operation for FortiOS updates the specified Addrgrp6.
// Returns the index value of the Addrgrp6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallAddrgrp6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/addrgrp6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallAddrgrp6 API operation for FortiOS deletes the specified Addrgrp6.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallAddrgrp6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/addrgrp6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallAddrgrp6 API operation for FortiOS gets the Addrgrp6
// with the specified index value.
// Returns the requested Addrgrp6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallAddrgrp6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/addrgrp6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallWildcardFqdnCustom API operation for FortiOS creates a new Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.wildcard-fqdn - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallWildcardFqdnCustom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.wildcard-fqdn/custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallWildcardFqdnCustom API operation for FortiOS updates the specified Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.wildcard-fqdn - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallWildcardFqdnCustom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.wildcard-fqdn/custom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallWildcardFqdnCustom API operation for FortiOS deletes the specified Custom.
// Returns error for service API and SDK errors.
// See the firewall.wildcard-fqdn - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallWildcardFqdnCustom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.wildcard-fqdn/custom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallWildcardFqdnCustom API operation for FortiOS gets the Custom
// with the specified index value.
// Returns the requested Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.wildcard-fqdn - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallWildcardFqdnCustom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.wildcard-fqdn/custom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallWildcardFqdnGroup API operation for FortiOS creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.wildcard-fqdn - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallWildcardFqdnGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.wildcard-fqdn/group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallWildcardFqdnGroup API operation for FortiOS updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.wildcard-fqdn - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallWildcardFqdnGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.wildcard-fqdn/group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallWildcardFqdnGroup API operation for FortiOS deletes the specified Group.
// Returns error for service API and SDK errors.
// See the firewall.wildcard-fqdn - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallWildcardFqdnGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.wildcard-fqdn/group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallWildcardFqdnGroup API operation for FortiOS gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.wildcard-fqdn - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallWildcardFqdnGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.wildcard-fqdn/group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallServiceCategory API operation for FortiOS creates a new Category.
// Returns the index value of the Category and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - category chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallServiceCategory(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.service/category"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallServiceCategory API operation for FortiOS updates the specified Category.
// Returns the index value of the Category and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - category chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallServiceCategory(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.service/category"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallServiceCategory API operation for FortiOS deletes the specified Category.
// Returns error for service API and SDK errors.
// See the firewall.service - category chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallServiceCategory(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.service/category"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallServiceCategory API operation for FortiOS gets the Category
// with the specified index value.
// Returns the requested Category value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - category chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallServiceCategory(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.service/category"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallServiceCustom API operation for FortiOS creates a new Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallServiceCustom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.service/custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallServiceCustom API operation for FortiOS updates the specified Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallServiceCustom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.service/custom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallServiceCustom API operation for FortiOS deletes the specified Custom.
// Returns error for service API and SDK errors.
// See the firewall.service - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallServiceCustom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.service/custom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallServiceCustom API operation for FortiOS gets the Custom
// with the specified index value.
// Returns the requested Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallServiceCustom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.service/custom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallServiceGroup API operation for FortiOS creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallServiceGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.service/group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallServiceGroup API operation for FortiOS updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallServiceGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.service/group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallServiceGroup API operation for FortiOS deletes the specified Group.
// Returns error for service API and SDK errors.
// See the firewall.service - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallServiceGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.service/group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallServiceGroup API operation for FortiOS gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.service - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallServiceGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.service/group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallInternetService API operation for FortiOS creates a new Internet Service.
// Returns the index value of the Internet Service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallInternetService(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/internet-service"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallInternetService API operation for FortiOS updates the specified Internet Service.
// Returns the index value of the Internet Service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallInternetService(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/internet-service"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallInternetService API operation for FortiOS deletes the specified Internet Service.
// Returns error for service API and SDK errors.
// See the firewall - internet-service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallInternetService(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/internet-service"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallInternetService API operation for FortiOS gets the Internet Service
// with the specified index value.
// Returns the requested Internet Service value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallInternetService(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/internet-service"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallInternetServiceGroup API operation for FortiOS creates a new Internet Service Group.
// Returns the index value of the Internet Service Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallInternetServiceGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/internet-service-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallInternetServiceGroup API operation for FortiOS updates the specified Internet Service Group.
// Returns the index value of the Internet Service Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallInternetServiceGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/internet-service-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallInternetServiceGroup API operation for FortiOS deletes the specified Internet Service Group.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallInternetServiceGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/internet-service-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallInternetServiceGroup API operation for FortiOS gets the Internet Service Group
// with the specified index value.
// Returns the requested Internet Service Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallInternetServiceGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/internet-service-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallInternetServiceExtension API operation for FortiOS creates a new Internet Service Extension.
// Returns the index value of the Internet Service Extension and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-extension chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallInternetServiceExtension(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/internet-service-extension"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallInternetServiceExtension API operation for FortiOS updates the specified Internet Service Extension.
// Returns the index value of the Internet Service Extension and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-extension chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallInternetServiceExtension(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/internet-service-extension"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallInternetServiceExtension API operation for FortiOS deletes the specified Internet Service Extension.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-extension chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallInternetServiceExtension(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/internet-service-extension"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallInternetServiceExtension API operation for FortiOS gets the Internet Service Extension
// with the specified index value.
// Returns the requested Internet Service Extension value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-extension chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallInternetServiceExtension(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/internet-service-extension"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallInternetServiceCustom API operation for FortiOS creates a new Internet Service Custom.
// Returns the index value of the Internet Service Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallInternetServiceCustom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/internet-service-custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallInternetServiceCustom API operation for FortiOS updates the specified Internet Service Custom.
// Returns the index value of the Internet Service Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallInternetServiceCustom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/internet-service-custom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallInternetServiceCustom API operation for FortiOS deletes the specified Internet Service Custom.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallInternetServiceCustom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/internet-service-custom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallInternetServiceCustom API operation for FortiOS gets the Internet Service Custom
// with the specified index value.
// Returns the requested Internet Service Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallInternetServiceCustom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/internet-service-custom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallInternetServiceCustomGroup API operation for FortiOS creates a new Internet Service Custom Group.
// Returns the index value of the Internet Service Custom Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-custom-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallInternetServiceCustomGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/internet-service-custom-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallInternetServiceCustomGroup API operation for FortiOS updates the specified Internet Service Custom Group.
// Returns the index value of the Internet Service Custom Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-custom-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallInternetServiceCustomGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/internet-service-custom-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallInternetServiceCustomGroup API operation for FortiOS deletes the specified Internet Service Custom Group.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-custom-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallInternetServiceCustomGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/internet-service-custom-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallInternetServiceCustomGroup API operation for FortiOS gets the Internet Service Custom Group
// with the specified index value.
// Returns the requested Internet Service Custom Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-custom-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallInternetServiceCustomGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/internet-service-custom-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallInternetServiceDefinition API operation for FortiOS creates a new Internet Service Definition.
// Returns the index value of the Internet Service Definition and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-definition chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallInternetServiceDefinition(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/internet-service-definition"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallInternetServiceDefinition API operation for FortiOS updates the specified Internet Service Definition.
// Returns the index value of the Internet Service Definition and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-definition chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallInternetServiceDefinition(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/internet-service-definition"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallInternetServiceDefinition API operation for FortiOS deletes the specified Internet Service Definition.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-definition chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallInternetServiceDefinition(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/internet-service-definition"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallInternetServiceDefinition API operation for FortiOS gets the Internet Service Definition
// with the specified index value.
// Returns the requested Internet Service Definition value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - internet-service-definition chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallInternetServiceDefinition(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/internet-service-definition"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallShaperTrafficShaper API operation for FortiOS creates a new Traffic Shaper.
// Returns the index value of the Traffic Shaper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.shaper - traffic-shaper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallShaperTrafficShaper(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.shaper/traffic-shaper"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallShaperTrafficShaper API operation for FortiOS updates the specified Traffic Shaper.
// Returns the index value of the Traffic Shaper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.shaper - traffic-shaper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallShaperTrafficShaper(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.shaper/traffic-shaper"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallShaperTrafficShaper API operation for FortiOS deletes the specified Traffic Shaper.
// Returns error for service API and SDK errors.
// See the firewall.shaper - traffic-shaper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallShaperTrafficShaper(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.shaper/traffic-shaper"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallShaperTrafficShaper API operation for FortiOS gets the Traffic Shaper
// with the specified index value.
// Returns the requested Traffic Shaper value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.shaper - traffic-shaper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallShaperTrafficShaper(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.shaper/traffic-shaper"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallShaperPerIpShaper API operation for FortiOS creates a new Per Ip Shaper.
// Returns the index value of the Per Ip Shaper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.shaper - per-ip-shaper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallShaperPerIpShaper(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.shaper/per-ip-shaper"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallShaperPerIpShaper API operation for FortiOS updates the specified Per Ip Shaper.
// Returns the index value of the Per Ip Shaper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.shaper - per-ip-shaper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallShaperPerIpShaper(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.shaper/per-ip-shaper"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallShaperPerIpShaper API operation for FortiOS deletes the specified Per Ip Shaper.
// Returns error for service API and SDK errors.
// See the firewall.shaper - per-ip-shaper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallShaperPerIpShaper(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.shaper/per-ip-shaper"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallShaperPerIpShaper API operation for FortiOS gets the Per Ip Shaper
// with the specified index value.
// Returns the requested Per Ip Shaper value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.shaper - per-ip-shaper chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallShaperPerIpShaper(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.shaper/per-ip-shaper"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallProxyAddress API operation for FortiOS creates a new Proxy Address.
// Returns the index value of the Proxy Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallProxyAddress(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/proxy-address"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallProxyAddress API operation for FortiOS updates the specified Proxy Address.
// Returns the index value of the Proxy Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallProxyAddress(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/proxy-address"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallProxyAddress API operation for FortiOS deletes the specified Proxy Address.
// Returns error for service API and SDK errors.
// See the firewall - proxy-address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallProxyAddress(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/proxy-address"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallProxyAddress API operation for FortiOS gets the Proxy Address
// with the specified index value.
// Returns the requested Proxy Address value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallProxyAddress(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/proxy-address"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallProxyAddrgrp API operation for FortiOS creates a new Proxy Addrgrp.
// Returns the index value of the Proxy Addrgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallProxyAddrgrp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/proxy-addrgrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallProxyAddrgrp API operation for FortiOS updates the specified Proxy Addrgrp.
// Returns the index value of the Proxy Addrgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallProxyAddrgrp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/proxy-addrgrp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallProxyAddrgrp API operation for FortiOS deletes the specified Proxy Addrgrp.
// Returns error for service API and SDK errors.
// See the firewall - proxy-addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallProxyAddrgrp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/proxy-addrgrp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallProxyAddrgrp API operation for FortiOS gets the Proxy Addrgrp
// with the specified index value.
// Returns the requested Proxy Addrgrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallProxyAddrgrp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/proxy-addrgrp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallScheduleOnetime API operation for FortiOS creates a new Onetime.
// Returns the index value of the Onetime and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - onetime chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallScheduleOnetime(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.schedule/onetime"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallScheduleOnetime API operation for FortiOS updates the specified Onetime.
// Returns the index value of the Onetime and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - onetime chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallScheduleOnetime(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.schedule/onetime"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallScheduleOnetime API operation for FortiOS deletes the specified Onetime.
// Returns error for service API and SDK errors.
// See the firewall.schedule - onetime chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallScheduleOnetime(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.schedule/onetime"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallScheduleOnetime API operation for FortiOS gets the Onetime
// with the specified index value.
// Returns the requested Onetime value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - onetime chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallScheduleOnetime(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.schedule/onetime"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallScheduleRecurring API operation for FortiOS creates a new Recurring.
// Returns the index value of the Recurring and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - recurring chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallScheduleRecurring(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.schedule/recurring"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallScheduleRecurring API operation for FortiOS updates the specified Recurring.
// Returns the index value of the Recurring and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - recurring chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallScheduleRecurring(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.schedule/recurring"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallScheduleRecurring API operation for FortiOS deletes the specified Recurring.
// Returns error for service API and SDK errors.
// See the firewall.schedule - recurring chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallScheduleRecurring(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.schedule/recurring"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallScheduleRecurring API operation for FortiOS gets the Recurring
// with the specified index value.
// Returns the requested Recurring value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - recurring chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallScheduleRecurring(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.schedule/recurring"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallScheduleGroup API operation for FortiOS creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallScheduleGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.schedule/group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallScheduleGroup API operation for FortiOS updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallScheduleGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.schedule/group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallScheduleGroup API operation for FortiOS deletes the specified Group.
// Returns error for service API and SDK errors.
// See the firewall.schedule - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallScheduleGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.schedule/group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallScheduleGroup API operation for FortiOS gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.schedule - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallScheduleGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.schedule/group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallIppool API operation for FortiOS creates a new Ippool.
// Returns the index value of the Ippool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallIppool(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/ippool"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallIppool API operation for FortiOS updates the specified Ippool.
// Returns the index value of the Ippool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallIppool(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ippool"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallIppool API operation for FortiOS deletes the specified Ippool.
// Returns error for service API and SDK errors.
// See the firewall - ippool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallIppool(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ippool"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallIppool API operation for FortiOS gets the Ippool
// with the specified index value.
// Returns the requested Ippool value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallIppool(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ippool"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallIppool6 API operation for FortiOS creates a new Ippool6.
// Returns the index value of the Ippool6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallIppool6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/ippool6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallIppool6 API operation for FortiOS updates the specified Ippool6.
// Returns the index value of the Ippool6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallIppool6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ippool6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallIppool6 API operation for FortiOS deletes the specified Ippool6.
// Returns error for service API and SDK errors.
// See the firewall - ippool6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallIppool6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ippool6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallIppool6 API operation for FortiOS gets the Ippool6
// with the specified index value.
// Returns the requested Ippool6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallIppool6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ippool6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallLdbMonitor API operation for FortiOS creates a new Ldb Monitor.
// Returns the index value of the Ldb Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ldb-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallLdbMonitor(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/ldb-monitor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallLdbMonitor API operation for FortiOS updates the specified Ldb Monitor.
// Returns the index value of the Ldb Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ldb-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallLdbMonitor(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ldb-monitor"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallLdbMonitor API operation for FortiOS deletes the specified Ldb Monitor.
// Returns error for service API and SDK errors.
// See the firewall - ldb-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallLdbMonitor(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ldb-monitor"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallLdbMonitor API operation for FortiOS gets the Ldb Monitor
// with the specified index value.
// Returns the requested Ldb Monitor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ldb-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallLdbMonitor(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ldb-monitor"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallVip API operation for FortiOS creates a new Vip.
// Returns the index value of the Vip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallVip(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vip"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallVip API operation for FortiOS updates the specified Vip.
// Returns the index value of the Vip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallVip(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vip"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallVip API operation for FortiOS deletes the specified Vip.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallVip(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vip"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallVip API operation for FortiOS gets the Vip
// with the specified index value.
// Returns the requested Vip value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallVip(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vip"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallVip46 API operation for FortiOS creates a new Vip46.
// Returns the index value of the Vip46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallVip46(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vip46"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallVip46 API operation for FortiOS updates the specified Vip46.
// Returns the index value of the Vip46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallVip46(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vip46"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallVip46 API operation for FortiOS deletes the specified Vip46.
// Returns error for service API and SDK errors.
// See the firewall - vip46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallVip46(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vip46"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallVip46 API operation for FortiOS gets the Vip46
// with the specified index value.
// Returns the requested Vip46 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallVip46(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vip46"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallVip6 API operation for FortiOS creates a new Vip6.
// Returns the index value of the Vip6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallVip6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vip6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallVip6 API operation for FortiOS updates the specified Vip6.
// Returns the index value of the Vip6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallVip6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vip6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallVip6 API operation for FortiOS deletes the specified Vip6.
// Returns error for service API and SDK errors.
// See the firewall - vip6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallVip6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vip6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallVip6 API operation for FortiOS gets the Vip6
// with the specified index value.
// Returns the requested Vip6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallVip6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vip6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallVip64 API operation for FortiOS creates a new Vip64.
// Returns the index value of the Vip64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallVip64(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vip64"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallVip64 API operation for FortiOS updates the specified Vip64.
// Returns the index value of the Vip64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallVip64(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vip64"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallVip64 API operation for FortiOS deletes the specified Vip64.
// Returns error for service API and SDK errors.
// See the firewall - vip64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallVip64(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vip64"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallVip64 API operation for FortiOS gets the Vip64
// with the specified index value.
// Returns the requested Vip64 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallVip64(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vip64"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallVipgrp API operation for FortiOS creates a new Vipgrp.
// Returns the index value of the Vipgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallVipgrp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vipgrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallVipgrp API operation for FortiOS updates the specified Vipgrp.
// Returns the index value of the Vipgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallVipgrp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vipgrp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallVipgrp API operation for FortiOS deletes the specified Vipgrp.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallVipgrp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vipgrp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallVipgrp API operation for FortiOS gets the Vipgrp
// with the specified index value.
// Returns the requested Vipgrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallVipgrp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vipgrp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallVipgrp46 API operation for FortiOS creates a new Vipgrp46.
// Returns the index value of the Vipgrp46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallVipgrp46(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vipgrp46"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallVipgrp46 API operation for FortiOS updates the specified Vipgrp46.
// Returns the index value of the Vipgrp46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallVipgrp46(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vipgrp46"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallVipgrp46 API operation for FortiOS deletes the specified Vipgrp46.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallVipgrp46(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vipgrp46"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallVipgrp46 API operation for FortiOS gets the Vipgrp46
// with the specified index value.
// Returns the requested Vipgrp46 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallVipgrp46(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vipgrp46"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallVipgrp6 API operation for FortiOS creates a new Vipgrp6.
// Returns the index value of the Vipgrp6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallVipgrp6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vipgrp6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallVipgrp6 API operation for FortiOS updates the specified Vipgrp6.
// Returns the index value of the Vipgrp6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallVipgrp6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vipgrp6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallVipgrp6 API operation for FortiOS deletes the specified Vipgrp6.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallVipgrp6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vipgrp6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallVipgrp6 API operation for FortiOS gets the Vipgrp6
// with the specified index value.
// Returns the requested Vipgrp6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallVipgrp6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vipgrp6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallVipgrp64 API operation for FortiOS creates a new Vipgrp64.
// Returns the index value of the Vipgrp64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallVipgrp64(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vipgrp64"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallVipgrp64 API operation for FortiOS updates the specified Vipgrp64.
// Returns the index value of the Vipgrp64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallVipgrp64(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vipgrp64"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallVipgrp64 API operation for FortiOS deletes the specified Vipgrp64.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallVipgrp64(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vipgrp64"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallVipgrp64 API operation for FortiOS gets the Vipgrp64
// with the specified index value.
// Returns the requested Vipgrp64 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vipgrp64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallVipgrp64(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vipgrp64"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateFirewallIpmacbindingSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ipmacbinding - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallIpmacbindingSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.ipmacbinding/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallIpmacbindingSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the firewall.ipmacbinding - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallIpmacbindingSetting(mkey string) (err error) {

	//No unset API for firewall.ipmacbinding - setting
	return
}

// ReadFirewallIpmacbindingSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ipmacbinding - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallIpmacbindingSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.ipmacbinding/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateFirewallIpmacbindingTable API operation for FortiOS creates a new Table.
// Returns the index value of the Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ipmacbinding - table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallIpmacbindingTable(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.ipmacbinding/table"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallIpmacbindingTable API operation for FortiOS updates the specified Table.
// Returns the index value of the Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ipmacbinding - table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallIpmacbindingTable(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.ipmacbinding/table"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallIpmacbindingTable API operation for FortiOS deletes the specified Table.
// Returns error for service API and SDK errors.
// See the firewall.ipmacbinding - table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallIpmacbindingTable(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.ipmacbinding/table"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallIpmacbindingTable API operation for FortiOS gets the Table
// with the specified index value.
// Returns the requested Table value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ipmacbinding - table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallIpmacbindingTable(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.ipmacbinding/table"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallProfileProtocolOptions API operation for FortiOS creates a new Profile Protocol Options.
// Returns the index value of the Profile Protocol Options and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - profile-protocol-options chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallProfileProtocolOptions(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/profile-protocol-options"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallProfileProtocolOptions API operation for FortiOS updates the specified Profile Protocol Options.
// Returns the index value of the Profile Protocol Options and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - profile-protocol-options chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallProfileProtocolOptions(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/profile-protocol-options"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallProfileProtocolOptions API operation for FortiOS deletes the specified Profile Protocol Options.
// Returns error for service API and SDK errors.
// See the firewall - profile-protocol-options chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallProfileProtocolOptions(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/profile-protocol-options"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallProfileProtocolOptions API operation for FortiOS gets the Profile Protocol Options
// with the specified index value.
// Returns the requested Profile Protocol Options value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - profile-protocol-options chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallProfileProtocolOptions(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/profile-protocol-options"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallSslSshProfile API operation for FortiOS creates a new Ssl Ssh Profile.
// Returns the index value of the Ssl Ssh Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ssl-ssh-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallSslSshProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/ssl-ssh-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallSslSshProfile API operation for FortiOS updates the specified Ssl Ssh Profile.
// Returns the index value of the Ssl Ssh Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ssl-ssh-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSslSshProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ssl-ssh-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallSslSshProfile API operation for FortiOS deletes the specified Ssl Ssh Profile.
// Returns error for service API and SDK errors.
// See the firewall - ssl-ssh-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSslSshProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ssl-ssh-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallSslSshProfile API operation for FortiOS gets the Ssl Ssh Profile
// with the specified index value.
// Returns the requested Ssl Ssh Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ssl-ssh-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSslSshProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ssl-ssh-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallProfileGroup API operation for FortiOS creates a new Profile Group.
// Returns the index value of the Profile Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - profile-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallProfileGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/profile-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallProfileGroup API operation for FortiOS updates the specified Profile Group.
// Returns the index value of the Profile Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - profile-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallProfileGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/profile-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallProfileGroup API operation for FortiOS deletes the specified Profile Group.
// Returns error for service API and SDK errors.
// See the firewall - profile-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallProfileGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/profile-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallProfileGroup API operation for FortiOS gets the Profile Group
// with the specified index value.
// Returns the requested Profile Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - profile-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallProfileGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/profile-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallSslServer API operation for FortiOS creates a new Ssl Server.
// Returns the index value of the Ssl Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ssl-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallSslServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/ssl-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallSslServer API operation for FortiOS updates the specified Ssl Server.
// Returns the index value of the Ssl Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ssl-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSslServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ssl-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallSslServer API operation for FortiOS deletes the specified Ssl Server.
// Returns error for service API and SDK errors.
// See the firewall - ssl-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSslServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ssl-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallSslServer API operation for FortiOS gets the Ssl Server
// with the specified index value.
// Returns the requested Ssl Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ssl-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSslServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ssl-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallIdentityBasedRoute API operation for FortiOS creates a new Identity Based Route.
// Returns the index value of the Identity Based Route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - identity-based-route chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallIdentityBasedRoute(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/identity-based-route"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallIdentityBasedRoute API operation for FortiOS updates the specified Identity Based Route.
// Returns the index value of the Identity Based Route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - identity-based-route chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallIdentityBasedRoute(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/identity-based-route"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallIdentityBasedRoute API operation for FortiOS deletes the specified Identity Based Route.
// Returns error for service API and SDK errors.
// See the firewall - identity-based-route chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallIdentityBasedRoute(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/identity-based-route"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallIdentityBasedRoute API operation for FortiOS gets the Identity Based Route
// with the specified index value.
// Returns the requested Identity Based Route value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - identity-based-route chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallIdentityBasedRoute(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/identity-based-route"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateFirewallAuthPortal API operation for FortiOS updates the specified Auth Portal.
// Returns the index value of the Auth Portal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - auth-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallAuthPortal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/auth-portal"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallAuthPortal API operation for FortiOS deletes the specified Auth Portal.
// Returns error for service API and SDK errors.
// See the firewall - auth-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallAuthPortal(mkey string) (err error) {

	//No unset API for firewall - auth-portal
	return
}

// ReadFirewallAuthPortal API operation for FortiOS gets the Auth Portal
// with the specified index value.
// Returns the requested Auth Portal value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - auth-portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallAuthPortal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/auth-portal"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateFirewallConsolidatedPolicy API operation for FortiOS creates a new Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.consolidated - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallConsolidatedPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.consolidated/policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallConsolidatedPolicy API operation for FortiOS updates the specified Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.consolidated - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallConsolidatedPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.consolidated/policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallConsolidatedPolicy API operation for FortiOS deletes the specified Policy.
// Returns error for service API and SDK errors.
// See the firewall.consolidated - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallConsolidatedPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.consolidated/policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallConsolidatedPolicy API operation for FortiOS gets the Policy
// with the specified index value.
// Returns the requested Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.consolidated - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallConsolidatedPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.consolidated/policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallPolicy API operation for FortiOS creates a new Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallPolicy API operation for FortiOS updates the specified Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallPolicy API operation for FortiOS deletes the specified Policy.
// Returns error for service API and SDK errors.
// See the firewall - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallPolicy API operation for FortiOS gets the Policy
// with the specified index value.
// Returns the requested Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallShapingPolicy API operation for FortiOS creates a new Shaping Policy.
// Returns the index value of the Shaping Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - shaping-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallShapingPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/shaping-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallShapingPolicy API operation for FortiOS updates the specified Shaping Policy.
// Returns the index value of the Shaping Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - shaping-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallShapingPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/shaping-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallShapingPolicy API operation for FortiOS deletes the specified Shaping Policy.
// Returns error for service API and SDK errors.
// See the firewall - shaping-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallShapingPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/shaping-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallShapingPolicy API operation for FortiOS gets the Shaping Policy
// with the specified index value.
// Returns the requested Shaping Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - shaping-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallShapingPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/shaping-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallShapingProfile API operation for FortiOS creates a new Shaping Profile.
// Returns the index value of the Shaping Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - shaping-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallShapingProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/shaping-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallShapingProfile API operation for FortiOS updates the specified Shaping Profile.
// Returns the index value of the Shaping Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - shaping-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallShapingProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/shaping-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallShapingProfile API operation for FortiOS deletes the specified Shaping Profile.
// Returns error for service API and SDK errors.
// See the firewall - shaping-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallShapingProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/shaping-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallShapingProfile API operation for FortiOS gets the Shaping Profile
// with the specified index value.
// Returns the requested Shaping Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - shaping-profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallShapingProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/shaping-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallLocalInPolicy API operation for FortiOS creates a new Local In Policy.
// Returns the index value of the Local In Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - local-in-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallLocalInPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/local-in-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallLocalInPolicy API operation for FortiOS updates the specified Local In Policy.
// Returns the index value of the Local In Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - local-in-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallLocalInPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/local-in-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallLocalInPolicy API operation for FortiOS deletes the specified Local In Policy.
// Returns error for service API and SDK errors.
// See the firewall - local-in-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallLocalInPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/local-in-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallLocalInPolicy API operation for FortiOS gets the Local In Policy
// with the specified index value.
// Returns the requested Local In Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - local-in-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallLocalInPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/local-in-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallPolicy6 API operation for FortiOS creates a new Policy6.
// Returns the index value of the Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallPolicy6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/policy6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallPolicy6 API operation for FortiOS updates the specified Policy6.
// Returns the index value of the Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallPolicy6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/policy6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallPolicy6 API operation for FortiOS deletes the specified Policy6.
// Returns error for service API and SDK errors.
// See the firewall - policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallPolicy6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/policy6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallPolicy6 API operation for FortiOS gets the Policy6
// with the specified index value.
// Returns the requested Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallPolicy6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/policy6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallLocalInPolicy6 API operation for FortiOS creates a new Local In Policy6.
// Returns the index value of the Local In Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - local-in-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallLocalInPolicy6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/local-in-policy6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallLocalInPolicy6 API operation for FortiOS updates the specified Local In Policy6.
// Returns the index value of the Local In Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - local-in-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallLocalInPolicy6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/local-in-policy6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallLocalInPolicy6 API operation for FortiOS deletes the specified Local In Policy6.
// Returns error for service API and SDK errors.
// See the firewall - local-in-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallLocalInPolicy6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/local-in-policy6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallLocalInPolicy6 API operation for FortiOS gets the Local In Policy6
// with the specified index value.
// Returns the requested Local In Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - local-in-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallLocalInPolicy6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/local-in-policy6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallTtlPolicy API operation for FortiOS creates a new Ttl Policy.
// Returns the index value of the Ttl Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ttl-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallTtlPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/ttl-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallTtlPolicy API operation for FortiOS updates the specified Ttl Policy.
// Returns the index value of the Ttl Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ttl-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallTtlPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ttl-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallTtlPolicy API operation for FortiOS deletes the specified Ttl Policy.
// Returns error for service API and SDK errors.
// See the firewall - ttl-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallTtlPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ttl-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallTtlPolicy API operation for FortiOS gets the Ttl Policy
// with the specified index value.
// Returns the requested Ttl Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ttl-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallTtlPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ttl-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallPolicy64 API operation for FortiOS creates a new Policy64.
// Returns the index value of the Policy64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallPolicy64(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/policy64"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallPolicy64 API operation for FortiOS updates the specified Policy64.
// Returns the index value of the Policy64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallPolicy64(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/policy64"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallPolicy64 API operation for FortiOS deletes the specified Policy64.
// Returns error for service API and SDK errors.
// See the firewall - policy64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallPolicy64(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/policy64"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallPolicy64 API operation for FortiOS gets the Policy64
// with the specified index value.
// Returns the requested Policy64 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy64 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallPolicy64(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/policy64"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallPolicy46 API operation for FortiOS creates a new Policy46.
// Returns the index value of the Policy46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallPolicy46(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/policy46"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallPolicy46 API operation for FortiOS updates the specified Policy46.
// Returns the index value of the Policy46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallPolicy46(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/policy46"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallPolicy46 API operation for FortiOS deletes the specified Policy46.
// Returns error for service API and SDK errors.
// See the firewall - policy46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallPolicy46(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/policy46"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallPolicy46 API operation for FortiOS gets the Policy46
// with the specified index value.
// Returns the requested Policy46 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy46 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallPolicy46(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/policy46"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallProxyPolicy API operation for FortiOS creates a new Proxy Policy.
// Returns the index value of the Proxy Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallProxyPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/proxy-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallProxyPolicy API operation for FortiOS updates the specified Proxy Policy.
// Returns the index value of the Proxy Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallProxyPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/proxy-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallProxyPolicy API operation for FortiOS deletes the specified Proxy Policy.
// Returns error for service API and SDK errors.
// See the firewall - proxy-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallProxyPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/proxy-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallProxyPolicy API operation for FortiOS gets the Proxy Policy
// with the specified index value.
// Returns the requested Proxy Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - proxy-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallProxyPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/proxy-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallDnstranslation API operation for FortiOS creates a new Dnstranslation.
// Returns the index value of the Dnstranslation and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - dnstranslation chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallDnstranslation(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/dnstranslation"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallDnstranslation API operation for FortiOS updates the specified Dnstranslation.
// Returns the index value of the Dnstranslation and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - dnstranslation chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallDnstranslation(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/dnstranslation"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallDnstranslation API operation for FortiOS deletes the specified Dnstranslation.
// Returns error for service API and SDK errors.
// See the firewall - dnstranslation chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallDnstranslation(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/dnstranslation"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallDnstranslation API operation for FortiOS gets the Dnstranslation
// with the specified index value.
// Returns the requested Dnstranslation value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - dnstranslation chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallDnstranslation(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/dnstranslation"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallMulticastPolicy API operation for FortiOS creates a new Multicast Policy.
// Returns the index value of the Multicast Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallMulticastPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/multicast-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallMulticastPolicy API operation for FortiOS updates the specified Multicast Policy.
// Returns the index value of the Multicast Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallMulticastPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/multicast-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallMulticastPolicy API operation for FortiOS deletes the specified Multicast Policy.
// Returns error for service API and SDK errors.
// See the firewall - multicast-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallMulticastPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/multicast-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallMulticastPolicy API operation for FortiOS gets the Multicast Policy
// with the specified index value.
// Returns the requested Multicast Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallMulticastPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/multicast-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallMulticastPolicy6 API operation for FortiOS creates a new Multicast Policy6.
// Returns the index value of the Multicast Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallMulticastPolicy6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/multicast-policy6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallMulticastPolicy6 API operation for FortiOS updates the specified Multicast Policy6.
// Returns the index value of the Multicast Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallMulticastPolicy6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/multicast-policy6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallMulticastPolicy6 API operation for FortiOS deletes the specified Multicast Policy6.
// Returns error for service API and SDK errors.
// See the firewall - multicast-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallMulticastPolicy6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/multicast-policy6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallMulticastPolicy6 API operation for FortiOS gets the Multicast Policy6
// with the specified index value.
// Returns the requested Multicast Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - multicast-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallMulticastPolicy6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/multicast-policy6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallInterfacePolicy API operation for FortiOS creates a new Interface Policy.
// Returns the index value of the Interface Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - interface-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallInterfacePolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/interface-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallInterfacePolicy API operation for FortiOS updates the specified Interface Policy.
// Returns the index value of the Interface Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - interface-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallInterfacePolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/interface-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallInterfacePolicy API operation for FortiOS deletes the specified Interface Policy.
// Returns error for service API and SDK errors.
// See the firewall - interface-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallInterfacePolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/interface-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallInterfacePolicy API operation for FortiOS gets the Interface Policy
// with the specified index value.
// Returns the requested Interface Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - interface-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallInterfacePolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/interface-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallInterfacePolicy6 API operation for FortiOS creates a new Interface Policy6.
// Returns the index value of the Interface Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - interface-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallInterfacePolicy6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/interface-policy6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallInterfacePolicy6 API operation for FortiOS updates the specified Interface Policy6.
// Returns the index value of the Interface Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - interface-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallInterfacePolicy6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/interface-policy6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallInterfacePolicy6 API operation for FortiOS deletes the specified Interface Policy6.
// Returns error for service API and SDK errors.
// See the firewall - interface-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallInterfacePolicy6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/interface-policy6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallInterfacePolicy6 API operation for FortiOS gets the Interface Policy6
// with the specified index value.
// Returns the requested Interface Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - interface-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallInterfacePolicy6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/interface-policy6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallDosPolicy API operation for FortiOS creates a new Dos Policy.
// Returns the index value of the Dos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - DoS-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallDosPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/DoS-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallDosPolicy API operation for FortiOS updates the specified Dos Policy.
// Returns the index value of the Dos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - DoS-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallDosPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/DoS-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallDosPolicy API operation for FortiOS deletes the specified Dos Policy.
// Returns error for service API and SDK errors.
// See the firewall - DoS-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallDosPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/DoS-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallDosPolicy API operation for FortiOS gets the Dos Policy
// with the specified index value.
// Returns the requested Dos Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - DoS-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallDosPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/DoS-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallDosPolicy6 API operation for FortiOS creates a new Dos Policy6.
// Returns the index value of the Dos Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - DoS-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallDosPolicy6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/DoS-policy6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallDosPolicy6 API operation for FortiOS updates the specified Dos Policy6.
// Returns the index value of the Dos Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - DoS-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallDosPolicy6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/DoS-policy6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallDosPolicy6 API operation for FortiOS deletes the specified Dos Policy6.
// Returns error for service API and SDK errors.
// See the firewall - DoS-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallDosPolicy6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/DoS-policy6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallDosPolicy6 API operation for FortiOS gets the Dos Policy6
// with the specified index value.
// Returns the requested Dos Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - DoS-policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallDosPolicy6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/DoS-policy6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallSniffer API operation for FortiOS creates a new Sniffer.
// Returns the index value of the Sniffer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - sniffer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallSniffer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/sniffer"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallSniffer API operation for FortiOS updates the specified Sniffer.
// Returns the index value of the Sniffer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - sniffer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSniffer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/sniffer"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallSniffer API operation for FortiOS deletes the specified Sniffer.
// Returns error for service API and SDK errors.
// See the firewall - sniffer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSniffer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/sniffer"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallSniffer API operation for FortiOS gets the Sniffer
// with the specified index value.
// Returns the requested Sniffer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - sniffer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSniffer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/sniffer"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallCentralSnatMap API operation for FortiOS creates a new Central Snat Map.
// Returns the index value of the Central Snat Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - central-snat-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallCentralSnatMap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/central-snat-map"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallCentralSnatMap API operation for FortiOS updates the specified Central Snat Map.
// Returns the index value of the Central Snat Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - central-snat-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallCentralSnatMap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/central-snat-map"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallCentralSnatMap API operation for FortiOS deletes the specified Central Snat Map.
// Returns error for service API and SDK errors.
// See the firewall - central-snat-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallCentralSnatMap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/central-snat-map"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallCentralSnatMap API operation for FortiOS gets the Central Snat Map
// with the specified index value.
// Returns the requested Central Snat Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - central-snat-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallCentralSnatMap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/central-snat-map"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateFirewallSslSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssl - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSslSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.ssl/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallSslSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the firewall.ssl - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSslSetting(mkey string) (err error) {

	//No unset API for firewall.ssl - setting
	return
}

// ReadFirewallSslSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssl - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSslSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.ssl/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateFirewallSshLocalKey API operation for FortiOS creates a new Local Key.
// Returns the index value of the Local Key and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - local-key chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallSshLocalKey(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.ssh/local-key"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallSshLocalKey API operation for FortiOS updates the specified Local Key.
// Returns the index value of the Local Key and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - local-key chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSshLocalKey(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.ssh/local-key"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallSshLocalKey API operation for FortiOS deletes the specified Local Key.
// Returns error for service API and SDK errors.
// See the firewall.ssh - local-key chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSshLocalKey(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.ssh/local-key"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallSshLocalKey API operation for FortiOS gets the Local Key
// with the specified index value.
// Returns the requested Local Key value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - local-key chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSshLocalKey(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.ssh/local-key"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallSshLocalCa API operation for FortiOS creates a new Local Ca.
// Returns the index value of the Local Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - local-ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallSshLocalCa(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.ssh/local-ca"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallSshLocalCa API operation for FortiOS updates the specified Local Ca.
// Returns the index value of the Local Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - local-ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSshLocalCa(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.ssh/local-ca"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallSshLocalCa API operation for FortiOS deletes the specified Local Ca.
// Returns error for service API and SDK errors.
// See the firewall.ssh - local-ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSshLocalCa(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.ssh/local-ca"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallSshLocalCa API operation for FortiOS gets the Local Ca
// with the specified index value.
// Returns the requested Local Ca value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - local-ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSshLocalCa(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.ssh/local-ca"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateFirewallSshSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSshSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.ssh/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallSshSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the firewall.ssh - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSshSetting(mkey string) (err error) {

	//No unset API for firewall.ssh - setting
	return
}

// ReadFirewallSshSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSshSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.ssh/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateFirewallSshHostKey API operation for FortiOS creates a new Host Key.
// Returns the index value of the Host Key and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - host-key chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallSshHostKey(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.ssh/host-key"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallSshHostKey API operation for FortiOS updates the specified Host Key.
// Returns the index value of the Host Key and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - host-key chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSshHostKey(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.ssh/host-key"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallSshHostKey API operation for FortiOS deletes the specified Host Key.
// Returns error for service API and SDK errors.
// See the firewall.ssh - host-key chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSshHostKey(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.ssh/host-key"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallSshHostKey API operation for FortiOS gets the Host Key
// with the specified index value.
// Returns the requested Host Key value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall.ssh - host-key chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSshHostKey(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.ssh/host-key"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateFirewallIpTranslation API operation for FortiOS creates a new Ip Translation.
// Returns the index value of the Ip Translation and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ip-translation chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallIpTranslation(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/ip-translation"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateFirewallIpTranslation API operation for FortiOS updates the specified Ip Translation.
// Returns the index value of the Ip Translation and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ip-translation chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallIpTranslation(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ip-translation"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallIpTranslation API operation for FortiOS deletes the specified Ip Translation.
// Returns error for service API and SDK errors.
// See the firewall - ip-translation chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallIpTranslation(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ip-translation"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadFirewallIpTranslation API operation for FortiOS gets the Ip Translation
// with the specified index value.
// Returns the requested Ip Translation value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ip-translation chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallIpTranslation(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ip-translation"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateFirewallIpv6EhFilter API operation for FortiOS updates the specified Ipv6 Eh Filter.
// Returns the index value of the Ipv6 Eh Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ipv6-eh-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallIpv6EhFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ipv6-eh-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFirewallIpv6EhFilter API operation for FortiOS deletes the specified Ipv6 Eh Filter.
// Returns error for service API and SDK errors.
// See the firewall - ipv6-eh-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallIpv6EhFilter(mkey string) (err error) {

	//No unset API for firewall - ipv6-eh-filter
	return
}

// ReadFirewallIpv6EhFilter API operation for FortiOS gets the Ipv6 Eh Filter
// with the specified index value.
// Returns the requested Ipv6 Eh Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ipv6-eh-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallIpv6EhFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ipv6-eh-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateVpnCertificateCa API operation for FortiOS creates a new Ca.
// Returns the index value of the Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnCertificateCa(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.certificate/ca"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnCertificateCa API operation for FortiOS updates the specified Ca.
// Returns the index value of the Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnCertificateCa(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.certificate/ca"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnCertificateCa API operation for FortiOS deletes the specified Ca.
// Returns error for service API and SDK errors.
// See the vpn.certificate - ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnCertificateCa(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.certificate/ca"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnCertificateCa API operation for FortiOS gets the Ca
// with the specified index value.
// Returns the requested Ca value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnCertificateCa(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.certificate/ca"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnCertificateLocal API operation for FortiOS creates a new Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnCertificateLocal(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.certificate/local"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnCertificateLocal API operation for FortiOS updates the specified Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnCertificateLocal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.certificate/local"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnCertificateLocal API operation for FortiOS deletes the specified Local.
// Returns error for service API and SDK errors.
// See the vpn.certificate - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnCertificateLocal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.certificate/local"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnCertificateLocal API operation for FortiOS gets the Local
// with the specified index value.
// Returns the requested Local value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnCertificateLocal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.certificate/local"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnCertificateCrl API operation for FortiOS creates a new Crl.
// Returns the index value of the Crl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - crl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnCertificateCrl(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.certificate/crl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnCertificateCrl API operation for FortiOS updates the specified Crl.
// Returns the index value of the Crl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - crl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnCertificateCrl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.certificate/crl"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnCertificateCrl API operation for FortiOS deletes the specified Crl.
// Returns error for service API and SDK errors.
// See the vpn.certificate - crl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnCertificateCrl(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.certificate/crl"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnCertificateCrl API operation for FortiOS gets the Crl
// with the specified index value.
// Returns the requested Crl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - crl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnCertificateCrl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.certificate/crl"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnCertificateRemote API operation for FortiOS creates a new Remote.
// Returns the index value of the Remote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - remote chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnCertificateRemote(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.certificate/remote"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnCertificateRemote API operation for FortiOS updates the specified Remote.
// Returns the index value of the Remote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - remote chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnCertificateRemote(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.certificate/remote"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnCertificateRemote API operation for FortiOS deletes the specified Remote.
// Returns error for service API and SDK errors.
// See the vpn.certificate - remote chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnCertificateRemote(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.certificate/remote"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnCertificateRemote API operation for FortiOS gets the Remote
// with the specified index value.
// Returns the requested Remote value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - remote chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnCertificateRemote(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.certificate/remote"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnCertificateOcspServer API operation for FortiOS creates a new Ocsp Server.
// Returns the index value of the Ocsp Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - ocsp-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnCertificateOcspServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.certificate/ocsp-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnCertificateOcspServer API operation for FortiOS updates the specified Ocsp Server.
// Returns the index value of the Ocsp Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - ocsp-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnCertificateOcspServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.certificate/ocsp-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnCertificateOcspServer API operation for FortiOS deletes the specified Ocsp Server.
// Returns error for service API and SDK errors.
// See the vpn.certificate - ocsp-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnCertificateOcspServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.certificate/ocsp-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnCertificateOcspServer API operation for FortiOS gets the Ocsp Server
// with the specified index value.
// Returns the requested Ocsp Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - ocsp-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnCertificateOcspServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.certificate/ocsp-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateVpnCertificateSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnCertificateSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.certificate/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnCertificateSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the vpn.certificate - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnCertificateSetting(mkey string) (err error) {

	//No unset API for vpn.certificate - setting
	return
}

// ReadVpnCertificateSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.certificate - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnCertificateSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.certificate/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateVpnSslWebRealm API operation for FortiOS creates a new Realm.
// Returns the index value of the Realm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - realm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnSslWebRealm(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ssl.web/realm"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnSslWebRealm API operation for FortiOS updates the specified Realm.
// Returns the index value of the Realm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - realm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnSslWebRealm(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ssl.web/realm"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnSslWebRealm API operation for FortiOS deletes the specified Realm.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - realm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnSslWebRealm(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ssl.web/realm"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnSslWebRealm API operation for FortiOS gets the Realm
// with the specified index value.
// Returns the requested Realm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - realm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnSslWebRealm(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ssl.web/realm"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnSslWebHostCheckSoftware API operation for FortiOS creates a new Host Check Software.
// Returns the index value of the Host Check Software and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - host-check-software chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnSslWebHostCheckSoftware(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ssl.web/host-check-software"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnSslWebHostCheckSoftware API operation for FortiOS updates the specified Host Check Software.
// Returns the index value of the Host Check Software and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - host-check-software chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnSslWebHostCheckSoftware(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ssl.web/host-check-software"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnSslWebHostCheckSoftware API operation for FortiOS deletes the specified Host Check Software.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - host-check-software chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnSslWebHostCheckSoftware(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ssl.web/host-check-software"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnSslWebHostCheckSoftware API operation for FortiOS gets the Host Check Software
// with the specified index value.
// Returns the requested Host Check Software value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - host-check-software chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnSslWebHostCheckSoftware(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ssl.web/host-check-software"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnSslWebPortal API operation for FortiOS creates a new Portal.
// Returns the index value of the Portal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnSslWebPortal(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ssl.web/portal"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnSslWebPortal API operation for FortiOS updates the specified Portal.
// Returns the index value of the Portal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnSslWebPortal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ssl.web/portal"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnSslWebPortal API operation for FortiOS deletes the specified Portal.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnSslWebPortal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ssl.web/portal"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnSslWebPortal API operation for FortiOS gets the Portal
// with the specified index value.
// Returns the requested Portal value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - portal chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnSslWebPortal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ssl.web/portal"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnSslWebUserGroupBookmark API operation for FortiOS creates a new User Group Bookmark.
// Returns the index value of the User Group Bookmark and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - user-group-bookmark chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnSslWebUserGroupBookmark(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ssl.web/user-group-bookmark"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnSslWebUserGroupBookmark API operation for FortiOS updates the specified User Group Bookmark.
// Returns the index value of the User Group Bookmark and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - user-group-bookmark chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnSslWebUserGroupBookmark(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ssl.web/user-group-bookmark"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnSslWebUserGroupBookmark API operation for FortiOS deletes the specified User Group Bookmark.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - user-group-bookmark chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnSslWebUserGroupBookmark(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ssl.web/user-group-bookmark"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnSslWebUserGroupBookmark API operation for FortiOS gets the User Group Bookmark
// with the specified index value.
// Returns the requested User Group Bookmark value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - user-group-bookmark chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnSslWebUserGroupBookmark(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ssl.web/user-group-bookmark"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnSslWebUserBookmark API operation for FortiOS creates a new User Bookmark.
// Returns the index value of the User Bookmark and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - user-bookmark chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnSslWebUserBookmark(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ssl.web/user-bookmark"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnSslWebUserBookmark API operation for FortiOS updates the specified User Bookmark.
// Returns the index value of the User Bookmark and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - user-bookmark chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnSslWebUserBookmark(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ssl.web/user-bookmark"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnSslWebUserBookmark API operation for FortiOS deletes the specified User Bookmark.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - user-bookmark chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnSslWebUserBookmark(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ssl.web/user-bookmark"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnSslWebUserBookmark API operation for FortiOS gets the User Bookmark
// with the specified index value.
// Returns the requested User Bookmark value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl.web - user-bookmark chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnSslWebUserBookmark(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ssl.web/user-bookmark"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateVpnSslSettings API operation for FortiOS updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnSslSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ssl/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnSslSettings API operation for FortiOS deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the vpn.ssl - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnSslSettings(mkey string) (err error) {

	//No unset API for vpn.ssl - settings
	return
}

// ReadVpnSslSettings API operation for FortiOS gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ssl - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnSslSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ssl/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateVpnIpsecPhase1 API operation for FortiOS creates a new Phase1.
// Returns the index value of the Phase1 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase1 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnIpsecPhase1(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/phase1"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnIpsecPhase1 API operation for FortiOS updates the specified Phase1.
// Returns the index value of the Phase1 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase1 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnIpsecPhase1(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/phase1"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnIpsecPhase1 API operation for FortiOS deletes the specified Phase1.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase1 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnIpsecPhase1(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/phase1"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnIpsecPhase1 API operation for FortiOS gets the Phase1
// with the specified index value.
// Returns the requested Phase1 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase1 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnIpsecPhase1(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/phase1"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnIpsecPhase2 API operation for FortiOS creates a new Phase2.
// Returns the index value of the Phase2 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase2 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnIpsecPhase2(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/phase2"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnIpsecPhase2 API operation for FortiOS updates the specified Phase2.
// Returns the index value of the Phase2 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase2 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnIpsecPhase2(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/phase2"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnIpsecPhase2 API operation for FortiOS deletes the specified Phase2.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase2 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnIpsecPhase2(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/phase2"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnIpsecPhase2 API operation for FortiOS gets the Phase2
// with the specified index value.
// Returns the requested Phase2 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase2 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnIpsecPhase2(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/phase2"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnIpsecManualkey API operation for FortiOS creates a new Manualkey.
// Returns the index value of the Manualkey and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - manualkey chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnIpsecManualkey(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/manualkey"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnIpsecManualkey API operation for FortiOS updates the specified Manualkey.
// Returns the index value of the Manualkey and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - manualkey chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnIpsecManualkey(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/manualkey"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnIpsecManualkey API operation for FortiOS deletes the specified Manualkey.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - manualkey chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnIpsecManualkey(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/manualkey"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnIpsecManualkey API operation for FortiOS gets the Manualkey
// with the specified index value.
// Returns the requested Manualkey value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - manualkey chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnIpsecManualkey(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/manualkey"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnIpsecConcentrator API operation for FortiOS creates a new Concentrator.
// Returns the index value of the Concentrator and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - concentrator chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnIpsecConcentrator(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/concentrator"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnIpsecConcentrator API operation for FortiOS updates the specified Concentrator.
// Returns the index value of the Concentrator and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - concentrator chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnIpsecConcentrator(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/concentrator"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnIpsecConcentrator API operation for FortiOS deletes the specified Concentrator.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - concentrator chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnIpsecConcentrator(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/concentrator"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnIpsecConcentrator API operation for FortiOS gets the Concentrator
// with the specified index value.
// Returns the requested Concentrator value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - concentrator chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnIpsecConcentrator(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/concentrator"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnIpsecPhase1Interface API operation for FortiOS creates a new Phase1 Interface.
// Returns the index value of the Phase1 Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase1-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnIpsecPhase1Interface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/phase1-interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnIpsecPhase1Interface API operation for FortiOS updates the specified Phase1 Interface.
// Returns the index value of the Phase1 Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase1-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnIpsecPhase1Interface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/phase1-interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnIpsecPhase1Interface API operation for FortiOS deletes the specified Phase1 Interface.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase1-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnIpsecPhase1Interface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/phase1-interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnIpsecPhase1Interface API operation for FortiOS gets the Phase1 Interface
// with the specified index value.
// Returns the requested Phase1 Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase1-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnIpsecPhase1Interface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/phase1-interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnIpsecPhase2Interface API operation for FortiOS creates a new Phase2 Interface.
// Returns the index value of the Phase2 Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase2-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnIpsecPhase2Interface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnIpsecPhase2Interface API operation for FortiOS updates the specified Phase2 Interface.
// Returns the index value of the Phase2 Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase2-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnIpsecPhase2Interface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnIpsecPhase2Interface API operation for FortiOS deletes the specified Phase2 Interface.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase2-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnIpsecPhase2Interface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnIpsecPhase2Interface API operation for FortiOS gets the Phase2 Interface
// with the specified index value.
// Returns the requested Phase2 Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - phase2-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnIpsecPhase2Interface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnIpsecManualkeyInterface API operation for FortiOS creates a new Manualkey Interface.
// Returns the index value of the Manualkey Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - manualkey-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnIpsecManualkeyInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/manualkey-interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnIpsecManualkeyInterface API operation for FortiOS updates the specified Manualkey Interface.
// Returns the index value of the Manualkey Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - manualkey-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnIpsecManualkeyInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/manualkey-interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnIpsecManualkeyInterface API operation for FortiOS deletes the specified Manualkey Interface.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - manualkey-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnIpsecManualkeyInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/manualkey-interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnIpsecManualkeyInterface API operation for FortiOS gets the Manualkey Interface
// with the specified index value.
// Returns the requested Manualkey Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - manualkey-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnIpsecManualkeyInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/manualkey-interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVpnIpsecForticlient API operation for FortiOS creates a new Forticlient.
// Returns the index value of the Forticlient and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - forticlient chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVpnIpsecForticlient(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/forticlient"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVpnIpsecForticlient API operation for FortiOS updates the specified Forticlient.
// Returns the index value of the Forticlient and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - forticlient chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnIpsecForticlient(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/forticlient"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnIpsecForticlient API operation for FortiOS deletes the specified Forticlient.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - forticlient chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnIpsecForticlient(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/forticlient"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVpnIpsecForticlient API operation for FortiOS gets the Forticlient
// with the specified index value.
// Returns the requested Forticlient value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn.ipsec - forticlient chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnIpsecForticlient(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/forticlient"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateVpnPptp API operation for FortiOS updates the specified Pptp.
// Returns the index value of the Pptp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - pptp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnPptp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn/pptp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnPptp API operation for FortiOS deletes the specified Pptp.
// Returns error for service API and SDK errors.
// See the vpn - pptp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnPptp(mkey string) (err error) {

	//No unset API for vpn - pptp
	return
}

// ReadVpnPptp API operation for FortiOS gets the Pptp
// with the specified index value.
// Returns the requested Pptp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - pptp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnPptp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn/pptp"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateVpnL2Tp API operation for FortiOS updates the specified L2Tp.
// Returns the index value of the L2Tp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - l2tp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVpnL2Tp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn/l2tp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVpnL2Tp API operation for FortiOS deletes the specified L2Tp.
// Returns error for service API and SDK errors.
// See the vpn - l2tp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVpnL2Tp(mkey string) (err error) {

	//No unset API for vpn - l2tp
	return
}

// ReadVpnL2Tp API operation for FortiOS gets the L2Tp
// with the specified index value.
// Returns the requested L2Tp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - l2tp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVpnL2Tp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn/l2tp"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateCertificateCa API operation for FortiOS creates a new Ca.
// Returns the index value of the Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateCertificateCa(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/certificate/ca"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateCertificateCa API operation for FortiOS updates the specified Ca.
// Returns the index value of the Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateCertificateCa(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/certificate/ca"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteCertificateCa API operation for FortiOS deletes the specified Ca.
// Returns error for service API and SDK errors.
// See the certificate - ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteCertificateCa(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/certificate/ca"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadCertificateCa API operation for FortiOS gets the Ca
// with the specified index value.
// Returns the requested Ca value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - ca chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadCertificateCa(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/certificate/ca"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateCertificateLocal API operation for FortiOS creates a new Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateCertificateLocal(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/certificate/local"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateCertificateLocal API operation for FortiOS updates the specified Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateCertificateLocal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/certificate/local"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteCertificateLocal API operation for FortiOS deletes the specified Local.
// Returns error for service API and SDK errors.
// See the certificate - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteCertificateLocal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/certificate/local"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadCertificateLocal API operation for FortiOS gets the Local
// with the specified index value.
// Returns the requested Local value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadCertificateLocal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/certificate/local"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateCertificateCrl API operation for FortiOS creates a new Crl.
// Returns the index value of the Crl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - crl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateCertificateCrl(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/certificate/crl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateCertificateCrl API operation for FortiOS updates the specified Crl.
// Returns the index value of the Crl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - crl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateCertificateCrl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/certificate/crl"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteCertificateCrl API operation for FortiOS deletes the specified Crl.
// Returns error for service API and SDK errors.
// See the certificate - crl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteCertificateCrl(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/certificate/crl"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadCertificateCrl API operation for FortiOS gets the Crl
// with the specified index value.
// Returns the requested Crl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the certificate - crl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadCertificateCrl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/certificate/crl"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebfilterFtgdLocalCat API operation for FortiOS creates a new Ftgd Local Cat.
// Returns the index value of the Ftgd Local Cat and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ftgd-local-cat chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebfilterFtgdLocalCat(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/webfilter/ftgd-local-cat"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebfilterFtgdLocalCat API operation for FortiOS updates the specified Ftgd Local Cat.
// Returns the index value of the Ftgd Local Cat and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ftgd-local-cat chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterFtgdLocalCat(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/ftgd-local-cat"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterFtgdLocalCat API operation for FortiOS deletes the specified Ftgd Local Cat.
// Returns error for service API and SDK errors.
// See the webfilter - ftgd-local-cat chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterFtgdLocalCat(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/webfilter/ftgd-local-cat"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebfilterFtgdLocalCat API operation for FortiOS gets the Ftgd Local Cat
// with the specified index value.
// Returns the requested Ftgd Local Cat value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ftgd-local-cat chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterFtgdLocalCat(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/ftgd-local-cat"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebfilterContent API operation for FortiOS creates a new Content.
// Returns the index value of the Content and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - content chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebfilterContent(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/webfilter/content"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebfilterContent API operation for FortiOS updates the specified Content.
// Returns the index value of the Content and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - content chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterContent(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/content"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterContent API operation for FortiOS deletes the specified Content.
// Returns error for service API and SDK errors.
// See the webfilter - content chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterContent(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/webfilter/content"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebfilterContent API operation for FortiOS gets the Content
// with the specified index value.
// Returns the requested Content value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - content chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterContent(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/content"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebfilterContentHeader API operation for FortiOS creates a new Content Header.
// Returns the index value of the Content Header and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - content-header chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebfilterContentHeader(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/webfilter/content-header"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebfilterContentHeader API operation for FortiOS updates the specified Content Header.
// Returns the index value of the Content Header and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - content-header chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterContentHeader(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/content-header"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterContentHeader API operation for FortiOS deletes the specified Content Header.
// Returns error for service API and SDK errors.
// See the webfilter - content-header chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterContentHeader(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/webfilter/content-header"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebfilterContentHeader API operation for FortiOS gets the Content Header
// with the specified index value.
// Returns the requested Content Header value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - content-header chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterContentHeader(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/content-header"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebfilterUrlfilter API operation for FortiOS creates a new Urlfilter.
// Returns the index value of the Urlfilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - urlfilter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebfilterUrlfilter(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/webfilter/urlfilter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebfilterUrlfilter API operation for FortiOS updates the specified Urlfilter.
// Returns the index value of the Urlfilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - urlfilter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterUrlfilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/urlfilter"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterUrlfilter API operation for FortiOS deletes the specified Urlfilter.
// Returns error for service API and SDK errors.
// See the webfilter - urlfilter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterUrlfilter(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/webfilter/urlfilter"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebfilterUrlfilter API operation for FortiOS gets the Urlfilter
// with the specified index value.
// Returns the requested Urlfilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - urlfilter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterUrlfilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/urlfilter"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateWebfilterIpsUrlfilterSetting API operation for FortiOS updates the specified Ips Urlfilter Setting.
// Returns the index value of the Ips Urlfilter Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterIpsUrlfilterSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/ips-urlfilter-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterIpsUrlfilterSetting API operation for FortiOS deletes the specified Ips Urlfilter Setting.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterIpsUrlfilterSetting(mkey string) (err error) {

	//No unset API for webfilter - ips-urlfilter-setting
	return
}

// ReadWebfilterIpsUrlfilterSetting API operation for FortiOS gets the Ips Urlfilter Setting
// with the specified index value.
// Returns the requested Ips Urlfilter Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterIpsUrlfilterSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/ips-urlfilter-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateWebfilterIpsUrlfilterSetting6 API operation for FortiOS updates the specified Ips Urlfilter Setting6.
// Returns the index value of the Ips Urlfilter Setting6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-setting6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterIpsUrlfilterSetting6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/ips-urlfilter-setting6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterIpsUrlfilterSetting6 API operation for FortiOS deletes the specified Ips Urlfilter Setting6.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-setting6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterIpsUrlfilterSetting6(mkey string) (err error) {

	//No unset API for webfilter - ips-urlfilter-setting6
	return
}

// ReadWebfilterIpsUrlfilterSetting6 API operation for FortiOS gets the Ips Urlfilter Setting6
// with the specified index value.
// Returns the requested Ips Urlfilter Setting6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-setting6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterIpsUrlfilterSetting6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/ips-urlfilter-setting6"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateWebfilterIpsUrlfilterCacheSetting API operation for FortiOS updates the specified Ips Urlfilter Cache Setting.
// Returns the index value of the Ips Urlfilter Cache Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-cache-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterIpsUrlfilterCacheSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/ips-urlfilter-cache-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterIpsUrlfilterCacheSetting API operation for FortiOS deletes the specified Ips Urlfilter Cache Setting.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-cache-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterIpsUrlfilterCacheSetting(mkey string) (err error) {

	//No unset API for webfilter - ips-urlfilter-cache-setting
	return
}

// ReadWebfilterIpsUrlfilterCacheSetting API operation for FortiOS gets the Ips Urlfilter Cache Setting
// with the specified index value.
// Returns the requested Ips Urlfilter Cache Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ips-urlfilter-cache-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterIpsUrlfilterCacheSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/ips-urlfilter-cache-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateWebfilterProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebfilterProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/webfilter/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebfilterProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the webfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/webfilter/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebfilterProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateWebfilterFortiguard API operation for FortiOS updates the specified Fortiguard.
// Returns the index value of the Fortiguard and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterFortiguard(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/fortiguard"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterFortiguard API operation for FortiOS deletes the specified Fortiguard.
// Returns error for service API and SDK errors.
// See the webfilter - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterFortiguard(mkey string) (err error) {

	//No unset API for webfilter - fortiguard
	return
}

// ReadWebfilterFortiguard API operation for FortiOS gets the Fortiguard
// with the specified index value.
// Returns the requested Fortiguard value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterFortiguard(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/fortiguard"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateWebfilterOverride API operation for FortiOS creates a new Override.
// Returns the index value of the Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - override chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebfilterOverride(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/webfilter/override"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebfilterOverride API operation for FortiOS updates the specified Override.
// Returns the index value of the Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - override chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterOverride(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/override"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterOverride API operation for FortiOS deletes the specified Override.
// Returns error for service API and SDK errors.
// See the webfilter - override chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterOverride(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/webfilter/override"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebfilterOverride API operation for FortiOS gets the Override
// with the specified index value.
// Returns the requested Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - override chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterOverride(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/override"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebfilterFtgdLocalRating API operation for FortiOS creates a new Ftgd Local Rating.
// Returns the index value of the Ftgd Local Rating and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ftgd-local-rating chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebfilterFtgdLocalRating(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/webfilter/ftgd-local-rating"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebfilterFtgdLocalRating API operation for FortiOS updates the specified Ftgd Local Rating.
// Returns the index value of the Ftgd Local Rating and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ftgd-local-rating chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterFtgdLocalRating(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/ftgd-local-rating"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterFtgdLocalRating API operation for FortiOS deletes the specified Ftgd Local Rating.
// Returns error for service API and SDK errors.
// See the webfilter - ftgd-local-rating chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterFtgdLocalRating(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/webfilter/ftgd-local-rating"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebfilterFtgdLocalRating API operation for FortiOS gets the Ftgd Local Rating
// with the specified index value.
// Returns the requested Ftgd Local Rating value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - ftgd-local-rating chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterFtgdLocalRating(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/ftgd-local-rating"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebfilterSearchEngine API operation for FortiOS creates a new Search Engine.
// Returns the index value of the Search Engine and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - search-engine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebfilterSearchEngine(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/webfilter/search-engine"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebfilterSearchEngine API operation for FortiOS updates the specified Search Engine.
// Returns the index value of the Search Engine and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - search-engine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebfilterSearchEngine(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/webfilter/search-engine"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebfilterSearchEngine API operation for FortiOS deletes the specified Search Engine.
// Returns error for service API and SDK errors.
// See the webfilter - search-engine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebfilterSearchEngine(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/webfilter/search-engine"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebfilterSearchEngine API operation for FortiOS gets the Search Engine
// with the specified index value.
// Returns the requested Search Engine value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the webfilter - search-engine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebfilterSearchEngine(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/webfilter/search-engine"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateIpsSensor API operation for FortiOS creates a new Sensor.
// Returns the index value of the Sensor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - sensor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateIpsSensor(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/ips/sensor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateIpsSensor API operation for FortiOS updates the specified Sensor.
// Returns the index value of the Sensor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - sensor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIpsSensor(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ips/sensor"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIpsSensor API operation for FortiOS deletes the specified Sensor.
// Returns error for service API and SDK errors.
// See the ips - sensor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIpsSensor(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/ips/sensor"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadIpsSensor API operation for FortiOS gets the Sensor
// with the specified index value.
// Returns the requested Sensor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - sensor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIpsSensor(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ips/sensor"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateIpsDecoder API operation for FortiOS creates a new Decoder.
// Returns the index value of the Decoder and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - decoder chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateIpsDecoder(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/ips/decoder"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateIpsDecoder API operation for FortiOS updates the specified Decoder.
// Returns the index value of the Decoder and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - decoder chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIpsDecoder(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ips/decoder"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIpsDecoder API operation for FortiOS deletes the specified Decoder.
// Returns error for service API and SDK errors.
// See the ips - decoder chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIpsDecoder(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/ips/decoder"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadIpsDecoder API operation for FortiOS gets the Decoder
// with the specified index value.
// Returns the requested Decoder value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - decoder chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIpsDecoder(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ips/decoder"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateIpsRule API operation for FortiOS creates a new Rule.
// Returns the index value of the Rule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateIpsRule(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/ips/rule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateIpsRule API operation for FortiOS updates the specified Rule.
// Returns the index value of the Rule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIpsRule(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ips/rule"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIpsRule API operation for FortiOS deletes the specified Rule.
// Returns error for service API and SDK errors.
// See the ips - rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIpsRule(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/ips/rule"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadIpsRule API operation for FortiOS gets the Rule
// with the specified index value.
// Returns the requested Rule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIpsRule(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ips/rule"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateIpsRuleSettings API operation for FortiOS creates a new Rule Settings.
// Returns the index value of the Rule Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - rule-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateIpsRuleSettings(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/ips/rule-settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateIpsRuleSettings API operation for FortiOS updates the specified Rule Settings.
// Returns the index value of the Rule Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - rule-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIpsRuleSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ips/rule-settings"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIpsRuleSettings API operation for FortiOS deletes the specified Rule Settings.
// Returns error for service API and SDK errors.
// See the ips - rule-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIpsRuleSettings(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/ips/rule-settings"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadIpsRuleSettings API operation for FortiOS gets the Rule Settings
// with the specified index value.
// Returns the requested Rule Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - rule-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIpsRuleSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ips/rule-settings"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateIpsCustom API operation for FortiOS creates a new Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateIpsCustom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/ips/custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateIpsCustom API operation for FortiOS updates the specified Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIpsCustom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ips/custom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIpsCustom API operation for FortiOS deletes the specified Custom.
// Returns error for service API and SDK errors.
// See the ips - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIpsCustom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/ips/custom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadIpsCustom API operation for FortiOS gets the Custom
// with the specified index value.
// Returns the requested Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIpsCustom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ips/custom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateIpsGlobal API operation for FortiOS updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIpsGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ips/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIpsGlobal API operation for FortiOS deletes the specified Global.
// Returns error for service API and SDK errors.
// See the ips - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIpsGlobal(mkey string) (err error) {

	//No unset API for ips - global
	return
}

// ReadIpsGlobal API operation for FortiOS gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIpsGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ips/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateIpsSettings API operation for FortiOS updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIpsSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ips/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIpsSettings API operation for FortiOS deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the ips - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIpsSettings(mkey string) (err error) {

	//No unset API for ips - settings
	return
}

// ReadIpsSettings API operation for FortiOS gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ips - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIpsSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ips/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateWebProxyProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebProxyProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/web-proxy/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebProxyProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebProxyProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/web-proxy/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebProxyProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the web-proxy - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebProxyProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/web-proxy/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebProxyProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebProxyProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/web-proxy/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateWebProxyGlobal API operation for FortiOS updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebProxyGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/web-proxy/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebProxyGlobal API operation for FortiOS deletes the specified Global.
// Returns error for service API and SDK errors.
// See the web-proxy - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebProxyGlobal(mkey string) (err error) {

	//No unset API for web-proxy - global
	return
}

// ReadWebProxyGlobal API operation for FortiOS gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebProxyGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/web-proxy/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateWebProxyExplicit API operation for FortiOS updates the specified Explicit.
// Returns the index value of the Explicit and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - explicit chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebProxyExplicit(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/web-proxy/explicit"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebProxyExplicit API operation for FortiOS deletes the specified Explicit.
// Returns error for service API and SDK errors.
// See the web-proxy - explicit chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebProxyExplicit(mkey string) (err error) {

	//No unset API for web-proxy - explicit
	return
}

// ReadWebProxyExplicit API operation for FortiOS gets the Explicit
// with the specified index value.
// Returns the requested Explicit value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - explicit chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebProxyExplicit(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/web-proxy/explicit"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateWebProxyForwardServer API operation for FortiOS creates a new Forward Server.
// Returns the index value of the Forward Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - forward-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebProxyForwardServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/web-proxy/forward-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebProxyForwardServer API operation for FortiOS updates the specified Forward Server.
// Returns the index value of the Forward Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - forward-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebProxyForwardServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/web-proxy/forward-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebProxyForwardServer API operation for FortiOS deletes the specified Forward Server.
// Returns error for service API and SDK errors.
// See the web-proxy - forward-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebProxyForwardServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/web-proxy/forward-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebProxyForwardServer API operation for FortiOS gets the Forward Server
// with the specified index value.
// Returns the requested Forward Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - forward-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebProxyForwardServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/web-proxy/forward-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebProxyForwardServerGroup API operation for FortiOS creates a new Forward Server Group.
// Returns the index value of the Forward Server Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - forward-server-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebProxyForwardServerGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/web-proxy/forward-server-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebProxyForwardServerGroup API operation for FortiOS updates the specified Forward Server Group.
// Returns the index value of the Forward Server Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - forward-server-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebProxyForwardServerGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/web-proxy/forward-server-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebProxyForwardServerGroup API operation for FortiOS deletes the specified Forward Server Group.
// Returns error for service API and SDK errors.
// See the web-proxy - forward-server-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebProxyForwardServerGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/web-proxy/forward-server-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebProxyForwardServerGroup API operation for FortiOS gets the Forward Server Group
// with the specified index value.
// Returns the requested Forward Server Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - forward-server-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebProxyForwardServerGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/web-proxy/forward-server-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebProxyDebugUrl API operation for FortiOS creates a new Debug Url.
// Returns the index value of the Debug Url and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - debug-url chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebProxyDebugUrl(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/web-proxy/debug-url"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebProxyDebugUrl API operation for FortiOS updates the specified Debug Url.
// Returns the index value of the Debug Url and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - debug-url chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebProxyDebugUrl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/web-proxy/debug-url"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebProxyDebugUrl API operation for FortiOS deletes the specified Debug Url.
// Returns error for service API and SDK errors.
// See the web-proxy - debug-url chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebProxyDebugUrl(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/web-proxy/debug-url"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebProxyDebugUrl API operation for FortiOS gets the Debug Url
// with the specified index value.
// Returns the requested Debug Url value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - debug-url chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebProxyDebugUrl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/web-proxy/debug-url"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebProxyWisp API operation for FortiOS creates a new Wisp.
// Returns the index value of the Wisp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - wisp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebProxyWisp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/web-proxy/wisp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebProxyWisp API operation for FortiOS updates the specified Wisp.
// Returns the index value of the Wisp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - wisp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebProxyWisp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/web-proxy/wisp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebProxyWisp API operation for FortiOS deletes the specified Wisp.
// Returns error for service API and SDK errors.
// See the web-proxy - wisp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebProxyWisp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/web-proxy/wisp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebProxyWisp API operation for FortiOS gets the Wisp
// with the specified index value.
// Returns the requested Wisp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - wisp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebProxyWisp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/web-proxy/wisp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWebProxyUrlMatch API operation for FortiOS creates a new Url Match.
// Returns the index value of the Url Match and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - url-match chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWebProxyUrlMatch(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/web-proxy/url-match"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWebProxyUrlMatch API operation for FortiOS updates the specified Url Match.
// Returns the index value of the Url Match and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - url-match chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWebProxyUrlMatch(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/web-proxy/url-match"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWebProxyUrlMatch API operation for FortiOS deletes the specified Url Match.
// Returns error for service API and SDK errors.
// See the web-proxy - url-match chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWebProxyUrlMatch(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/web-proxy/url-match"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWebProxyUrlMatch API operation for FortiOS gets the Url Match
// with the specified index value.
// Returns the requested Url Match value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the web-proxy - url-match chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWebProxyUrlMatch(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/web-proxy/url-match"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateWanoptWebcache API operation for FortiOS updates the specified Webcache.
// Returns the index value of the Webcache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - webcache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWanoptWebcache(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wanopt/webcache"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWanoptWebcache API operation for FortiOS deletes the specified Webcache.
// Returns error for service API and SDK errors.
// See the wanopt - webcache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWanoptWebcache(mkey string) (err error) {

	//No unset API for wanopt - webcache
	return
}

// ReadWanoptWebcache API operation for FortiOS gets the Webcache
// with the specified index value.
// Returns the requested Webcache value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - webcache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWanoptWebcache(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wanopt/webcache"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateWanoptSettings API operation for FortiOS updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWanoptSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wanopt/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWanoptSettings API operation for FortiOS deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the wanopt - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWanoptSettings(mkey string) (err error) {

	//No unset API for wanopt - settings
	return
}

// ReadWanoptSettings API operation for FortiOS gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWanoptSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wanopt/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateWanoptPeer API operation for FortiOS creates a new Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - peer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWanoptPeer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wanopt/peer"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWanoptPeer API operation for FortiOS updates the specified Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - peer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWanoptPeer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wanopt/peer"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWanoptPeer API operation for FortiOS deletes the specified Peer.
// Returns error for service API and SDK errors.
// See the wanopt - peer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWanoptPeer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wanopt/peer"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWanoptPeer API operation for FortiOS gets the Peer
// with the specified index value.
// Returns the requested Peer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - peer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWanoptPeer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wanopt/peer"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWanoptAuthGroup API operation for FortiOS creates a new Auth Group.
// Returns the index value of the Auth Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - auth-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWanoptAuthGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wanopt/auth-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWanoptAuthGroup API operation for FortiOS updates the specified Auth Group.
// Returns the index value of the Auth Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - auth-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWanoptAuthGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wanopt/auth-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWanoptAuthGroup API operation for FortiOS deletes the specified Auth Group.
// Returns error for service API and SDK errors.
// See the wanopt - auth-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWanoptAuthGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wanopt/auth-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWanoptAuthGroup API operation for FortiOS gets the Auth Group
// with the specified index value.
// Returns the requested Auth Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - auth-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWanoptAuthGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wanopt/auth-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWanoptProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWanoptProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wanopt/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWanoptProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWanoptProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wanopt/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWanoptProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the wanopt - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWanoptProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wanopt/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWanoptProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWanoptProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wanopt/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWanoptContentDeliveryNetworkRule API operation for FortiOS creates a new Content Delivery Network Rule.
// Returns the index value of the Content Delivery Network Rule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - content-delivery-network-rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWanoptContentDeliveryNetworkRule(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/wanopt/content-delivery-network-rule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWanoptContentDeliveryNetworkRule API operation for FortiOS updates the specified Content Delivery Network Rule.
// Returns the index value of the Content Delivery Network Rule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - content-delivery-network-rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWanoptContentDeliveryNetworkRule(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wanopt/content-delivery-network-rule"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWanoptContentDeliveryNetworkRule API operation for FortiOS deletes the specified Content Delivery Network Rule.
// Returns error for service API and SDK errors.
// See the wanopt - content-delivery-network-rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWanoptContentDeliveryNetworkRule(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/wanopt/content-delivery-network-rule"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWanoptContentDeliveryNetworkRule API operation for FortiOS gets the Content Delivery Network Rule
// with the specified index value.
// Returns the requested Content Delivery Network Rule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - content-delivery-network-rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWanoptContentDeliveryNetworkRule(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wanopt/content-delivery-network-rule"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateWanoptCacheService API operation for FortiOS updates the specified Cache Service.
// Returns the index value of the Cache Service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - cache-service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWanoptCacheService(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wanopt/cache-service"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWanoptCacheService API operation for FortiOS deletes the specified Cache Service.
// Returns error for service API and SDK errors.
// See the wanopt - cache-service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWanoptCacheService(mkey string) (err error) {

	//No unset API for wanopt - cache-service
	return
}

// ReadWanoptCacheService API operation for FortiOS gets the Cache Service
// with the specified index value.
// Returns the requested Cache Service value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - cache-service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWanoptCacheService(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wanopt/cache-service"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateWanoptRemoteStorage API operation for FortiOS updates the specified Remote Storage.
// Returns the index value of the Remote Storage and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - remote-storage chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWanoptRemoteStorage(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/wanopt/remote-storage"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWanoptRemoteStorage API operation for FortiOS deletes the specified Remote Storage.
// Returns error for service API and SDK errors.
// See the wanopt - remote-storage chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWanoptRemoteStorage(mkey string) (err error) {

	//No unset API for wanopt - remote-storage
	return
}

// ReadWanoptRemoteStorage API operation for FortiOS gets the Remote Storage
// with the specified index value.
// Returns the requested Remote Storage value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the wanopt - remote-storage chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWanoptRemoteStorage(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/wanopt/remote-storage"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateFtpProxyExplicit API operation for FortiOS updates the specified Explicit.
// Returns the index value of the Explicit and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ftp-proxy - explicit chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFtpProxyExplicit(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ftp-proxy/explicit"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteFtpProxyExplicit API operation for FortiOS deletes the specified Explicit.
// Returns error for service API and SDK errors.
// See the ftp-proxy - explicit chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFtpProxyExplicit(mkey string) (err error) {

	//No unset API for ftp-proxy - explicit
	return
}

// ReadFtpProxyExplicit API operation for FortiOS gets the Explicit
// with the specified index value.
// Returns the requested Explicit value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ftp-proxy - explicit chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFtpProxyExplicit(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ftp-proxy/explicit"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateApplicationName API operation for FortiOS creates a new Name.
// Returns the index value of the Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateApplicationName(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/application/name"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateApplicationName API operation for FortiOS updates the specified Name.
// Returns the index value of the Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateApplicationName(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/application/name"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteApplicationName API operation for FortiOS deletes the specified Name.
// Returns error for service API and SDK errors.
// See the application - name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteApplicationName(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/application/name"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadApplicationName API operation for FortiOS gets the Name
// with the specified index value.
// Returns the requested Name value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - name chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadApplicationName(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/application/name"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateApplicationCustom API operation for FortiOS creates a new Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateApplicationCustom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/application/custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateApplicationCustom API operation for FortiOS updates the specified Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateApplicationCustom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/application/custom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteApplicationCustom API operation for FortiOS deletes the specified Custom.
// Returns error for service API and SDK errors.
// See the application - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteApplicationCustom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/application/custom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadApplicationCustom API operation for FortiOS gets the Custom
// with the specified index value.
// Returns the requested Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - custom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadApplicationCustom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/application/custom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateApplicationRuleSettings API operation for FortiOS creates a new Rule Settings.
// Returns the index value of the Rule Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - rule-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateApplicationRuleSettings(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/application/rule-settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateApplicationRuleSettings API operation for FortiOS updates the specified Rule Settings.
// Returns the index value of the Rule Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - rule-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateApplicationRuleSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/application/rule-settings"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteApplicationRuleSettings API operation for FortiOS deletes the specified Rule Settings.
// Returns error for service API and SDK errors.
// See the application - rule-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteApplicationRuleSettings(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/application/rule-settings"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadApplicationRuleSettings API operation for FortiOS gets the Rule Settings
// with the specified index value.
// Returns the requested Rule Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - rule-settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadApplicationRuleSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/application/rule-settings"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateApplicationList API operation for FortiOS creates a new List.
// Returns the index value of the List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateApplicationList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/application/list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateApplicationList API operation for FortiOS updates the specified List.
// Returns the index value of the List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateApplicationList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/application/list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteApplicationList API operation for FortiOS deletes the specified List.
// Returns error for service API and SDK errors.
// See the application - list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteApplicationList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/application/list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadApplicationList API operation for FortiOS gets the List
// with the specified index value.
// Returns the requested List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadApplicationList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/application/list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateApplicationGroup API operation for FortiOS creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateApplicationGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/application/group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateApplicationGroup API operation for FortiOS updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateApplicationGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/application/group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteApplicationGroup API operation for FortiOS deletes the specified Group.
// Returns error for service API and SDK errors.
// See the application - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteApplicationGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/application/group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadApplicationGroup API operation for FortiOS gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the application - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadApplicationGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/application/group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateDlpFilepattern API operation for FortiOS creates a new Filepattern.
// Returns the index value of the Filepattern and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - filepattern chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDlpFilepattern(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/dlp/filepattern"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateDlpFilepattern API operation for FortiOS updates the specified Filepattern.
// Returns the index value of the Filepattern and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - filepattern chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDlpFilepattern(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/dlp/filepattern"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteDlpFilepattern API operation for FortiOS deletes the specified Filepattern.
// Returns error for service API and SDK errors.
// See the dlp - filepattern chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDlpFilepattern(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/dlp/filepattern"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadDlpFilepattern API operation for FortiOS gets the Filepattern
// with the specified index value.
// Returns the requested Filepattern value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - filepattern chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDlpFilepattern(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/dlp/filepattern"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateDlpFpSensitivity API operation for FortiOS creates a new Fp Sensitivity.
// Returns the index value of the Fp Sensitivity and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - fp-sensitivity chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDlpFpSensitivity(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/dlp/fp-sensitivity"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateDlpFpSensitivity API operation for FortiOS updates the specified Fp Sensitivity.
// Returns the index value of the Fp Sensitivity and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - fp-sensitivity chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDlpFpSensitivity(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/dlp/fp-sensitivity"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteDlpFpSensitivity API operation for FortiOS deletes the specified Fp Sensitivity.
// Returns error for service API and SDK errors.
// See the dlp - fp-sensitivity chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDlpFpSensitivity(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/dlp/fp-sensitivity"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadDlpFpSensitivity API operation for FortiOS gets the Fp Sensitivity
// with the specified index value.
// Returns the requested Fp Sensitivity value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - fp-sensitivity chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDlpFpSensitivity(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/dlp/fp-sensitivity"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateDlpFpDocSource API operation for FortiOS creates a new Fp Doc Source.
// Returns the index value of the Fp Doc Source and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - fp-doc-source chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDlpFpDocSource(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/dlp/fp-doc-source"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateDlpFpDocSource API operation for FortiOS updates the specified Fp Doc Source.
// Returns the index value of the Fp Doc Source and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - fp-doc-source chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDlpFpDocSource(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/dlp/fp-doc-source"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteDlpFpDocSource API operation for FortiOS deletes the specified Fp Doc Source.
// Returns error for service API and SDK errors.
// See the dlp - fp-doc-source chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDlpFpDocSource(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/dlp/fp-doc-source"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadDlpFpDocSource API operation for FortiOS gets the Fp Doc Source
// with the specified index value.
// Returns the requested Fp Doc Source value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - fp-doc-source chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDlpFpDocSource(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/dlp/fp-doc-source"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateDlpSensor API operation for FortiOS creates a new Sensor.
// Returns the index value of the Sensor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - sensor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDlpSensor(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/dlp/sensor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateDlpSensor API operation for FortiOS updates the specified Sensor.
// Returns the index value of the Sensor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - sensor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDlpSensor(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/dlp/sensor"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteDlpSensor API operation for FortiOS deletes the specified Sensor.
// Returns error for service API and SDK errors.
// See the dlp - sensor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDlpSensor(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/dlp/sensor"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadDlpSensor API operation for FortiOS gets the Sensor
// with the specified index value.
// Returns the requested Sensor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - sensor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDlpSensor(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/dlp/sensor"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateDlpSettings API operation for FortiOS updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDlpSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/dlp/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteDlpSettings API operation for FortiOS deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the dlp - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDlpSettings(mkey string) (err error) {

	//No unset API for dlp - settings
	return
}

// ReadDlpSettings API operation for FortiOS gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dlp - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDlpSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/dlp/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSpamfilterBword API operation for FortiOS creates a new Bword.
// Returns the index value of the Bword and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - bword chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSpamfilterBword(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/spamfilter/bword"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSpamfilterBword API operation for FortiOS updates the specified Bword.
// Returns the index value of the Bword and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - bword chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSpamfilterBword(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/spamfilter/bword"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSpamfilterBword API operation for FortiOS deletes the specified Bword.
// Returns error for service API and SDK errors.
// See the spamfilter - bword chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSpamfilterBword(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/spamfilter/bword"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSpamfilterBword API operation for FortiOS gets the Bword
// with the specified index value.
// Returns the requested Bword value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - bword chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSpamfilterBword(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/spamfilter/bword"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSpamfilterBwl API operation for FortiOS creates a new Bwl.
// Returns the index value of the Bwl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - bwl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSpamfilterBwl(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/spamfilter/bwl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSpamfilterBwl API operation for FortiOS updates the specified Bwl.
// Returns the index value of the Bwl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - bwl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSpamfilterBwl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/spamfilter/bwl"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSpamfilterBwl API operation for FortiOS deletes the specified Bwl.
// Returns error for service API and SDK errors.
// See the spamfilter - bwl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSpamfilterBwl(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/spamfilter/bwl"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSpamfilterBwl API operation for FortiOS gets the Bwl
// with the specified index value.
// Returns the requested Bwl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - bwl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSpamfilterBwl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/spamfilter/bwl"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSpamfilterMheader API operation for FortiOS creates a new Mheader.
// Returns the index value of the Mheader and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - mheader chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSpamfilterMheader(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/spamfilter/mheader"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSpamfilterMheader API operation for FortiOS updates the specified Mheader.
// Returns the index value of the Mheader and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - mheader chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSpamfilterMheader(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/spamfilter/mheader"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSpamfilterMheader API operation for FortiOS deletes the specified Mheader.
// Returns error for service API and SDK errors.
// See the spamfilter - mheader chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSpamfilterMheader(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/spamfilter/mheader"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSpamfilterMheader API operation for FortiOS gets the Mheader
// with the specified index value.
// Returns the requested Mheader value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - mheader chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSpamfilterMheader(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/spamfilter/mheader"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSpamfilterDnsbl API operation for FortiOS creates a new Dnsbl.
// Returns the index value of the Dnsbl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - dnsbl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSpamfilterDnsbl(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/spamfilter/dnsbl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSpamfilterDnsbl API operation for FortiOS updates the specified Dnsbl.
// Returns the index value of the Dnsbl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - dnsbl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSpamfilterDnsbl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/spamfilter/dnsbl"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSpamfilterDnsbl API operation for FortiOS deletes the specified Dnsbl.
// Returns error for service API and SDK errors.
// See the spamfilter - dnsbl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSpamfilterDnsbl(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/spamfilter/dnsbl"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSpamfilterDnsbl API operation for FortiOS gets the Dnsbl
// with the specified index value.
// Returns the requested Dnsbl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - dnsbl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSpamfilterDnsbl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/spamfilter/dnsbl"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSpamfilterIptrust API operation for FortiOS creates a new Iptrust.
// Returns the index value of the Iptrust and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - iptrust chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSpamfilterIptrust(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/spamfilter/iptrust"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSpamfilterIptrust API operation for FortiOS updates the specified Iptrust.
// Returns the index value of the Iptrust and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - iptrust chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSpamfilterIptrust(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/spamfilter/iptrust"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSpamfilterIptrust API operation for FortiOS deletes the specified Iptrust.
// Returns error for service API and SDK errors.
// See the spamfilter - iptrust chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSpamfilterIptrust(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/spamfilter/iptrust"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSpamfilterIptrust API operation for FortiOS gets the Iptrust
// with the specified index value.
// Returns the requested Iptrust value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - iptrust chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSpamfilterIptrust(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/spamfilter/iptrust"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSpamfilterProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSpamfilterProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/spamfilter/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSpamfilterProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSpamfilterProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/spamfilter/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSpamfilterProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the spamfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSpamfilterProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/spamfilter/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSpamfilterProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSpamfilterProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/spamfilter/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSpamfilterFortishield API operation for FortiOS updates the specified Fortishield.
// Returns the index value of the Fortishield and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - fortishield chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSpamfilterFortishield(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/spamfilter/fortishield"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSpamfilterFortishield API operation for FortiOS deletes the specified Fortishield.
// Returns error for service API and SDK errors.
// See the spamfilter - fortishield chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSpamfilterFortishield(mkey string) (err error) {

	//No unset API for spamfilter - fortishield
	return
}

// ReadSpamfilterFortishield API operation for FortiOS gets the Fortishield
// with the specified index value.
// Returns the requested Fortishield value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - fortishield chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSpamfilterFortishield(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/spamfilter/fortishield"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSpamfilterOptions API operation for FortiOS updates the specified Options.
// Returns the index value of the Options and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - options chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSpamfilterOptions(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/spamfilter/options"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSpamfilterOptions API operation for FortiOS deletes the specified Options.
// Returns error for service API and SDK errors.
// See the spamfilter - options chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSpamfilterOptions(mkey string) (err error) {

	//No unset API for spamfilter - options
	return
}

// ReadSpamfilterOptions API operation for FortiOS gets the Options
// with the specified index value.
// Returns the requested Options value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the spamfilter - options chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSpamfilterOptions(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/spamfilter/options"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogThreatWeight API operation for FortiOS updates the specified Threat Weight.
// Returns the index value of the Threat Weight and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - threat-weight chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogThreatWeight(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log/threat-weight"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogThreatWeight API operation for FortiOS deletes the specified Threat Weight.
// Returns error for service API and SDK errors.
// See the log - threat-weight chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogThreatWeight(mkey string) (err error) {

	//No unset API for log - threat-weight
	return
}

// ReadLogThreatWeight API operation for FortiOS gets the Threat Weight
// with the specified index value.
// Returns the requested Threat Weight value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - threat-weight chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogThreatWeight(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log/threat-weight"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateLogCustomField API operation for FortiOS creates a new Custom Field.
// Returns the index value of the Custom Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - custom-field chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateLogCustomField(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/log/custom-field"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateLogCustomField API operation for FortiOS updates the specified Custom Field.
// Returns the index value of the Custom Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - custom-field chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogCustomField(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log/custom-field"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogCustomField API operation for FortiOS deletes the specified Custom Field.
// Returns error for service API and SDK errors.
// See the log - custom-field chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogCustomField(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/log/custom-field"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadLogCustomField API operation for FortiOS gets the Custom Field
// with the specified index value.
// Returns the requested Custom Field value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - custom-field chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogCustomField(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log/custom-field"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateLogSyslogdSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogdSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogdSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogdSetting(mkey string) (err error) {

	//No unset API for log.syslogd - setting
	return
}

// ReadLogSyslogdSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogdSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogdOverrideSetting API operation for FortiOS updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogdOverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogdOverrideSetting API operation for FortiOS deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogdOverrideSetting(mkey string) (err error) {

	//No unset API for log.syslogd - override-setting
	return
}

// ReadLogSyslogdOverrideSetting API operation for FortiOS gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogdOverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogdFilter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogdFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogdFilter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogdFilter(mkey string) (err error) {

	//No unset API for log.syslogd - filter
	return
}

// ReadLogSyslogdFilter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogdFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogdOverrideFilter API operation for FortiOS updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogdOverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogdOverrideFilter API operation for FortiOS deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogdOverrideFilter(mkey string) (err error) {

	//No unset API for log.syslogd - override-filter
	return
}

// ReadLogSyslogdOverrideFilter API operation for FortiOS gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogdOverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd2Setting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd2Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd2/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd2Setting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd2Setting(mkey string) (err error) {

	//No unset API for log.syslogd2 - setting
	return
}

// ReadLogSyslogd2Setting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd2Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd2/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd2OverrideSetting API operation for FortiOS updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd2OverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd2/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd2OverrideSetting API operation for FortiOS deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd2OverrideSetting(mkey string) (err error) {

	//No unset API for log.syslogd2 - override-setting
	return
}

// ReadLogSyslogd2OverrideSetting API operation for FortiOS gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd2OverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd2/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd2Filter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd2Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd2/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd2Filter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd2Filter(mkey string) (err error) {

	//No unset API for log.syslogd2 - filter
	return
}

// ReadLogSyslogd2Filter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd2Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd2/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd2OverrideFilter API operation for FortiOS updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd2OverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd2/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd2OverrideFilter API operation for FortiOS deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd2OverrideFilter(mkey string) (err error) {

	//No unset API for log.syslogd2 - override-filter
	return
}

// ReadLogSyslogd2OverrideFilter API operation for FortiOS gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd2OverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd2/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd3Setting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd3Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd3/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd3Setting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd3Setting(mkey string) (err error) {

	//No unset API for log.syslogd3 - setting
	return
}

// ReadLogSyslogd3Setting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd3Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd3/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd3OverrideSetting API operation for FortiOS updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd3OverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd3/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd3OverrideSetting API operation for FortiOS deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd3OverrideSetting(mkey string) (err error) {

	//No unset API for log.syslogd3 - override-setting
	return
}

// ReadLogSyslogd3OverrideSetting API operation for FortiOS gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd3OverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd3/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd3Filter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd3Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd3/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd3Filter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd3Filter(mkey string) (err error) {

	//No unset API for log.syslogd3 - filter
	return
}

// ReadLogSyslogd3Filter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd3Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd3/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd3OverrideFilter API operation for FortiOS updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd3OverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd3/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd3OverrideFilter API operation for FortiOS deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd3OverrideFilter(mkey string) (err error) {

	//No unset API for log.syslogd3 - override-filter
	return
}

// ReadLogSyslogd3OverrideFilter API operation for FortiOS gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd3OverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd3/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd4Setting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd4Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd4/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd4Setting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd4Setting(mkey string) (err error) {

	//No unset API for log.syslogd4 - setting
	return
}

// ReadLogSyslogd4Setting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd4Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd4/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd4OverrideSetting API operation for FortiOS updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd4OverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd4/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd4OverrideSetting API operation for FortiOS deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd4OverrideSetting(mkey string) (err error) {

	//No unset API for log.syslogd4 - override-setting
	return
}

// ReadLogSyslogd4OverrideSetting API operation for FortiOS gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd4OverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd4/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd4Filter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd4Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd4/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd4Filter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd4Filter(mkey string) (err error) {

	//No unset API for log.syslogd4 - filter
	return
}

// ReadLogSyslogd4Filter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd4Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd4/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSyslogd4OverrideFilter API operation for FortiOS updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd4OverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd4/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd4OverrideFilter API operation for FortiOS deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd4OverrideFilter(mkey string) (err error) {

	//No unset API for log.syslogd4 - override-filter
	return
}

// ReadLogSyslogd4OverrideFilter API operation for FortiOS gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd4 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd4OverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd4/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogWebtrendsSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.webtrends - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogWebtrendsSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.webtrends/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogWebtrendsSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.webtrends - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogWebtrendsSetting(mkey string) (err error) {

	//No unset API for log.webtrends - setting
	return
}

// ReadLogWebtrendsSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.webtrends - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogWebtrendsSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.webtrends/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogWebtrendsFilter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.webtrends - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogWebtrendsFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.webtrends/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogWebtrendsFilter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.webtrends - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogWebtrendsFilter(mkey string) (err error) {

	//No unset API for log.webtrends - filter
	return
}

// ReadLogWebtrendsFilter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.webtrends - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogWebtrendsFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.webtrends/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogMemoryGlobalSetting API operation for FortiOS updates the specified Global Setting.
// Returns the index value of the Global Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - global-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogMemoryGlobalSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.memory/global-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogMemoryGlobalSetting API operation for FortiOS deletes the specified Global Setting.
// Returns error for service API and SDK errors.
// See the log.memory - global-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogMemoryGlobalSetting(mkey string) (err error) {

	//No unset API for log.memory - global-setting
	return
}

// ReadLogMemoryGlobalSetting API operation for FortiOS gets the Global Setting
// with the specified index value.
// Returns the requested Global Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - global-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogMemoryGlobalSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.memory/global-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogMemorySetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogMemorySetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.memory/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogMemorySetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.memory - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogMemorySetting(mkey string) (err error) {

	//No unset API for log.memory - setting
	return
}

// ReadLogMemorySetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogMemorySetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.memory/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogMemoryFilter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogMemoryFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.memory/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogMemoryFilter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.memory - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogMemoryFilter(mkey string) (err error) {

	//No unset API for log.memory - filter
	return
}

// ReadLogMemoryFilter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogMemoryFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.memory/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogDiskSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.disk - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogDiskSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.disk/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogDiskSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.disk - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogDiskSetting(mkey string) (err error) {

	//No unset API for log.disk - setting
	return
}

// ReadLogDiskSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.disk - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogDiskSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.disk/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogDiskFilter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.disk - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogDiskFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.disk/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogDiskFilter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.disk - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogDiskFilter(mkey string) (err error) {

	//No unset API for log.disk - filter
	return
}

// ReadLogDiskFilter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.disk - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogDiskFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.disk/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogEventfilter API operation for FortiOS updates the specified Eventfilter.
// Returns the index value of the Eventfilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - eventfilter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogEventfilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log/eventfilter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogEventfilter API operation for FortiOS deletes the specified Eventfilter.
// Returns error for service API and SDK errors.
// See the log - eventfilter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogEventfilter(mkey string) (err error) {

	//No unset API for log - eventfilter
	return
}

// ReadLogEventfilter API operation for FortiOS gets the Eventfilter
// with the specified index value.
// Returns the requested Eventfilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - eventfilter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogEventfilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log/eventfilter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortiguardSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortiguardSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortiguard/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortiguardSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.fortiguard - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortiguardSetting(mkey string) (err error) {

	//No unset API for log.fortiguard - setting
	return
}

// ReadLogFortiguardSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortiguardSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortiguard/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortiguardOverrideSetting API operation for FortiOS updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortiguardOverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortiguard/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortiguardOverrideSetting API operation for FortiOS deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.fortiguard - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortiguardOverrideSetting(mkey string) (err error) {

	//No unset API for log.fortiguard - override-setting
	return
}

// ReadLogFortiguardOverrideSetting API operation for FortiOS gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortiguardOverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortiguard/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortiguardFilter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortiguardFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortiguard/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortiguardFilter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.fortiguard - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortiguardFilter(mkey string) (err error) {

	//No unset API for log.fortiguard - filter
	return
}

// ReadLogFortiguardFilter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortiguardFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortiguard/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortiguardOverrideFilter API operation for FortiOS updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortiguardOverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortiguard/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortiguardOverrideFilter API operation for FortiOS deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.fortiguard - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortiguardOverrideFilter(mkey string) (err error) {

	//No unset API for log.fortiguard - override-filter
	return
}

// ReadLogFortiguardOverrideFilter API operation for FortiOS gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortiguardOverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortiguard/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogNullDeviceSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.null-device - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogNullDeviceSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.null-device/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogNullDeviceSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.null-device - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogNullDeviceSetting(mkey string) (err error) {

	//No unset API for log.null-device - setting
	return
}

// ReadLogNullDeviceSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.null-device - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogNullDeviceSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.null-device/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogNullDeviceFilter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.null-device - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogNullDeviceFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.null-device/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogNullDeviceFilter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.null-device - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogNullDeviceFilter(mkey string) (err error) {

	//No unset API for log.null-device - filter
	return
}

// ReadLogNullDeviceFilter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.null-device - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogNullDeviceFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.null-device/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSetting(mkey string) (err error) {

	//No unset API for log - setting
	return
}

// ReadLogSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogGuiDisplay API operation for FortiOS updates the specified Gui Display.
// Returns the index value of the Gui Display and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - gui-display chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogGuiDisplay(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log/gui-display"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogGuiDisplay API operation for FortiOS deletes the specified Gui Display.
// Returns error for service API and SDK errors.
// See the log - gui-display chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogGuiDisplay(mkey string) (err error) {

	//No unset API for log - gui-display
	return
}

// ReadLogGuiDisplay API operation for FortiOS gets the Gui Display
// with the specified index value.
// Returns the requested Gui Display value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - gui-display chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogGuiDisplay(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log/gui-display"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzerSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzerSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzerSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzerSetting(mkey string) (err error) {

	//No unset API for log.fortianalyzer - setting
	return
}

// ReadLogFortianalyzerSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzerSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzerOverrideSetting API operation for FortiOS updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzerOverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzerOverrideSetting API operation for FortiOS deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzerOverrideSetting(mkey string) (err error) {

	//No unset API for log.fortianalyzer - override-setting
	return
}

// ReadLogFortianalyzerOverrideSetting API operation for FortiOS gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzerOverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzerFilter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzerFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzerFilter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzerFilter(mkey string) (err error) {

	//No unset API for log.fortianalyzer - filter
	return
}

// ReadLogFortianalyzerFilter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzerFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzerOverrideFilter API operation for FortiOS updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzerOverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzerOverrideFilter API operation for FortiOS deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzerOverrideFilter(mkey string) (err error) {

	//No unset API for log.fortianalyzer - override-filter
	return
}

// ReadLogFortianalyzerOverrideFilter API operation for FortiOS gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzerOverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzer2Setting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer2Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer2/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer2Setting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer2Setting(mkey string) (err error) {

	//No unset API for log.fortianalyzer2 - setting
	return
}

// ReadLogFortianalyzer2Setting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer2Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer2/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzer2OverrideSetting API operation for FortiOS updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer2OverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer2/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer2OverrideSetting API operation for FortiOS deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer2OverrideSetting(mkey string) (err error) {

	//No unset API for log.fortianalyzer2 - override-setting
	return
}

// ReadLogFortianalyzer2OverrideSetting API operation for FortiOS gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer2OverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer2/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzer2Filter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer2Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer2/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer2Filter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer2Filter(mkey string) (err error) {

	//No unset API for log.fortianalyzer2 - filter
	return
}

// ReadLogFortianalyzer2Filter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer2Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer2/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzer2OverrideFilter API operation for FortiOS updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer2OverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer2/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer2OverrideFilter API operation for FortiOS deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer2OverrideFilter(mkey string) (err error) {

	//No unset API for log.fortianalyzer2 - override-filter
	return
}

// ReadLogFortianalyzer2OverrideFilter API operation for FortiOS gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer2OverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer2/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzer3Setting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer3Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer3/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer3Setting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer3Setting(mkey string) (err error) {

	//No unset API for log.fortianalyzer3 - setting
	return
}

// ReadLogFortianalyzer3Setting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer3Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer3/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzer3OverrideSetting API operation for FortiOS updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer3OverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer3/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer3OverrideSetting API operation for FortiOS deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer3OverrideSetting(mkey string) (err error) {

	//No unset API for log.fortianalyzer3 - override-setting
	return
}

// ReadLogFortianalyzer3OverrideSetting API operation for FortiOS gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - override-setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer3OverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer3/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzer3Filter API operation for FortiOS updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer3Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer3/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer3Filter API operation for FortiOS deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer3Filter(mkey string) (err error) {

	//No unset API for log.fortianalyzer3 - filter
	return
}

// ReadLogFortianalyzer3Filter API operation for FortiOS gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer3Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer3/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateLogFortianalyzer3OverrideFilter API operation for FortiOS updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer3OverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer3/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer3OverrideFilter API operation for FortiOS deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer3OverrideFilter(mkey string) (err error) {

	//No unset API for log.fortianalyzer3 - override-filter
	return
}

// ReadLogFortianalyzer3OverrideFilter API operation for FortiOS gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - override-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer3OverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer3/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateIcapServer API operation for FortiOS creates a new Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the icap - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateIcapServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/icap/server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateIcapServer API operation for FortiOS updates the specified Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the icap - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIcapServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/icap/server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIcapServer API operation for FortiOS deletes the specified Server.
// Returns error for service API and SDK errors.
// See the icap - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIcapServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/icap/server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadIcapServer API operation for FortiOS gets the Server
// with the specified index value.
// Returns the requested Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the icap - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIcapServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/icap/server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateIcapProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the icap - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateIcapProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/icap/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateIcapProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the icap - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateIcapProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/icap/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteIcapProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the icap - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteIcapProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/icap/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadIcapProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the icap - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadIcapProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/icap/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSshFilterProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ssh-filter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSshFilterProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/ssh-filter/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSshFilterProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ssh-filter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSshFilterProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/ssh-filter/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSshFilterProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the ssh-filter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSshFilterProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/ssh-filter/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSshFilterProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ssh-filter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSshFilterProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/ssh-filter/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserRadius API operation for FortiOS creates a new Radius.
// Returns the index value of the Radius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - radius chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserRadius(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/radius"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserRadius API operation for FortiOS updates the specified Radius.
// Returns the index value of the Radius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - radius chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserRadius(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/radius"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserRadius API operation for FortiOS deletes the specified Radius.
// Returns error for service API and SDK errors.
// See the user - radius chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserRadius(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/radius"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserRadius API operation for FortiOS gets the Radius
// with the specified index value.
// Returns the requested Radius value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - radius chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserRadius(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/radius"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserTacacs API operation for FortiOS creates a new Tacacs+.
// Returns the index value of the Tacacs+ and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - tacacs+ chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserTacacs(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/tacacs+"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserTacacs API operation for FortiOS updates the specified Tacacs+.
// Returns the index value of the Tacacs+ and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - tacacs+ chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserTacacs(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/tacacs+"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserTacacs API operation for FortiOS deletes the specified Tacacs+.
// Returns error for service API and SDK errors.
// See the user - tacacs+ chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserTacacs(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/tacacs+"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserTacacs API operation for FortiOS gets the Tacacs+
// with the specified index value.
// Returns the requested Tacacs+ value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - tacacs+ chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserTacacs(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/tacacs+"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserLdap API operation for FortiOS creates a new Ldap.
// Returns the index value of the Ldap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - ldap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserLdap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/ldap"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserLdap API operation for FortiOS updates the specified Ldap.
// Returns the index value of the Ldap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - ldap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserLdap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/ldap"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserLdap API operation for FortiOS deletes the specified Ldap.
// Returns error for service API and SDK errors.
// See the user - ldap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserLdap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/ldap"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserLdap API operation for FortiOS gets the Ldap
// with the specified index value.
// Returns the requested Ldap value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - ldap chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserLdap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/ldap"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserKrbKeytab API operation for FortiOS creates a new Krb Keytab.
// Returns the index value of the Krb Keytab and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - krb-keytab chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserKrbKeytab(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/krb-keytab"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserKrbKeytab API operation for FortiOS updates the specified Krb Keytab.
// Returns the index value of the Krb Keytab and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - krb-keytab chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserKrbKeytab(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/krb-keytab"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserKrbKeytab API operation for FortiOS deletes the specified Krb Keytab.
// Returns error for service API and SDK errors.
// See the user - krb-keytab chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserKrbKeytab(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/krb-keytab"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserKrbKeytab API operation for FortiOS gets the Krb Keytab
// with the specified index value.
// Returns the requested Krb Keytab value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - krb-keytab chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserKrbKeytab(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/krb-keytab"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserDomainController API operation for FortiOS creates a new Domain Controller.
// Returns the index value of the Domain Controller and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - domain-controller chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserDomainController(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/domain-controller"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserDomainController API operation for FortiOS updates the specified Domain Controller.
// Returns the index value of the Domain Controller and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - domain-controller chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserDomainController(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/domain-controller"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserDomainController API operation for FortiOS deletes the specified Domain Controller.
// Returns error for service API and SDK errors.
// See the user - domain-controller chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserDomainController(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/domain-controller"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserDomainController API operation for FortiOS gets the Domain Controller
// with the specified index value.
// Returns the requested Domain Controller value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - domain-controller chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserDomainController(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/domain-controller"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserPop3 API operation for FortiOS creates a new Pop3.
// Returns the index value of the Pop3 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - pop3 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserPop3(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/pop3"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserPop3 API operation for FortiOS updates the specified Pop3.
// Returns the index value of the Pop3 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - pop3 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserPop3(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/pop3"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserPop3 API operation for FortiOS deletes the specified Pop3.
// Returns error for service API and SDK errors.
// See the user - pop3 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserPop3(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/pop3"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserPop3 API operation for FortiOS gets the Pop3
// with the specified index value.
// Returns the requested Pop3 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - pop3 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserPop3(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/pop3"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserFsso API operation for FortiOS creates a new Fsso.
// Returns the index value of the Fsso and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fsso chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserFsso(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/fsso"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserFsso API operation for FortiOS updates the specified Fsso.
// Returns the index value of the Fsso and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fsso chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserFsso(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/fsso"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserFsso API operation for FortiOS deletes the specified Fsso.
// Returns error for service API and SDK errors.
// See the user - fsso chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserFsso(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/fsso"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserFsso API operation for FortiOS gets the Fsso
// with the specified index value.
// Returns the requested Fsso value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fsso chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserFsso(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/fsso"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserAdgrp API operation for FortiOS creates a new Adgrp.
// Returns the index value of the Adgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - adgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserAdgrp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/adgrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserAdgrp API operation for FortiOS updates the specified Adgrp.
// Returns the index value of the Adgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - adgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserAdgrp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/adgrp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserAdgrp API operation for FortiOS deletes the specified Adgrp.
// Returns error for service API and SDK errors.
// See the user - adgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserAdgrp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/adgrp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserAdgrp API operation for FortiOS gets the Adgrp
// with the specified index value.
// Returns the requested Adgrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - adgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserAdgrp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/adgrp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserFssoPolling API operation for FortiOS creates a new Fsso Polling.
// Returns the index value of the Fsso Polling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserFssoPolling(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/fsso-polling"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserFssoPolling API operation for FortiOS updates the specified Fsso Polling.
// Returns the index value of the Fsso Polling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserFssoPolling(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/fsso-polling"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserFssoPolling API operation for FortiOS deletes the specified Fsso Polling.
// Returns error for service API and SDK errors.
// See the user - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserFssoPolling(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/fsso-polling"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserFssoPolling API operation for FortiOS gets the Fsso Polling
// with the specified index value.
// Returns the requested Fsso Polling value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserFssoPolling(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/fsso-polling"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserFortitoken API operation for FortiOS creates a new Fortitoken.
// Returns the index value of the Fortitoken and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fortitoken chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserFortitoken(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/fortitoken"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserFortitoken API operation for FortiOS updates the specified Fortitoken.
// Returns the index value of the Fortitoken and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fortitoken chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserFortitoken(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/fortitoken"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserFortitoken API operation for FortiOS deletes the specified Fortitoken.
// Returns error for service API and SDK errors.
// See the user - fortitoken chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserFortitoken(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/fortitoken"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserFortitoken API operation for FortiOS gets the Fortitoken
// with the specified index value.
// Returns the requested Fortitoken value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - fortitoken chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserFortitoken(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/fortitoken"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserPasswordPolicy API operation for FortiOS creates a new Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserPasswordPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/password-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserPasswordPolicy API operation for FortiOS updates the specified Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserPasswordPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/password-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserPasswordPolicy API operation for FortiOS deletes the specified Password Policy.
// Returns error for service API and SDK errors.
// See the user - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserPasswordPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/password-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserPasswordPolicy API operation for FortiOS gets the Password Policy
// with the specified index value.
// Returns the requested Password Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserPasswordPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/password-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserLocal API operation for FortiOS creates a new Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserLocal(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/local"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserLocal API operation for FortiOS updates the specified Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserLocal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/local"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserLocal API operation for FortiOS deletes the specified Local.
// Returns error for service API and SDK errors.
// See the user - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserLocal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/local"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserLocal API operation for FortiOS gets the Local
// with the specified index value.
// Returns the requested Local value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - local chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserLocal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/local"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateUserSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the user - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserSetting(mkey string) (err error) {

	//No unset API for user - setting
	return
}

// ReadUserSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateUserPeer API operation for FortiOS creates a new Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserPeer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/peer"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserPeer API operation for FortiOS updates the specified Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserPeer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/peer"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserPeer API operation for FortiOS deletes the specified Peer.
// Returns error for service API and SDK errors.
// See the user - peer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserPeer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/peer"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserPeer API operation for FortiOS gets the Peer
// with the specified index value.
// Returns the requested Peer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peer chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserPeer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/peer"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserPeergrp API operation for FortiOS creates a new Peergrp.
// Returns the index value of the Peergrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peergrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserPeergrp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/peergrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserPeergrp API operation for FortiOS updates the specified Peergrp.
// Returns the index value of the Peergrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peergrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserPeergrp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/peergrp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserPeergrp API operation for FortiOS deletes the specified Peergrp.
// Returns error for service API and SDK errors.
// See the user - peergrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserPeergrp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/peergrp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserPeergrp API operation for FortiOS gets the Peergrp
// with the specified index value.
// Returns the requested Peergrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peergrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserPeergrp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/peergrp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateUserQuarantine API operation for FortiOS updates the specified Quarantine.
// Returns the index value of the Quarantine and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserQuarantine(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/quarantine"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserQuarantine API operation for FortiOS deletes the specified Quarantine.
// Returns error for service API and SDK errors.
// See the user - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserQuarantine(mkey string) (err error) {

	//No unset API for user - quarantine
	return
}

// ReadUserQuarantine API operation for FortiOS gets the Quarantine
// with the specified index value.
// Returns the requested Quarantine value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserQuarantine(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/quarantine"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateUserGroup API operation for FortiOS creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserGroup API operation for FortiOS updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserGroup API operation for FortiOS deletes the specified Group.
// Returns error for service API and SDK errors.
// See the user - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserGroup API operation for FortiOS gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserDeviceCategory API operation for FortiOS creates a new Device Category.
// Returns the index value of the Device Category and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-category chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserDeviceCategory(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/device-category"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserDeviceCategory API operation for FortiOS updates the specified Device Category.
// Returns the index value of the Device Category and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-category chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserDeviceCategory(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/device-category"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserDeviceCategory API operation for FortiOS deletes the specified Device Category.
// Returns error for service API and SDK errors.
// See the user - device-category chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserDeviceCategory(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/device-category"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserDeviceCategory API operation for FortiOS gets the Device Category
// with the specified index value.
// Returns the requested Device Category value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-category chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserDeviceCategory(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/device-category"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserDevice API operation for FortiOS creates a new Device.
// Returns the index value of the Device and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserDevice(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/device"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserDevice API operation for FortiOS updates the specified Device.
// Returns the index value of the Device and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserDevice(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/device"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserDevice API operation for FortiOS deletes the specified Device.
// Returns error for service API and SDK errors.
// See the user - device chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserDevice(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/device"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserDevice API operation for FortiOS gets the Device
// with the specified index value.
// Returns the requested Device value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserDevice(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/device"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserDeviceGroup API operation for FortiOS creates a new Device Group.
// Returns the index value of the Device Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserDeviceGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/device-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserDeviceGroup API operation for FortiOS updates the specified Device Group.
// Returns the index value of the Device Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserDeviceGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/device-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserDeviceGroup API operation for FortiOS deletes the specified Device Group.
// Returns error for service API and SDK errors.
// See the user - device-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserDeviceGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/device-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserDeviceGroup API operation for FortiOS gets the Device Group
// with the specified index value.
// Returns the requested Device Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserDeviceGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/device-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserDeviceAccessList API operation for FortiOS creates a new Device Access List.
// Returns the index value of the Device Access List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-access-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserDeviceAccessList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/device-access-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserDeviceAccessList API operation for FortiOS updates the specified Device Access List.
// Returns the index value of the Device Access List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-access-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserDeviceAccessList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/device-access-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserDeviceAccessList API operation for FortiOS deletes the specified Device Access List.
// Returns error for service API and SDK errors.
// See the user - device-access-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserDeviceAccessList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/device-access-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserDeviceAccessList API operation for FortiOS gets the Device Access List
// with the specified index value.
// Returns the requested Device Access List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - device-access-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserDeviceAccessList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/device-access-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserSecurityExemptList API operation for FortiOS creates a new Security Exempt List.
// Returns the index value of the Security Exempt List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - security-exempt-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserSecurityExemptList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/security-exempt-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserSecurityExemptList API operation for FortiOS updates the specified Security Exempt List.
// Returns the index value of the Security Exempt List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - security-exempt-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserSecurityExemptList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/security-exempt-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserSecurityExemptList API operation for FortiOS deletes the specified Security Exempt List.
// Returns error for service API and SDK errors.
// See the user - security-exempt-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserSecurityExemptList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/security-exempt-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserSecurityExemptList API operation for FortiOS gets the Security Exempt List
// with the specified index value.
// Returns the requested Security Exempt List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - security-exempt-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserSecurityExemptList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/security-exempt-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateVoipProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the voip - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVoipProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/voip/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateVoipProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the voip - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVoipProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/voip/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteVoipProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the voip - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVoipProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/voip/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadVoipProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the voip - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVoipProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/voip/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateDnsfilterDomainFilter API operation for FortiOS creates a new Domain Filter.
// Returns the index value of the Domain Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dnsfilter - domain-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDnsfilterDomainFilter(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/dnsfilter/domain-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateDnsfilterDomainFilter API operation for FortiOS updates the specified Domain Filter.
// Returns the index value of the Domain Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dnsfilter - domain-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDnsfilterDomainFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/dnsfilter/domain-filter"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteDnsfilterDomainFilter API operation for FortiOS deletes the specified Domain Filter.
// Returns error for service API and SDK errors.
// See the dnsfilter - domain-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDnsfilterDomainFilter(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/dnsfilter/domain-filter"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadDnsfilterDomainFilter API operation for FortiOS gets the Domain Filter
// with the specified index value.
// Returns the requested Domain Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dnsfilter - domain-filter chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDnsfilterDomainFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/dnsfilter/domain-filter"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateDnsfilterProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dnsfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDnsfilterProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/dnsfilter/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateDnsfilterProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dnsfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDnsfilterProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/dnsfilter/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteDnsfilterProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the dnsfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDnsfilterProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/dnsfilter/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadDnsfilterProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dnsfilter - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDnsfilterProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/dnsfilter/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateAntivirusSettings API operation for FortiOS updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAntivirusSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/antivirus/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAntivirusSettings API operation for FortiOS deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the antivirus - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAntivirusSettings(mkey string) (err error) {

	//No unset API for antivirus - settings
	return
}

// ReadAntivirusSettings API operation for FortiOS gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAntivirusSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/antivirus/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateAntivirusHeuristic API operation for FortiOS updates the specified Heuristic.
// Returns the index value of the Heuristic and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - heuristic chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAntivirusHeuristic(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/antivirus/heuristic"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAntivirusHeuristic API operation for FortiOS deletes the specified Heuristic.
// Returns error for service API and SDK errors.
// See the antivirus - heuristic chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAntivirusHeuristic(mkey string) (err error) {

	//No unset API for antivirus - heuristic
	return
}

// ReadAntivirusHeuristic API operation for FortiOS gets the Heuristic
// with the specified index value.
// Returns the requested Heuristic value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - heuristic chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAntivirusHeuristic(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/antivirus/heuristic"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateAntivirusQuarantine API operation for FortiOS updates the specified Quarantine.
// Returns the index value of the Quarantine and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAntivirusQuarantine(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/antivirus/quarantine"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAntivirusQuarantine API operation for FortiOS deletes the specified Quarantine.
// Returns error for service API and SDK errors.
// See the antivirus - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAntivirusQuarantine(mkey string) (err error) {

	//No unset API for antivirus - quarantine
	return
}

// ReadAntivirusQuarantine API operation for FortiOS gets the Quarantine
// with the specified index value.
// Returns the requested Quarantine value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - quarantine chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAntivirusQuarantine(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/antivirus/quarantine"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateAntivirusProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateAntivirusProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/antivirus/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateAntivirusProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAntivirusProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/antivirus/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAntivirusProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the antivirus - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAntivirusProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/antivirus/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadAntivirusProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the antivirus - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAntivirusProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/antivirus/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateReportDataset API operation for FortiOS creates a new Dataset.
// Returns the index value of the Dataset and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - dataset chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateReportDataset(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/report/dataset"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateReportDataset API operation for FortiOS updates the specified Dataset.
// Returns the index value of the Dataset and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - dataset chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateReportDataset(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/report/dataset"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteReportDataset API operation for FortiOS deletes the specified Dataset.
// Returns error for service API and SDK errors.
// See the report - dataset chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteReportDataset(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/report/dataset"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadReportDataset API operation for FortiOS gets the Dataset
// with the specified index value.
// Returns the requested Dataset value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - dataset chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadReportDataset(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/report/dataset"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateReportChart API operation for FortiOS creates a new Chart.
// Returns the index value of the Chart and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - chart chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateReportChart(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/report/chart"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateReportChart API operation for FortiOS updates the specified Chart.
// Returns the index value of the Chart and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - chart chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateReportChart(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/report/chart"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteReportChart API operation for FortiOS deletes the specified Chart.
// Returns error for service API and SDK errors.
// See the report - chart chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteReportChart(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/report/chart"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadReportChart API operation for FortiOS gets the Chart
// with the specified index value.
// Returns the requested Chart value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - chart chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadReportChart(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/report/chart"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateReportStyle API operation for FortiOS creates a new Style.
// Returns the index value of the Style and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - style chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateReportStyle(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/report/style"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateReportStyle API operation for FortiOS updates the specified Style.
// Returns the index value of the Style and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - style chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateReportStyle(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/report/style"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteReportStyle API operation for FortiOS deletes the specified Style.
// Returns error for service API and SDK errors.
// See the report - style chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteReportStyle(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/report/style"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadReportStyle API operation for FortiOS gets the Style
// with the specified index value.
// Returns the requested Style value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - style chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadReportStyle(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/report/style"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateReportTheme API operation for FortiOS creates a new Theme.
// Returns the index value of the Theme and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - theme chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateReportTheme(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/report/theme"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateReportTheme API operation for FortiOS updates the specified Theme.
// Returns the index value of the Theme and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - theme chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateReportTheme(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/report/theme"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteReportTheme API operation for FortiOS deletes the specified Theme.
// Returns error for service API and SDK errors.
// See the report - theme chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteReportTheme(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/report/theme"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadReportTheme API operation for FortiOS gets the Theme
// with the specified index value.
// Returns the requested Theme value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - theme chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadReportTheme(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/report/theme"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateReportLayout API operation for FortiOS creates a new Layout.
// Returns the index value of the Layout and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - layout chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateReportLayout(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/report/layout"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateReportLayout API operation for FortiOS updates the specified Layout.
// Returns the index value of the Layout and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - layout chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateReportLayout(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/report/layout"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteReportLayout API operation for FortiOS deletes the specified Layout.
// Returns error for service API and SDK errors.
// See the report - layout chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteReportLayout(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/report/layout"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadReportLayout API operation for FortiOS gets the Layout
// with the specified index value.
// Returns the requested Layout value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - layout chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadReportLayout(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/report/layout"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateReportSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateReportSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/report/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteReportSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the report - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteReportSetting(mkey string) (err error) {

	//No unset API for report - setting
	return
}

// ReadReportSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the report - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadReportSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/report/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateWafMainClass API operation for FortiOS creates a new Main Class.
// Returns the index value of the Main Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - main-class chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWafMainClass(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/waf/main-class"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWafMainClass API operation for FortiOS updates the specified Main Class.
// Returns the index value of the Main Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - main-class chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWafMainClass(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/waf/main-class"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWafMainClass API operation for FortiOS deletes the specified Main Class.
// Returns error for service API and SDK errors.
// See the waf - main-class chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWafMainClass(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/waf/main-class"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWafMainClass API operation for FortiOS gets the Main Class
// with the specified index value.
// Returns the requested Main Class value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - main-class chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWafMainClass(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/waf/main-class"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWafSubClass API operation for FortiOS creates a new Sub Class.
// Returns the index value of the Sub Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - sub-class chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWafSubClass(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/waf/sub-class"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWafSubClass API operation for FortiOS updates the specified Sub Class.
// Returns the index value of the Sub Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - sub-class chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWafSubClass(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/waf/sub-class"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWafSubClass API operation for FortiOS deletes the specified Sub Class.
// Returns error for service API and SDK errors.
// See the waf - sub-class chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWafSubClass(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/waf/sub-class"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWafSubClass API operation for FortiOS gets the Sub Class
// with the specified index value.
// Returns the requested Sub Class value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - sub-class chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWafSubClass(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/waf/sub-class"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWafSignature API operation for FortiOS creates a new Signature.
// Returns the index value of the Signature and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - signature chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWafSignature(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/waf/signature"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWafSignature API operation for FortiOS updates the specified Signature.
// Returns the index value of the Signature and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - signature chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWafSignature(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/waf/signature"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWafSignature API operation for FortiOS deletes the specified Signature.
// Returns error for service API and SDK errors.
// See the waf - signature chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWafSignature(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/waf/signature"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWafSignature API operation for FortiOS gets the Signature
// with the specified index value.
// Returns the requested Signature value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - signature chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWafSignature(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/waf/signature"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateWafProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateWafProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/waf/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateWafProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateWafProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/waf/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteWafProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the waf - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteWafProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/waf/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadWafProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the waf - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadWafProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/waf/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateAuthenticationScheme API operation for FortiOS creates a new Scheme.
// Returns the index value of the Scheme and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the authentication - scheme chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateAuthenticationScheme(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/authentication/scheme"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateAuthenticationScheme API operation for FortiOS updates the specified Scheme.
// Returns the index value of the Scheme and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the authentication - scheme chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAuthenticationScheme(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/authentication/scheme"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAuthenticationScheme API operation for FortiOS deletes the specified Scheme.
// Returns error for service API and SDK errors.
// See the authentication - scheme chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAuthenticationScheme(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/authentication/scheme"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadAuthenticationScheme API operation for FortiOS gets the Scheme
// with the specified index value.
// Returns the requested Scheme value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the authentication - scheme chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAuthenticationScheme(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/authentication/scheme"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateAuthenticationRule API operation for FortiOS creates a new Rule.
// Returns the index value of the Rule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the authentication - rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateAuthenticationRule(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/authentication/rule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateAuthenticationRule API operation for FortiOS updates the specified Rule.
// Returns the index value of the Rule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the authentication - rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAuthenticationRule(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/authentication/rule"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAuthenticationRule API operation for FortiOS deletes the specified Rule.
// Returns error for service API and SDK errors.
// See the authentication - rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAuthenticationRule(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/authentication/rule"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadAuthenticationRule API operation for FortiOS gets the Rule
// with the specified index value.
// Returns the requested Rule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the authentication - rule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAuthenticationRule(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/authentication/rule"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateAuthenticationSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the authentication - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAuthenticationSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/authentication/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAuthenticationSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the authentication - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAuthenticationSetting(mkey string) (err error) {

	//No unset API for authentication - setting
	return
}

// ReadAuthenticationSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the authentication - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAuthenticationSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/authentication/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateEndpointControlSettings API operation for FortiOS updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateEndpointControlSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/endpoint-control/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteEndpointControlSettings API operation for FortiOS deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the endpoint-control - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteEndpointControlSettings(mkey string) (err error) {

	//No unset API for endpoint-control - settings
	return
}

// ReadEndpointControlSettings API operation for FortiOS gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - settings chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadEndpointControlSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/endpoint-control/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateEndpointControlForticlientEms API operation for FortiOS creates a new Forticlient Ems.
// Returns the index value of the Forticlient Ems and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - forticlient-ems chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateEndpointControlForticlientEms(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/endpoint-control/forticlient-ems"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateEndpointControlForticlientEms API operation for FortiOS updates the specified Forticlient Ems.
// Returns the index value of the Forticlient Ems and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - forticlient-ems chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateEndpointControlForticlientEms(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/endpoint-control/forticlient-ems"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteEndpointControlForticlientEms API operation for FortiOS deletes the specified Forticlient Ems.
// Returns error for service API and SDK errors.
// See the endpoint-control - forticlient-ems chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteEndpointControlForticlientEms(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/endpoint-control/forticlient-ems"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadEndpointControlForticlientEms API operation for FortiOS gets the Forticlient Ems
// with the specified index value.
// Returns the requested Forticlient Ems value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - forticlient-ems chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadEndpointControlForticlientEms(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/endpoint-control/forticlient-ems"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateEndpointControlProfile API operation for FortiOS creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateEndpointControlProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/endpoint-control/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateEndpointControlProfile API operation for FortiOS updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateEndpointControlProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/endpoint-control/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteEndpointControlProfile API operation for FortiOS deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the endpoint-control - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteEndpointControlProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/endpoint-control/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadEndpointControlProfile API operation for FortiOS gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - profile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadEndpointControlProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/endpoint-control/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateEndpointControlForticlientRegistrationSync API operation for FortiOS creates a new Forticlient Registration Sync.
// Returns the index value of the Forticlient Registration Sync and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - forticlient-registration-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateEndpointControlForticlientRegistrationSync(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/endpoint-control/forticlient-registration-sync"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateEndpointControlForticlientRegistrationSync API operation for FortiOS updates the specified Forticlient Registration Sync.
// Returns the index value of the Forticlient Registration Sync and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - forticlient-registration-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateEndpointControlForticlientRegistrationSync(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/endpoint-control/forticlient-registration-sync"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteEndpointControlForticlientRegistrationSync API operation for FortiOS deletes the specified Forticlient Registration Sync.
// Returns error for service API and SDK errors.
// See the endpoint-control - forticlient-registration-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteEndpointControlForticlientRegistrationSync(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/endpoint-control/forticlient-registration-sync"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadEndpointControlForticlientRegistrationSync API operation for FortiOS gets the Forticlient Registration Sync
// with the specified index value.
// Returns the requested Forticlient Registration Sync value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - forticlient-registration-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadEndpointControlForticlientRegistrationSync(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/endpoint-control/forticlient-registration-sync"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateEndpointControlClient API operation for FortiOS creates a new Client.
// Returns the index value of the Client and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - client chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateEndpointControlClient(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/endpoint-control/client"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateEndpointControlClient API operation for FortiOS updates the specified Client.
// Returns the index value of the Client and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - client chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateEndpointControlClient(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/endpoint-control/client"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteEndpointControlClient API operation for FortiOS deletes the specified Client.
// Returns error for service API and SDK errors.
// See the endpoint-control - client chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteEndpointControlClient(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/endpoint-control/client"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadEndpointControlClient API operation for FortiOS gets the Client
// with the specified index value.
// Returns the requested Client value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - client chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadEndpointControlClient(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/endpoint-control/client"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateEndpointControlRegisteredForticlient API operation for FortiOS creates a new Registered Forticlient.
// Returns the index value of the Registered Forticlient and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - registered-forticlient chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateEndpointControlRegisteredForticlient(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/endpoint-control/registered-forticlient"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateEndpointControlRegisteredForticlient API operation for FortiOS updates the specified Registered Forticlient.
// Returns the index value of the Registered Forticlient and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - registered-forticlient chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateEndpointControlRegisteredForticlient(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/endpoint-control/registered-forticlient"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteEndpointControlRegisteredForticlient API operation for FortiOS deletes the specified Registered Forticlient.
// Returns error for service API and SDK errors.
// See the endpoint-control - registered-forticlient chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteEndpointControlRegisteredForticlient(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/endpoint-control/registered-forticlient"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadEndpointControlRegisteredForticlient API operation for FortiOS gets the Registered Forticlient
// with the specified index value.
// Returns the requested Registered Forticlient value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the endpoint-control - registered-forticlient chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadEndpointControlRegisteredForticlient(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/endpoint-control/registered-forticlient"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateAlertemailSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the alertemail - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAlertemailSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/alertemail/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAlertemailSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the alertemail - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAlertemailSetting(mkey string) (err error) {

	//No unset API for alertemail - setting
	return
}

// ReadAlertemailSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the alertemail - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAlertemailSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/alertemail/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterAccessList API operation for FortiOS creates a new Access List.
// Returns the index value of the Access List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterAccessList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/access-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterAccessList API operation for FortiOS updates the specified Access List.
// Returns the index value of the Access List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterAccessList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/access-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterAccessList API operation for FortiOS deletes the specified Access List.
// Returns error for service API and SDK errors.
// See the router - access-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterAccessList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/access-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterAccessList API operation for FortiOS gets the Access List
// with the specified index value.
// Returns the requested Access List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterAccessList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/access-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterAccessList6 API operation for FortiOS creates a new Access List6.
// Returns the index value of the Access List6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterAccessList6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/access-list6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterAccessList6 API operation for FortiOS updates the specified Access List6.
// Returns the index value of the Access List6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterAccessList6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/access-list6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterAccessList6 API operation for FortiOS deletes the specified Access List6.
// Returns error for service API and SDK errors.
// See the router - access-list6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterAccessList6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/access-list6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterAccessList6 API operation for FortiOS gets the Access List6
// with the specified index value.
// Returns the requested Access List6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterAccessList6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/access-list6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterAspathList API operation for FortiOS creates a new Aspath List.
// Returns the index value of the Aspath List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - aspath-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterAspathList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/aspath-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterAspathList API operation for FortiOS updates the specified Aspath List.
// Returns the index value of the Aspath List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - aspath-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterAspathList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/aspath-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterAspathList API operation for FortiOS deletes the specified Aspath List.
// Returns error for service API and SDK errors.
// See the router - aspath-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterAspathList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/aspath-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterAspathList API operation for FortiOS gets the Aspath List
// with the specified index value.
// Returns the requested Aspath List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - aspath-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterAspathList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/aspath-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterPrefixList API operation for FortiOS creates a new Prefix List.
// Returns the index value of the Prefix List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterPrefixList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/prefix-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterPrefixList API operation for FortiOS updates the specified Prefix List.
// Returns the index value of the Prefix List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterPrefixList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/prefix-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterPrefixList API operation for FortiOS deletes the specified Prefix List.
// Returns error for service API and SDK errors.
// See the router - prefix-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterPrefixList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/prefix-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterPrefixList API operation for FortiOS gets the Prefix List
// with the specified index value.
// Returns the requested Prefix List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterPrefixList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/prefix-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterPrefixList6 API operation for FortiOS creates a new Prefix List6.
// Returns the index value of the Prefix List6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterPrefixList6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/prefix-list6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterPrefixList6 API operation for FortiOS updates the specified Prefix List6.
// Returns the index value of the Prefix List6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterPrefixList6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/prefix-list6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterPrefixList6 API operation for FortiOS deletes the specified Prefix List6.
// Returns error for service API and SDK errors.
// See the router - prefix-list6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterPrefixList6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/prefix-list6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterPrefixList6 API operation for FortiOS gets the Prefix List6
// with the specified index value.
// Returns the requested Prefix List6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterPrefixList6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/prefix-list6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterKeyChain API operation for FortiOS creates a new Key Chain.
// Returns the index value of the Key Chain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - key-chain chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterKeyChain(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/key-chain"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterKeyChain API operation for FortiOS updates the specified Key Chain.
// Returns the index value of the Key Chain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - key-chain chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterKeyChain(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/key-chain"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterKeyChain API operation for FortiOS deletes the specified Key Chain.
// Returns error for service API and SDK errors.
// See the router - key-chain chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterKeyChain(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/key-chain"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterKeyChain API operation for FortiOS gets the Key Chain
// with the specified index value.
// Returns the requested Key Chain value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - key-chain chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterKeyChain(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/key-chain"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterCommunityList API operation for FortiOS creates a new Community List.
// Returns the index value of the Community List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - community-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterCommunityList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/community-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterCommunityList API operation for FortiOS updates the specified Community List.
// Returns the index value of the Community List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - community-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterCommunityList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/community-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterCommunityList API operation for FortiOS deletes the specified Community List.
// Returns error for service API and SDK errors.
// See the router - community-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterCommunityList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/community-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterCommunityList API operation for FortiOS gets the Community List
// with the specified index value.
// Returns the requested Community List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - community-list chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterCommunityList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/community-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterRouteMap API operation for FortiOS creates a new Route Map.
// Returns the index value of the Route Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - route-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterRouteMap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/route-map"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterRouteMap API operation for FortiOS updates the specified Route Map.
// Returns the index value of the Route Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - route-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterRouteMap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/route-map"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterRouteMap API operation for FortiOS deletes the specified Route Map.
// Returns error for service API and SDK errors.
// See the router - route-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterRouteMap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/route-map"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterRouteMap API operation for FortiOS gets the Route Map
// with the specified index value.
// Returns the requested Route Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - route-map chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterRouteMap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/route-map"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateRouterRip API operation for FortiOS updates the specified Rip.
// Returns the index value of the Rip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - rip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterRip(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/rip"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterRip API operation for FortiOS deletes the specified Rip.
// Returns error for service API and SDK errors.
// See the router - rip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterRip(mkey string) (err error) {

	//No unset API for router - rip
	return
}

// ReadRouterRip API operation for FortiOS gets the Rip
// with the specified index value.
// Returns the requested Rip value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - rip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterRip(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/rip"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateRouterRipng API operation for FortiOS updates the specified Ripng.
// Returns the index value of the Ripng and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ripng chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterRipng(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ripng"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterRipng API operation for FortiOS deletes the specified Ripng.
// Returns error for service API and SDK errors.
// See the router - ripng chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterRipng(mkey string) (err error) {

	//No unset API for router - ripng
	return
}

// ReadRouterRipng API operation for FortiOS gets the Ripng
// with the specified index value.
// Returns the requested Ripng value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ripng chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterRipng(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ripng"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterStatic API operation for FortiOS creates a new Static.
// Returns the index value of the Static and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterStatic(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/static"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterStatic API operation for FortiOS updates the specified Static.
// Returns the index value of the Static and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterStatic(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/static"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterStatic API operation for FortiOS deletes the specified Static.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterStatic(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/static"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterStatic API operation for FortiOS gets the Static
// with the specified index value.
// Returns the requested Static value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterStatic(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/static"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterPolicy API operation for FortiOS creates a new Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterPolicy API operation for FortiOS updates the specified Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterPolicy API operation for FortiOS deletes the specified Policy.
// Returns error for service API and SDK errors.
// See the router - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterPolicy API operation for FortiOS gets the Policy
// with the specified index value.
// Returns the requested Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterPolicy6 API operation for FortiOS creates a new Policy6.
// Returns the index value of the Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterPolicy6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/policy6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterPolicy6 API operation for FortiOS updates the specified Policy6.
// Returns the index value of the Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterPolicy6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/policy6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterPolicy6 API operation for FortiOS deletes the specified Policy6.
// Returns error for service API and SDK errors.
// See the router - policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterPolicy6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/policy6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterPolicy6 API operation for FortiOS gets the Policy6
// with the specified index value.
// Returns the requested Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - policy6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterPolicy6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/policy6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterStatic6 API operation for FortiOS creates a new Static6.
// Returns the index value of the Static6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterStatic6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/static6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterStatic6 API operation for FortiOS updates the specified Static6.
// Returns the index value of the Static6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterStatic6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/static6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterStatic6 API operation for FortiOS deletes the specified Static6.
// Returns error for service API and SDK errors.
// See the router - static6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterStatic6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/static6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterStatic6 API operation for FortiOS gets the Static6
// with the specified index value.
// Returns the requested Static6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterStatic6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/static6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateRouterOspf API operation for FortiOS updates the specified Ospf.
// Returns the index value of the Ospf and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ospf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterOspf(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterOspf API operation for FortiOS deletes the specified Ospf.
// Returns error for service API and SDK errors.
// See the router - ospf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterOspf(mkey string) (err error) {

	//No unset API for router - ospf
	return
}

// ReadRouterOspf API operation for FortiOS gets the Ospf
// with the specified index value.
// Returns the requested Ospf value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ospf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterOspf(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateRouterOspf6 API operation for FortiOS updates the specified Ospf6.
// Returns the index value of the Ospf6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ospf6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterOspf6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterOspf6 API operation for FortiOS deletes the specified Ospf6.
// Returns error for service API and SDK errors.
// See the router - ospf6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterOspf6(mkey string) (err error) {

	//No unset API for router - ospf6
	return
}

// ReadRouterOspf6 API operation for FortiOS gets the Ospf6
// with the specified index value.
// Returns the requested Ospf6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ospf6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterOspf6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf6"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateRouterBgp API operation for FortiOS updates the specified Bgp.
// Returns the index value of the Bgp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - bgp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterBgp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/bgp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterBgp API operation for FortiOS deletes the specified Bgp.
// Returns error for service API and SDK errors.
// See the router - bgp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterBgp(mkey string) (err error) {

	//No unset API for router - bgp
	return
}

// ReadRouterBgp API operation for FortiOS gets the Bgp
// with the specified index value.
// Returns the requested Bgp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - bgp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterBgp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/bgp"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateRouterIsis API operation for FortiOS updates the specified Isis.
// Returns the index value of the Isis and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - isis chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterIsis(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/isis"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterIsis API operation for FortiOS deletes the specified Isis.
// Returns error for service API and SDK errors.
// See the router - isis chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterIsis(mkey string) (err error) {

	//No unset API for router - isis
	return
}

// ReadRouterIsis API operation for FortiOS gets the Isis
// with the specified index value.
// Returns the requested Isis value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - isis chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterIsis(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/isis"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterMulticastFlow API operation for FortiOS creates a new Multicast Flow.
// Returns the index value of the Multicast Flow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast-flow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterMulticastFlow(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/multicast-flow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterMulticastFlow API operation for FortiOS updates the specified Multicast Flow.
// Returns the index value of the Multicast Flow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast-flow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterMulticastFlow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/multicast-flow"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterMulticastFlow API operation for FortiOS deletes the specified Multicast Flow.
// Returns error for service API and SDK errors.
// See the router - multicast-flow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterMulticastFlow(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/multicast-flow"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterMulticastFlow API operation for FortiOS gets the Multicast Flow
// with the specified index value.
// Returns the requested Multicast Flow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast-flow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterMulticastFlow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/multicast-flow"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateRouterMulticast API operation for FortiOS updates the specified Multicast.
// Returns the index value of the Multicast and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterMulticast(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/multicast"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterMulticast API operation for FortiOS deletes the specified Multicast.
// Returns error for service API and SDK errors.
// See the router - multicast chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterMulticast(mkey string) (err error) {

	//No unset API for router - multicast
	return
}

// ReadRouterMulticast API operation for FortiOS gets the Multicast
// with the specified index value.
// Returns the requested Multicast value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterMulticast(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/multicast"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateRouterMulticast6 API operation for FortiOS updates the specified Multicast6.
// Returns the index value of the Multicast6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterMulticast6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/multicast6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterMulticast6 API operation for FortiOS deletes the specified Multicast6.
// Returns error for service API and SDK errors.
// See the router - multicast6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterMulticast6(mkey string) (err error) {

	//No unset API for router - multicast6
	return
}

// ReadRouterMulticast6 API operation for FortiOS gets the Multicast6
// with the specified index value.
// Returns the requested Multicast6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterMulticast6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/multicast6"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterAuthPath API operation for FortiOS creates a new Auth Path.
// Returns the index value of the Auth Path and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - auth-path chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterAuthPath(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/auth-path"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterAuthPath API operation for FortiOS updates the specified Auth Path.
// Returns the index value of the Auth Path and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - auth-path chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterAuthPath(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/auth-path"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterAuthPath API operation for FortiOS deletes the specified Auth Path.
// Returns error for service API and SDK errors.
// See the router - auth-path chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterAuthPath(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/auth-path"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterAuthPath API operation for FortiOS gets the Auth Path
// with the specified index value.
// Returns the requested Auth Path value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - auth-path chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterAuthPath(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/auth-path"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateRouterSetting API operation for FortiOS updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterSetting API operation for FortiOS deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the router - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterSetting(mkey string) (err error) {

	//No unset API for router - setting
	return
}

// ReadRouterSetting API operation for FortiOS gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateRouterBfd API operation for FortiOS updates the specified Bfd.
// Returns the index value of the Bfd and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - bfd chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterBfd(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/bfd"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterBfd API operation for FortiOS deletes the specified Bfd.
// Returns error for service API and SDK errors.
// See the router - bfd chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterBfd(mkey string) (err error) {

	//No unset API for router - bfd
	return
}

// ReadRouterBfd API operation for FortiOS gets the Bfd
// with the specified index value.
// Returns the requested Bfd value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - bfd chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterBfd(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/bfd"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateRouterBfd6 API operation for FortiOS updates the specified Bfd6.
// Returns the index value of the Bfd6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - bfd6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterBfd6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/bfd6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterBfd6 API operation for FortiOS deletes the specified Bfd6.
// Returns error for service API and SDK errors.
// See the router - bfd6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterBfd6(mkey string) (err error) {

	//No unset API for router - bfd6
	return
}

// ReadRouterBfd6 API operation for FortiOS gets the Bfd6
// with the specified index value.
// Returns the requested Bfd6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - bfd6 chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterBfd6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/bfd6"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}
