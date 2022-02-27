package elliptic

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wumansgy/goEncrypt"
)

// crypto/elliptic: 在素数域上实现了几个标准椭圆曲线。

// func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error)
// func Marshal(curve Curve, x, y *big.Int) []byte
// func MarshalCompressed(curve Curve, x, y *big.Int) []byte
// func Unmarshal(curve Curve, data []byte) (x, y *big.Int)
// func UnmarshalCompressed(curve Curve, data []byte) (x, y *big.Int)
// type Curve
//    func P224() Curve
//    func P256() Curve
//    func P384() Curve
//    func P521() Curve
// type CurveParams
//    func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int)
//    func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int)
//    func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool
//    func (curve *CurveParams) Params() *CurveParams
//    func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int)
//    func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int)

func TestECC(t *testing.T) {
	curve := elliptic.P256()
	ecPrivateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(ecPrivateKey)
	if err != nil {
		panic(err)
	}
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privKey := pem.EncodeToMemory(block)
	fmt.Println(privKey)

	publicKey := &ecPrivateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	pubKey := pem.EncodeToMemory(block)

	data := []byte("123456")
	encryptData, err := goEncrypt.EccEncrypt(data, pubKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(encryptData), base64.StdEncoding.EncodeToString(encryptData))

	decryptData, err := goEncrypt.EccDecrypt(encryptData, privKey)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, data, decryptData)

	// Output:
	// [45 45 45 45 45 66 69 71 73 78 32 80 82 73 86 65 84 69 32 75 69 89 45 45 45 45 45 10 77 72 99 67 65 81 69 69 73 78 111 79 101 77 114 52 86 103 102 70 75 70 57 105 74 97 111 53 68 111 49 107 101 82 47 43 115 77 86 105 116 98 81 75 82 67 109 117 84 111 83 43 111 65 111 71 67 67 113 71 83 77 52 57 10 65 119 69 72 111 85 81 68 81 103 65 69 99 84 68 109 85 119 85 106 99 84 79 81 47 70 80 113 115 98 87 55 49 73 49 52 67 104 67 68 98 76 81 71 86 65 115 53 110 48 102 116 114 69 84 99 117 105 103 78 55 102 74 52 10 104 102 85 89 102 111 88 65 82 47 76 67 87 100 75 73 86 73 122 103 80 82 87 69 100 56 90 49 97 56 81 66 121 103 61 61 10 45 45 45 45 45 69 78 68 32 80 82 73 86 65 84 69 32 75 69 89 45 45 45 45 45 10]
	// 119 BA1TU1SLIcPihChuqWXn/7b4wX0z4P1Gj6+8jv9Kk/qdJiDPBTKyV35Tl8iXIDyWTTu9YRv/xf3y5mdEzA0emleFPN9opOvEaBwWbfoPlMFz49kRwr8LO7E7hHvTlnF2YHf0P+rQ58ibLDzTTQI9ngHIV1b9JL8=
}
