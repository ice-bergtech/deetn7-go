package bpv7

import "github.com/fxamacker/cbor/v2"

// https://www.rfc-editor.org/rfc/rfc9171.html#name-primary-bundle-block
type BlockPrimary struct {
	_ struct{} `cbor:",toarray"`
	// Version:
	// An unsigned integer value indicating the version of the Bundle Protocol that constructed this block.
	// The present document describes BPv7.
	// This version number SHALL be represented as a CBOR unsigned integer item.
	Version uint
	//Bundle Processing Control Flags:
	// The bundle processing control flags are discussed in Section 4.2.3.
	BundleFlags BlockProcessingControl
	// CRC Type:
	// CRC type codes are discussed in Section 4.2.1.
	// The CRC type code for the primary block MAY be zero if the bundle contains a BPSec Block Integrity Block [BPSEC] whose target is the primary block; otherwise, the CRC type code for the primary block MUST be non-zero.
	CRCType uint
	// Destination EID:
	// The Destination EID field identifies the bundle endpoint that is the bundle's destination, i.e., the endpoint that contains the node(s) at which the bundle is to be delivered.
	Destination EID
	// Source node ID:
	// The Source node ID field identifies the bundle node at which the bundle was initially transmitted, except that source node ID may be the null endpoint ID in the event that the bundle's source chooses to remain anonymous.
	Source EID
	// Report-to EID:
	// The Report-to EID field identifies the bundle endpoint to which status reports pertaining to the forwarding and delivery of this bundle are to be transmitted.
	ReportTo EID
	// Creation Timestamp:
	// The creation timestamp comprises two unsigned integers that, together with the source node ID and (if the bundle is a fragment) the fragment offset and payload length, serve to identify the bundle. See Section 4.2.7 for the definition of this field.
	Creation CreationTimestamp
	// Lifetime:
	// The Lifetime field is an unsigned integer that indicates the time at which the bundle's payload will no longer be useful, encoded as a number of milliseconds past the creation time. (For high-rate deployments with very brief disruptions, fine-grained expression of bundle lifetime may be useful.) When a bundle's age exceeds its lifetime, bundle nodes need no longer retain or forward the bundle; the bundle SHOULD be deleted from the network.
	// If the asserted lifetime for a received bundle is so lengthy that retention of the bundle until its expiration time might degrade operation of the node at which the bundle is received, or if the BPA of that node determines that the bundle must be deleted in order to prevent network performance degradation (e.g., the bundle appears to be part of a denial-of-service attack), then that BPA MAY impose a temporary overriding lifetime of shorter duration; such an overriding lifetime SHALL NOT replace the lifetime asserted in the bundle but SHALL serve as the bundle's effective lifetime while the bundle resides at that node. Procedures for imposing lifetime overrides are beyond the scope of this specification. For bundles originating at nodes that lack accurate clocks, it is recommended that bundle age be obtained from the Bundle Age extension block (see Section 4.4.2) rather than from the difference between current time and bundle creation time. Bundle lifetime SHALL be represented as a CBOR unsigned integer item.
	Lifetime uint
	// Fragment offset:
	// If and only if the bundle processing control flags of this primary block indicate that the bundle is a fragment, fragment offset SHALL be present in the primary block. Fragment offset SHALL be represented as a CBOR unsigned integer indicating the offset from the start of the original ADU at which the bytes comprising the payload of this bundle were located.
	FragmentOffset uint `cbor:",omitempty"`
	// Total Application Data Unit Length:
	// If and only if the bundle processing control flags of this primary block indicate that the bundle is a fragment, total application data unit length SHALL be present in the primary block. Total application data unit length SHALL be represented as a CBOR unsigned integer indicating the total length of the original ADU of which this bundle's payload is a part.
	ADULength uint `cbor:",omitempty"`
	// CRC:
	// A CRC SHALL be present in the primary block unless the bundle includes a BPSec Block Integrity Block [BPSEC] whose target is the primary block, in which case a CRC MAY be present in the primary block. The length and nature of the CRC SHALL be as indicated by the CRC type. The CRC SHALL be computed over the concatenation of all bytes (including CBOR "break" characters) of the primary block including the CRC field itself, which, for this purpose, SHALL be temporarily populated with all bytes set to zero.
	CRC uint `cbor:",omitempty"`
}

func (block BlockPrimary) SetCRC() error {
	block.CRC = 0
	cbor, err := cbor.Marshal(block)
	// crc32 encode and store

}

func GetCBOR(block BlockPrimary) ([]bytes, error) {
	return cbor.Marshal(block)
}

func (block BlockPrimary) Verify() error {

}
