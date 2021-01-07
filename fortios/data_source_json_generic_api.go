package fortios

import (
	"fmt"

	forticlient "github.com/fortinetdev/forti-sdk-go/fortios/sdkcore"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceJSONGenericAPI() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJSONGenericAPIRead,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"specialparams": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJSONGenericAPIRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	i := &forticlient.JSONJSONGenericAPI{
		Path:          d.Get("path").(string),
		Method:        "GET",
		Specialparams: d.Get("specialparams").(string),
		Json:          "",
	}

	res, err := c.CreateJSONGenericAPI(i)
	if err != nil {
		return fmt.Errorf("Error reading json generic api: %v", err)
	}

	d.SetId("DataSourceJsonGenericApi" + uuid.New().String())
	d.Set("response", res)

	return nil
}
