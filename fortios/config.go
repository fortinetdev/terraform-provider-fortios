package fortios

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateConvIPMask2CIDR(oOldIP, oNewIP string) string {
	if oNewIP != oOldIP {
		newCvt := convertIP(oNewIP)
		oldCvt := convertIP(oOldIP)
		if newCvt == oldCvt {
			return oOldIP
		} else {
			return oNewIP
		}
	}
	return oOldIP
}

func convertIP(inputIP string) string {
	if inputIP == "" {
		return ""
	}
	f := func(c rune) bool {
		return c != '.' && c != '/' && !unicode.IsNumber(c)
	}
	line := strings.FieldsFunc(inputIP, f)
	if len(line) == 1 {
		return line[0]
	} else if len(line) == 2 {
		ip := line[0]
		mask := line[1]
		prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
		return ip + "/" + strconv.Itoa(prefixSize)
	}
	return inputIP
}

func fortiStringValue(t interface{}) string {
	if v, ok := t.(string); ok {
		return v
	} else {
		return ""
	}
}

func fortiIntValue(t interface{}) int {
	if v, ok := t.(float64); ok {
		return int(v)
	} else {
		return 0
	}
}

func escapeFilter(filter string) string {
	var rstSb strings.Builder
	andSlice := strings.Split(filter, "&")

	for i := 0; i < len(andSlice); i++ {
		orSlice := strings.Split(andSlice[i], ",")
		if i > 0 {
			rstSb.WriteString("&")
		}
		rstSb.WriteString("filter=")
		for j := 0; j < len(orSlice); j++ {
			reg := regexp.MustCompile(`([^=*!@><]+)([=*!@><]+)([^=*!@><]+)`)
			match := reg.FindStringSubmatch(orSlice[j])
			if j > 0 {
				rstSb.WriteString(",")
			}
			if match != nil {
				argName := match[1]
				argName = strings.ReplaceAll(argName, "_", "-")
				argName = strings.ReplaceAll(argName, "fosid", "id")
				argName = strings.ReplaceAll(argName, ".", "\\.")
				argName = strings.ReplaceAll(argName, "\\", "\\\\")
				argValue := url.QueryEscape(match[3])
				rstSb.WriteString(argName)
				rstSb.WriteString(match[2])
				rstSb.WriteString(argValue)
			}
		}
	}
	return rstSb.String()
}

func sortStringwithNumber(v string) string {
	i := len(v) - 1
	for ; i >= 0; i-- {
		if '0' > v[i] || v[i] > '9' {
			break
		}
	}
	i++

	b64 := make([]byte, 64/8)
	s64 := v[i:]
	if len(s64) > 0 {
		u64, err := strconv.ParseUint(s64, 10, 64)
		if err == nil {
			binary.BigEndian.PutUint64(b64, u64+1)
		}
	}

	return v[:i] + string(b64)
}

func dynamic_sort_subtable(result []map[string]interface{}, fieldname string, d *schema.ResourceData) {
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		vs := v.(string)
		if vs == "true" || vs == "natural" {
			sort.Slice(result, func(i, j int) bool {
				v1 := fmt.Sprintf("%v", result[i][fieldname])
				v2 := fmt.Sprintf("%v", result[j][fieldname])

				return sortStringwithNumber(v1) < sortStringwithNumber(v2)
			})
		} else if vs == "alphabetical" {
			sort.Slice(result, func(i, j int) bool {
				v1 := fmt.Sprintf("%v", result[i][fieldname])
				v2 := fmt.Sprintf("%v", result[j][fieldname])

				return v1 < v2
			})
		}
	}
}

func fortiAPIPatch(t interface{}) bool {
	if t == nil {
		return false
	} else if _, ok := t.(string); ok {
		return true
	} else if _, ok := t.(float64); ok {
		return true
	} else if _, ok := t.([]interface{}); ok {
		return true
	}

	return false
}

func isImportTable() bool {
	itable := os.Getenv("FORTIOS_IMPORT_TABLE")
	if itable == "false" {
		return false
	}
	return true
}

func convintf2i(v interface{}) interface{} {
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return v
		}
		return t[0]
	} else if t, ok := v.(string); ok {
		if t == "" {
			return 0
		} else if iVal, _ := strconv.Atoi(t); ok {
			return iVal
		}
	}
	return v
}

func convintflist2str(v interface{}) interface{} {
	res := ""
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return res
		}

		bFirst := true
		for _, v1 := range t {
			if t1, ok := v1.(float64); ok {
				if bFirst == true {
					res += strconv.Itoa(int(t1))
					bFirst = false
				} else {
					res += " "
					res += strconv.Itoa(int(t1))
				}
			}
		}
	}
	return res
}

func convmap2str(v, tfv interface{}, target_key string) interface{} {
	if vMap, ok := v.([]interface{}); ok {
		if len(vMap) == 0 {
			return ""
		}
		vsList := make([]string, len(vMap))
		for i, r := range vMap {
			if item, ok := r.(map[string]interface{})[target_key]; ok {
				if ts, ok := item.(string); ok {
					vsList[i] = strings.TrimSpace(fmt.Sprintf("%v", ts))
					if strings.Contains(vsList[i], ",") {
						vsList[i] = "'" + vsList[i] + "'"
					}
				}
			}
		}
		if tfv != nil {
			if tfvs := fmt.Sprintf("%v", tfv); tfvs != "" {
				tfvList := flattenStringList(tfv).([]string)
				if len(tfvList) == len(vsList) {
					tfvDict := make(map[string]bool)
					for _, item := range tfvList {
						tfvDict[item] = true
					}
					for _, item := range vsList {
						item = strings.Trim(item, "'\" ")
						if _, ok := tfvDict[item]; !ok {
							return strings.Join(vsList[:], ", ")
						}
					}
					return tfv
				}
			}
		}
		return strings.Join(vsList[:], ", ")

	}
	return v
}

