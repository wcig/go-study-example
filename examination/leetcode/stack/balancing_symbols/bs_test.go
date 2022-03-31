package balancing_symbols

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * 20.有效的括号
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 * 有效字符串需满足：左括号必须用相同类型的右括号闭合。左括号必须以正确的顺序闭合。
 */
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if v, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || v != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func Test(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{
			input:  "()",
			expect: true,
		},
		{
			input:  "()[]{}",
			expect: true,
		},
		{
			input:  "(]",
			expect: false,
		},
		{
			input:  "([)]",
			expect: false,
		},
		{
			input:  "{[]}",
			expect: true,
		},
	}

	for _, v := range cases {
		result := isValid(v.input)
		assert.Equal(t, v.expect, result)
	}
}
