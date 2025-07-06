package bpv7

// 6.1.1. Bundle Status Reports

type ReasonCode int

const (
	// No additional information.
	ReasonCode_NoInfo = iota
	// Lifetime expired.
	ReasonCode_LifetimeExpired
	// Forwarded over unidirectional link.
	ReasonCode_ForwardedUnidirectional
	// Transmission canceled.
	ReasonCode_Canceled
	// Depleted storage.
	ReasonCode_DepletedStorage
	// Destination endpoint ID unavailable.
	ReasonCode_DestinationIDUnavailable
	// No known route to destination from here.
	ReasonCode_NoKnownRoute
	// No timely contact with next node on route.
	ReasonCode_NoTimelyRoute
	// Block unintelligible.
	ReasonCode_BlockUnintelligible
	// Hop limit exceeded.
	ReasonCode_HopLimitExceeded
	// Traffic pared (e.g., status reports).
	ReasonCode_TrafficPared
	// Block unsupported.
	ReasonCode_BlockUnsupported
)

type BundleStatusReport struct {
	// Bundle Status Information
	BundleStatusInfo BundleStatusInformation
	// Bundle Status Reason Code
	ReasonCode ReasonCode
	// Node that was the source for the bundle being reported on
	SourceNode EID
	// Creation timestamp of the bundle being reported on
	BundleCreationTimestamp DTNTimestamp
	// Only present if fragment
	// Bundle's fragment offset
	FragmentOffset uint
	// Only present if fragment
	// Bundle's payload length
	LengthBundlePayload uint
}

type BundleStatusInformation struct {
	ReportingNodeReceivedBundle  BundleStatusItem
	ReportingNodeForwardedBundle BundleStatusItem
	ReportingNodeDeliveredBundle BundleStatusItem
	ReportingNodeDeletedBundle   BundleStatusItem
}

type BundleStatusItem struct {
	Asserted bool
	//Optional, only set if reporting node recieved bundle AND "Report status time" flag was set to 1 in the bundle processing control flags
	AssertedTime DTNTimestamp
}
