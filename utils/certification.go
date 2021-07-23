package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	ErrPemDecodeFailed = errors.New("pem decode failed")
)

func ParsePrivateKey(privateKeyStr string) (*rsa.PrivateKey, error) {
	apiPrivateKeyBlock, rest := pem.Decode([]byte(privateKeyStr))
	_ = rest
	if apiPrivateKeyBlock == nil {
		return nil, ErrPemDecodeFailed
	}
	apiPrivateKey, err := x509.ParsePKCS8PrivateKey(apiPrivateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return apiPrivateKey.(*rsa.PrivateKey), nil
}

func ParseCertification(certKey string) (*x509.Certificate, error) {
	apiCertBlock, _ := pem.Decode([]byte(certKey))
	if apiCertBlock == nil {
		return nil, ErrPemDecodeFailed
	}
	apiCert, err := x509.ParseCertificate(apiCertBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return apiCert, nil
}
