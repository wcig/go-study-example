package ch3_uuid

import (
	"fmt"
	"strings"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUUID(t *testing.T) {
	fmt.Println(uuid.NewV4().String())                              // 3939f666-893f-46ff-82e9-61db93cd2679
	fmt.Println(strings.ReplaceAll(uuid.NewV4().String(), "-", "")) // 26a44457b2834fe7be4d20e756e032ab
}
