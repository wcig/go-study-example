package bitmap

import (
	"fmt"
	"testing"
)

// bitmap应用: 权限系统
const (
	PermNone   Permission = iota
	PermInsert Permission = 1 << (iota - 1)
	PermDelete
	PermUpdate
	PermSelect
)

type Permission uint8

func (p *Permission) Add(perm Permission) {
	*p = *p | perm
}

func (p *Permission) Remove(perm Permission) {
	*p = *p & (^perm)
}

func (p Permission) String() string {
	return fmt.Sprintf("%04b", p)
}

func TestPermission(t *testing.T) {
	fmt.Println("none:  ", PermNone)
	fmt.Println("insert:", PermInsert)
	fmt.Println("delete:", PermDelete)
	fmt.Println("update:", PermUpdate)
	fmt.Println("select:", PermSelect)

	p := PermNone
	fmt.Println("init:", p)
	p.Add(PermInsert)
	fmt.Println("add insert:", p)
	p.Add(PermDelete)
	fmt.Println("add delete:", p)
	p.Add(PermUpdate)
	fmt.Println("add update:", p)
	p.Add(PermSelect)
	fmt.Println("add select:", p)

	p.Remove(PermSelect)
	fmt.Println("remove select:", p)
	p.Remove(PermUpdate)
	fmt.Println("remove update:", p)
	p.Remove(PermDelete)
	fmt.Println("remove delete:", p)
	p.Remove(PermInsert)
	fmt.Println("remove insert:", p)

	// Output:
	// none:   0000
	// insert: 0001
	// delete: 0010
	// update: 0100
	// select: 1000
	// init: 0000
	// add insert: 0001
	// add delete: 0011
	// add update: 0111
	// add select: 1111
	// remove select: 0111
	// remove update: 0011
	// remove delete: 0001
	// remove insert: 0000
}
