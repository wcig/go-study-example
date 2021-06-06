package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"
)

// rsa-sha256测试
func TestRsaSha256(t *testing.T) {
	privKey, pubKey := GenRsaKey()
	fmt.Println(string(privKey))
	fmt.Println(string(pubKey))

	data := []byte("hello world.")
	sign, err := SignWithRasSha256(data, privKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("sign:", base64.StdEncoding.EncodeToString(sign))

	flag, err := verifySignWithRsaSha256(data, sign, pubKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("verify sign result:", flag)
	// output:
	// -----BEGIN PRIVATE KEY-----
	// MIICXQIBAAKBgQDS7WvQcsA/R5Ho8dgVcQCdhBVVzZ+XfIQ/NfYmjkKZr2rWrDjK
	// s2MAMCWVFcDgTI4UdH98wed8aq9OlYoY7+UrBjZSVTydDbrSC/bvGL9P/0HjzN8E
	// 1YkY70Fp4hS2vZ77FJTVeOBflJvnke3lA+7KaKHnhkUp9eq305Mc55pxXQIDAQAB
	// AoGAUSFxbuQ0fiKRRQCBFP46NlwvuKk20ZXbaFThMUNW6rZGafwkfBMz6UabrhSH
	// 875+h++ned07pOeENr3eOGlNw+PXYZxGGx/ABbLqTS/P7Yt2w3CZCDOHz8+BXFZ4
	// iaA+sbsgbSl96n/HOBLT2UFKmcheLiK6WUPmposALXCNUQECQQDeQHoXcOPJkohZ
	// w9XJW9XDA0uYp74J2RoBKzvwUl7CbvuvqXRsXFFiwji5P7ibSZKOgc0DJQYt6EdL
	// Zq8VS2OFAkEA8vS7sAcehKIZPqup+/rxd5cUmFgsGezlGfJP0ylw/xV99DZIfmOE
	// baDlGL35f2+Q3Ym8bCDpYnSQCbCG//Ch+QJBAIv/TWjqoXHdrantpAMcqAAZWUsl
	// jp2hdbH6y9zL5gjP5BZ4xW9kj5eIBDccnmqi44CE3uD2N3g/cnlaizr8rAkCQB9d
	// w0F1mdYLy3CQ3xA+FKCY67rDX2Y1dFqGMxIS/pigILO/slpG5Nj2db1dAXIXaD+B
	// wca/9QGWBOaLWNKAMZkCQQDAdGDVMoz2uskoCgGF6rRb2gjRT9P7WhAvRaEonh41
	// PjjRQLOcmiBZgvu0J0JYjNu9CjIVLOxAVjZYO7y5/pUi
	// -----END PRIVATE KEY-----
	//
	// -----BEGIN PUBLIC KEY-----
	// MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDS7WvQcsA/R5Ho8dgVcQCdhBVV
	// zZ+XfIQ/NfYmjkKZr2rWrDjKs2MAMCWVFcDgTI4UdH98wed8aq9OlYoY7+UrBjZS
	// VTydDbrSC/bvGL9P/0HjzN8E1YkY70Fp4hS2vZ77FJTVeOBflJvnke3lA+7KaKHn
	// hkUp9eq305Mc55pxXQIDAQAB
	// -----END PUBLIC KEY-----
	//
	// sign: RhYCMYFrHjKTOTeE+UTa6d/QMDXxdwfdRcuRQrLTk7CUoCM1eoXSUkPulFH3Ns3mEcVk+ZNzdBFmsKLDwVB9Gy1ots6vkJ/VQrljTlyT9GrPHyhFODtwqmaSj02JU+OCSHxADp0cmixn+/syG9ZNXIXoxA6ZKnq9Vw27xXHwBEg=
	// verify sign result: true
}

// 生成RSA私钥和公钥
func GenRsaKey() (privKey, pubKey []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	privKey = pem.EncodeToMemory(block)

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKey = pem.EncodeToMemory(block)
	return
}

// 根据私钥对data数据生成签名
func SignWithRasSha256(data []byte, privKeyBytes []byte) (sign []byte, err error) {
	block, _ := pem.Decode(privKeyBytes)
	if block == nil {
		return nil, errors.New("decode private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hashed := sha256.Sum256(data)
	sign, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	return sign, err
}

// 根据data数据和公钥校验签名
func verifySignWithRsaSha256(data []byte, sign []byte, pubKeyBytes []byte) (ok bool, err error) {
	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		return false, errors.New("decode private key error")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], sign)
	if err != nil {
		return false, err
	}
	return true, nil
}
