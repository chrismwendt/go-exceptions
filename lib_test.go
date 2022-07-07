package exceptions

import (
	"fmt"
	"testing"
)

func TestThrow(t *testing.T) {
	func() {
		defer func() {
			if v := recover(); v != nil {
				t.Fatalf("should not have panicked")
			}
		}()
		Throw(nil)
	}()

	func() {
		defer func() {
			if v := recover(); v == nil {
				t.Fatalf("should have panicked")
			}
		}()
		Throw(fmt.Errorf("💥"))
	}()
}

func TestCatch(t *testing.T) {
	err := func() (err error) {
		defer Catch(&err)
		panic("💥")
	}()

	if err == nil {
		t.Fatalf("should have caught an error")
	}
}
