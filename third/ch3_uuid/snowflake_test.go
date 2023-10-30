package ch3_uuid

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/snowflake"
)

func TestSnowflake(t *testing.T) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", node.Generate().Int64())

	// Output:
	// Int64  ID: 1718899741695152128
	// String ID: 1718899741695152128
	// Base2  ID: 1011111011010110000100011010100111100010000000001000000000000
	// Base64 ID: MTcxODg5OTc0MTY5NTE1MjEyOA==
	// ID Time  : 1698652608690
	// ID Node  : 1
	// ID Step  : 0
	// ID       : 1718899741695152129
}

func Test(t *testing.T) {
	node, _ := snowflake.NewNode(1)
	for i := 0; i < 10; i++ {
		id := node.Generate()
		fmt.Printf("%d %s\n", id.Int64(), id.Base2())
	}
}
