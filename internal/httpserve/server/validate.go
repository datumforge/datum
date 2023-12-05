package server

// GetCertFiles for https enabled echo server and ensure the values are set
func GetCertFiles(certFile, keyFile string) (string, string, error) {
	if certFile == "" {
		return "", "", ErrCertFileMissing
	}

	if keyFile == "" {
		return "", "", ErrKeyFileMissing
	}

	return certFile, keyFile, nil
}
