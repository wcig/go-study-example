package main

import (
	"errors"
	"fmt"
)

var baseErr = errors.New("base error")

func main() {
	{
		// %w warp error
		fooErr := fmt.Errorf("foo failed: %w", baseErr)
		fmt.Println(fooErr)
		fmt.Println(errors.Is(fooErr, baseErr))

		err1 := errors.Unwrap(fooErr)
		fmt.Println(err1 == baseErr, errors.Is(err1, baseErr), errors.Is(err1, fooErr))
		// Output:
		// foo failed: base error
		// true
		// true true false
	}

	{
		// %v transform error
		barErr := fmt.Errorf("bar failed: %v", baseErr)
		fmt.Println(barErr)
		fmt.Println(errors.Is(barErr, baseErr))

		err2 := errors.Unwrap(barErr)
		fmt.Println(err2 == baseErr, errors.Is(err2, baseErr), errors.Is(err2, barErr))
		// Output:
		// bar failed: base error
		// false
		// false false false
	}
}
