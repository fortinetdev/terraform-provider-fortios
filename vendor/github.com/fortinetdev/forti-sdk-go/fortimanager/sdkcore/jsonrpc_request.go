package fmgclient

import (
	"fmt"
	"encoding/json"
)

// HandleJSONRPCRequest is for handling json rpc request
// Input:
//   @p: json rpc data
// Output:
//   @data: return data when executing "read" operation
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) HandleJSONRPCRequest(p map[string]interface{}) ([]byte, error) {
	defer c.Trace("HandleJSONRPCRequest")()

	method := p["method"].(string)
	params := p["params"].([]interface{})[0].(map[string]interface{})

	result, err := c.Do(method, params)
	if err != nil {
		t := "~"
		if result["result"] != nil {
			data := result["result"]
			t1, err1 := json.MarshalIndent(data, "", "    ")
			if err1 == nil {
				t = string(t1)
			}
		}

		return nil, fmt.Errorf("jsonrpc request failed: %s\n%s", err, t)
	}

	if result["result"] == nil {
		return nil, fmt.Errorf("No result fields found:\n%v", result)
	}

	data := result["result"]

	t, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Error handling JSON RPC Request: %s", err)
	}

	return t, nil
}
