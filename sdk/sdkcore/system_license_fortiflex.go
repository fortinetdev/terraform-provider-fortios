package forticlient

// CreateSystemLicenseFortiFlex API operation for FortiOS download license file by FortiFlex token.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateSystemLicenseFortiFlex(params *map[string]interface{}) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/system/vmlicense/download"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, "")
	return
}
