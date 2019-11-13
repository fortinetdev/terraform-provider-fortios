package fortios

import (
	"fmt"
	"time"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMScriptExecute() *schema.Resource {
	return &schema.Resource{
		Create: createUpdateFMGDVMScriptExecute,
		Read:   readFMGDVMScriptExecute,
		Update: createUpdateFMGDVMScriptExecute,
		Delete: deleteFMGDVMScriptExecute,

		Schema: map[string]*schema.Schema{
			"script_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"target_devname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
		},
	}
}

func createUpdateFMGDVMScriptExecute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createUpdateFMGDVMScriptExecute")()

	var err error
	maxTry := 3

	i := &fmgclient.JSONDVMScriptExecute{
		ScriptName:    d.Get("script_name").(string),
		TargetDevName: d.Get("target_devname").(string),
		Timeout:       d.Get("timeout").(int),
	}

	for j := 0; j < maxTry; j++ {
		err = c.ExecuteDVMScript(i)
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			d.SetId(i.ScriptName)
			return nil
		}
	}

	return fmt.Errorf("Error executing devicemanager script: %s", err)
}

func readFMGDVMScriptExecute(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFMGDVMScriptExecute(d *schema.ResourceData, m interface{}) error {
	return nil
}
