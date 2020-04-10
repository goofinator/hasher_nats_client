package api

// Hasher wraps the RequestHashes function
type Hasher interface {
	// RequestHashes returns slice of hashes
	// produced from message
	RequestHashes(message []byte) ([][]byte, error)
}
