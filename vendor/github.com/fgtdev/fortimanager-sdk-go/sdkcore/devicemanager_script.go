package fortimngclient

import (
	"fmt"
	"strconv"
)

type JSONDVMScript struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Content     string `json:"content"`
	Type        string `json:"type"`
}

type JSONDVMScriptExecute struct {
	Target     string `json:"name"`
	ScriptName string `json:"script"`
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateDVMScript(params *JSONDVMScript, method string) (err error) {
	defer c.Trace("CreateUpdateDVMScript")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/dvmdb/adom/root/script",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateDVMScript failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadDVMScript(id string) (out *JSONDVMScript, err error) {
	defer c.Trace("ReadDVMScript")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/root/script/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadDVMScript failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONDVMScript{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["desc"] != nil {
		out.Description = data["desc"].(string)
	}
	if data["content"] != nil {
		out.Content = data["content"].(string)
	}

	return
}

func (c *FortiMngClient) DeleteDVMScript(id string) (err error) {
	defer c.Trace("DeleteDVMScript")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/root/script/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteDVMScript failed :%s", err)
		return
	}

	return
}

func (c *FortiMngClient) ExecuteDVMScript(params *JSONDVMScriptExecute) (err error) {
	defer c.Trace("ExecuteDVMScript")()

	d := map[string]interface{}{
		"adom": "root",
		"scope": map[string]string{
			"vdom": "root",
			"name": params.Target,
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

	/*
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
				//query result, how long should be wait? or need another way to this goal .....
				time.Sleep(3 * time.Second)
				percent, err := c.QueryTask(taskid)
				if err != nil {
					return err
				}

				log.Printf("Execute script '%s' success and is %d percentage now", params.ScriptName, percent)
			} else {
				err = fmt.Errorf("cannot get the task id after executing the script")
				return
			}
	*/

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
