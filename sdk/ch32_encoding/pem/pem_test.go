package pem

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"testing"
)

// encoding/pem
// 程序包pem实现了PEM数据编码，该数据编码起源于“隐私增强邮件”。 今天，PEM编码最常见的用途是在TLS密钥和证书中。 请参阅RFC 1421。

// func Decode(data []byte) (p *Block, rest []byte)
// 解码将在输入中找到下一个PEM格式化的块（证书，私钥等）。 它返回该块和输入的其余部分。 如果未找到PEM数据，则p为nil，并且整个输入均以静态返回。
func TestDecode(t *testing.T) {
	var pubPEMData = []byte(`
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAlRuRnThUjU8/prwYxbty
WPT9pURI3lbsKMiB6Fn/VHOKE13p4D8xgOCADpdRagdT6n4etr9atzDKUSvpMtR3
CP5noNc97WiNCggBjVWhs7szEe8ugyqF23XwpHQ6uV1LKH50m92MbOWfCtjU9p/x
qhNpQQ1AZhqNy5Gevap5k8XzRmjSldNAFZMY7Yv3Gi+nyCwGwpVtBUwhuLzgNFK/
yDtw2WcWmUU7NuC8Q6MWvPebxVtCfVp/iQU6q60yyt6aGOBkhAX0LpKAEhKidixY
nP9PNVBvxgu3XZ4P36gZV6+ummKdBVnc3NqwBLu5+CcdRdusmHPHd5pHf4/38Z3/
6qU2a/fPvWzceVTEgZ47QjFMTCTmCwNt29cvi7zZeQzjtwQgn4ipN9NibRH/Ax/q
TbIzHfrJ1xa2RteWSdFjwtxi9C20HUkjXSeI4YlzQMH0fPX6KCE7aVePTOnB69I/
a9/q96DiXZajwlpq3wFctrs1oXqBp5DVrCIj8hU2wNgB7LtQ1mCtsYz//heai0K9
PhE4X6hiE0YmeAZjR0uHl8M/5aW9xCoJ72+12kKpWAa0SFRWLy6FejNYCYpkupVJ
yecLk/4L1W0l6jQQZnWErXZYe0PNFcmwGXy1Rep83kfBRNKRy5tvocalLlwXLdUk
AIU+2GKjyT3iMuzZxxFxPFMCAwEAAQ==
-----END PUBLIC KEY-----
and some more`)

	block, rest := pem.Decode(pubPEMData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got a %T, with remaining data: %q\n", pub, rest)
	// output:
	// Got a *rsa.PublicKey, with remaining data: "and some more"
}

// func Encode(out io.Writer, b *Block) error
// 编码将b的PEM编码写出。
func TestEncode(t *testing.T) {
	block := &pem.Block{
		Type: "MESSAGE",
		Headers: map[string]string{
			"Animal": "Gopher",
		},
		Bytes: []byte("test"),
	}

	if err := pem.Encode(os.Stdout, block); err != nil {
		log.Fatal(err)
	}
	// output:
	// -----BEGIN MESSAGE-----
	// Animal: Gopher
	//
	// dGVzdA==
	// -----END MESSAGE-----
}

// func EncodeToMemory(b *Block) []byte
// EncodeToMemory返回b的PEM编码。
// 如果b的标头无效且无法编码，则EncodeToMemory返回nil。 如果重要的是报告有关此错误情况的详细信息，请改用“编码”。
func TestEncodeToMemory(t *testing.T) {
	block := &pem.Block{
		Type: "MESSAGE",
		Headers: map[string]string{
			"Animal": "Gopher",
		},
		Bytes: []byte("test"),
	}

	b := pem.EncodeToMemory(block)
	fmt.Println(string(b))
	// output:
	// -----BEGIN MESSAGE-----
	// Animal: Gopher
	//
	// dGVzdA==
	// -----END MESSAGE-----
}
