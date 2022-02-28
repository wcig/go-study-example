package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/crypto/ssh"

	"github.com/stretchr/testify/assert"
)

// RSA生成公钥私钥对

const bits = 2048 // 当前已破解到768位秘钥, 所以建议秘钥长度最少1024位. 普遍建议秘钥长度为 2048~4096 位.

func TestGenRsaKeyPair(t *testing.T) {
	privateKey, publicKey, err := GenRsaKeyPair()
	assert.Nil(t, err)
	assert.NotNil(t, privateKey, publicKey)
}

func TestGenRsaKeyPairBytes(t *testing.T) {
	privateKeyBytes, publicKeyBytes, err := GenRsaKeyPairBytes()
	assert.Nil(t, err)
	fmt.Println(string(privateKeyBytes))
	fmt.Println(string(publicKeyBytes))
	// Output:
	// -----BEGIN RSA PRIVATE KEY-----
	// MIIEpgIBAAKCAQEAw1uSkpkuTX58oxcbcbO577pN0L6IALKQKFJJfTLdC1GG2NOu
	// Yei9DVvR4oe2MjOkVrwvR3lHEgAJj3BrBNnBDGDugaEIYXD0fo6NNMIjj5wftnky
	// yV3cPqJMGOMYBic6S6wZuUAE6v6uxzEM6hwIazmjaUL2OTGkxKY7owuAVPUivTLY
	// C0KhoQrMivwaQok9Rv6ggmV0VF9HniDExupcTLlj1RDQ5BHmbFD0UuDu6uyzVFk5
	// bxg7rgLXNJ86b+UxBdL6aRU3gyNWCGZfMxY1LhzRXlrtNu7itgDLmh5N9mlu4JoI
	// mbUOKQjXqpTzj/vGO/dfp5wm+hG6U5+7yneYFwIDAQABAoIBAQCfcO1/mGWh+cnf
	// kcxkUDtA3N0FsXtulk3xmrc/53Ng+XkFHj8Eyd8COB5ktD6PWnaWE6DCxMiseDtV
	// RnsRymc0oygFxFgCrx3c7Jtyo8BCRTncw64fiutmcy+IYFALNZjuide0yfs5kAM1
	// r2LK73YvjStBjdVHeKzxknq553kl3D8dhrL7NcOs1jWhKWzVqoh/xYwENs52v6vo
	// aSU+9okfUw3RJPQ8929k+Xr7ZuWndhRsAfd4GGfkrprBgvPq99s1fL1lgI7EHBQA
	// TU8Hx1E9ADPwQ9QbWxaFsK7RJZBOq30ecd0k3RZLgy1XcMZtxsp6tEPjQB9pUG9D
	// 2I0VefMBAoGBAOIvQQhktaiDWZKHSia5t+GbLCEg/P8aELppl+NUfqhpudB9AO60
	// BY9dvkpAVz35WTafLT4EK6i70d1uAnURx9fYHlm/jWodqRwd6gXgyvwZEAcoi+jP
	// 0E5Zdrl9Kx3+pgg+tIeeOaroMMohyZhltpTzD21C8y+CrRW4z1PIFdvvAoGBAN0c
	// C52POqg8vFGQrfqVfpYNAk1nrIXbW/CkKT78gspfYRKsSNTuNj7gH8wz6rXaPYKE
	// g6Wtxvlr17BCrfmajgqfTyXlp2iY9mWVK/H+VYSRQKDutQlOxYKrEuTYxk5FvV8U
	// idx19St2ogXm/oU2nNWwwHB7J50ixnfwd6Xkyf5ZAoGBAI7g2MzriVFMv+1u5ul1
	// kfHC6up2gzrGhWCsSVVCT8wlFUrdXBRtxcFdr1IKONIeTDvT62lsiwPIuXaOAqFf
	// PQVdMTbyaOtqlyP6xZU2/iKn1lcfRtonRGfp3bSom+8TqGA29CVjTvxmNkPJeiZ5
	// wg9vRwR3MVWGz72UA+vhzeUjAoGBAKpCa0tTPEbKJfTk5KNdtFEVCV3TZEzawixG
	// LKCczpXI5jX7xkV4fPrjmHikzapkuBKlub29V8c+XKMUZWUbtx7E16yF7+giSTGl
	// sWYxg2aGdu61gV/+cSaSVwUzS1iAaTcq2JGTn0ttmYvT+M9ZY2FFLpouqy1b1toK
	// qAdtaoFZAoGBALOAxMY4lAvBdc8cTwTHemcpCjlxBaLTGUKKAPq0FQJM00amVclF
	// ThEB2mmtjmBTfQFeAN7692YvftRgAkQRvFgIpswJ4S6jh7hvgJUXi/kzzCGtoKrd
	// OE8W53TLaOH8fz7qQRQQcHFr2zpmna+Ak9BwHpo8lu71AC6NupVUSTuC
	// -----END RSA PRIVATE KEY-----
	//
	// -----BEGIN RSA PUBLIC KEY-----
	// MIIBCgKCAQEAw1uSkpkuTX58oxcbcbO577pN0L6IALKQKFJJfTLdC1GG2NOuYei9
	// DVvR4oe2MjOkVrwvR3lHEgAJj3BrBNnBDGDugaEIYXD0fo6NNMIjj5wftnkyyV3c
	// PqJMGOMYBic6S6wZuUAE6v6uxzEM6hwIazmjaUL2OTGkxKY7owuAVPUivTLYC0Kh
	// oQrMivwaQok9Rv6ggmV0VF9HniDExupcTLlj1RDQ5BHmbFD0UuDu6uyzVFk5bxg7
	// rgLXNJ86b+UxBdL6aRU3gyNWCGZfMxY1LhzRXlrtNu7itgDLmh5N9mlu4JoImbUO
	// KQjXqpTzj/vGO/dfp5wm+hG6U5+7yneYFwIDAQAB
	// -----END RSA PUBLIC KEY-----
	//
}

