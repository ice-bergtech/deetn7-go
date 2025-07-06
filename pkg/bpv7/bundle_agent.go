package bpv7

// https://www.rfc-editor.org/rfc/rfc9171.html#name-bundle-processing

type BundleTransmissionRequest struct{}

// https://www.rfc-editor.org/rfc/rfc9171.html#name-bundle-transmission
func GenerateBundle(btr BundleTransmissionRequest, original Bundle) (*Bundle, error) {
	// The source node ID of the bundle MUST be either (a) the null endpoint ID, indicating that the source of the bundle is anonymous or (b) the EID of a singleton endpoint whose only member is the node of which the BPA is a component.
	return nil, nil
}

// 5.2 Bundle Transmission
// https://www.rfc-editor.org/rfc/rfc9171.html#name-bundle-transmission
func Transmission(bun Bundle, btr BundleTransmissionRequest) error {
	// Step 1:
	// Transmission of the bundle is initiated. An outbound bundle MUST be created per the parameters of the bundle transmission request, with the retention constraint "Dispatch pending".
	outboundBundle, err := GenerateBundle(btr, bun)

	// Step 2:
	// Processing proceeds from Step 1 of Section 5.4.
	return Forwarding(bun)
}

// 5.3 Bundle Dispatching
// https://www.rfc-editor.org/rfc/rfc9171.html#name-bundle-dispatching
func Dispatching(bun Bundle) error {}

// 5.4 Bundle Forwarding
// https://www.rfc-editor.org/rfc/rfc9171.html#name-bundle-forwarding
func Forwarding(bun Bundle) error {}

// 5.5 Bundle Expiration
