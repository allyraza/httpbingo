package assert

import (
	"strings"
	"testing"
)

// Contains asserts if strings contains pattern
func Contains(t *testing.T, want string, got string) {
	if !strings.Contains(want, got) {
		t.Errorf("Want %v, Got %v", want, got)
	}
}

// Equal asserts if want and got are equal
func Equal(t *testing.T, want interface{}, got interface{}) {
	if want != got {
		t.Errorf("Want %v, Got %v\n", want, got)
	}
}

func NotEqual(t *testing.T, want interface{}, got interface{}) {
	if want == got {
		t.Errorf("Want %v, Got %v", want, got)
	}
}
