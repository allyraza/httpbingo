package assert

import (
	"strings"
	"testing"
)

func Contains(t *testing.T, want string, got string) {
	if !strings.Contains(want, got) {
		t.Errorf("Want %v, Got %v", want, got)
	}
}
