package bpv7

// BlockType == 6
type ExtBlockPreviousNode struct {
	PreviousNode EID
}

// BlockType == 7
type ExtBlockBundleAge struct {
	Age uint
}

// BlockType == 10
type ExtBlockHopCount struct {
	HopLimit byte
	HopCount byte
}
