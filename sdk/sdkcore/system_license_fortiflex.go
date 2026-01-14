package forticlient

// CreateSystemLicenseFortiFlex API operation for FortiOS download license file by FortiFlex token.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateSystemLicenseFortiFlex(params *map[string]interface{}) (output map[string]interface{}, err error) {
	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "POST"
	requestInput.path = "/api/v2/monitor/system/vmlicense/download"
	output = make(map[string]interface{})

	output, err = createUpdate(requestInput)
	return
}
