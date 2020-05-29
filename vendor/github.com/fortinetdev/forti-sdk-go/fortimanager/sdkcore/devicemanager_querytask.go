package fmgclient

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// QueryResult is for query result
type QueryResult struct {
	err, percent int
	detail       string
}

// QueryTask do the query task work
// Input:
//   @task: task id
//   @timeout: timeout for the querying operation
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) QueryTask(task, timeout int) (err error) {
	defer c.Trace("QueryTask")()
	// interval time: 3s
	for i := 0; i < (timeout*60)/3; i++ {
		time.Sleep(3 * time.Second)
		ret, err := c.queryTask(task)
		if err != nil {
			return err
		}
		if ret == nil {
			continue
		}
		if ret.err != 0 {
			return fmt.Errorf("Executing script failed, detail: %s", ret.detail)
		}
		if ret.percent == 100 {
			log.Printf("Task complete")
			return nil
		}
		log.Printf("Task in progress, %d percentage now", ret.percent)
	}

	// if error here, maybe need more time for execting the script
	return fmt.Errorf("%d minutes has passed, timeout", timeout)
}

func (c *FmgSDKClient) queryTask(task int) (query_ret *QueryResult, err error) {
	if task <= 0 {
		return nil, fmt.Errorf("task id(%d) is not right, should be > 0", task)
	}

	p := map[string]interface{}{
		"url": "/task/task/" + strconv.Itoa(task) + "/line",
	}

	result, err := c.Do("get", p)
	if err != nil {
		return nil, fmt.Errorf("queryTask failed: %s", err)
	}

	// data maybe empty if query too quickly
	array_size := len((result["result"].([]interface{}))[0].(map[string]interface{})["data"].([]interface{}))
	if array_size == 0 {
		return nil, nil
	}

	data := ((result["result"].([]interface{}))[0].(map[string]interface{})["data"].([]interface{}))[array_size-1].(map[string]interface{})
	if data == nil {
		return nil, fmt.Errorf("cannot get the results from the response")
	}

	query_ret = &QueryResult{}
	if data["percent"] != nil {
		query_ret.percent = int(data["percent"].(float64))
	} else {
		return nil, fmt.Errorf("cannot get the percent data")
	}
	if data["err"] != nil {
		query_ret.err = int(data["err"].(float64))
	} else {
		return nil, fmt.Errorf("cannot get the err data")
	}
	if data["detail"] != nil {
		query_ret.detail = data["detail"].(string)
	}

	return query_ret, nil
}
