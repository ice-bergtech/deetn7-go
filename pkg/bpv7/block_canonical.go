package bpv7

type Block interface {
	// Must be:
	// 0: no Cyclic Redundancy check (CRC) present
	// 1: standard X-25 CRC-16 [CRC16]
	// 2: standard CRC32C CRC-32 present [RFC4960]
	//
	// https://www.rfc-editor.org/rfc/rfc9171.html#name-crc-type
	GetCRCType() uint
	GetCRC() uint
	GetType() uint
	GetCBOR() interface{}
	GetData() interface{}
}

// https://www.rfc-editor.org/rfc/rfc9171.html#name-canonical-bundle-block-form
type BlockCanonical struct {
	// Block type code:
	// An unsigned integer. Bundle block type code 1 indicates that the block is a Bundle Payload Block. Other block type codes are described in Section 9.1. Block type codes 192 through 255 are not reserved and are available for private and/or experimental use. All other block type code values are reserved for future use.
	BlockType uint
	// Block number:
	// An unsigned integer as discussed in Section 4.1. The block number SHALL be represented as a CBOR unsigned integer.
	BlockNum uint
	// Block processing control flags:
	// As discussed in Section 4.2.4.
	ControlFlags BlockProcessingControl
	// CRC type:
	// As discussed in Section 4.2.1.
	CRCType uint
	// Block-type-specific data:
	// Represented as a single definite-length CBOR byte string, i.e., a CBOR byte string that is not of indefinite length. For each type of block, the block-type-specific data byte string is the serialization, in a block-type-specific manner, of the data conveyed by that type of block; definitions of blocks are required to define the manner in which block-type-specific data are serialized within the block-type-specific data field. For the Bundle Payload Block in particular (block type 1), the block-type-specific data field, termed the "payload", SHALL be an ADU, or some contiguous extent thereof, represented as a definite-length CBOR byte string.
	BlockData interface{}
	// If and only if the value of the CRC type field of this block is non-zero:
	// A CRC. If present, the length and nature of the CRC SHALL be as indicated by the CRC type and the CRC SHALL be computed over the concatenation of all bytes of the block (including CBOR "break" characters) including the CRC field itself, which, for this purpose, SHALL be temporarily populated with all bytes set to zero.
	CRC uint
}
