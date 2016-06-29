package tests

import "testing"

func Fail(t *testing.T) {
	t.Errorf("Whoops - I am not listed in vim's QuickFix window")
}
