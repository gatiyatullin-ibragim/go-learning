package foo

import (
	"testing"
)

func TestFooFunc(t *testing.T){
	expectedFooResult := "bar"
	if actualResult := Foo(); actualResult != expectedFooResult{
		t.Errorf("expected %s; got: %s", expectedFooResult, actualResult)
	}
}

