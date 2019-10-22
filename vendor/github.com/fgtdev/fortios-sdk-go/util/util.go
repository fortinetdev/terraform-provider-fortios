package util

import "log"

// HttpStatus2Str converts http status code to string
func HttpStatus2Str(c int) (s string) {
	switch c {
	case 200:
		s = "OK: Request returns successful"
	case 400:
		s = "Bad Request: Request cannot be processed by the API"
	case 401:
		s = "Not Authorized: Request without successful login session"
	case 403:
		s = "Forbidden: Request is missing CSRF token or administrator is missing access profile permissions"
	case 404:
		s = "Resource Not Found: Unable to find the specified resource"
	case 405:
		s = "Method Not Allowed: Specified HTTP method is not allowed for this resource"
	case 413:
		s = "Request Entity Too Large: Request cannot be processed due to large entity"
	case 424:
		s = "Failed Dependency: Fail dependency can be duplicate resource, missing required parameter, missing required attribute, invalid attribute value"
	case 429:
		s = "Access temporarily blocked: Maximum failed authentications reached. The offended source is temporarily blocked for certain amount of time."
	case 500:
		s = "Internal Server Error: Internal error when processing the request"
	default:
		log.Printf("[op2Str][Warning] not support number")
	}

	return
}