func flattenStringList(v interface{}) interface{} {
	if v == nil {
		return v
	}
	vsList := []string{}
	if cv, ok := v.(string); ok {
		if strings.Contains(cv, "'") || strings.Contains(cv, "\"") {
			re := regexp.MustCompile(`['\"].*?['\"]`)
			comma := re.FindAllString(cv, -1)
			non_comma := re.Split(cv, -1)
			for i := range non_comma {
				cur_list := strings.Split(non_comma[i], ",")
				for _, item := range cur_list {
					item = strings.TrimSpace(item)
					if item != "" {
						vsList = append(vsList, item)
					}
				}
				if i < len(comma) {
					cur_item := strings.Trim(comma[i], "'\" ")
					vsList = append(vsList, cur_item)
				}
			}
		} else {
			vsList = strings.Split(cv, ",")
		}
	} else if vList, ok := v.([]interface{}); ok {
		for _, item := range vList {
			vsList = append(vsList, fmt.Sprintf("%v", item))
		}
	}
	if len(vsList) == 0 {
		return vsList
	}
	for i, item := range vsList {
		vsList[i] = strings.TrimSpace(item)
	}

	return vsList
}

func checkVersionMatch(v string, new_version_map map[string][]string) (bool, error) {
	v1, err := version.NewVersion(v)
	if err != nil {
		return false, err
	}

	for operator, version_list := range new_version_map {
		if operator == "=" {
			for _, cur_version := range version_list {
				if cur_version == v {
					return true, nil
				}
			}
		} else if operator == ">=" {
			min_version, err := version.NewVersion(version_list[0])
			if err != nil {
				continue
			}
			if v1.GreaterThanOrEqual(min_version) {
				return true, nil
			}
		} else if operator == "<=" {
			max_version, err := version.NewVersion(version_list[0])
			if err != nil {
				continue
			}
			if v1.LessThanOrEqual(max_version) {
				return true, nil
			}
		}
	}
	var supported_version_list []string
	for operator, version_list := range new_version_map {
		supported_version_list = append(supported_version_list, operator+strings.Join(version_list, ","))
	}
	err = fmt.Errorf("requires FortiOS version: %s, your device version is: %s.", strings.Join(supported_version_list, ""), v)
	return false, err
}

func intBetweenWithZero(min, max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if (v >= min && v <= max) || (v == 0) {
			return warnings, errors
		}

		errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d) or equal to 0, got %d", k, min, max, v))

		return warnings, errors
	}
}

func intBetweenWithMinus1(min, max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if (v >= min && v <= max) || (v == -1) {
			return warnings, errors
		}

		errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d) or equal to -1, got %d", k, min, max, v))

		return warnings, errors
	}
}

func toCertFormat(v interface{}) interface{} {
	if t, ok := v.(string); ok {
		if t != "" && !strings.HasPrefix(t, "\"") {
			t = strings.TrimRight(t, "\n")
			return "\"" + t + "\""
		}
	}
	return v
}

func remove_quote(v interface{}) interface{} {
	if t, ok := v.(string); ok {
		t = strings.ReplaceAll(t, "\"", "")
		t = strings.TrimSpace(t)
		return t
	}
	return v
}

func bZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}

func mergeBlock(tf_list, rsp_list []interface{}, tf_mkey, api_mkey string) []interface{} {
	result := []interface{}{}
	mkey_index_map := make(map[string]int)

	// create mkey to index map
	for i, raw := range rsp_list {
		if raw == nil {
			continue
		}
		item := raw.(map[string]interface{})

		keyStr := fmt.Sprintf("%v", item[api_mkey])
		mkey_index_map[keyStr] = i
	}

	// parse item in terraform configuration
	zeroCount := 0
	for _, raw := range tf_list {
		if raw == nil {
			continue
		}
		item := raw.(map[string]interface{})

		keyStr := fmt.Sprintf("%v", item[tf_mkey])

		if rspIndex, ok := mkey_index_map[keyStr]; ok {
			rspItem := rsp_list[rspIndex].(map[string]interface{})
			rspItem["tf_exist"] = true
			result = append(result, rspItem)

			delete(mkey_index_map, keyStr)
		} else {
			if keyStr == "0" {
				zeroCount += 1
				continue
			}
			emptyMap := map[string]interface{}{
				"tf_exist": true,
			}
			result = append(result, emptyMap)
		}
	}

	// add items only in response
	for _, idx := range mkey_index_map {
		v := rsp_list[idx].(map[string]interface{})
		if zeroCount > 0 {
			v["tf_exist"] = true
			zeroCount -= 1
		} else {
			v["tf_exist"] = false
		}
		result = append(result, v)
	}

	return result
}

func isValidSubnet(i interface{}, k string) (warnings []string, errors []error) {
	value := i.(string)
	cidrValue := convertIP(value)
	ip, ipNet, err := net.ParseCIDR(cidrValue)
	if err != nil {
		errors = append(errors, fmt.Errorf(
			"Variable %q is not a valid subnet: %v", k, value,
		))
		return
	}
	// Check if IP equals the network base address
	if !ip.Equal(ipNet.IP) {
		warnings = append(warnings, fmt.Sprintf(
			"Variable %q is not a valid subnet: %v, do you mean %v?", k, value, ipNet,
		))
		return
	}
	return
}
