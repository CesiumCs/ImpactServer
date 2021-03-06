package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func GenerateRsa() *rsa.PrivateKey {
	// Should never error
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	return key
}

func RsaPubToStr(key *rsa.PublicKey) string {
	bytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func RsaToStr(key *rsa.PrivateKey) string {
	bytes := x509.MarshalPKCS1PrivateKey(key)
	return base64.StdEncoding.EncodeToString(bytes)
}

func StrToRsa(keyString string) (key *rsa.PrivateKey, err error) {
	bytes, err := base64.StdEncoding.DecodeString(keyString)
	if err != nil {
		return
	}
	key, err = x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return
	}

	return
}
