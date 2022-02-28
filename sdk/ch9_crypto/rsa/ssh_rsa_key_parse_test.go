package rsa

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/ssh"
)

// ssh-keygen 生成rsa秘钥对解析 (ssh-keygen -t rsa -C "your comment" -f id_rsa_test)

func TestParsePrivateKey(t *testing.T) {
	privateBytes, err := ioutil.ReadFile("id_rsa_test")
	assert.Nil(t, err)

	privateKey, err := ssh.ParsePrivateKey(privateBytes)
	assert.Nil(t, err)
	assert.NotNil(t, privateKey)
}

func TestParsePublicKey(t *testing.T) {
	publicBytes, err := ioutil.ReadFile("id_rsa_test.pub")
	assert.Nil(t, err)

	publicKey, comment, _, _, err := ssh.ParseAuthorizedKey(publicBytes)
	assert.Nil(t, err)
	assert.NotNil(t, publicKey)
	fmt.Println("comment:", comment)
}
