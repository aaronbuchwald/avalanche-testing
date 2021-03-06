package certs

import "bytes"

// StaticGeckoCertProvider implements GeckoCertProvider and provides the same cert every time
type StaticGeckoCertProvider struct {
	key  bytes.Buffer
	cert bytes.Buffer
}

// NewStaticGeckoCertProvider creates an instance of StaticGeckoCertProvider using the given key and cert
// Args:
// 	key: The private key that the StaticGeckoCertProvider will return on every call to GetCertAndKey
// 	cert: The cert that will be returned on every call to GetCertAndKey
func NewStaticGeckoCertProvider(key bytes.Buffer, cert bytes.Buffer) *StaticGeckoCertProvider {
	return &StaticGeckoCertProvider{key: key, cert: cert}
}

// GetCertAndKey returns the same cert and key that was configured at the time of construction
func (s StaticGeckoCertProvider) GetCertAndKey() (certPemBytes bytes.Buffer, keyPemBytes bytes.Buffer, err error) {
	return s.cert, s.key, nil
}
