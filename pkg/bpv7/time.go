package bpv7

// https://www.rfc-editor.org/rfc/rfc9171.html#name-dtn-time
// A DTN time is an unsigned integer indicating the number of milliseconds that have elapsed since the DTN Epoch, 2000-01-01 00:00:00 +0000 (UTC).
// DTN time is not affected by leap seconds.
type DTNTimestamp uint

// 4.2.7
type CreationTimestamp struct {
	_                  struct{} `cbor:",toarray"`
	BundleCreationTime DTNTimestamp
	SequenceValue      DTNTimestamp
}
