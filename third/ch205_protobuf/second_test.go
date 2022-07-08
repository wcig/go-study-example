package ch205_protobuf

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// 测试允许null值字段
// .proto生成go文件命令: protoc --go_out=. proto/second.proto (注意: protoc编译器include目录拷贝至环境变量path下)
func TestSecond(t *testing.T) {
	list := []*SecondMsg{
		{
			Id:      1,
			Name:    "tom",
			Number:  &wrapperspb.Int32Value{Value: 100},
			Address: &wrapperspb.StringValue{Value: "beijing"},
			Cell:    &wrapperspb.BytesValue{Value: []byte("hello-你好")},
		},
		{
			Id:      2,
			Name:    "jerry",
			Number:  nil,
			Address: nil,
			Cell:    nil,
		},
	}

	for i, msg := range list {
		fmt.Println(">> index:", i)
		fmt.Println("raw data:", msg)

		// protobuf
		protoData, err := proto.Marshal(msg)
		assert.Nil(t, err)

		var newProtoSecond SecondMsg
		err = proto.Unmarshal(protoData, &newProtoSecond)
		assert.Nil(t, err)
		fmt.Println("proto bytes size:", len(protoData))

		// json
		jsonData, err := json.Marshal(msg)
		assert.Nil(t, err)

		var newJsonSecond SecondMsg
		err = json.Unmarshal(jsonData, &newJsonSecond)
		assert.Nil(t, err)
		fmt.Println("json bytes size:", len(jsonData))
		fmt.Println("json data:", string(jsonData))
	}

	// Output:
	// >> index: 0
	// raw data: id:1  name:"tom"  number:{value:100}  address:{value:"beijing"}  cell:{value:"hello-你好"}
	// proto bytes size: 38
	// json bytes size: 110
	// json data: {"id":1,"name":"tom","number":{"value":100},"address":{"value":"beijing"},"cell":{"value":"aGVsbG8t5L2g5aW9"}}
	// >> index: 1
	// raw data: id:2  name:"jerry"
	// proto bytes size: 9
	// json bytes size: 23
	// json data: {"id":2,"name":"jerry"}
}
