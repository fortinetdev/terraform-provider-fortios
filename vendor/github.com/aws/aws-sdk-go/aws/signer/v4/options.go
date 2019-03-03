package v4

import "log"

// WithUnsignedPayload will enable and set the UnsignedPayload field to
// true of the signer.
func WithUnsignedPayload(v4 *Signer) {
	log.Printf("shengh.............v4.go 18")

	v4.UnsignedPayload = true
}
