// +build go1.5

package v4

import (
	"log"
	"net/url"
	"strings"
)

func getURIPath(u *url.URL) string {
	log.Printf("shengh.............v4.go 17")

	var uri string

	if len(u.Opaque) > 0 {
		uri = "/" + strings.Join(strings.Split(u.Opaque, "/")[3:], "/")
	} else {
		uri = u.EscapedPath()
	}

	if len(uri) == 0 {
		uri = "/"
	}

	return uri
}
