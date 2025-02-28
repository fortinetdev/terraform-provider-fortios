package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

// sortFirewallLocalinpolicy6Item contains the parameters for each Policy item
type sortFirewallLocalinpolicy6Item struct {
	policyid int
	action   string
}

func getPolicyListFirewallLocalinpolicy6(c *FortiSDKClient, vdomparam string) (itemList []sortFirewallLocalinpolicy6Item, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/local-in-policy6"

	specialparams := "format=policyid|action"

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.SendWithSpecialParams(specialparams, vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp := result["results"].([]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		var members []sortFirewallLocalinpolicy6Item
		for _, v := range mapTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortFirewallLocalinpolicy6Item{
					policyid: int(c["policyid"].(float64)),
					action:   c["action"].(string),
				})
		}

		itemList = members
	}

	return
}

func bPolicyListSortedFirewallLocalinpolicy6(itemList []sortFirewallLocalinpolicy6Item, sortby, sortdirection string, manual_order []string) (bsorted bool) {
	bsorted = true
	if sortby == "policyid" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].policyid > itemList[i+1].policyid {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].policyid < itemList[i+1].policyid {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := strconv.Itoa(item.policyid)
					curItemMap[curKeyValue] = index
				}
				for j := 0; j < len(manual_order)-1; j++ {
					indexL, okL := curItemMap[manual_order[j]]
					indexR, okR := curItemMap[manual_order[j+1]]
					if okL && okR && indexL > indexR {
						bsorted = false
						return
					}
				}
			}
		}
	}

	return
}

func moveAfterFirewallLocalinpolicy6(idbefore, idafter int, c *FortiSDKClient, vdomparam string) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)

	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/local-in-policy6/"
	path += idbefores

	specialparams := "action=move&after="
	specialparams += idafters

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.SendWithSpecialParams(specialparams, vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))
	return
}

func sortPolicyListFirewallLocalinpolicy6(itemList []sortFirewallLocalinpolicy6Item, sortby, sortdirection string, c *FortiSDKClient, vdomparam string, manual_order []string) (err error) {
	var targetItemOrder []sortFirewallLocalinpolicy6Item
	if sortby == "policyid" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].policyid < itemList[j].policyid
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].policyid > itemList[j].policyid
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortFirewallLocalinpolicy6Item)
			for _, item := range itemList {
				curIndex := strconv.Itoa(item.policyid)
				curItemMap[curIndex] = item
			}
			for _, val := range manual_order {
				if item, ok := curItemMap[val]; ok {
					targetItemOrder = append(targetItemOrder, item)
				}
			}
		}
	}

	for i := 0; i < len(targetItemOrder)-1; i++ {
		err = moveAfterFirewallLocalinpolicy6(targetItemOrder[i+1].policyid, targetItemOrder[i].policyid, c, vdomparam)
		if err != nil {
			err = fmt.Errorf("sort err %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateFirewallLocalinpolicy6Sort API operation for FortiOS to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateFirewallLocalinpolicy6Sort(sortby, sortdirection, vdomparam string, manual_order []string) (err error) {
	itemList, err := getPolicyListFirewallLocalinpolicy6(c, vdomparam)
	log.Printf("[INFO] Firewall policy id list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("sort err %s", err)
		return
	}

	bsorted := bPolicyListSortedFirewallLocalinpolicy6(itemList, sortby, sortdirection, manual_order)
	if bsorted == true {
		return
	}

	err = sortPolicyListFirewallLocalinpolicy6(itemList, sortby, sortdirection, c, vdomparam, manual_order)
	if err != nil {
		err = fmt.Errorf("sort err %s", err)
		return
	}

	return
}

// ReadFirewallLocalinpolicy6Sort API operation for FortiOS to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadFirewallLocalinpolicy6Sort(sortby, sortdirection string, vdomparam string, manual_order []string) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getPolicyListFirewallLocalinpolicy6(c, vdomparam)
	if err != nil {
		err = fmt.Errorf("sort err %s", err)
		return
	}

	bsorted := bPolicyListSortedFirewallLocalinpolicy6(itemList, sortby, sortdirection, manual_order)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["policyid"] = item.policyid
		curItemMap["action"] = item.action
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
