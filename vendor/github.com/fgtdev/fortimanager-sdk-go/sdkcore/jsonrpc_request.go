package fmgclient

import (
	"fmt"
)

// HandleJSONRPCRequest is for handling json rpc request
// Input:
//   @p: json rpc data
// Output:
//   @data: return data when executing "read" operation
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) HandleJSONRPCRequest(p map[string]interface{}) (data map[string]interface{}, err error) {
	defer c.Trace("HandleJSONRPCRequest")()

	method := p["method"].(string)
	params := p["params"].([]interface{})[0].(map[string]interface{})

	if method != "get" {
		_, err = c.Do(method, params)
		if err != nil {
			return nil, fmt.Errorf("HandleJSONRPCRequest failed: %s", err)
		}
		return nil, nil
	} else {
		result, err := c.Do(method, params)
		if err != nil {
			return nil, fmt.Errorf("HandleJSONRPCRequest(READ) failed :%s", err)
		}

		data = (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
		if data == nil {
			return nil, fmt.Errorf("cannot get the results from the response")
		}
		return data, nil
	}
}
