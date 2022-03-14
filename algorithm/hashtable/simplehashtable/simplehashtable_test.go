package simplehashtable

import (
	"fmt"
	"testing"
)

type user struct {
	Id   int
	Name string
}

func userHash(v interface{}) int {
	intHash := func(i int) int {
		return i
	}
	stringHash := func(s string) int {
		h := 0
		if s != "" {
			for _, r := range []rune(s) {
				h += 31*h + int(r)
			}
		}
		return h
	}
	u := v.(user)
	h := 1
	h += 31*h + intHash(u.Id)
	h += 31*h + stringHash(u.Name)
	return h
}

func TestSimple(t *testing.T) {
	u := user{1, "tom"}
	fmt.Println("user hash:", userHash(u)) // user hash: 123501

	ht := New(userHash)
	fmt.Printf("after init, size: %d, empty: %t\n", ht.size, ht.IsEmpty())

	ht.Add(u)
	fmt.Printf("after add, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u))

	ht.Remove(u)
	fmt.Printf("after remove, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u))

	_ = ht
}

func TestAdd(t *testing.T) {
	u1 := user{1, "tom"}
	u2 := user{1, "to}"}
	u3 := user{1, "to]"}
	fmt.Println(userHash(u1)%16, userHash(u2)%16, userHash(u3)%16)

	ht := New(userHash)
	ht.Add(u1)
	fmt.Printf("add u1, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u1))
	ht.Add(u2)
	fmt.Printf("add u2, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u2))
	ht.Add(u3)
	fmt.Printf("add u3, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u3))
}

func TestRemove(t *testing.T) {
	u1 := user{1, "tom"}
	u2 := user{1, "to}"}
	u3 := user{1, "to]"}
	fmt.Println(userHash(u1)%16, userHash(u2)%16, userHash(u3)%16)

	ht := New(userHash)
	ht.Add(u1)
	fmt.Printf("add u1, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u1))
	ht.Add(u2)
	fmt.Printf("add u2, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u2))
	ht.Add(u3)
	fmt.Printf("add u3, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u3))

	ht.Remove(u1)
	fmt.Printf("remove u1, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u1))
	ht.Remove(u2)
	fmt.Printf("remove u2, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u2))
	ht.Remove(u3)
	fmt.Printf("remove u3, size: %d, empty: %t, contain: %t\n", ht.size, ht.IsEmpty(), ht.Contain(u3))
}

func TestResize(t *testing.T) {
	ht := New(userHash)
	for i := 1; i <= 16; i++ {
		u := user{i, fmt.Sprintf("tom-%d", i)}
		ht.Add(u)
	}
	fmt.Printf("size: %d, empty: %t\n", ht.size, ht.IsEmpty())
}