func TestGenRsaKeyPairFile(t *testing.T) {
	var (
		privateKeyFileName = "private.pem"
		publicKeyFileName  = "public.pem"
	)

	err := GenRsaKeyPairFile(privateKeyFileName, publicKeyFileName)
	assert.Nil(t, err)
	defer func() {
		_ = os.Remove(privateKeyFileName)
		_ = os.Remove(publicKeyFileName)
	}()
}

// --------------------------------------------------------------------------- //

func GenRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func GenRsaKeyPairBytes() (privateKeyBytes, publicKeyBytes []byte, err error) {
	privateKey, publicKey, err := GenRsaKeyPair()
	if err != nil {
		return nil, nil, err
	}

	privateKeyStream := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyStream,
	}
	privateKeyBytes = pem.EncodeToMemory(privateBlock)

	publicKeyStream := x509.MarshalPKCS1PublicKey(publicKey)
	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyStream,
	}
	publicKeyBytes = pem.EncodeToMemory(publicBlock)
	return privateKeyBytes, publicKeyBytes, nil
}

func GenRsaKeyPairFile(privateKeyFileName, publicKeyFileName string) error {
	privateKeyBytes, publicKeyBytes, err := GenRsaKeyPairBytes()
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(privateKeyFileName, privateKeyBytes, 0600); err != nil {
		return err
	}
	if err = ioutil.WriteFile(publicKeyFileName, publicKeyBytes, 0644); err != nil {
		return err
	}
	return nil
}

func Test2(t *testing.T) {
	privateBytes, err := ioutil.ReadFile("id_rsa_test")
	if err != nil {
		panic(err)
	}
	// privateBlock, _ := pem.Decode(privateBytes)
	// privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	// privateKey, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	// privateKey, err := x509.ParseECPrivateKey(privateBlock.Bytes)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(privateKey)

	privateKey, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(privateKey)
}

func Test3(t *testing.T) {
	publicBytes, err := ioutil.ReadFile("id_rsa_test.pub")
	if err != nil {
		panic(err)
	}

	key, comment, options, rest, err := ssh.ParseAuthorizedKey(publicBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(key, comment, options, rest)
}
