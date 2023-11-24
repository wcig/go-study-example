package main

import (
	"bytes"
	"io"
	"strings"
)

// bad
func getBytes1(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return []byte(sanitize1(string(b))), nil
}

func sanitize1(s string) string {
	return strings.TrimSpace(s)
}

// good
func getBytes2(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return sanitize2(b), nil
}

func sanitize2(b []byte) []byte {
	return bytes.TrimSpace(b)
}
