package ch205_protobuf

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	first := &FirstMsg{
		Id:   1,
		Name: "tom",
	}

	// protobuf
	protoData, err := proto.Marshal(first)
	assert.Nil(t, err)

	var newProtoFirst FirstMsg
	err = proto.Unmarshal(protoData, &newProtoFirst)
	assert.Nil(t, err)
	fmt.Println("proto data:", first)
	fmt.Println("proto bytes:", len(protoData))

	// json
	jsonData, err := json.Marshal(first)
	assert.Nil(t, err)

	var newJsonFirst FirstMsg
	err = json.Unmarshal(jsonData, &newJsonFirst)
	assert.Nil(t, err)
	fmt.Println("json:", first)
	fmt.Println("json bytes:", len(jsonData))

	// Output:
	// proto data: id:1 name:"tom"
	// proto bytes: 7
	// json: id:1 name:"tom"
	// json bytes: 21
}
