package totp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateValidateCode(t *testing.T) {
	secret := "URMHV4KWPZBXUTTG"
	code, err := GenerateCode(secret)
	if err != nil {
		panic(err)
	}
	fmt.Println(code)

	time.Sleep(time.Second * 3)
	result := ValidateCode(secret, code)
	assert.True(t, result)
}
