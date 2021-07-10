package rand

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"testing"
)

func TestPerm(t *testing.T) {
	s := rand.Perm(5)
	for _, v := range s {
		fmt.Println(v)
	}
	// output:
	// 0
	// 4
	// 2
	// 3
	// 1
}

func TestRead(t *testing.T) {
	s := make([]byte, 5)
	n, err := rand.Read(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("write byte size:", n)
	for _, v := range s {
		fmt.Println(v)
	}
	// output:
	// write byte size: 5
	// 82
	// 253
	// 252
	// 7
	// 33
}

func TestShuffle(t *testing.T) {
	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}

func TestInt(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int())
	}
	// output:
	// 5577006791947779410
	// 8674665223082153551
	// 6129484611666145821
	// 4037200794235010051
	// 3916589616287113937
}

func TestIntn(t *testing.T) {
	n := 10
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(n))
	}
	// output:
	// 1
	// 7
	// 7
	// 9
	// 1
}
