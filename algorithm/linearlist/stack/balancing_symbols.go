package stack

import (
	"bufio"
	"errors"
	"go-app/algorithm/linearlist/stack/arraystack"
	"io"
	"os"
)

// 栈应用: 平衡符号(balancing symbols)
const (
	openSymbols  = "([{"
	closeSymbols = ")]}"
)

var (
	EmptyError    = errors.New("check balancing symbols error: stack empty")
	NotMatchError = errors.New("check balancing symbols error: not match")
)

func CheckStringBalancingSymbols(src string) error {
	s := arraystack.New()
	for _, r := range []rune(src) {
		if isOpenSymbol(r) > -1 {
			s.Push(r)
		} else if isCloseSymbol(r) > -1 {
			if s.IsEmpty() {
				return EmptyError
			}
			v, _ := s.Pop()
			if !isMatchSymbol(v.(rune), r) {
				return NotMatchError
			}
		}
	}
	if !s.IsEmpty() {
		return NotMatchError
	}
	s.Clear()
	return nil
}

func CheckFileBalancingSymbols(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	s := arraystack.New()
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if isOpenSymbol(r) > -1 {
			s.Push(r)
		} else if isCloseSymbol(r) > -1 {
			if s.IsEmpty() {
				return EmptyError
			}
			v, _ := s.Pop()
			if !isMatchSymbol(v.(rune), r) {
				return NotMatchError
			}
		}
	}
	if !s.IsEmpty() {
		return NotMatchError
	}
	s.Clear()
	return nil
}

func isOpenSymbol(r rune) int {
	for i, v := range []rune(openSymbols) {
		if r == v {
			return i
		}
	}
	return -1
}

func isCloseSymbol(r rune) int {
	for i, v := range []rune(closeSymbols) {
		if r == v {
			return i
		}
	}
	return -1
}

func isMatchSymbol(openRune, closeRune rune) bool {
	return isOpenSymbol(openRune) == isCloseSymbol(closeRune)
}
