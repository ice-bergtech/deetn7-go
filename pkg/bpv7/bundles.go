package bpv7

type Bundle struct {
	_       struct{} `cbor:",toarray"`
	Primary *BlockPrimary
	Blocks  []*Block // CBOR encoded
	Payload *BlockCanonical
}

func (bun Bundle) Verify() error {

	return nil
}

// https://www.rfc-editor.org/rfc/rfc9171.html#name-block-processing-control-fl
type BlockProcessingControl struct {
	// Bit 0 (the low-order bit, 0x01):    Block must be replicated in every fragment.
	BlockInEveryFragment bool
	// Bit 1 (0x02):    Transmit status report if block can't be processed.
	TransmitFailedStatus bool
	// Bit 2 (0x04):    Delete bundle if block can't be processed.
	DeleteBundleOnFailure bool
	// Bit 3 (0x08):    Reserved.
	_ bool
	// Bit 4 (0x10):    Discard block if it can't be processed.
	DiscardBlockOnFailure bool
	// Bit 5 (0x20):    Reserved.
	_ bool
	// Bit 6 (0x40):    Reserved.
	_ bool
	// Bits 7-63:    Unassigned.
	_ [63 - 7 + 1]bool
}

// https://www.rfc-editor.org/rfc/rfc9171.html#name-bundle-processing-control-f
type BundleProcessControl struct {
	// Bit 0 (the low-order bit, 0x000001):	Bundle is a fragment.
	IsFragment bool
	// Bit 1 (0x000002):    ADU is an administrative record.
	IsPayloadAdministrative bool
	// Bit 2 (0x000004):    Bundle must not be fragmented.
	NoFragment bool
	// Bit 3 (0x000008):    Reserved.
	_ bool
	// Bit 4 (0x000010):    Reserved.
	_ bool
	// Bit 5 (0x000020):    Acknowledgement by application is requested.
	AwkRequested bool
	// Bit 6 (0x000040):    Status time requested in reports.
	StatusTimeRequested bool
	// Bit 7 (0x000080):    Reserved.
	_ bool
	// Bit 8 (0x000100):    Reserved.
	_ bool
	// Bit 9 (0x000200):    Reserved.
	_ bool
	// Bit 10 (0x000400):    Reserved.
	_ bool
	// Bit 11 (0x000800):    Reserved.
	_ bool
	// Bit 12 (0x001000):    Reserved.
	_ bool
	// Bit 13 (0x002000):    Reserved.
	_ bool
	// Request reporting of bundle reception.
	// Bit 14 (0x004000):    Request reporting of bundle reception.
	Reception bool
	// Bit 15 (0x008000):    Reserved.
	_ bool
	// Request reporting of bundle forwarding.
	// Bit 16 (0x010000):    Request reporting of bundle forwarding.
	ReportForwarding bool
	// Request reporting of bundle delivery.
	// Bit 17 (0x020000):    Request reporting of bundle delivery.
	ReportDelivery bool
	// Request reporting of bundle deletion.
	// Bit 18 (0x040000):    Request reporting of bundle deletion.
	ReportDeletion bool
	// Bits 19-20:    Reserved.
	_ [2]bool
	// Bits 21-63:    Unassigned.
	_ [63 - 21 + 1]bool
}
