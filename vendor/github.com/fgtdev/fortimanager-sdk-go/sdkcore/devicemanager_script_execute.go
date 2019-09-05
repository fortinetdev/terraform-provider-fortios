package fortimngclient

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type JSONDVMScriptExecute struct {
	ScriptName    string `json:"script"`
	TargetDevName string `json:"name"`
	Timeout       int
}

// Create and Update function
func (c *FortiMngClient) ExecuteDVMScript(params *JSONDVMScriptExecute) (err error) {
	defer c.Trace("ExecuteDVMScript")()

	d := map[string]interface{}{
		"adom": "root",
		"scope": map[string]string{
			"vdom": "root",
			"name": params.TargetDevName,
		},
		"script": params.ScriptName,
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/dvmdb/adom/root/script/execute",
	}

	_, err = c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("ExecuteDVMScript failed: %s", err)
		return
	}

	// This piece of code works well, but there is sleep() here, should be checked whether it is needed or not.
	result, err := c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("ExecuteDVMScript failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	if data["task"] != nil {
		taskid := int(data["task"].(float64))
		// interval time: 3s
		for i := 0; i < (params.Timeout*60)/3; i++ {
			time.Sleep(3 * time.Second)
			percent, err := c.QueryTask(taskid)
			if err != nil {
				return err
			}
			if percent == 100 {
				log.Printf("[complete] script '%s' execution success", params.ScriptName)
				return nil
			}
			log.Printf("[in progress] script '%s' execution %d percentage complete", params.ScriptName, percent)
		}

		// if error here, maybe need more time for execting the script
		return fmt.Errorf("%d minutes has passed, timeout", params.Timeout)

	} else {
		err = fmt.Errorf("cannot get the task id after executing the script")
		return
	}

	return
}

func (c *FortiMngClient) QueryTask(task int) (percent int, err error) {
	defer c.Trace("QueryTask")()

	if task <= 0 {
		err = fmt.Errorf("task id is not right, please have a check:%s", err)
	}

	p := map[string]interface{}{
		"url": "/task/task/" + strconv.Itoa(task) + "/line",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("QueryTask failed :%s", err)
		return
	}

	data := ((result["result"].([]interface{}))[0].(map[string]interface{})["data"].([]interface{}))[0].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	if data["percent"] != nil {
		percent = int(data["percent"].(float64))
	} else {
		err = fmt.Errorf("cannot get the percent data of executing the script")
		return
	}

	return
}
