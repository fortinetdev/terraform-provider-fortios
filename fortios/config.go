package fortios

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func validateConvIPMask2CIDR(oNewIP, oOldIP string) string {
	if oNewIP != oOldIP && strings.Contains(oNewIP, "/") && strings.Contains(oOldIP, " ") {
		line := strings.Split(oOldIP, " ")
		if len(line) >= 2 {
			ip := line[0]
			mask := line[1]
			prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
			return ip + "/" + strconv.Itoa(prefixSize)
		}
	}
	return oOldIP
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
	filter = strings.ReplaceAll(filter, "_", "-")
	filter = strings.ReplaceAll(filter, "fosid", "id")
	filter = strings.ReplaceAll(filter, "&", "&filter=")
	filter = strings.ReplaceAll(filter, ".", "\\.")
	filter = strings.ReplaceAll(filter, "\\", "\\\\")
	filter = "filter=" + filter
	return filter
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
		if v.(string) == "true" {
			sort.Slice(result, func(i, j int) bool {
				v1 := fmt.Sprintf("%v", result[i][fieldname])
				v2 := fmt.Sprintf("%v", result[j][fieldname])

				return sortStringwithNumber(v1) < sortStringwithNumber(v2)
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
	if itable == "true" {
		return true
	}
	return false
}

func convintflist2i(v interface{}) interface{} {
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return v
		}
		return t[0]
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

func i2ss2arrFortiAPIUpgrade(v string, splitv string) bool {
	splitv = strings.ReplaceAll(splitv, "v", "")

	v1, err := version.NewVersion(v)
	if err != nil {
		return false
	}

	v2, err := version.NewVersion(splitv)
	if err != nil {
		return false
	}

	if v1.GreaterThanOrEqual(v2) {
		return true
	}

	return false
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

func parseDownloadedPemCertificate(v string) (*map[string]interface{}, error) {
	der, rest := pem.Decode([]byte(v))

	if der == nil {
		return nil, fmt.Errorf("error, valid certificate not found: %v", rest)
	}

	cert, err := x509.ParseCertificate(der.Bytes)

	if err != nil {
		return nil, fmt.Errorf("error parsing certificate: %v", err)
	}

	o := map[string]interface{}{
		"signature_algorithm":  cert.SignatureAlgorithm.String(),
		"public_key_algorithm": cert.PublicKeyAlgorithm.String(),
		"serial_number":        cert.SerialNumber.String(),
		"is_ca":                cert.IsCA,
		"version":              cert.Version,
		"issuer":               cert.Issuer.String(),
		"subject":              cert.Subject.String(),
		"not_before":           cert.NotBefore.Format(time.RFC3339),
		"not_after":            cert.NotAfter.Format(time.RFC3339),
		"sha1_fingerprint":     fmt.Sprintf("%x", sha1.Sum(cert.Raw)),
		"is_valid":             time.Now().Before(cert.NotAfter) && time.Now().After(cert.NotBefore),
	}

	return &o, nil
}
