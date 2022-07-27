package fortios

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFortimanagerJSONRPCRequest() *schema.Resource {
	return &schema.Resource{
		Create: handleFMGJSONRPCRequest,
		Read:   readFMGJSONRPCRequest,
		Update: handleFMGJSONRPCRequest,
		Delete: deleteFMGJSONRPCRequest,

		Schema: map[string]*schema.Schema{
			"json_content": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsJSON,
			},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func handleFMGJSONRPCRequest(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("handleFMGJSONRPCRequest")()

	jsonContent := d.Get("json_content").(string)

	input := map[string]interface{}{}
	json.Unmarshal([]byte(jsonContent), &input)

	data, err := c.HandleJSONRPCRequest(input)
	if err != nil {
		res_data := ""
		if data != nil {
			res_data = string(data)
		}
		return fmt.Errorf("Error handling JSON RPC Request : %s\n%s", err, res_data)
	}

	d.SetId("JSONRPC-Requst-" + uuid.New().String())
	d.Set("response", string(data))

	return nil
}

func readFMGJSONRPCRequest(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFMGJSONRPCRequest(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}
