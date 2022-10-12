package totp

import (
	"fmt"
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
)

// TOTP实现 (https://github.com/xlzd/gotp第三方库不建议使用,校验存在问题)

// quick first
func TestTOTP(t *testing.T) {
	const (
		issuer  = "GitHub"
		account = "wcig"
	)

	// generate key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(key)

	// generate code
	secret := key.Secret()
	now := time.Now()
	passcode, err := totp.GenerateCode(secret, now)
	if err != nil {
		panic(err)
	}
	fmt.Println(secret, now.Unix(), passcode)

	// validate code
	check := totp.Validate(passcode, secret)
	fmt.Println(check)

	// Output:
	// otpauth://totp/GitHub:wcig?algorithm=SHA1&digits=6&issuer=GitHub&period=30&secret=QV2EHRLR3KQ52AWESMHKNGEXBZP76P3A
	// QV2EHRLR3KQ52AWESMHKNGEXBZP76P3A 1665561398 644899
	// true
}

// github two-factor authentication (2FA)
func TestGithub2FA(t *testing.T) {
	// otpauth://totp/GitHub:wcig?secret=URMHV4KWPZBXUTTG&issuer=GitHub
	secert := "URMHV4KWPZBXUTTG"
	code, err := totp.GenerateCode(secert, time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(code)

	check := totp.Validate(code, secert)
	fmt.Println(check)

	// Output:
	// 383400
	// true
}
