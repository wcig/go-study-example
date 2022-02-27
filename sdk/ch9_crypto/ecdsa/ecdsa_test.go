package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// crypto/ecdsa: 实现了ECC签名算法

// func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
// func SignASN1(rand io.Reader, priv *PrivateKey, hash []byte) ([]byte, error)
// func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
// func VerifyASN1(pub *PublicKey, hash, sig []byte) bool
// type PrivateKey
//    func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)
//    func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool
//    func (priv *PrivateKey) Public() crypto.PublicKey
//    func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
// type PublicKey
//    func (pub *PublicKey) Equal(x crypto.PublicKey) bool

func TestEcdsaSign(t *testing.T) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	data := []byte("123456")
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	fmt.Println(s)

	check := ecdsa.Verify(publicKey, data, r, s)
	assert.True(t, check)
	// Output:
	// 85812117771020857812328478020250853263393707781944532767897547445834195597521
	// 73595711572807974164527307495041575520888129750040143204257559408373336421866
}
