package main

import (
	"fmt"
	"log"
)

type PageReq struct {
	PageIndex int
	PageSize  int
}

// bad
func validPageReq1(r PageReq) error {
	if r.PageIndex < 0 || r.PageIndex > 100 {
		log.Printf("invalid page index: %d", r.PageIndex)
		return fmt.Errorf("invalid page index: %d", r.PageIndex)
	}
	if r.PageSize < 0 || r.PageSize > 100 {
		log.Printf("invalid page size: %d", r.PageSize)
		return fmt.Errorf("invalid page size: %d", r.PageSize)
	}
	return nil
}

// good
func validPageReq2(r PageReq) error {
	if r.PageIndex < 0 || r.PageIndex > 100 {
		return fmt.Errorf("invalid page index: %d", r.PageIndex)
	}
	if r.PageSize < 0 || r.PageSize > 100 {
		return fmt.Errorf("invalid page size: %d", r.PageSize)
	}
	return nil
}
